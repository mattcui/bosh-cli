// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/cloudfoundry/bosh-init/installation (interfaces: Installation,Installer,InstallerFactory)

package mocks

import (
	gomock "code.google.com/p/gomock/gomock"
	installation "github.com/cloudfoundry/bosh-init/installation"
	job "github.com/cloudfoundry/bosh-init/installation/job"
	manifest "github.com/cloudfoundry/bosh-init/installation/manifest"
	ui "github.com/cloudfoundry/bosh-init/ui"
	"github.com/cloudfoundry/bosh-utils/logger" // Mock of Installation interface
)

type MockInstallation struct {
	ctrl     *gomock.Controller
	recorder *_MockInstallationRecorder
}

// Recorder for MockInstallation (not exported)
type _MockInstallationRecorder struct {
	mock *MockInstallation
}

func NewMockInstallation(ctrl *gomock.Controller) *MockInstallation {
	mock := &MockInstallation{ctrl: ctrl}
	mock.recorder = &_MockInstallationRecorder{mock}
	return mock
}

func (_m *MockInstallation) EXPECT() *_MockInstallationRecorder {
	return _m.recorder
}

func (_m *MockInstallation) Job() job.InstalledJob {
	ret := _m.ctrl.Call(_m, "Job")
	ret0, _ := ret[0].(job.InstalledJob)
	return ret0
}

func (_mr *_MockInstallationRecorder) Job() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Job")
}

func (_m *MockInstallation) StartRegistry() error {
	ret := _m.ctrl.Call(_m, "StartRegistry")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockInstallationRecorder) StartRegistry() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartRegistry")
}

func (_m *MockInstallation) StopRegistry() error {
	ret := _m.ctrl.Call(_m, "StopRegistry")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockInstallationRecorder) StopRegistry() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopRegistry")
}

func (_m *MockInstallation) Target() installation.Target {
	ret := _m.ctrl.Call(_m, "Target")
	ret0, _ := ret[0].(installation.Target)
	return ret0
}

func (_mr *_MockInstallationRecorder) Target() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Target")
}

func (_m *MockInstallation) WithRunningRegistry(_param0 logger.Logger, _param1 ui.Stage, _param2 func() error) error {
	ret := _m.ctrl.Call(_m, "WithRunningRegistry", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockInstallationRecorder) WithRunningRegistry(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WithRunningRegistry", arg0, arg1, arg2)
}

// Mock of Installer interface
type MockInstaller struct {
	ctrl     *gomock.Controller
	recorder *_MockInstallerRecorder
}

// Recorder for MockInstaller (not exported)
type _MockInstallerRecorder struct {
	mock *MockInstaller
}

func NewMockInstaller(ctrl *gomock.Controller) *MockInstaller {
	mock := &MockInstaller{ctrl: ctrl}
	mock.recorder = &_MockInstallerRecorder{mock}
	return mock
}

func (_m *MockInstaller) EXPECT() *_MockInstallerRecorder {
	return _m.recorder
}

func (_m *MockInstaller) Install(_param0 manifest.Manifest, _param1 ui.Stage) (installation.Installation, error) {
	ret := _m.ctrl.Call(_m, "Install", _param0, _param1)
	ret0, _ := ret[0].(installation.Installation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstallerRecorder) Install(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Install", arg0, arg1)
}

// Mock of InstallerFactory interface
type MockInstallerFactory struct {
	ctrl     *gomock.Controller
	recorder *_MockInstallerFactoryRecorder
}

// Recorder for MockInstallerFactory (not exported)
type _MockInstallerFactoryRecorder struct {
	mock *MockInstallerFactory
}

func NewMockInstallerFactory(ctrl *gomock.Controller) *MockInstallerFactory {
	mock := &MockInstallerFactory{ctrl: ctrl}
	mock.recorder = &_MockInstallerFactoryRecorder{mock}
	return mock
}

func (_m *MockInstallerFactory) EXPECT() *_MockInstallerFactoryRecorder {
	return _m.recorder
}

func (_m *MockInstallerFactory) NewInstaller() (installation.Installer, error) {
	ret := _m.ctrl.Call(_m, "NewInstaller")
	ret0, _ := ret[0].(installation.Installer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstallerFactoryRecorder) NewInstaller() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewInstaller")
}
