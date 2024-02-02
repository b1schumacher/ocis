// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	v0 "github.com/owncloud/ocis/v2/protogen/gen/ocis/messages/settings/v0"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

type Manager_Expecter struct {
	mock *mock.Mock
}

func (_m *Manager) EXPECT() *Manager_Expecter {
	return &Manager_Expecter{mock: &_m.Mock}
}

// AddSettingToBundle provides a mock function with given fields: bundleID, setting
func (_m *Manager) AddSettingToBundle(bundleID string, setting *v0.Setting) (*v0.Setting, error) {
	ret := _m.Called(bundleID, setting)

	if len(ret) == 0 {
		panic("no return value specified for AddSettingToBundle")
	}

	var r0 *v0.Setting
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *v0.Setting) (*v0.Setting, error)); ok {
		return rf(bundleID, setting)
	}
	if rf, ok := ret.Get(0).(func(string, *v0.Setting) *v0.Setting); ok {
		r0 = rf(bundleID, setting)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.Setting)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *v0.Setting) error); ok {
		r1 = rf(bundleID, setting)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_AddSettingToBundle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSettingToBundle'
type Manager_AddSettingToBundle_Call struct {
	*mock.Call
}

// AddSettingToBundle is a helper method to define mock.On call
//   - bundleID string
//   - setting *v0.Setting
func (_e *Manager_Expecter) AddSettingToBundle(bundleID interface{}, setting interface{}) *Manager_AddSettingToBundle_Call {
	return &Manager_AddSettingToBundle_Call{Call: _e.mock.On("AddSettingToBundle", bundleID, setting)}
}

func (_c *Manager_AddSettingToBundle_Call) Run(run func(bundleID string, setting *v0.Setting)) *Manager_AddSettingToBundle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*v0.Setting))
	})
	return _c
}

func (_c *Manager_AddSettingToBundle_Call) Return(_a0 *v0.Setting, _a1 error) *Manager_AddSettingToBundle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_AddSettingToBundle_Call) RunAndReturn(run func(string, *v0.Setting) (*v0.Setting, error)) *Manager_AddSettingToBundle_Call {
	_c.Call.Return(run)
	return _c
}

// ListBundles provides a mock function with given fields: bundleType, bundleIDs
func (_m *Manager) ListBundles(bundleType v0.Bundle_Type, bundleIDs []string) ([]*v0.Bundle, error) {
	ret := _m.Called(bundleType, bundleIDs)

	if len(ret) == 0 {
		panic("no return value specified for ListBundles")
	}

	var r0 []*v0.Bundle
	var r1 error
	if rf, ok := ret.Get(0).(func(v0.Bundle_Type, []string) ([]*v0.Bundle, error)); ok {
		return rf(bundleType, bundleIDs)
	}
	if rf, ok := ret.Get(0).(func(v0.Bundle_Type, []string) []*v0.Bundle); ok {
		r0 = rf(bundleType, bundleIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v0.Bundle)
		}
	}

	if rf, ok := ret.Get(1).(func(v0.Bundle_Type, []string) error); ok {
		r1 = rf(bundleType, bundleIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_ListBundles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListBundles'
type Manager_ListBundles_Call struct {
	*mock.Call
}

// ListBundles is a helper method to define mock.On call
//   - bundleType v0.Bundle_Type
//   - bundleIDs []string
func (_e *Manager_Expecter) ListBundles(bundleType interface{}, bundleIDs interface{}) *Manager_ListBundles_Call {
	return &Manager_ListBundles_Call{Call: _e.mock.On("ListBundles", bundleType, bundleIDs)}
}

func (_c *Manager_ListBundles_Call) Run(run func(bundleType v0.Bundle_Type, bundleIDs []string)) *Manager_ListBundles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(v0.Bundle_Type), args[1].([]string))
	})
	return _c
}

func (_c *Manager_ListBundles_Call) Return(_a0 []*v0.Bundle, _a1 error) *Manager_ListBundles_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_ListBundles_Call) RunAndReturn(run func(v0.Bundle_Type, []string) ([]*v0.Bundle, error)) *Manager_ListBundles_Call {
	_c.Call.Return(run)
	return _c
}

