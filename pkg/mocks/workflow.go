// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package mocks is a generated GoMock package.
package mocks

import (
	log "log"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	zerolog "github.com/rs/zerolog"
	analytics "github.com/snyk/go-application-framework/pkg/analytics"
	configuration "github.com/snyk/go-application-framework/pkg/configuration"
	networking "github.com/snyk/go-application-framework/pkg/networking"
	runtimeinfo "github.com/snyk/go-application-framework/pkg/runtimeinfo"
	ui "github.com/snyk/go-application-framework/pkg/ui"
	workflow "github.com/snyk/go-application-framework/pkg/workflow"
)

// MockData is a mock of Data interface.
type MockData struct {
	ctrl     *gomock.Controller
	recorder *MockDataMockRecorder
}

// MockDataMockRecorder is the mock recorder for MockData.
type MockDataMockRecorder struct {
	mock *MockData
}

// NewMockData creates a new mock instance.
func NewMockData(ctrl *gomock.Controller) *MockData {
	mock := &MockData{ctrl: ctrl}
	mock.recorder = &MockDataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockData) EXPECT() *MockDataMockRecorder {
	return m.recorder
}

// GetContentLocation mocks base method.
func (m *MockData) GetContentLocation() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContentLocation")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetContentLocation indicates an expected call of GetContentLocation.
func (mr *MockDataMockRecorder) GetContentLocation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContentLocation", reflect.TypeOf((*MockData)(nil).GetContentLocation))
}

// GetContentType mocks base method.
func (m *MockData) GetContentType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContentType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetContentType indicates an expected call of GetContentType.
func (mr *MockDataMockRecorder) GetContentType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContentType", reflect.TypeOf((*MockData)(nil).GetContentType))
}

// GetIdentifier mocks base method.
func (m *MockData) GetIdentifier() workflow.Identifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdentifier")
	ret0, _ := ret[0].(workflow.Identifier)
	return ret0
}

// GetIdentifier indicates an expected call of GetIdentifier.
func (mr *MockDataMockRecorder) GetIdentifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdentifier", reflect.TypeOf((*MockData)(nil).GetIdentifier))
}

// GetMetaData mocks base method.
func (m *MockData) GetMetaData(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetaData", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetaData indicates an expected call of GetMetaData.
func (mr *MockDataMockRecorder) GetMetaData(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetaData", reflect.TypeOf((*MockData)(nil).GetMetaData), key)
}

// GetPayload mocks base method.
func (m *MockData) GetPayload() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayload")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetPayload indicates an expected call of GetPayload.
func (mr *MockDataMockRecorder) GetPayload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayload", reflect.TypeOf((*MockData)(nil).GetPayload))
}

// SetContentLocation mocks base method.
func (m *MockData) SetContentLocation(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetContentLocation", arg0)
}

// SetContentLocation indicates an expected call of SetContentLocation.
func (mr *MockDataMockRecorder) SetContentLocation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContentLocation", reflect.TypeOf((*MockData)(nil).SetContentLocation), arg0)
}

// SetMetaData mocks base method.
func (m *MockData) SetMetaData(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMetaData", key, value)
}

// SetMetaData indicates an expected call of SetMetaData.
func (mr *MockDataMockRecorder) SetMetaData(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMetaData", reflect.TypeOf((*MockData)(nil).SetMetaData), key, value)
}

// SetPayload mocks base method.
func (m *MockData) SetPayload(payload interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPayload", payload)
}

// SetPayload indicates an expected call of SetPayload.
func (mr *MockDataMockRecorder) SetPayload(payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPayload", reflect.TypeOf((*MockData)(nil).SetPayload), payload)
}

// MockInvocationContext is a mock of InvocationContext interface.
type MockInvocationContext struct {
	ctrl     *gomock.Controller
	recorder *MockInvocationContextMockRecorder
}

// MockInvocationContextMockRecorder is the mock recorder for MockInvocationContext.
type MockInvocationContextMockRecorder struct {
	mock *MockInvocationContext
}

// NewMockInvocationContext creates a new mock instance.
func NewMockInvocationContext(ctrl *gomock.Controller) *MockInvocationContext {
	mock := &MockInvocationContext{ctrl: ctrl}
	mock.recorder = &MockInvocationContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvocationContext) EXPECT() *MockInvocationContextMockRecorder {
	return m.recorder
}

// GetAnalytics mocks base method.
func (m *MockInvocationContext) GetAnalytics() analytics.Analytics {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnalytics")
	ret0, _ := ret[0].(analytics.Analytics)
	return ret0
}

// GetAnalytics indicates an expected call of GetAnalytics.
func (mr *MockInvocationContextMockRecorder) GetAnalytics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnalytics", reflect.TypeOf((*MockInvocationContext)(nil).GetAnalytics))
}

