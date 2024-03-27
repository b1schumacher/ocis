package svc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"path"
	"time"

	gateway "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	cs3rpc "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	rpc "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	collaboration "github.com/cs3org/go-cs3apis/cs3/sharing/collaboration/v1beta1"
	link "github.com/cs3org/go-cs3apis/cs3/sharing/link/v1beta1"
	storageprovider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	types "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	"github.com/cs3org/reva/v2/pkg/rgrpc/todo/pool"
	"github.com/cs3org/reva/v2/pkg/share"
	"github.com/cs3org/reva/v2/pkg/storagespace"
	"github.com/cs3org/reva/v2/pkg/utils"
	libregraph "github.com/owncloud/libre-graph-api-go"
	"github.com/owncloud/ocis/v2/ocis-pkg/log"
	"github.com/owncloud/ocis/v2/services/graph/pkg/config"
	"github.com/owncloud/ocis/v2/services/graph/pkg/errorcode"
	"github.com/owncloud/ocis/v2/services/graph/pkg/identity"
	"github.com/owncloud/ocis/v2/services/graph/pkg/linktype"
	"github.com/owncloud/ocis/v2/services/graph/pkg/unifiedrole"
)

// BaseGraphService implements a couple of helper functions that are
// shared between the different graph services
type BaseGraphService struct {
	logger          *log.Logger
	gatewaySelector pool.Selectable[gateway.GatewayAPIClient]
	identityCache   identity.IdentityCache
	config          *config.Config
}

func (g BaseGraphService) getSpaceRootPermissions(ctx context.Context, spaceID *storageprovider.StorageSpaceId) ([]libregraph.Permission, error) {
	gatewayClient, err := g.gatewaySelector.Next()

	if err != nil {
		g.logger.Debug().Err(err).Msg("selecting gatewaySelector failed")
		return nil, err
	}
	space, err := utils.GetSpace(ctx, spaceID.GetOpaqueId(), gatewayClient)
	if err != nil {
		return nil, err
	}

	return g.cs3SpacePermissionsToLibreGraph(ctx, space, APIVersion_1_Beta_1), nil
}

func (g BaseGraphService) getDriveItem(ctx context.Context, ref storageprovider.Reference) (*libregraph.DriveItem, error) {
	gatewayClient, err := g.gatewaySelector.Next()
	if err != nil {
		return nil, err
	}

	res, err := gatewayClient.Stat(ctx, &storageprovider.StatRequest{Ref: &ref})
	if err != nil {
		return nil, err
	}
	if res.GetStatus().GetCode() != cs3rpc.Code_CODE_OK {
		refStr, _ := storagespace.FormatReference(&ref)
		return nil, fmt.Errorf("could not stat %s: %s", refStr, res.GetStatus().GetMessage())
	}
	return cs3ResourceToDriveItem(g.logger, res.GetInfo())
}

