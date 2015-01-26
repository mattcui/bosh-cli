// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/cloudfoundry/bosh-micro-cli/deployment/instance (interfaces: Instance,Manager,StateBuilderFactory,StateBuilder,State)

package mocks

import (
	gomock "code.google.com/p/gomock/gomock"
	blobstore "github.com/cloudfoundry/bosh-micro-cli/blobstore"
	applyspec "github.com/cloudfoundry/bosh-micro-cli/deployment/applyspec"
	disk "github.com/cloudfoundry/bosh-micro-cli/deployment/disk"
	instance "github.com/cloudfoundry/bosh-micro-cli/deployment/instance"
	manifest0 "github.com/cloudfoundry/bosh-micro-cli/deployment/manifest"
	stemcell "github.com/cloudfoundry/bosh-micro-cli/deployment/stemcell"
	eventlogger "github.com/cloudfoundry/bosh-micro-cli/eventlogger"
	manifest "github.com/cloudfoundry/bosh-micro-cli/installation/manifest"
	time "time"
)

// Mock of Instance interface
type MockInstance struct {
	ctrl     *gomock.Controller
	recorder *_MockInstanceRecorder
}

// Recorder for MockInstance (not exported)
type _MockInstanceRecorder struct {
	mock *MockInstance
}

func NewMockInstance(ctrl *gomock.Controller) *MockInstance {
	mock := &MockInstance{ctrl: ctrl}
	mock.recorder = &_MockInstanceRecorder{mock}
	return mock
}

func (_m *MockInstance) EXPECT() *_MockInstanceRecorder {
	return _m.recorder
}

func (_m *MockInstance) Delete(_param0 time.Duration, _param1 time.Duration, _param2 eventlogger.Stage) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockInstanceRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0, arg1, arg2)
}

func (_m *MockInstance) Disks() ([]disk.Disk, error) {
	ret := _m.ctrl.Call(_m, "Disks")
	ret0, _ := ret[0].([]disk.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstanceRecorder) Disks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Disks")
}

func (_m *MockInstance) ID() int {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockInstanceRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockInstance) JobName() string {
	ret := _m.ctrl.Call(_m, "JobName")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockInstanceRecorder) JobName() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "JobName")
}

func (_m *MockInstance) UpdateDisks(_param0 manifest0.Manifest, _param1 eventlogger.Stage) ([]disk.Disk, error) {
	ret := _m.ctrl.Call(_m, "UpdateDisks", _param0, _param1)
	ret0, _ := ret[0].([]disk.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstanceRecorder) UpdateDisks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateDisks", arg0, arg1)
}

func (_m *MockInstance) UpdateJobs(_param0 manifest0.Manifest, _param1 stemcell.ApplySpec, _param2 eventlogger.Stage) error {
	ret := _m.ctrl.Call(_m, "UpdateJobs", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockInstanceRecorder) UpdateJobs(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateJobs", arg0, arg1, arg2)
}

func (_m *MockInstance) WaitUntilReady(_param0 manifest.Registry, _param1 manifest.SSHTunnel, _param2 eventlogger.Stage) error {
	ret := _m.ctrl.Call(_m, "WaitUntilReady", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockInstanceRecorder) WaitUntilReady(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WaitUntilReady", arg0, arg1, arg2)
}

// Mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *_MockManagerRecorder
}

// Recorder for MockManager (not exported)
type _MockManagerRecorder struct {
	mock *MockManager
}

func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &_MockManagerRecorder{mock}
	return mock
}

func (_m *MockManager) EXPECT() *_MockManagerRecorder {
	return _m.recorder
}

