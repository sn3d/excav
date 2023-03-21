package provider

type Typ string

const (
	TypGitHub Typ = "github"
	TypGitLab Typ = "gitlab"
)

type Options struct {
	Provider Typ    `yaml:"type"`
	Host     string `yaml:"host"`
	ApiHost  string `yaml:"api_host"`
	Token    string `yaml:"token"`
	User     string `yaml:"user"`
}

type Provider interface {
	GetRepositoryURL(repoName string) string
	GetToken() string
	GetUser() string
	CreateMergeRequest(repoName string, branch string, commitMsg string) (string, error)
	DeleteBranch(repoName string, branch string) error
}

func New(opts Options) Provider {
	if opts.Provider == TypGitLab {
		return &gitlabProvider{
			host:  opts.Host,
			token: opts.Token,
		}
	} else if opts.Provider == TypGitHub {
		return &githubProvider{
			gitHost: opts.Host,
			apiHost: opts.ApiHost,
			user:    opts.User,
			token:   opts.Token,
		}
	} else {
		//panic("no provider configured")
		return nil
	}
}