func (g BaseGraphService) cs3SpacePermissionsToLibreGraph(ctx context.Context, space *storageprovider.StorageSpace, apiVersion APIVersion) []libregraph.Permission {
	if space.Opaque == nil {
		return nil
	}
	logger := g.logger.SubloggerWithRequestID(ctx)

	var permissionsMap map[string]*storageprovider.ResourcePermissions
	opaqueGrants, ok := space.Opaque.Map["grants"]
	if ok {
		err := json.Unmarshal(opaqueGrants.Value, &permissionsMap)
		if err != nil {
			logger.Debug().
				Err(err).
				Interface("space", space.Root).
				Bytes("grants", opaqueGrants.Value).
				Msg("unable to parse space: failed to read spaces grants")
		}
	}
	if len(permissionsMap) == 0 {
		return nil
	}

	var permissionsExpirations map[string]*types.Timestamp
	opaqueGrantsExpirations, ok := space.Opaque.Map["grants_expirations"]
	if ok {
		err := json.Unmarshal(opaqueGrantsExpirations.Value, &permissionsExpirations)
		if err != nil {
			logger.Debug().
				Err(err).
				Interface("space", space.Root).
				Bytes("grants_expirations", opaqueGrantsExpirations.Value).
				Msg("unable to parse space: failed to read spaces grants expirations")
		}
	}

	var groupsMap map[string]struct{}
	opaqueGroups, ok := space.Opaque.Map["groups"]
	if ok {
		err := json.Unmarshal(opaqueGroups.Value, &groupsMap)
		if err != nil {
			logger.Debug().
				Err(err).
				Interface("space", space.Root).
				Bytes("groups", opaqueGroups.Value).
				Msg("unable to parse space: failed to read spaces groups")
		}
	}

	permissions := make([]libregraph.Permission, 0, len(permissionsMap))
	for id, perm := range permissionsMap {
		// This temporary variable is necessary since we need to pass a pointer to the
		// libregraph.Identity and if we pass the pointer from the loop every identity
		// will have the same id.
		tmp := id
		isGroup := false
		var identity libregraph.Identity
		var err error
		var p libregraph.Permission
		if _, ok := groupsMap[id]; ok {
			identity, err = groupIdToIdentity(ctx, g.identityCache, tmp)
			if err != nil {
				g.logger.Warn().Str("groupid", tmp).Msg("Group not found by id")
			}
			isGroup = true
		} else {
			identity, err = userIdToIdentity(ctx, g.identityCache, tmp)
			if err != nil {
				g.logger.Warn().Str("userid", tmp).Msg("User not found by id")
			}
		}
		switch apiVersion {
		case APIVersion_1:
			var identitySet libregraph.IdentitySet
			if isGroup {
				identitySet.SetGroup(identity)
			} else {
				identitySet.SetUser(identity)
			}
			p.SetGrantedToIdentities([]libregraph.IdentitySet{identitySet})
		case APIVersion_1_Beta_1:
			var identitySet libregraph.SharePointIdentitySet
			if isGroup {
				identitySet.SetGroup(identity)
			} else {
				identitySet.SetUser(identity)
			}
			p.SetId(identitySetToSpacePermissionID(identitySet))
			p.SetGrantedToV2(identitySet)
		}

		if exp := permissionsExpirations[id]; exp != nil {
			p.SetExpirationDateTime(time.Unix(int64(exp.GetSeconds()), int64(exp.GetNanos())))
		}

		if role := unifiedrole.CS3ResourcePermissionsToUnifiedRole(*perm, unifiedrole.UnifiedRoleConditionOwner, false); role != nil {
			switch apiVersion {
			case APIVersion_1:
				if r := unifiedrole.GetLegacyName(*role); r != "" {
					p.SetRoles([]string{r})
				}
			case APIVersion_1_Beta_1:
				p.SetRoles([]string{role.GetId()})
			}
		}

		permissions = append(permissions, p)
	}
	return permissions
}

func (g BaseGraphService) libreGraphPermissionFromCS3PublicShare(createdLink *link.PublicShare) (*libregraph.Permission, error) {
	webURL, err := url.Parse(g.config.Spaces.WebDavBase)
	if err != nil {
		g.logger.Error().
			Err(err).
			Str("url", g.config.Spaces.WebDavBase).
			Msg("failed to parse webURL base url")
		return nil, err
	}
	lt, actions := linktype.SharingLinkTypeFromCS3Permissions(createdLink.GetPermissions())
	perm := libregraph.NewPermission()
	perm.Id = libregraph.PtrString(createdLink.GetId().GetOpaqueId())
	perm.Link = &libregraph.SharingLink{
		Type:                  lt,
		PreventsDownload:      libregraph.PtrBool(false),
		LibreGraphDisplayName: libregraph.PtrString(createdLink.GetDisplayName()),
		LibreGraphQuickLink:   libregraph.PtrBool(createdLink.GetQuicklink()),
	}
	perm.LibreGraphPermissionsActions = actions
	webURL.Path = path.Join(webURL.Path, "s", createdLink.GetToken())
	perm.Link.SetWebUrl(webURL.String())

	// set expiration date
	if createdLink.GetExpiration() != nil {
		perm.SetExpirationDateTime(cs3TimestampToTime(createdLink.GetExpiration()).UTC())
	}

	perm.SetHasPassword(createdLink.GetPasswordProtected())

	return perm, nil
}