// GetConfiguration mocks base method.
func (m *MockInvocationContext) GetConfiguration() configuration.Configuration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfiguration")
	ret0, _ := ret[0].(configuration.Configuration)
	return ret0
}

// GetConfiguration indicates an expected call of GetConfiguration.
func (mr *MockInvocationContextMockRecorder) GetConfiguration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfiguration", reflect.TypeOf((*MockInvocationContext)(nil).GetConfiguration))
}

// GetEngine mocks base method.
func (m *MockInvocationContext) GetEngine() workflow.Engine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEngine")
	ret0, _ := ret[0].(workflow.Engine)
	return ret0
}

// GetEngine indicates an expected call of GetEngine.
func (mr *MockInvocationContextMockRecorder) GetEngine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEngine", reflect.TypeOf((*MockInvocationContext)(nil).GetEngine))
}

// GetEnhancedLogger mocks base method.
func (m *MockInvocationContext) GetEnhancedLogger() *zerolog.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnhancedLogger")
	ret0, _ := ret[0].(*zerolog.Logger)
	return ret0
}

// GetEnhancedLogger indicates an expected call of GetEnhancedLogger.
func (mr *MockInvocationContextMockRecorder) GetEnhancedLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnhancedLogger", reflect.TypeOf((*MockInvocationContext)(nil).GetEnhancedLogger))
}

// GetLogger mocks base method.
func (m *MockInvocationContext) GetLogger() *log.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(*log.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockInvocationContextMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockInvocationContext)(nil).GetLogger))
}

// GetNetworkAccess mocks base method.
func (m *MockInvocationContext) GetNetworkAccess() networking.NetworkAccess {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkAccess")
	ret0, _ := ret[0].(networking.NetworkAccess)
	return ret0
}

// GetNetworkAccess indicates an expected call of GetNetworkAccess.
func (mr *MockInvocationContextMockRecorder) GetNetworkAccess() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkAccess", reflect.TypeOf((*MockInvocationContext)(nil).GetNetworkAccess))
}

// GetRuntimeInfo mocks base method.
func (m *MockInvocationContext) GetRuntimeInfo() runtimeinfo.RuntimeInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRuntimeInfo")
	ret0, _ := ret[0].(runtimeinfo.RuntimeInfo)
	return ret0
}

// GetRuntimeInfo indicates an expected call of GetRuntimeInfo.
func (mr *MockInvocationContextMockRecorder) GetRuntimeInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuntimeInfo", reflect.TypeOf((*MockInvocationContext)(nil).GetRuntimeInfo))
}

// GetUserInterface mocks base method.
func (m *MockInvocationContext) GetUserInterface() ui.UserInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInterface")
	ret0, _ := ret[0].(ui.UserInterface)
	return ret0
}

// GetUserInterface indicates an expected call of GetUserInterface.
func (mr *MockInvocationContextMockRecorder) GetUserInterface() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInterface", reflect.TypeOf((*MockInvocationContext)(nil).GetUserInterface))
}

// GetWorkflowIdentifier mocks base method.
func (m *MockInvocationContext) GetWorkflowIdentifier() workflow.Identifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflowIdentifier")
	ret0, _ := ret[0].(workflow.Identifier)
	return ret0
}

// GetWorkflowIdentifier indicates an expected call of GetWorkflowIdentifier.
func (mr *MockInvocationContextMockRecorder) GetWorkflowIdentifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowIdentifier", reflect.TypeOf((*MockInvocationContext)(nil).GetWorkflowIdentifier))
}

// MockConfigurationOptions is a mock of ConfigurationOptions interface.
type MockConfigurationOptions struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationOptionsMockRecorder
}

// MockConfigurationOptionsMockRecorder is the mock recorder for MockConfigurationOptions.
type MockConfigurationOptionsMockRecorder struct {
	mock *MockConfigurationOptions
}

// NewMockConfigurationOptions creates a new mock instance.
func NewMockConfigurationOptions(ctrl *gomock.Controller) *MockConfigurationOptions {
	mock := &MockConfigurationOptions{ctrl: ctrl}
	mock.recorder = &MockConfigurationOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurationOptions) EXPECT() *MockConfigurationOptionsMockRecorder {
	return m.recorder
}

// MockEntry is a mock of Entry interface.
type MockEntry struct {
	ctrl     *gomock.Controller
	recorder *MockEntryMockRecorder
}

// MockEntryMockRecorder is the mock recorder for MockEntry.
type MockEntryMockRecorder struct {
	mock *MockEntry
}