// ListPermissionsByResource provides a mock function with given fields: resource, roleIDs
func (_m *Manager) ListPermissionsByResource(resource *v0.Resource, roleIDs []string) ([]*v0.Permission, error) {
	ret := _m.Called(resource, roleIDs)

	if len(ret) == 0 {
		panic("no return value specified for ListPermissionsByResource")
	}

	var r0 []*v0.Permission
	var r1 error
	if rf, ok := ret.Get(0).(func(*v0.Resource, []string) ([]*v0.Permission, error)); ok {
		return rf(resource, roleIDs)
	}
	if rf, ok := ret.Get(0).(func(*v0.Resource, []string) []*v0.Permission); ok {
		r0 = rf(resource, roleIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v0.Permission)
		}
	}

	if rf, ok := ret.Get(1).(func(*v0.Resource, []string) error); ok {
		r1 = rf(resource, roleIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_ListPermissionsByResource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPermissionsByResource'
type Manager_ListPermissionsByResource_Call struct {
	*mock.Call
}

// ListPermissionsByResource is a helper method to define mock.On call
//   - resource *v0.Resource
//   - roleIDs []string
func (_e *Manager_Expecter) ListPermissionsByResource(resource interface{}, roleIDs interface{}) *Manager_ListPermissionsByResource_Call {
	return &Manager_ListPermissionsByResource_Call{Call: _e.mock.On("ListPermissionsByResource", resource, roleIDs)}
}

func (_c *Manager_ListPermissionsByResource_Call) Run(run func(resource *v0.Resource, roleIDs []string)) *Manager_ListPermissionsByResource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v0.Resource), args[1].([]string))
	})
	return _c
}

func (_c *Manager_ListPermissionsByResource_Call) Return(_a0 []*v0.Permission, _a1 error) *Manager_ListPermissionsByResource_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_ListPermissionsByResource_Call) RunAndReturn(run func(*v0.Resource, []string) ([]*v0.Permission, error)) *Manager_ListPermissionsByResource_Call {
	_c.Call.Return(run)
	return _c
}

// ListRoleAssignments provides a mock function with given fields: accountUUID
func (_m *Manager) ListRoleAssignments(accountUUID string) ([]*v0.UserRoleAssignment, error) {
	ret := _m.Called(accountUUID)

	if len(ret) == 0 {
		panic("no return value specified for ListRoleAssignments")
	}

	var r0 []*v0.UserRoleAssignment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*v0.UserRoleAssignment, error)); ok {
		return rf(accountUUID)
	}
	if rf, ok := ret.Get(0).(func(string) []*v0.UserRoleAssignment); ok {
		r0 = rf(accountUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v0.UserRoleAssignment)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accountUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_ListRoleAssignments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListRoleAssignments'
type Manager_ListRoleAssignments_Call struct {
	*mock.Call
}

// ListRoleAssignments is a helper method to define mock.On call
//   - accountUUID string
func (_e *Manager_Expecter) ListRoleAssignments(accountUUID interface{}) *Manager_ListRoleAssignments_Call {
	return &Manager_ListRoleAssignments_Call{Call: _e.mock.On("ListRoleAssignments", accountUUID)}
}

func (_c *Manager_ListRoleAssignments_Call) Run(run func(accountUUID string)) *Manager_ListRoleAssignments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Manager_ListRoleAssignments_Call) Return(_a0 []*v0.UserRoleAssignment, _a1 error) *Manager_ListRoleAssignments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_ListRoleAssignments_Call) RunAndReturn(run func(string) ([]*v0.UserRoleAssignment, error)) *Manager_ListRoleAssignments_Call {
	_c.Call.Return(run)
	return _c
}

// ListValues provides a mock function with given fields: bundleID, accountUUID
func (_m *Manager) ListValues(bundleID string, accountUUID string) ([]*v0.Value, error) {
	ret := _m.Called(bundleID, accountUUID)

	if len(ret) == 0 {
		panic("no return value specified for ListValues")
	}

	var r0 []*v0.Value
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]*v0.Value, error)); ok {
		return rf(bundleID, accountUUID)
	}
	if rf, ok := ret.Get(0).(func(string, string) []*v0.Value); ok {
		r0 = rf(bundleID, accountUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v0.Value)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(bundleID, accountUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_ListValues_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListValues'
type Manager_ListValues_Call struct {
	*mock.Call
}

// ListValues is a helper method to define mock.On call
//   - bundleID string
//   - accountUUID string
func (_e *Manager_Expecter) ListValues(bundleID interface{}, accountUUID interface{}) *Manager_ListValues_Call {
	return &Manager_ListValues_Call{Call: _e.mock.On("ListValues", bundleID, accountUUID)}
}

func (_c *Manager_ListValues_Call) Run(run func(bundleID string, accountUUID string)) *Manager_ListValues_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Manager_ListValues_Call) Return(_a0 []*v0.Value, _a1 error) *Manager_ListValues_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_ListValues_Call) RunAndReturn(run func(string, string) ([]*v0.Value, error)) *Manager_ListValues_Call {
	_c.Call.Return(run)
	return _c
}

