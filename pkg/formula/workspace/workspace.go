package workspace

const (
	workspacesFile       = "/formula_workspaces.json"
	DefaultWorkspaceName = "Default"
	DefaultWorkspaceDir  = "/ritchie-formulas-local"
)

type Workspaces map[string]string

type Workspace struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
}

type Adder interface {
	Add(workspace Workspace) error
}

type Lister interface {
	List() (Workspaces, error)
}

type Validator interface {
	Validate(workspace Workspace) error
}

type AddListValidator interface {
	Adder
	Lister
	Validator
}
