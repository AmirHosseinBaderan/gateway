package base

type Config struct {
	App            AppConfig    `yaml:"app"`
	Server         ServerConfig `yaml:"server"`
	MiddlewarePath string       `yaml:"middleware_path"`
}

type AppConfig struct {
	Name     string   `yaml:"name"`
	Env      string   `yaml:"env"`
	Upstream Upstream `yaml:"upstream"`
}

type ServerConfig struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	ReadTimeoutMs  int    `yaml:"read_timeout_ms"`
	WriteTimeoutMs int    `yaml:"write_timeout_ms"`
}

type Upstream struct {
	ConfigPath string   `yaml:"config_path"`
	Servers    []string `yaml:"servers"`
}