// ReadBundle provides a mock function with given fields: bundleID
func (_m *Manager) ReadBundle(bundleID string) (*v0.Bundle, error) {
	ret := _m.Called(bundleID)

	if len(ret) == 0 {
		panic("no return value specified for ReadBundle")
	}

	var r0 *v0.Bundle
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*v0.Bundle, error)); ok {
		return rf(bundleID)
	}
	if rf, ok := ret.Get(0).(func(string) *v0.Bundle); ok {
		r0 = rf(bundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.Bundle)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_ReadBundle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadBundle'
type Manager_ReadBundle_Call struct {
	*mock.Call
}

// ReadBundle is a helper method to define mock.On call
//   - bundleID string
func (_e *Manager_Expecter) ReadBundle(bundleID interface{}) *Manager_ReadBundle_Call {
	return &Manager_ReadBundle_Call{Call: _e.mock.On("ReadBundle", bundleID)}
}

func (_c *Manager_ReadBundle_Call) Run(run func(bundleID string)) *Manager_ReadBundle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Manager_ReadBundle_Call) Return(_a0 *v0.Bundle, _a1 error) *Manager_ReadBundle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_ReadBundle_Call) RunAndReturn(run func(string) (*v0.Bundle, error)) *Manager_ReadBundle_Call {
	_c.Call.Return(run)
	return _c
}

// ReadPermissionByID provides a mock function with given fields: permissionID, roleIDs
func (_m *Manager) ReadPermissionByID(permissionID string, roleIDs []string) (*v0.Permission, error) {
	ret := _m.Called(permissionID, roleIDs)

	if len(ret) == 0 {
		panic("no return value specified for ReadPermissionByID")
	}

	var r0 *v0.Permission
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string) (*v0.Permission, error)); ok {
		return rf(permissionID, roleIDs)
	}
	if rf, ok := ret.Get(0).(func(string, []string) *v0.Permission); ok {
		r0 = rf(permissionID, roleIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.Permission)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(permissionID, roleIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_ReadPermissionByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadPermissionByID'
type Manager_ReadPermissionByID_Call struct {
	*mock.Call
}

// ReadPermissionByID is a helper method to define mock.On call
//   - permissionID string
//   - roleIDs []string
func (_e *Manager_Expecter) ReadPermissionByID(permissionID interface{}, roleIDs interface{}) *Manager_ReadPermissionByID_Call {
	return &Manager_ReadPermissionByID_Call{Call: _e.mock.On("ReadPermissionByID", permissionID, roleIDs)}
}

func (_c *Manager_ReadPermissionByID_Call) Run(run func(permissionID string, roleIDs []string)) *Manager_ReadPermissionByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string))
	})
	return _c
}

func (_c *Manager_ReadPermissionByID_Call) Return(_a0 *v0.Permission, _a1 error) *Manager_ReadPermissionByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_ReadPermissionByID_Call) RunAndReturn(run func(string, []string) (*v0.Permission, error)) *Manager_ReadPermissionByID_Call {
	_c.Call.Return(run)
	return _c
}

