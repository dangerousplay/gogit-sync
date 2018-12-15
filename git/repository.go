package git

type RepoType string
type Storage string
type AuthMode string

const (
	GITLAB    RepoType = "GITLAB"
	GITHUB    RepoType = "GITHUB"
	BITBUCKET RepoType = "BITBUCKET"
)

const (
	NONE AuthMode = "NONE"
	USER AuthMode = "USER"
	SSH  AuthMode = "SSH"
)

const (
	REMOTE Storage = "REMOTE"
	LOCAL  Storage = "LOCAL"
)

type Location struct {
	Type  Storage
	Value string
}

type Auth struct {
	Type        AuthMode
	Credentials []string
}

type Repository struct {
	Name     string
	Type     RepoType
	Location Location
	Auth     Auth
}

type ServiceRepository struct {
	Repository Repository
	Mirrors    []Repository
}
