package config

import "time"

type Config struct {
	Debug           bool                  `yaml:"debug,omitempty"`
	Server          ServerConfig          `yaml:"server"`
	Database        DatabaseConfig        `yaml:"database"`
	Initialization  InitializationConfig  `yaml:"initialization"`
	Providers       ProvidersConfig       `yaml:"providers,omitempty"`
	Auth            AuthConfig            `yaml:"auth"`
	Deployments     DeploymentsConfig     `yaml:"deployments"`
	ProjectDefaults ProjectDefaultsConfig `yaml:"project_defaults"`
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
	DefaultAdmin   DefaultAdminConfig `yaml:"default_admin"`
	DefaultProject string             `yaml:"default_project"`
	AdminGroup     string             `yaml:"admin_group"`
}

type DefaultAdminConfig struct {
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Email     string `yaml:"email"`
	FirstName string `yaml:"first_name"`
	LastName  string `yaml:"last_name"`
}

type ProvidersConfig struct {
	AutoLoad *bool  `yaml:"auto_load,omitempty"`
	CacheDir string `yaml:"cache"`
}

type AuthConfig struct {
	JWTKey         string        `yaml:"jwt_key"`
	SessionTimeout time.Duration `yaml:"session_timeout"`
}

type DeploymentsConfig struct {
	AutoSuspendTime time.Duration `yaml:"auto_suspend_time"`
	LeaseTime       time.Duration `yaml:"lease_time"`
}

type ProjectDefaultsConfig struct {
	QuotaCPU     int `yaml:"quota_cpu"`
	QuotaRAM     int `yaml:"quota_ram"`
	QuotaDisk    int `yaml:"quota_disk"`
	QuotaNetwork int `yaml:"quota_network"`
	QuotaRouter  int `yaml:"quota_router"`
}
