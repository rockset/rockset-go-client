package option

// TODO: this should be done use go:generate instead
// https://pkg.go.dev/golang.org/x/tools/cmd/stringer

const unknown = "unknown"

// GlobalAction is the type for RBAC actions that operate on global resources.
type GlobalAction int

// String returns the string representation used for the REST API call
func (a GlobalAction) String() string {
	for k, v := range globalActions {
		if v == a {
			return k
		}
	}
	return unknown
}

// GetGlobalAction returns the corresponding GlobalAction
func GetGlobalAction(action string) GlobalAction {
	if a, found := globalActions[action]; found {
		return a
	}

	return UnknownGlobalAction
}

// IsGlobalAction returns true if action is a GlobalAction
func IsGlobalAction(action string) bool {
	return GetGlobalAction(action) != UnknownGlobalAction
}

const (
	UnknownGlobalAction GlobalAction = iota
	AllGlobalActions
	GetOrgGlobal
	GetCurrentUserGlobal
	InviteUserGlobal
	DeleteUserGlobal
	ListUsersGlobal
	GetBillingGlobal
	UpdateBillingGlobal
	UpdateSettingsGlobal
	GetMetricsGlobal
	UpdateViGlobal
	ListViGlobal
	CreateWsGlobal
	ListWsGlobal
	CreateIntegrationGlobal
	DeleteIntegrationGlobal
	ListIntegrationsGlobal
	UpdateResourceOwnerGlobal
	CreateAPIKeyGlobal
	CreateRoleGlobal
	UpdateRoleGlobal
	DeleteRoleGlobal
	ListRolesGlobal
	GrantRevokeRoleGlobal
	CreateVirtualInstanceGlobal
	CreateQueryLogsCollectionGlobal
)

var globalActions = map[string]GlobalAction{
	"ALL_GLOBAL_ACTIONS":                  AllGlobalActions,
	"GET_ORG_GLOBAL":                      GetOrgGlobal,
	"GET_CURRENT_USER_GLOBAL":             GetCurrentUserGlobal,
	"INVITE_USER_GLOBAL":                  InviteUserGlobal,
	"DELETE_USER_GLOBAL":                  DeleteUserGlobal,
	"LIST_USERS_GLOBAL":                   ListUsersGlobal,
	"GET_BILLING_GLOBAL":                  GetBillingGlobal,
	"UPDATE_BILLING_GLOBAL":               UpdateBillingGlobal,
	"UPDATE_SETTINGS_GLOBAL":              UpdateSettingsGlobal,
	"GET_METRICS_GLOBAL":                  GetMetricsGlobal,
	"UPDATE_VI_GLOBAL":                    UpdateViGlobal,
	"LIST_VI_GLOBAL":                      ListViGlobal,
	"CREATE_WS_GLOBAL":                    CreateWsGlobal,
	"LIST_WS_GLOBAL":                      ListWsGlobal,
	"CREATE_INTEGRATION_GLOBAL":           CreateIntegrationGlobal,
	"DELETE_INTEGRATION_GLOBAL":           DeleteIntegrationGlobal,
	"LIST_INTEGRATIONS_GLOBAL":            ListIntegrationsGlobal,
	"UPDATE_RESOURCE_OWNER_GLOBAL":        UpdateResourceOwnerGlobal,
	"CREATE_API_KEY_GLOBAL":               CreateAPIKeyGlobal,
	"CREATE_ROLE_GLOBAL":                  CreateRoleGlobal,
	"UPDATE_ROLE_GLOBAL":                  UpdateRoleGlobal,
	"DELETE_ROLE_GLOBAL":                  DeleteRoleGlobal,
	"LIST_ROLES_GLOBAL":                   ListRolesGlobal,
	"GRANT_REVOKE_ROLE_GLOBAL":            GrantRevokeRoleGlobal,
	"CREATE_VI_GLOBAL":                    CreateVirtualInstanceGlobal,
	"CREATE_QUERY_LOGS_COLLECTION_GLOBAL": CreateQueryLogsCollectionGlobal,
}

// IntegrationAction is the type for actions that operate on integrations.
type IntegrationAction int

// GetIntegrationAction returns the corresponding IntegrationAction
func GetIntegrationAction(action string) IntegrationAction {
	if a, found := integrationActions[action]; found {
		return a
	}

	return UnknownIntegrationAction
}

// IsIntegrationAction returns true if action is an IntegrationAction
func IsIntegrationAction(action string) bool {
	return GetIntegrationAction(action) != UnknownIntegrationAction
}

// String returns the string representation used for the REST API call
func (a IntegrationAction) String() string {
	return actionName(integrationActions, a)
}

const (
	UnknownIntegrationAction IntegrationAction = iota
	AllIntegrationActions
	CreateCollectionIntegration
)

