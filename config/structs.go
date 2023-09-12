package config

type Config struct {
	Database       DatabaseConfig       `yaml:"database"`
	Initialization InitializationConfig `yaml:"initialization"`
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
}

type DefaultAdminConfig struct {
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Email     string `yaml:"email"`
	FirstName string `yaml:"first_name"`
	LastName  string `yaml:"last_name"`
}
