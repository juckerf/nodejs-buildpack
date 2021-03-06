// Code generated by MockGen. DO NOT EDIT.
// Source: supply.go

package supply_test

import (
	libbuildpack "github.com/cloudfoundry/libbuildpack"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockCache is a mock of Cache interface
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockCache) EXPECT() *MockCacheMockRecorder {
	return _m.recorder
}

// Initialize mocks base method
func (_m *MockCache) Initialize() error {
	ret := _m.ctrl.Call(_m, "Initialize")
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (_mr *MockCacheMockRecorder) Initialize() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Initialize", reflect.TypeOf((*MockCache)(nil).Initialize))
}

// Restore mocks base method
func (_m *MockCache) Restore() error {
	ret := _m.ctrl.Call(_m, "Restore")
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore
func (_mr *MockCacheMockRecorder) Restore() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Restore", reflect.TypeOf((*MockCache)(nil).Restore))
}

// Save mocks base method
func (_m *MockCache) Save() error {
	ret := _m.ctrl.Call(_m, "Save")
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (_mr *MockCacheMockRecorder) Save() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Save", reflect.TypeOf((*MockCache)(nil).Save))
}

// MockCommand is a mock of Command interface
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return _m.recorder
}

// Execute mocks base method
func (_m *MockCommand) Execute(_param0 string, _param1 io.Writer, _param2 io.Writer, _param3 string, _param4 ...string) error {
	_s := []interface{}{_param0, _param1, _param2, _param3}
	for _, _x := range _param4 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Execute", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (_mr *MockCommandMockRecorder) Execute(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Execute", reflect.TypeOf((*MockCommand)(nil).Execute), _s...)
}

// MockManifest is a mock of Manifest interface
type MockManifest struct {
	ctrl     *gomock.Controller
	recorder *MockManifestMockRecorder
}

// MockManifestMockRecorder is the mock recorder for MockManifest
type MockManifestMockRecorder struct {
	mock *MockManifest
}

// NewMockManifest creates a new mock instance
func NewMockManifest(ctrl *gomock.Controller) *MockManifest {
	mock := &MockManifest{ctrl: ctrl}
	mock.recorder = &MockManifestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockManifest) EXPECT() *MockManifestMockRecorder {
	return _m.recorder
}

// AllDependencyVersions mocks base method
func (_m *MockManifest) AllDependencyVersions(_param0 string) []string {
	ret := _m.ctrl.Call(_m, "AllDependencyVersions", _param0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// AllDependencyVersions indicates an expected call of AllDependencyVersions
func (_mr *MockManifestMockRecorder) AllDependencyVersions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "AllDependencyVersions", reflect.TypeOf((*MockManifest)(nil).AllDependencyVersions), arg0)
}

