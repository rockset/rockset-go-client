package option

import (
	"github.com/rockset/rockset-go-client/openapi"
)

// GlobalAction is the type for RBAC actions that operate on global resources.
type GlobalAction int

// String returns the string representation used for the REST API call
func (a GlobalAction) String() string {
	return globalActions[a]
}

const (
	AllGlobalActions GlobalAction = iota
	GetOrg
	GetCurrentUser
	InviteUser
	DeleteUser
	ListUsers
	GetBilling
	UpdateBilling
	UpdateSettings
	GetMetrics
	UpdateVi
	ListVi
	CreateWs
	ListWs
	CreateIntegration
	DeleteIntegration
	ListIntegrations
	UpdateResourceOwner
	CreateApiKey
	CreateRole
	UpdateRole
	DeleteRole
	ListRolesGlobal
	GrantRevokeRole
)

var globalActions = []string{
	"ALL_GLOBAL_ACTIONS",
	"GET_ORG_GLOBAL",
	"GET_CURRENT_USER_GLOBAL",
	"INVITE_USER_GLOBAL",
	"DELETE_USER_GLOBAL",
	"LIST_USERS_GLOBAL",
	"GET_BILLING_GLOBAL",
	"UPDATE_BILLING_GLOBAL",
	"UPDATE_SETTINGS_GLOBAL",
	"GET_METRICS_GLOBAL",
	"UPDATE_VI_GLOBAL",
	"LIST_VI_GLOBAL",
	"CREATE_WS_GLOBAL",
	"LIST_WS_GLOBAL",
	"CREATE_INTEGRATION_GLOBAL",
	"DELETE_INTEGRATION_GLOBAL",
	"LIST_INTEGRATIONS_GLOBAL",
	"UPDATE_RESOURCE_OWNER_GLOBAL",
	"CREATE_API_KEY_GLOBAL",
	"CREATE_ROLE_GLOBAL",
	"UPDATE_ROLE_GLOBAL",
	"DELETE_ROLE_GLOBAL",
	"LIST_ROLES_GLOBAL",
	"GRANT_REVOKE_ROLE_GLOBAL",
}

// IntegrationAction is the type for actions that operate on integrations.
type IntegrationAction int

// String returns the string representation used for the REST API call
func (a IntegrationAction) String() string {
	return integrationActions[a]
}

const (
	AllIntegrationActions IntegrationAction = iota
	CreateCollectionIntegration
)

var integrationActions = []string{
	"ALL_INTEGRATION_ACTIONS",
	"CREATE_COLLECTION_INTEGRATION",
}

// WorkspaceAction is the type for actions that operate on workspaces.
type WorkspaceAction int

// String returns the string representation used for the REST API call
func (a WorkspaceAction) String() string {
	return wsActions[a]
}

const (
	AllWorkspaceActions WorkspaceAction = iota
	DeleteWorkspace
	QueryData
	WriteData
	CreateCollection
	DeleteCollection
	CreateAlias
	DeleteAlias
	ListResources
	CreateQueryLambda
	DeleteQueryLambda
	ExecuteQueryLambda
	CreateView
	DeleteView
)

var wsActions = []string{
	"ALL_WORKSPACE_ACTIONS",
	"DELETE_WS",
	"QUERY_DATA_WS",
	"WRITE_DATA_WS",
	"CREATE_COLLECTION_WS",
	"DELETE_COLLECTION_WS",
	"CREATE_ALIAS_WS",
	"DELETE_ALIAS_WS",
	"LIST_RESOURCES_WS",
	"CREATE_QUERY_LAMBDA_WS",
	"DELETE_QUERY_LAMBDA_WS",
	"EXECUTE_QUERY_LAMBDA_WS",
	"CREATE_VIEW_WS",
	"DELETE_VIEW_WS",
}

// RoleOptions is used to hold optional workspace settings
type RoleOptions struct {
	// Description of the role
	Description *string
	Privileges  []openapi.Privilege
}

type RoleOption func(p *RoleOptions)

// WithRoleDescription is used to optionally set a role description.
func WithRoleDescription(desc string) RoleOption {
	return func(o *RoleOptions) {
		o.Description = &desc
	}
}

// WithGlobalPrivilege is used to add a global action to a role.
func WithGlobalPrivilege(action GlobalAction) RoleOption {
	return func(o *RoleOptions) {
		a := action.String()
		o.Privileges = append(o.Privileges, openapi.Privilege{
			Action: &a,
		})
	}
}

// WithIntegrationPrivilege is used to add an integration action to a role.
func WithIntegrationPrivilege(action IntegrationAction, integration string) RoleOption {
	return func(o *RoleOptions) {
		a := action.String()
		o.Privileges = append(o.Privileges, openapi.Privilege{
			Action:       &a,
			ResourceName: &integration,
		})
	}
}

// WithWorkspacePrivilege is used to add a workspace action to a role. If WithCluster isn't specified, the privilege
// applied to all clusters. Only the last WithCluster is what is used.
func WithWorkspacePrivilege(action WorkspaceAction, workspace string, options ...func(*openapi.Privilege)) RoleOption {
	return func(o *RoleOptions) {
		a := action.String()

		p := openapi.Privilege{
			Action:       &a,
			ResourceName: &workspace,
		}

		for _, opt := range options {
			opt(&p)
		}

		o.Privileges = append(o.Privileges, p)
	}
}

// WithCluster is used to specify a cluster for a WithWorkspacePrivilege.
func WithCluster(name string) func(*openapi.Privilege) {
	return func(p *openapi.Privilege) {
		p.Cluster = &name
	}
}