// ReadPermissionByName provides a mock function with given fields: name, roleIDs
func (_m *Manager) ReadPermissionByName(name string, roleIDs []string) (*v0.Permission, error) {
	ret := _m.Called(name, roleIDs)

	if len(ret) == 0 {
		panic("no return value specified for ReadPermissionByName")
	}

	var r0 *v0.Permission
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string) (*v0.Permission, error)); ok {
		return rf(name, roleIDs)
	}
	if rf, ok := ret.Get(0).(func(string, []string) *v0.Permission); ok {
		r0 = rf(name, roleIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.Permission)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(name, roleIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_ReadPermissionByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadPermissionByName'
type Manager_ReadPermissionByName_Call struct {
	*mock.Call
}

// ReadPermissionByName is a helper method to define mock.On call
//   - name string
//   - roleIDs []string
func (_e *Manager_Expecter) ReadPermissionByName(name interface{}, roleIDs interface{}) *Manager_ReadPermissionByName_Call {
	return &Manager_ReadPermissionByName_Call{Call: _e.mock.On("ReadPermissionByName", name, roleIDs)}
}

func (_c *Manager_ReadPermissionByName_Call) Run(run func(name string, roleIDs []string)) *Manager_ReadPermissionByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string))
	})
	return _c
}

func (_c *Manager_ReadPermissionByName_Call) Return(_a0 *v0.Permission, _a1 error) *Manager_ReadPermissionByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_ReadPermissionByName_Call) RunAndReturn(run func(string, []string) (*v0.Permission, error)) *Manager_ReadPermissionByName_Call {
	_c.Call.Return(run)
	return _c
}

// ReadSetting provides a mock function with given fields: settingID
func (_m *Manager) ReadSetting(settingID string) (*v0.Setting, error) {
	ret := _m.Called(settingID)

	if len(ret) == 0 {
		panic("no return value specified for ReadSetting")
	}

	var r0 *v0.Setting
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*v0.Setting, error)); ok {
		return rf(settingID)
	}
	if rf, ok := ret.Get(0).(func(string) *v0.Setting); ok {
		r0 = rf(settingID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.Setting)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(settingID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_ReadSetting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadSetting'
type Manager_ReadSetting_Call struct {
	*mock.Call
}

// ReadSetting is a helper method to define mock.On call
//   - settingID string
func (_e *Manager_Expecter) ReadSetting(settingID interface{}) *Manager_ReadSetting_Call {
	return &Manager_ReadSetting_Call{Call: _e.mock.On("ReadSetting", settingID)}
}

func (_c *Manager_ReadSetting_Call) Run(run func(settingID string)) *Manager_ReadSetting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Manager_ReadSetting_Call) Return(_a0 *v0.Setting, _a1 error) *Manager_ReadSetting_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_ReadSetting_Call) RunAndReturn(run func(string) (*v0.Setting, error)) *Manager_ReadSetting_Call {
	_c.Call.Return(run)
	return _c
}

// ReadValue provides a mock function with given fields: valueID
func (_m *Manager) ReadValue(valueID string) (*v0.Value, error) {
	ret := _m.Called(valueID)

	if len(ret) == 0 {
		panic("no return value specified for ReadValue")
	}

	var r0 *v0.Value
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*v0.Value, error)); ok {
		return rf(valueID)
	}
	if rf, ok := ret.Get(0).(func(string) *v0.Value); ok {
		r0 = rf(valueID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.Value)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(valueID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_ReadValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadValue'
type Manager_ReadValue_Call struct {
	*mock.Call
}

// ReadValue is a helper method to define mock.On call
//   - valueID string
func (_e *Manager_Expecter) ReadValue(valueID interface{}) *Manager_ReadValue_Call {
	return &Manager_ReadValue_Call{Call: _e.mock.On("ReadValue", valueID)}
}

func (_c *Manager_ReadValue_Call) Run(run func(valueID string)) *Manager_ReadValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Manager_ReadValue_Call) Return(_a0 *v0.Value, _a1 error) *Manager_ReadValue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_ReadValue_Call) RunAndReturn(run func(string) (*v0.Value, error)) *Manager_ReadValue_Call {
	_c.Call.Return(run)
	return _c
}