var integrationActions = map[string]IntegrationAction{
	"ALL_INTEGRATION_ACTIONS":       AllIntegrationActions,
	"CREATE_COLLECTION_INTEGRATION": CreateCollectionIntegration,
}

// WorkspaceAction is the type for actions that operate on workspaces.
type WorkspaceAction int

// GetWorkspaceAction returns the corresponding WorkspaceAction
func GetWorkspaceAction(action string) WorkspaceAction {
	if a, found := wsActions[action]; found {
		return a
	}

	return UnknownWorkspaceAction
}

// IsWorkspaceAction returns true if action is an WorkspaceAction
func IsWorkspaceAction(action string) bool {
	return GetWorkspaceAction(action) != UnknownWorkspaceAction
}

// String returns the string representation used for the REST API call
func (a WorkspaceAction) String() string {
	return actionName(wsActions, a)
}

const (
	UnknownWorkspaceAction WorkspaceAction = iota
	AllWorkspaceActions
	DeleteWs
	QueryDataWs
	WriteDataWs
	CreateCollectionWs
	DeleteCollectionWs
	CreateAliasWs
	DeleteAliasWs
	ListResourcesWs
	CreateQueryLambdaWs
	DeleteQueryLambdaWs
	ExecuteQueryLambdaWs
	CreateViewWs
	DeleteViewWs
	CreateSnapshotWs
	CreateScheduledLambdaWs
	DeleteScheduledLambdaWs
	CreateSimilarityIndexWs
	DeleteSimilarityIndexWs
)

var wsActions = map[string]WorkspaceAction{
	"ALL_WORKSPACE_ACTIONS":      AllWorkspaceActions,
	"DELETE_WS":                  DeleteWs,
	"QUERY_DATA_WS":              QueryDataWs,
	"WRITE_DATA_WS":              WriteDataWs,
	"CREATE_COLLECTION_WS":       CreateCollectionWs,
	"DELETE_COLLECTION_WS":       DeleteCollectionWs,
	"CREATE_ALIAS_WS":            CreateAliasWs,
	"DELETE_ALIAS_WS":            DeleteAliasWs,
	"LIST_RESOURCES_WS":          ListResourcesWs,
	"CREATE_QUERY_LAMBDA_WS":     CreateQueryLambdaWs,
	"DELETE_QUERY_LAMBDA_WS":     DeleteQueryLambdaWs,
	"EXECUTE_QUERY_LAMBDA_WS":    ExecuteQueryLambdaWs,
	"CREATE_VIEW_WS":             CreateViewWs,
	"DELETE_VIEW_WS":             DeleteViewWs,
	"CREATE_SNAPSHOT_WS":         CreateSnapshotWs,
	"CREATE_SCHEDULED_LAMBDA_WS": CreateScheduledLambdaWs,
	"DELETE_SCHEDULED_LAMBDA_WS": DeleteScheduledLambdaWs,
	"CREATE_SIMILARITY_INDEX_WS": CreateSimilarityIndexWs,
	"DELETE_SIMILARITY_INDEX_WS": DeleteSimilarityIndexWs,
}

// VirtualInstanceAction is the type for actions that operate on virtual instances.
type VirtualInstanceAction int

// GetVirtualInstanceAction returns the corresponding VirtualInstanceAction
func GetVirtualInstanceAction(action string) VirtualInstanceAction {
	if a, found := viActions[action]; found {
		return a
	}

	return UnknownVirtualInstanceAction
}

// IsVirtualInstanceAction returns true if action is an VirtualInstanceAction
func IsVirtualInstanceAction(action string) bool {
	return GetVirtualInstanceAction(action) != UnknownVirtualInstanceAction
}

// String returns the string representation used for the REST API call
func (a VirtualInstanceAction) String() string {
	return actionName(viActions, a)
}

const (
	UnknownVirtualInstanceAction VirtualInstanceAction = iota
	AllVirtualInstanceAction
	QueryVirtualInstanceAction
	UpdateVirtualInstanceAction
	SuspendResumeVirtualInstanceAction
	DeleteVirtualInstanceAction
)

var viActions = map[string]VirtualInstanceAction{
	"ALL_VI_ACTIONS":    AllVirtualInstanceAction,
	"QUERY_VI":          QueryVirtualInstanceAction,
	"UPDATE_VI":         UpdateVirtualInstanceAction,
	"SUSPEND_RESUME_VI": SuspendResumeVirtualInstanceAction,
	"DELETE_VI":         DeleteVirtualInstanceAction,
}

func actionName[T comparable](m map[string]T, a T) string {
	for k, v := range m {
		if v == a {
			return k
		}
	}

	return unknown
}
