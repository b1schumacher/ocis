// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: ocis/services/thumbnails/v0/thumbnails.proto

package v0

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v0 "github.com/owncloud/ocis/protogen/gen/ocis/messages/thumbnails/v0"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A request to retrieve a thumbnail
type GetThumbnailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path to the source image
	Filepath string `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"`
	// The type to which the thumbnail should get encoded to.
	ThumbnailType v0.ThumbnailType `protobuf:"varint,2,opt,name=thumbnail_type,json=thumbnailType,proto3,enum=ocis.messages.thumbnails.v0.ThumbnailType" json:"thumbnail_type,omitempty"`
	// The width of the thumbnail
	Width int32 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	// The height of the thumbnail
	Height int32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	// Types that are assignable to Source:
	//	*GetThumbnailRequest_WebdavSource
	//	*GetThumbnailRequest_Cs3Source
	Source isGetThumbnailRequest_Source `protobuf_oneof:"source"`
}

func (x *GetThumbnailRequest) Reset() {
	*x = GetThumbnailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_services_thumbnails_v0_thumbnails_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThumbnailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThumbnailRequest) ProtoMessage() {}

func (x *GetThumbnailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_services_thumbnails_v0_thumbnails_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThumbnailRequest.ProtoReflect.Descriptor instead.
func (*GetThumbnailRequest) Descriptor() ([]byte, []int) {
	return file_ocis_services_thumbnails_v0_thumbnails_proto_rawDescGZIP(), []int{0}
}

func (x *GetThumbnailRequest) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *GetThumbnailRequest) GetThumbnailType() v0.ThumbnailType {
	if x != nil {
		return x.ThumbnailType
	}
	return v0.ThumbnailType(0)
}

func (x *GetThumbnailRequest) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *GetThumbnailRequest) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (m *GetThumbnailRequest) GetSource() isGetThumbnailRequest_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *GetThumbnailRequest) GetWebdavSource() *v0.WebdavSource {
	if x, ok := x.GetSource().(*GetThumbnailRequest_WebdavSource); ok {
		return x.WebdavSource
	}
	return nil
}

func (x *GetThumbnailRequest) GetCs3Source() *v0.CS3Source {
	if x, ok := x.GetSource().(*GetThumbnailRequest_Cs3Source); ok {
		return x.Cs3Source
	}
	return nil
}

type isGetThumbnailRequest_Source interface {
	isGetThumbnailRequest_Source()
}

type GetThumbnailRequest_WebdavSource struct {
	WebdavSource *v0.WebdavSource `protobuf:"bytes,5,opt,name=webdav_source,json=webdavSource,proto3,oneof"`
}

type GetThumbnailRequest_Cs3Source struct {
	Cs3Source *v0.CS3Source `protobuf:"bytes,6,opt,name=cs3_source,json=cs3Source,proto3,oneof"`
}

func (*GetThumbnailRequest_WebdavSource) isGetThumbnailRequest_Source() {}

func (*GetThumbnailRequest_Cs3Source) isGetThumbnailRequest_Source() {}

// The service response
type GetThumbnailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The endpoint where the thumbnail can be downloaded.
	DataEndpoint string `protobuf:"bytes,1,opt,name=data_endpoint,json=dataEndpoint,proto3" json:"data_endpoint,omitempty"`
	// The transfer token to be able to download the thumbnail.
	TransferToken string `protobuf:"bytes,2,opt,name=transfer_token,json=transferToken,proto3" json:"transfer_token,omitempty"`
	// The mimetype of the thumbnail
	Mimetype string `protobuf:"bytes,3,opt,name=mimetype,proto3" json:"mimetype,omitempty"`
}

func (x *GetThumbnailResponse) Reset() {
	*x = GetThumbnailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_services_thumbnails_v0_thumbnails_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThumbnailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThumbnailResponse) ProtoMessage() {}

