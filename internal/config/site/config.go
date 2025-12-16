package site

import "gateway/internal/config/middleware"

type Config struct {
	Servers      []string `json:"servers"`
	LoadBalancer string   `yaml:"load_balancer"`
	Routes       []Route  `yaml:"routes"`
}

type Route struct {
	Name        string                  `yaml:"name"`
	Upstream    Upstream                `yaml:"upstream"`
	Downstream  Downstream              `yaml:"downstream"`
	Middlewares []middleware.Middleware `yaml:"middlewares"`
}

type Downstream struct {
	Schema string `yaml:"schema"`
	Route  string `yaml:"route"`
}

type Upstream struct {
	Route  string `yaml:"route"`
	Method string `yaml:"method"`
}
