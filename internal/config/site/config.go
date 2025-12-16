package site

type Config struct {
	Servers      []string `json:"servers"`
	LoadBalancer string   `yaml:"load_balancer"`
	Routes       []Route  `yaml:"routes"`
}

type Route struct {
	Name   string      `yaml:"name"`
	Config RouteConfig `yaml:"config"`
}

type RouteConfig struct {
	Upstream   Upstream   `yaml:"upstream"`
	Downstream Downstream `yaml:"downstream"`
}

type Downstream struct {
	Schema string `yaml:"schema"`
	Route  string `yaml:"route"`
}

type Upstream struct {
	Route  string `yaml:"route"`
	Method string `yaml:"method"`
}