func (x *GetThumbnailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_services_thumbnails_v0_thumbnails_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThumbnailResponse.ProtoReflect.Descriptor instead.
func (*GetThumbnailResponse) Descriptor() ([]byte, []int) {
	return file_ocis_services_thumbnails_v0_thumbnails_proto_rawDescGZIP(), []int{1}
}

func (x *GetThumbnailResponse) GetDataEndpoint() string {
	if x != nil {
		return x.DataEndpoint
	}
	return ""
}

func (x *GetThumbnailResponse) GetTransferToken() string {
	if x != nil {
		return x.TransferToken
	}
	return ""
}

func (x *GetThumbnailResponse) GetMimetype() string {
	if x != nil {
		return x.Mimetype
	}
	return ""
}

var File_ocis_services_thumbnails_v0_thumbnails_proto protoreflect.FileDescriptor

var file_ocis_services_thumbnails_v0_thumbnails_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6f, 0x63, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x76, 0x30, 0x2f, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x6f, 0x63, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x30, 0x1a, 0x2c, 0x6f, 0x63, 0x69,
	0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x76, 0x30, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x02, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x12, 0x51, 0x0a,
	0x0e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73,
	0x2e, 0x76, 0x30, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x50,
	0x0a, 0x0d, 0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73,
	0x2e, 0x76, 0x30, 0x2e, 0x57, 0x65, 0x62, 0x64, 0x61, 0x76, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x48, 0x00, 0x52, 0x0c, 0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x47, 0x0a, 0x0a, 0x63, 0x73, 0x33, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2e,
	0x76, 0x30, 0x2e, 0x43, 0x53, 0x33, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x09,
	0x63, 0x73, 0x33, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0x7e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x32, 0x87, 0x01, 0x0a, 0x10, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x30, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6f, 0x63, 0x69,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xeb, 0x02,
	0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x77, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x63, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x63, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73,
	0x2f, 0x76, 0x30, 0x92, 0x41, 0xa4, 0x02, 0x12, 0xb8, 0x01, 0x0a, 0x22, 0x6f, 0x77, 0x6e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x20, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x20, 0x53, 0x63,
	0x61, 0x6c, 0x65, 0x20, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x47,
	0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x20, 0x47, 0x6d, 0x62, 0x48, 0x12,
	0x20, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x63, 0x69,
	0x73, 0x1a, 0x14, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x40, 0x6f, 0x77, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x42, 0x0a, 0x0a, 0x41, 0x70, 0x61, 0x63, 0x68,
	0x65, 0x2d, 0x32, 0x2e, 0x30, 0x12, 0x34, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x6f, 0x63, 0x69, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x32, 0x05, 0x31, 0x2e, 0x30,
	0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x72, 0x3f, 0x0a, 0x10, 0x44, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x20, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x12, 0x2b,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x76, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ocis_services_thumbnails_v0_thumbnails_proto_rawDescOnce sync.Once
	file_ocis_services_thumbnails_v0_thumbnails_proto_rawDescData = file_ocis_services_thumbnails_v0_thumbnails_proto_rawDesc
)

func file_ocis_services_thumbnails_v0_thumbnails_proto_rawDescGZIP() []byte {
	file_ocis_services_thumbnails_v0_thumbnails_proto_rawDescOnce.Do(func() {
		file_ocis_services_thumbnails_v0_thumbnails_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocis_services_thumbnails_v0_thumbnails_proto_rawDescData)
	})
	return file_ocis_services_thumbnails_v0_thumbnails_proto_rawDescData
}

var file_ocis_services_thumbnails_v0_thumbnails_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ocis_services_thumbnails_v0_thumbnails_proto_goTypes = []interface{}{
	(*GetThumbnailRequest)(nil),  // 0: ocis.services.thumbnails.v0.GetThumbnailRequest
	(*GetThumbnailResponse)(nil), // 1: ocis.services.thumbnails.v0.GetThumbnailResponse
	(v0.ThumbnailType)(0),        // 2: ocis.messages.thumbnails.v0.ThumbnailType
	(*v0.WebdavSource)(nil),      // 3: ocis.messages.thumbnails.v0.WebdavSource
	(*v0.CS3Source)(nil),         // 4: ocis.messages.thumbnails.v0.CS3Source
}
var file_ocis_services_thumbnails_v0_thumbnails_proto_depIdxs = []int32{
	2, // 0: ocis.services.thumbnails.v0.GetThumbnailRequest.thumbnail_type:type_name -> ocis.messages.thumbnails.v0.ThumbnailType
	3, // 1: ocis.services.thumbnails.v0.GetThumbnailRequest.webdav_source:type_name -> ocis.messages.thumbnails.v0.WebdavSource
	4, // 2: ocis.services.thumbnails.v0.GetThumbnailRequest.cs3_source:type_name -> ocis.messages.thumbnails.v0.CS3Source
	0, // 3: ocis.services.thumbnails.v0.ThumbnailService.GetThumbnail:input_type -> ocis.services.thumbnails.v0.GetThumbnailRequest
	1, // 4: ocis.services.thumbnails.v0.ThumbnailService.GetThumbnail:output_type -> ocis.services.thumbnails.v0.GetThumbnailResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ocis_services_thumbnails_v0_thumbnails_proto_init() }
func file_ocis_services_thumbnails_v0_thumbnails_proto_init() {
	if File_ocis_services_thumbnails_v0_thumbnails_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocis_services_thumbnails_v0_thumbnails_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThumbnailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ocis_services_thumbnails_v0_thumbnails_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThumbnailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_ocis_services_thumbnails_v0_thumbnails_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GetThumbnailRequest_WebdavSource)(nil),
		(*GetThumbnailRequest_Cs3Source)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ocis_services_thumbnails_v0_thumbnails_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ocis_services_thumbnails_v0_thumbnails_proto_goTypes,
		DependencyIndexes: file_ocis_services_thumbnails_v0_thumbnails_proto_depIdxs,
		MessageInfos:      file_ocis_services_thumbnails_v0_thumbnails_proto_msgTypes,
	}.Build()
	File_ocis_services_thumbnails_v0_thumbnails_proto = out.File
	file_ocis_services_thumbnails_v0_thumbnails_proto_rawDesc = nil
	file_ocis_services_thumbnails_v0_thumbnails_proto_goTypes = nil
	file_ocis_services_thumbnails_v0_thumbnails_proto_depIdxs = nil
}