func (g BaseGraphService) listUserShares(ctx context.Context, filters []*collaboration.Filter, driveItems driveItemsByResourceID) (driveItemsByResourceID, error) {
	gatewayClient, err := g.gatewaySelector.Next()
	if err != nil {
		g.logger.Error().Err(err).Msg("could not select next gateway client")
		return driveItems, errorcode.New(errorcode.GeneralException, err.Error())
	}

	concreteFilters := []*collaboration.Filter{
		share.UserGranteeFilter(),
		share.GroupGranteeFilter(),
	}
	concreteFilters = append(concreteFilters, filters...)

	lsUserSharesRequest := collaboration.ListSharesRequest{
		Filters: concreteFilters,
	}

	lsUserSharesResponse, err := gatewayClient.ListShares(ctx, &lsUserSharesRequest)
	if err != nil {
		return driveItems, errorcode.New(errorcode.GeneralException, err.Error())
	}
	if statusCode := lsUserSharesResponse.GetStatus().GetCode(); statusCode != rpc.Code_CODE_OK {
		return driveItems, errorcode.New(cs3StatusToErrCode(statusCode), lsUserSharesResponse.Status.Message)
	}
	driveItems, err = g.cs3UserSharesToDriveItems(ctx, lsUserSharesResponse.Shares, driveItems)
	if err != nil {
		return driveItems, errorcode.New(errorcode.GeneralException, err.Error())
	}
	return driveItems, nil
}

func (g BaseGraphService) listPublicShares(ctx context.Context, filters []*link.ListPublicSharesRequest_Filter, driveItems driveItemsByResourceID) (driveItemsByResourceID, error) {

	gatewayClient, err := g.gatewaySelector.Next()
	if err != nil {
		g.logger.Error().Err(err).Msg("could not select next gateway client")
		return driveItems, errorcode.New(errorcode.GeneralException, err.Error())
	}

	var concreteFilters []*link.ListPublicSharesRequest_Filter
	concreteFilters = append(concreteFilters, filters...)

	req := link.ListPublicSharesRequest{
		Filters: concreteFilters,
	}

	lsPublicSharesResponse, err := gatewayClient.ListPublicShares(ctx, &req)
	if err != nil {
		return driveItems, errorcode.New(errorcode.GeneralException, err.Error())
	}
	if statusCode := lsPublicSharesResponse.GetStatus().GetCode(); statusCode != rpc.Code_CODE_OK {
		return driveItems, errorcode.New(cs3StatusToErrCode(statusCode), lsPublicSharesResponse.Status.Message)
	}
	driveItems, err = g.cs3PublicSharesToDriveItems(ctx, lsPublicSharesResponse.Share, driveItems)
	if err != nil {
		return driveItems, errorcode.New(errorcode.GeneralException, err.Error())
	}
	return driveItems, nil

}

func (g BaseGraphService) cs3UserSharesToDriveItems(ctx context.Context, shares []*collaboration.Share, driveItems driveItemsByResourceID) (driveItemsByResourceID, error) {
	for _, s := range shares {
		g.logger.Debug().Interface("CS3 UserShare", s).Msg("Got Share")
		resIDStr := storagespace.FormatResourceID(*s.ResourceId)
		item, ok := driveItems[resIDStr]
		if !ok {
			itemptr, err := g.getDriveItem(ctx, storageprovider.Reference{ResourceId: s.ResourceId})
			if err != nil {
				g.logger.Debug().Err(err).Interface("Share", s.ResourceId).Msg("could not stat share, skipping")
				continue
			}
			item = *itemptr
		}
		perm, err := g.cs3UserShareToPermission(ctx, s, false)

		var errcode errorcode.Error
		switch {
		case errors.As(err, &errcode) && errcode.GetCode() == errorcode.ItemNotFound:
			// The Grantee couldn't be found (user/group does not exist anymore)
			continue
		case err != nil:
			return driveItems, err
		}
		item.Permissions = append(item.Permissions, *perm)
		driveItems[resIDStr] = item
	}
	return driveItems, nil
}

