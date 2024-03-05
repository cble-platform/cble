package providers

type ProviderType string

const (
	TypeDocker ProviderType = "docker"
	TypeShell  ProviderType = "shell"
)

type ProviderMetadata struct {
	Name        string          `yaml:"name"`
	Description string          `yaml:"description"`
	Author      string          `yaml:"author"`
	Version     string          `yaml:"version"`
	Type        ProviderType    `yaml:"type"`
	DockerMeta  *DockerMetadata `yaml:"docker,omitempty"`
	ShellMeta   *ShellMetadata  `yaml:"shell,omitempty"`
}

type DockerMetadata struct {
	Dockerfile string `yaml:"dockerfile"`
	Command    string `yaml:"cmd"`
}

type ShellMetadata struct {
	PrebuildCommand string `yaml:"prebuild_cmd"`
	BuildCommand    string `yaml:"build_cmd"`
	ExecCommand     string `yaml:"exec_cmd"`
}

// Docker build structs

type DockerErrorDetail struct {
	Message string `json:"message"`
}

type DockerErrorLine struct {
	Error       string            `json:"error"`
	ErrorDetail DockerErrorDetail `json:"errorDetail"`
}
