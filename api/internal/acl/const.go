package acl

const (
	RoleAdmin       = "role:admin"
	RoleUser        = "role:user"
	RoleClientAdmin = "role:client:admin"
	RoleClientUser  = "role:client:user"
)

const (
	ModuleOAuth2  = "module:oauth2"
	ModuleAccount = "module:account"
	ModuleProject = "module:project"
	ModuleACL     = "module:acl"
	ModuleCore    = "module:core"
)

const (
	ActionLogin   = "action:login"
	ActionLogout  = "action:logout"
	ActionRefresh = "action:refresh"
	ActionCreate  = "action:create"
	ActionUpdate  = "action:update"
	ActionDelete  = "action:delete"
	ActionRead    = "action:read"
	ActionList    = "action:list"
	ActionAll     = "action:all"
)