func (g BaseGraphService) cs3UserShareToPermission(ctx context.Context, share *collaboration.Share, isSpacePermission bool) (*libregraph.Permission, error) {
	perm := libregraph.Permission{}
	perm.SetRoles([]string{})
	if !isSpacePermission {
		perm.SetId(share.GetId().GetOpaqueId())
	}
	grantedTo := libregraph.SharePointIdentitySet{}
	switch share.GetGrantee().GetType() {
	case storageprovider.GranteeType_GRANTEE_TYPE_USER:
		user, err := cs3UserIdToIdentity(ctx, g.identityCache, share.Grantee.GetUserId())
		switch {
		case errors.Is(err, identity.ErrNotFound):
			g.logger.Warn().Str("userid", share.Grantee.GetUserId().GetOpaqueId()).Msg("User not found by id")
			// User does not seem to exist anymore, don't add a permission for this
			return nil, errorcode.New(errorcode.ItemNotFound, "grantee does not exist")
		case err != nil:
			return nil, errorcode.New(errorcode.GeneralException, err.Error())
		default:
			grantedTo.SetUser(user)
			if isSpacePermission {
				perm.SetId("u:" + user.GetId())
			}
		}
	case storageprovider.GranteeType_GRANTEE_TYPE_GROUP:
		group, err := groupIdToIdentity(ctx, g.identityCache, share.Grantee.GetGroupId().GetOpaqueId())
		switch {
		case errors.Is(err, identity.ErrNotFound):
			g.logger.Warn().Str("groupid", share.Grantee.GetGroupId().GetOpaqueId()).Msg("Group not found by id")
			// Group not seem to exist anymore, don't add a permission for this
			return nil, errorcode.New(errorcode.ItemNotFound, "grantee does not exist")
		case err != nil:
			return nil, errorcode.New(errorcode.GeneralException, err.Error())
		default:
			grantedTo.SetGroup(group)
			if isSpacePermission {
				perm.SetId("g:" + group.GetId())
			}
		}
	}

	// set expiration date
	if share.GetExpiration() != nil {
		perm.SetExpirationDateTime(cs3TimestampToTime(share.GetExpiration()))
	}
	condition := unifiedrole.UnifiedRoleConditionGrantee
	if isSpacePermission {
		condition = unifiedrole.UnifiedRoleConditionOwner
	}
	role := unifiedrole.CS3ResourcePermissionsToUnifiedRole(
		*share.GetPermissions().GetPermissions(),
		condition,
		g.config.FilesSharing.EnableResharing,
	)
	if role != nil {
		perm.SetRoles([]string{role.GetId()})
	} else {
		actions := unifiedrole.CS3ResourcePermissionsToLibregraphActions(*share.GetPermissions().GetPermissions())
		perm.SetLibreGraphPermissionsActions(actions)
		perm.SetRoles(nil)
	}
	perm.SetGrantedToV2(grantedTo)
	return &perm, nil
}

func (g BaseGraphService) cs3PublicSharesToDriveItems(ctx context.Context, shares []*link.PublicShare, driveItems driveItemsByResourceID) (driveItemsByResourceID, error) {
	for _, s := range shares {
		g.logger.Debug().Interface("CS3 PublicShare", s).Msg("Got Share")
		resIDStr := storagespace.FormatResourceID(*s.ResourceId)
		item, ok := driveItems[resIDStr]
		if !ok {
			itemptr, err := g.getDriveItem(ctx, storageprovider.Reference{ResourceId: s.ResourceId})
			if err != nil {
				g.logger.Debug().Err(err).Interface("Share", s.ResourceId).Msg("could not stat share, skipping")
				continue
			}
			item = *itemptr
		}
		perm, err := g.libreGraphPermissionFromCS3PublicShare(s)
		if err != nil {
			g.logger.Error().Err(err).Interface("Link", s.ResourceId).Msg("could not convert link to libregraph")
			return driveItems, err
		}

		item.Permissions = append(item.Permissions, *perm)
		driveItems[resIDStr] = item
	}

	return driveItems, nil
}

func (g BaseGraphService) getLinkPermissionResourceID(ctx context.Context, permissionID string) (*storageprovider.ResourceId, error) {
	share, err := g.getCS3PublicShareByID(ctx, permissionID)
	if err != nil {
		return nil, err
	}
	return share.GetResourceId(), nil
}