// ReadValueByUniqueIdentifiers provides a mock function with given fields: accountUUID, settingID
func (_m *Manager) ReadValueByUniqueIdentifiers(accountUUID string, settingID string) (*v0.Value, error) {
	ret := _m.Called(accountUUID, settingID)

	if len(ret) == 0 {
		panic("no return value specified for ReadValueByUniqueIdentifiers")
	}

	var r0 *v0.Value
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*v0.Value, error)); ok {
		return rf(accountUUID, settingID)
	}
	if rf, ok := ret.Get(0).(func(string, string) *v0.Value); ok {
		r0 = rf(accountUUID, settingID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.Value)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(accountUUID, settingID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_ReadValueByUniqueIdentifiers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadValueByUniqueIdentifiers'
type Manager_ReadValueByUniqueIdentifiers_Call struct {
	*mock.Call
}

// ReadValueByUniqueIdentifiers is a helper method to define mock.On call
//   - accountUUID string
//   - settingID string
func (_e *Manager_Expecter) ReadValueByUniqueIdentifiers(accountUUID interface{}, settingID interface{}) *Manager_ReadValueByUniqueIdentifiers_Call {
	return &Manager_ReadValueByUniqueIdentifiers_Call{Call: _e.mock.On("ReadValueByUniqueIdentifiers", accountUUID, settingID)}
}

func (_c *Manager_ReadValueByUniqueIdentifiers_Call) Run(run func(accountUUID string, settingID string)) *Manager_ReadValueByUniqueIdentifiers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Manager_ReadValueByUniqueIdentifiers_Call) Return(_a0 *v0.Value, _a1 error) *Manager_ReadValueByUniqueIdentifiers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_ReadValueByUniqueIdentifiers_Call) RunAndReturn(run func(string, string) (*v0.Value, error)) *Manager_ReadValueByUniqueIdentifiers_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveRoleAssignment provides a mock function with given fields: assignmentID
func (_m *Manager) RemoveRoleAssignment(assignmentID string) error {
	ret := _m.Called(assignmentID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveRoleAssignment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(assignmentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Manager_RemoveRoleAssignment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveRoleAssignment'
type Manager_RemoveRoleAssignment_Call struct {
	*mock.Call
}

// RemoveRoleAssignment is a helper method to define mock.On call
//   - assignmentID string
func (_e *Manager_Expecter) RemoveRoleAssignment(assignmentID interface{}) *Manager_RemoveRoleAssignment_Call {
	return &Manager_RemoveRoleAssignment_Call{Call: _e.mock.On("RemoveRoleAssignment", assignmentID)}
}

func (_c *Manager_RemoveRoleAssignment_Call) Run(run func(assignmentID string)) *Manager_RemoveRoleAssignment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Manager_RemoveRoleAssignment_Call) Return(_a0 error) *Manager_RemoveRoleAssignment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_RemoveRoleAssignment_Call) RunAndReturn(run func(string) error) *Manager_RemoveRoleAssignment_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveSettingFromBundle provides a mock function with given fields: bundleID, settingID
func (_m *Manager) RemoveSettingFromBundle(bundleID string, settingID string) error {
	ret := _m.Called(bundleID, settingID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveSettingFromBundle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(bundleID, settingID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Manager_RemoveSettingFromBundle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveSettingFromBundle'
type Manager_RemoveSettingFromBundle_Call struct {
	*mock.Call
}

// RemoveSettingFromBundle is a helper method to define mock.On call
//   - bundleID string
//   - settingID string
func (_e *Manager_Expecter) RemoveSettingFromBundle(bundleID interface{}, settingID interface{}) *Manager_RemoveSettingFromBundle_Call {
	return &Manager_RemoveSettingFromBundle_Call{Call: _e.mock.On("RemoveSettingFromBundle", bundleID, settingID)}
}

func (_c *Manager_RemoveSettingFromBundle_Call) Run(run func(bundleID string, settingID string)) *Manager_RemoveSettingFromBundle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Manager_RemoveSettingFromBundle_Call) Return(_a0 error) *Manager_RemoveSettingFromBundle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_RemoveSettingFromBundle_Call) RunAndReturn(run func(string, string) error) *Manager_RemoveSettingFromBundle_Call {
	_c.Call.Return(run)
	return _c
}

// WriteBundle provides a mock function with given fields: bundle
func (_m *Manager) WriteBundle(bundle *v0.Bundle) (*v0.Bundle, error) {
	ret := _m.Called(bundle)

	if len(ret) == 0 {
		panic("no return value specified for WriteBundle")
	}

	var r0 *v0.Bundle
	var r1 error
	if rf, ok := ret.Get(0).(func(*v0.Bundle) (*v0.Bundle, error)); ok {
		return rf(bundle)
	}
	if rf, ok := ret.Get(0).(func(*v0.Bundle) *v0.Bundle); ok {
		r0 = rf(bundle)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.Bundle)
		}
	}

	if rf, ok := ret.Get(1).(func(*v0.Bundle) error); ok {
		r1 = rf(bundle)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_WriteBundle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteBundle'
type Manager_WriteBundle_Call struct {
	*mock.Call
}

// WriteBundle is a helper method to define mock.On call
//   - bundle *v0.Bundle
func (_e *Manager_Expecter) WriteBundle(bundle interface{}) *Manager_WriteBundle_Call {
	return &Manager_WriteBundle_Call{Call: _e.mock.On("WriteBundle", bundle)}
}

func (_c *Manager_WriteBundle_Call) Run(run func(bundle *v0.Bundle)) *Manager_WriteBundle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v0.Bundle))
	})
	return _c
}