func (_m *MockManager) Create(_param0 string, _param1 int, _param2 manifest0.Manifest, _param3 stemcell.CloudStemcell, _param4 manifest.Registry, _param5 manifest.SSHTunnel, _param6 eventlogger.Stage) (instance.Instance, []disk.Disk, error) {
	ret := _m.ctrl.Call(_m, "Create", _param0, _param1, _param2, _param3, _param4, _param5, _param6)
	ret0, _ := ret[0].(instance.Instance)
	ret1, _ := ret[1].([]disk.Disk)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockManagerRecorder) Create(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

func (_m *MockManager) DeleteAll(_param0 time.Duration, _param1 time.Duration, _param2 eventlogger.Stage) error {
	ret := _m.ctrl.Call(_m, "DeleteAll", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) DeleteAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAll", arg0, arg1, arg2)
}

func (_m *MockManager) FindCurrent() ([]instance.Instance, error) {
	ret := _m.ctrl.Call(_m, "FindCurrent")
	ret0, _ := ret[0].([]instance.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) FindCurrent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindCurrent")
}

// Mock of StateBuilderFactory interface
type MockStateBuilderFactory struct {
	ctrl     *gomock.Controller
	recorder *_MockStateBuilderFactoryRecorder
}

// Recorder for MockStateBuilderFactory (not exported)
type _MockStateBuilderFactoryRecorder struct {
	mock *MockStateBuilderFactory
}

func NewMockStateBuilderFactory(ctrl *gomock.Controller) *MockStateBuilderFactory {
	mock := &MockStateBuilderFactory{ctrl: ctrl}
	mock.recorder = &_MockStateBuilderFactoryRecorder{mock}
	return mock
}

func (_m *MockStateBuilderFactory) EXPECT() *_MockStateBuilderFactoryRecorder {
	return _m.recorder
}

func (_m *MockStateBuilderFactory) NewStateBuilder(_param0 blobstore.Blobstore) instance.StateBuilder {
	ret := _m.ctrl.Call(_m, "NewStateBuilder", _param0)
	ret0, _ := ret[0].(instance.StateBuilder)
	return ret0
}

func (_mr *_MockStateBuilderFactoryRecorder) NewStateBuilder(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewStateBuilder", arg0)
}

// Mock of StateBuilder interface
type MockStateBuilder struct {
	ctrl     *gomock.Controller
	recorder *_MockStateBuilderRecorder
}

// Recorder for MockStateBuilder (not exported)
type _MockStateBuilderRecorder struct {
	mock *MockStateBuilder
}

func NewMockStateBuilder(ctrl *gomock.Controller) *MockStateBuilder {
	mock := &MockStateBuilder{ctrl: ctrl}
	mock.recorder = &_MockStateBuilderRecorder{mock}
	return mock
}

func (_m *MockStateBuilder) EXPECT() *_MockStateBuilderRecorder {
	return _m.recorder
}

func (_m *MockStateBuilder) Build(_param0 string, _param1 int, _param2 manifest0.Manifest, _param3 stemcell.ApplySpec) (instance.State, error) {
	ret := _m.ctrl.Call(_m, "Build", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(instance.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStateBuilderRecorder) Build(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Build", arg0, arg1, arg2, arg3)
}

// Mock of State interface
type MockState struct {
	ctrl     *gomock.Controller
	recorder *_MockStateRecorder
}

// Recorder for MockState (not exported)
type _MockStateRecorder struct {
	mock *MockState
}

func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &_MockStateRecorder{mock}
	return mock
}

func (_m *MockState) EXPECT() *_MockStateRecorder {
	return _m.recorder
}

func (_m *MockState) CompiledPackages() []instance.PackageRef {
	ret := _m.ctrl.Call(_m, "CompiledPackages")
	ret0, _ := ret[0].([]instance.PackageRef)
	return ret0
}

func (_mr *_MockStateRecorder) CompiledPackages() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CompiledPackages")
}

func (_m *MockState) NetworkInterfaces() []instance.NetworkRef {
	ret := _m.ctrl.Call(_m, "NetworkInterfaces")
	ret0, _ := ret[0].([]instance.NetworkRef)
	return ret0
}

func (_mr *_MockStateRecorder) NetworkInterfaces() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NetworkInterfaces")
}

func (_m *MockState) RenderedJobListArchive() instance.BlobRef {
	ret := _m.ctrl.Call(_m, "RenderedJobListArchive")
	ret0, _ := ret[0].(instance.BlobRef)
	return ret0
}

func (_mr *_MockStateRecorder) RenderedJobListArchive() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RenderedJobListArchive")
}

func (_m *MockState) RenderedJobs() []instance.JobRef {
	ret := _m.ctrl.Call(_m, "RenderedJobs")
	ret0, _ := ret[0].([]instance.JobRef)
	return ret0
}

func (_mr *_MockStateRecorder) RenderedJobs() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RenderedJobs")
}

func (_m *MockState) ToApplySpec() applyspec.ApplySpec {
	ret := _m.ctrl.Call(_m, "ToApplySpec")
	ret0, _ := ret[0].(applyspec.ApplySpec)
	return ret0
}

func (_mr *_MockStateRecorder) ToApplySpec() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ToApplySpec")
}