// NewMockEntry creates a new mock instance.
func NewMockEntry(ctrl *gomock.Controller) *MockEntry {
	mock := &MockEntry{ctrl: ctrl}
	mock.recorder = &MockEntryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntry) EXPECT() *MockEntryMockRecorder {
	return m.recorder
}

// GetConfigurationOptions mocks base method.
func (m *MockEntry) GetConfigurationOptions() workflow.ConfigurationOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigurationOptions")
	ret0, _ := ret[0].(workflow.ConfigurationOptions)
	return ret0
}

// GetConfigurationOptions indicates an expected call of GetConfigurationOptions.
func (mr *MockEntryMockRecorder) GetConfigurationOptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigurationOptions", reflect.TypeOf((*MockEntry)(nil).GetConfigurationOptions))
}

// GetEntryPoint mocks base method.
func (m *MockEntry) GetEntryPoint() workflow.Callback {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntryPoint")
	ret0, _ := ret[0].(workflow.Callback)
	return ret0
}

// GetEntryPoint indicates an expected call of GetEntryPoint.
func (mr *MockEntryMockRecorder) GetEntryPoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntryPoint", reflect.TypeOf((*MockEntry)(nil).GetEntryPoint))
}

// IsVisible mocks base method.
func (m *MockEntry) IsVisible() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsVisible")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsVisible indicates an expected call of IsVisible.
func (mr *MockEntryMockRecorder) IsVisible() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVisible", reflect.TypeOf((*MockEntry)(nil).IsVisible))
}

// SetVisibility mocks base method.
func (m *MockEntry) SetVisibility(visible bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVisibility", visible)
}

// SetVisibility indicates an expected call of SetVisibility.
func (mr *MockEntryMockRecorder) SetVisibility(visible interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVisibility", reflect.TypeOf((*MockEntry)(nil).SetVisibility), visible)
}

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// AddExtensionInitializer mocks base method.
func (m *MockEngine) AddExtensionInitializer(initializer workflow.ExtensionInit) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddExtensionInitializer", initializer)
}

// AddExtensionInitializer indicates an expected call of AddExtensionInitializer.
func (mr *MockEngineMockRecorder) AddExtensionInitializer(initializer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddExtensionInitializer", reflect.TypeOf((*MockEngine)(nil).AddExtensionInitializer), initializer)
}

// GetAnalytics mocks base method.
func (m *MockEngine) GetAnalytics() analytics.Analytics {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnalytics")
	ret0, _ := ret[0].(analytics.Analytics)
	return ret0
}

// GetAnalytics indicates an expected call of GetAnalytics.
func (mr *MockEngineMockRecorder) GetAnalytics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnalytics", reflect.TypeOf((*MockEngine)(nil).GetAnalytics))
}

// GetConfiguration mocks base method.
func (m *MockEngine) GetConfiguration() configuration.Configuration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfiguration")
	ret0, _ := ret[0].(configuration.Configuration)
	return ret0
}

// GetConfiguration indicates an expected call of GetConfiguration.
func (mr *MockEngineMockRecorder) GetConfiguration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfiguration", reflect.TypeOf((*MockEngine)(nil).GetConfiguration))
}

// GetLogger mocks base method.
func (m *MockEngine) GetLogger() *zerolog.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(*zerolog.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockEngineMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockEngine)(nil).GetLogger))
}

// GetNetworkAccess mocks base method.
func (m *MockEngine) GetNetworkAccess() networking.NetworkAccess {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkAccess")
	ret0, _ := ret[0].(networking.NetworkAccess)
	return ret0
}

// GetNetworkAccess indicates an expected call of GetNetworkAccess.
func (mr *MockEngineMockRecorder) GetNetworkAccess() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkAccess", reflect.TypeOf((*MockEngine)(nil).GetNetworkAccess))
}

// GetRuntimeInfo mocks base method.
func (m *MockEngine) GetRuntimeInfo() runtimeinfo.RuntimeInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRuntimeInfo")
	ret0, _ := ret[0].(runtimeinfo.RuntimeInfo)
	return ret0
}

// GetRuntimeInfo indicates an expected call of GetRuntimeInfo.
func (mr *MockEngineMockRecorder) GetRuntimeInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuntimeInfo", reflect.TypeOf((*MockEngine)(nil).GetRuntimeInfo))
}

// GetUserInterface mocks base method.
func (m *MockEngine) GetUserInterface() ui.UserInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInterface")
	ret0, _ := ret[0].(ui.UserInterface)
	return ret0
}

// GetUserInterface indicates an expected call of GetUserInterface.
func (mr *MockEngineMockRecorder) GetUserInterface() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInterface", reflect.TypeOf((*MockEngine)(nil).GetUserInterface))
}