func (g BaseGraphService) getCS3PublicShareByID(ctx context.Context, permissionID string) (*link.PublicShare, error) {
	gatewayClient, err := g.gatewaySelector.Next()
	if err != nil {
		g.logger.Debug().Err(err).Msg("selecting gatewaySelector failed")
		return nil, err
	}

	getPublicShareResp, err := gatewayClient.GetPublicShare(ctx,
		&link.GetPublicShareRequest{
			Ref: &link.PublicShareReference{
				Spec: &link.PublicShareReference_Id{
					Id: &link.PublicShareId{
						OpaqueId: permissionID,
					},
				},
			},
		},
	)
	if errCode := errorcode.FromCS3Status(getPublicShareResp.GetStatus(), err); errCode != nil {
		return nil, *errCode
	}
	return getPublicShareResp.GetShare(), nil
}

func (g BaseGraphService) removePublicShare(ctx context.Context, permissionID string) error {
	gatewayClient, err := g.gatewaySelector.Next()
	if err != nil {
		g.logger.Debug().Err(err).Msg("selecting gatewaySelector failed")
		return err
	}

	removePublicShareResp, err := gatewayClient.RemovePublicShare(ctx,
		&link.RemovePublicShareRequest{
			Ref: &link.PublicShareReference{
				Spec: &link.PublicShareReference_Id{
					Id: &link.PublicShareId{
						OpaqueId: permissionID,
					},
				},
			},
		})
	if errcode := errorcode.FromCS3Status(removePublicShareResp.GetStatus(), err); errcode != nil {
		return *errcode
	}
	// We need to return an untyped nil here otherwise the error==nil check won't work
	return nil
}

func (g BaseGraphService) removeUserShare(ctx context.Context, permissionID string) error {
	gatewayClient, err := g.gatewaySelector.Next()
	if err != nil {
		g.logger.Debug().Err(err).Msg("selecting gatewaySelector failed")
		return err
	}

	removeShareResp, err := gatewayClient.RemoveShare(ctx,
		&collaboration.RemoveShareRequest{
			Ref: &collaboration.ShareReference{
				Spec: &collaboration.ShareReference_Id{
					Id: &collaboration.ShareId{
						OpaqueId: permissionID,
					},
				},
			},
		})

	if errCode := errorcode.FromCS3Status(removeShareResp.GetStatus(), err); errCode != nil {
		return *errCode
	}
	// We need to return an untyped nil here otherwise the error==nil check won't work
	return nil
}

func (g BaseGraphService) removeSpacePermission(ctx context.Context, permissionID string, resourceId *storageprovider.ResourceId) error {
	grantee, err := spacePermissionIdToCS3Grantee(permissionID)
	if err != nil {
		return err
	}

	gatewayClient, err := g.gatewaySelector.Next()
	if err != nil {
		g.logger.Debug().Err(err).Msg("selecting gatewaySelector failed")
		return err
	}
	removeShareResp, err := gatewayClient.RemoveShare(ctx, &collaboration.RemoveShareRequest{
		Ref: &collaboration.ShareReference{
			Spec: &collaboration.ShareReference_Key{
				Key: &collaboration.ShareKey{
					ResourceId: resourceId,
					Grantee:    &grantee,
				},
			},
		},
	})

	if errCode := errorcode.FromCS3Status(removeShareResp.GetStatus(), err); errCode != nil {
		return *errCode
	}
	// We need to return an untyped nil here otherwise the error==nil check won't work
	return nil
}

func (g BaseGraphService) getUserPermissionResourceID(ctx context.Context, permissionID string) (*storageprovider.ResourceId, error) {
	share, err := g.getCS3UserShareByID(ctx, permissionID)
	if err != nil {
		return nil, err
	}
	return share.GetResourceId(), nil
}

func (g BaseGraphService) getCS3UserShareByID(ctx context.Context, permissionID string) (*collaboration.Share, error) {
	gatewayClient, err := g.gatewaySelector.Next()
	if err != nil {
		g.logger.Debug().Err(err).Msg("selecting gatewaySelector failed")
		return nil, err
	}

	getShareResp, err := gatewayClient.GetShare(ctx,
		&collaboration.GetShareRequest{
			Ref: &collaboration.ShareReference{
				Spec: &collaboration.ShareReference_Id{
					Id: &collaboration.ShareId{
						OpaqueId: permissionID,
					},
				},
			},
		})
	if errCode := errorcode.FromCS3Status(getShareResp.GetStatus(), err); errCode != nil {
		return nil, *errCode
	}
	return getShareResp.GetShare(), nil
}