func (_c *Manager_WriteBundle_Call) Return(_a0 *v0.Bundle, _a1 error) *Manager_WriteBundle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_WriteBundle_Call) RunAndReturn(run func(*v0.Bundle) (*v0.Bundle, error)) *Manager_WriteBundle_Call {
	_c.Call.Return(run)
	return _c
}

// WriteRoleAssignment provides a mock function with given fields: accountUUID, roleID
func (_m *Manager) WriteRoleAssignment(accountUUID string, roleID string) (*v0.UserRoleAssignment, error) {
	ret := _m.Called(accountUUID, roleID)

	if len(ret) == 0 {
		panic("no return value specified for WriteRoleAssignment")
	}

	var r0 *v0.UserRoleAssignment
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*v0.UserRoleAssignment, error)); ok {
		return rf(accountUUID, roleID)
	}
	if rf, ok := ret.Get(0).(func(string, string) *v0.UserRoleAssignment); ok {
		r0 = rf(accountUUID, roleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.UserRoleAssignment)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(accountUUID, roleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_WriteRoleAssignment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteRoleAssignment'
type Manager_WriteRoleAssignment_Call struct {
	*mock.Call
}

// WriteRoleAssignment is a helper method to define mock.On call
//   - accountUUID string
//   - roleID string
func (_e *Manager_Expecter) WriteRoleAssignment(accountUUID interface{}, roleID interface{}) *Manager_WriteRoleAssignment_Call {
	return &Manager_WriteRoleAssignment_Call{Call: _e.mock.On("WriteRoleAssignment", accountUUID, roleID)}
}

func (_c *Manager_WriteRoleAssignment_Call) Run(run func(accountUUID string, roleID string)) *Manager_WriteRoleAssignment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Manager_WriteRoleAssignment_Call) Return(_a0 *v0.UserRoleAssignment, _a1 error) *Manager_WriteRoleAssignment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_WriteRoleAssignment_Call) RunAndReturn(run func(string, string) (*v0.UserRoleAssignment, error)) *Manager_WriteRoleAssignment_Call {
	_c.Call.Return(run)
	return _c
}

// WriteValue provides a mock function with given fields: value
func (_m *Manager) WriteValue(value *v0.Value) (*v0.Value, error) {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for WriteValue")
	}

	var r0 *v0.Value
	var r1 error
	if rf, ok := ret.Get(0).(func(*v0.Value) (*v0.Value, error)); ok {
		return rf(value)
	}
	if rf, ok := ret.Get(0).(func(*v0.Value) *v0.Value); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.Value)
		}
	}

	if rf, ok := ret.Get(1).(func(*v0.Value) error); ok {
		r1 = rf(value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_WriteValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteValue'
type Manager_WriteValue_Call struct {
	*mock.Call
}

// WriteValue is a helper method to define mock.On call
//   - value *v0.Value
func (_e *Manager_Expecter) WriteValue(value interface{}) *Manager_WriteValue_Call {
	return &Manager_WriteValue_Call{Call: _e.mock.On("WriteValue", value)}
}

func (_c *Manager_WriteValue_Call) Run(run func(value *v0.Value)) *Manager_WriteValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v0.Value))
	})
	return _c
}

func (_c *Manager_WriteValue_Call) Return(_a0 *v0.Value, _a1 error) *Manager_WriteValue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_WriteValue_Call) RunAndReturn(run func(*v0.Value) (*v0.Value, error)) *Manager_WriteValue_Call {
	_c.Call.Return(run)
	return _c
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