// GetWorkflow mocks base method.
func (m *MockEngine) GetWorkflow(id workflow.Identifier) (workflow.Entry, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflow", id)
	ret0, _ := ret[0].(workflow.Entry)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetWorkflow indicates an expected call of GetWorkflow.
func (mr *MockEngineMockRecorder) GetWorkflow(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflow", reflect.TypeOf((*MockEngine)(nil).GetWorkflow), id)
}

// GetWorkflows mocks base method.
func (m *MockEngine) GetWorkflows() []workflow.Identifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflows")
	ret0, _ := ret[0].([]workflow.Identifier)
	return ret0
}

// GetWorkflows indicates an expected call of GetWorkflows.
func (mr *MockEngineMockRecorder) GetWorkflows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflows", reflect.TypeOf((*MockEngine)(nil).GetWorkflows))
}

// Init mocks base method.
func (m *MockEngine) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockEngineMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockEngine)(nil).Init))
}

// Invoke mocks base method.
func (m *MockEngine) Invoke(id workflow.Identifier) ([]workflow.Data, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Invoke", id)
	ret0, _ := ret[0].([]workflow.Data)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Invoke indicates an expected call of Invoke.
func (mr *MockEngineMockRecorder) Invoke(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invoke", reflect.TypeOf((*MockEngine)(nil).Invoke), id)
}

// InvokeWithConfig mocks base method.
func (m *MockEngine) InvokeWithConfig(id workflow.Identifier, config configuration.Configuration) ([]workflow.Data, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeWithConfig", id, config)
	ret0, _ := ret[0].([]workflow.Data)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeWithConfig indicates an expected call of InvokeWithConfig.
func (mr *MockEngineMockRecorder) InvokeWithConfig(id, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeWithConfig", reflect.TypeOf((*MockEngine)(nil).InvokeWithConfig), id, config)
}

// InvokeWithInput mocks base method.
func (m *MockEngine) InvokeWithInput(id workflow.Identifier, input []workflow.Data) ([]workflow.Data, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeWithInput", id, input)
	ret0, _ := ret[0].([]workflow.Data)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeWithInput indicates an expected call of InvokeWithInput.
func (mr *MockEngineMockRecorder) InvokeWithInput(id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeWithInput", reflect.TypeOf((*MockEngine)(nil).InvokeWithInput), id, input)
}

// InvokeWithInputAndConfig mocks base method.
func (m *MockEngine) InvokeWithInputAndConfig(id workflow.Identifier, input []workflow.Data, config configuration.Configuration) ([]workflow.Data, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeWithInputAndConfig", id, input, config)
	ret0, _ := ret[0].([]workflow.Data)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeWithInputAndConfig indicates an expected call of InvokeWithInputAndConfig.
func (mr *MockEngineMockRecorder) InvokeWithInputAndConfig(id, input, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeWithInputAndConfig", reflect.TypeOf((*MockEngine)(nil).InvokeWithInputAndConfig), id, input, config)
}

// Register mocks base method.
func (m *MockEngine) Register(id workflow.Identifier, config workflow.ConfigurationOptions, callback workflow.Callback) (workflow.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", id, config, callback)
	ret0, _ := ret[0].(workflow.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockEngineMockRecorder) Register(id, config, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockEngine)(nil).Register), id, config, callback)
}

// SetConfiguration mocks base method.
func (m *MockEngine) SetConfiguration(config configuration.Configuration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConfiguration", config)
}

// SetConfiguration indicates an expected call of SetConfiguration.
func (mr *MockEngineMockRecorder) SetConfiguration(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfiguration", reflect.TypeOf((*MockEngine)(nil).SetConfiguration), config)
}

// SetLogger mocks base method.
func (m *MockEngine) SetLogger(logger *zerolog.Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogger", logger)
}

// SetLogger indicates an expected call of SetLogger.
func (mr *MockEngineMockRecorder) SetLogger(logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogger", reflect.TypeOf((*MockEngine)(nil).SetLogger), logger)
}

// SetRuntimeInfo mocks base method.
func (m *MockEngine) SetRuntimeInfo(ri runtimeinfo.RuntimeInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRuntimeInfo", ri)
}

// SetRuntimeInfo indicates an expected call of SetRuntimeInfo.
func (mr *MockEngineMockRecorder) SetRuntimeInfo(ri interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRuntimeInfo", reflect.TypeOf((*MockEngine)(nil).SetRuntimeInfo), ri)
}

// SetUserInterface mocks base method.
func (m *MockEngine) SetUserInterface(ui ui.UserInterface) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUserInterface", ui)
}

// SetUserInterface indicates an expected call of SetUserInterface.
func (mr *MockEngineMockRecorder) SetUserInterface(ui interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserInterface", reflect.TypeOf((*MockEngine)(nil).SetUserInterface), ui)
}