// DefaultVersion mocks base method
func (_m *MockManifest) DefaultVersion(_param0 string) (libbuildpack.Dependency, error) {
	ret := _m.ctrl.Call(_m, "DefaultVersion", _param0)
	ret0, _ := ret[0].(libbuildpack.Dependency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultVersion indicates an expected call of DefaultVersion
func (_mr *MockManifestMockRecorder) DefaultVersion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DefaultVersion", reflect.TypeOf((*MockManifest)(nil).DefaultVersion), arg0)
}

// InstallDependency mocks base method
func (_m *MockManifest) InstallDependency(_param0 libbuildpack.Dependency, _param1 string) error {
	ret := _m.ctrl.Call(_m, "InstallDependency", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallDependency indicates an expected call of InstallDependency
func (_mr *MockManifestMockRecorder) InstallDependency(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InstallDependency", reflect.TypeOf((*MockManifest)(nil).InstallDependency), arg0, arg1)
}

// InstallOnlyVersion mocks base method
func (_m *MockManifest) InstallOnlyVersion(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "InstallOnlyVersion", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallOnlyVersion indicates an expected call of InstallOnlyVersion
func (_mr *MockManifestMockRecorder) InstallOnlyVersion(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InstallOnlyVersion", reflect.TypeOf((*MockManifest)(nil).InstallOnlyVersion), arg0, arg1)
}

// MockNPM is a mock of NPM interface
type MockNPM struct {
	ctrl     *gomock.Controller
	recorder *MockNPMMockRecorder
}

// MockNPMMockRecorder is the mock recorder for MockNPM
type MockNPMMockRecorder struct {
	mock *MockNPM
}

// NewMockNPM creates a new mock instance
func NewMockNPM(ctrl *gomock.Controller) *MockNPM {
	mock := &MockNPM{ctrl: ctrl}
	mock.recorder = &MockNPMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockNPM) EXPECT() *MockNPMMockRecorder {
	return _m.recorder
}

// Build mocks base method
func (_m *MockNPM) Build() error {
	ret := _m.ctrl.Call(_m, "Build")
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build
func (_mr *MockNPMMockRecorder) Build() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Build", reflect.TypeOf((*MockNPM)(nil).Build))
}

// Rebuild mocks base method
func (_m *MockNPM) Rebuild() error {
	ret := _m.ctrl.Call(_m, "Rebuild")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rebuild indicates an expected call of Rebuild
func (_mr *MockNPMMockRecorder) Rebuild() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Rebuild", reflect.TypeOf((*MockNPM)(nil).Rebuild))
}

// MockYarn is a mock of Yarn interface
type MockYarn struct {
	ctrl     *gomock.Controller
	recorder *MockYarnMockRecorder
}

// MockYarnMockRecorder is the mock recorder for MockYarn
type MockYarnMockRecorder struct {
	mock *MockYarn
}

// NewMockYarn creates a new mock instance
func NewMockYarn(ctrl *gomock.Controller) *MockYarn {
	mock := &MockYarn{ctrl: ctrl}
	mock.recorder = &MockYarnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockYarn) EXPECT() *MockYarnMockRecorder {
	return _m.recorder
}

// Build mocks base method
func (_m *MockYarn) Build() error {
	ret := _m.ctrl.Call(_m, "Build")
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build
func (_mr *MockYarnMockRecorder) Build() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Build", reflect.TypeOf((*MockYarn)(nil).Build))
}

// MockStager is a mock of Stager interface
type MockStager struct {
	ctrl     *gomock.Controller
	recorder *MockStagerMockRecorder
}

// MockStagerMockRecorder is the mock recorder for MockStager
type MockStagerMockRecorder struct {
	mock *MockStager
}

// NewMockStager creates a new mock instance
func NewMockStager(ctrl *gomock.Controller) *MockStager {
	mock := &MockStager{ctrl: ctrl}
	mock.recorder = &MockStagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockStager) EXPECT() *MockStagerMockRecorder {
	return _m.recorder
}

// BuildDir mocks base method
func (_m *MockStager) BuildDir() string {
	ret := _m.ctrl.Call(_m, "BuildDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildDir indicates an expected call of BuildDir
func (_mr *MockStagerMockRecorder) BuildDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BuildDir", reflect.TypeOf((*MockStager)(nil).BuildDir))
}

// DepDir mocks base method
func (_m *MockStager) DepDir() string {
	ret := _m.ctrl.Call(_m, "DepDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// DepDir indicates an expected call of DepDir
func (_mr *MockStagerMockRecorder) DepDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DepDir", reflect.TypeOf((*MockStager)(nil).DepDir))
}

// DepsIdx mocks base method
func (_m *MockStager) DepsIdx() string {
	ret := _m.ctrl.Call(_m, "DepsIdx")
	ret0, _ := ret[0].(string)
	return ret0
}

// DepsIdx indicates an expected call of DepsIdx
func (_mr *MockStagerMockRecorder) DepsIdx() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DepsIdx", reflect.TypeOf((*MockStager)(nil).DepsIdx))
}

// LinkDirectoryInDepDir mocks base method
func (_m *MockStager) LinkDirectoryInDepDir(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "LinkDirectoryInDepDir", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkDirectoryInDepDir indicates an expected call of LinkDirectoryInDepDir
func (_mr *MockStagerMockRecorder) LinkDirectoryInDepDir(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "LinkDirectoryInDepDir", reflect.TypeOf((*MockStager)(nil).LinkDirectoryInDepDir), arg0, arg1)
}

// WriteEnvFile mocks base method
func (_m *MockStager) WriteEnvFile(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "WriteEnvFile", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteEnvFile indicates an expected call of WriteEnvFile
func (_mr *MockStagerMockRecorder) WriteEnvFile(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "WriteEnvFile", reflect.TypeOf((*MockStager)(nil).WriteEnvFile), arg0, arg1)
}

// WriteProfileD mocks base method
func (_m *MockStager) WriteProfileD(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "WriteProfileD", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteProfileD indicates an expected call of WriteProfileD
func (_mr *MockStagerMockRecorder) WriteProfileD(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "WriteProfileD", reflect.TypeOf((*MockStager)(nil).WriteProfileD), arg0, arg1)
}

// SetStagingEnvironment mocks base method
func (_m *MockStager) SetStagingEnvironment() error {
	ret := _m.ctrl.Call(_m, "SetStagingEnvironment")
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStagingEnvironment indicates an expected call of SetStagingEnvironment
func (_mr *MockStagerMockRecorder) SetStagingEnvironment() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetStagingEnvironment", reflect.TypeOf((*MockStager)(nil).SetStagingEnvironment))
}
