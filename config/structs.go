package config

type Config struct {
	Debug          bool                 `yaml:"debug,omitempty"`
	Server         ServerConfig         `yaml:"server"`
	Database       DatabaseConfig       `yaml:"database"`
	Initialization InitializationConfig `yaml:"initialization"`
	Providers      ProvidersConfig      `yaml:"providers,omitempty"`
	Auth           AuthConfig           `yaml:"auth"`
}

type ServerConfig struct {
	Hostname         string   `yaml:"hostname"`
	Port             int      `yaml:"port"`
	SSL              bool     `yaml:"ssl"`
	AllowedOrigins   []string `yaml:"origins,omitempty"`
	GQlTrace         bool     `yaml:"gql_trace,omitempty"`
	GQlIntrospection bool     `yaml:"gql_introspection,omitempty"`
}

type DatabaseConfig struct {
	Username string  `yaml:"username"`
	Password string  `yaml:"password"`
	Host     string  `yaml:"host"`
	Database *string `yaml:"database,omitempty"`
	Port     *int    `yaml:"port,omitempty"`
	SSL      *bool   `yaml:"ssl,omitempty"`
}

type InitializationConfig struct {
	DefaultAdmin DefaultAdminConfig `yaml:"default_admin"`
	AdminGroup   string             `yaml:"admin_group"`
	Permissions  []PermissionConfig `yaml:"permissions,omitempty"`
}

type PermissionConfig struct {
	Key         string `yaml:"key"`
	Component   string `yaml:"component,omitempty"`
	Description string `yaml:"description,omitempty"`
}

type DefaultAdminConfig struct {
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Email     string `yaml:"email"`
	FirstName string `yaml:"first_name"`
	LastName  string `yaml:"last_name"`
}

type ProvidersConfig struct {
	CacheDir string `yaml:"cache"`
}

type AuthConfig struct {
	JWTKey         string `yaml:"jwt_key"`
	SessionTimeout int    `yaml:"session_timeout"`
}
