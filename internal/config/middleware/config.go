package middleware

type Middleware struct {
	Name           string                 `yaml:"name"`
	When           string                 `yaml:"when"`
	Type           string                 `yaml:"type"`
	Method         string                 `yaml:"method"`
	URL            string                 `yaml:"url"`
	Body           map[string]interface{} `yaml:"body"`
	ResponsePolicy ResponsePolicy         `yaml:"response_policy"`
}

type ResponsePolicy struct {
	SuccessIf string `yaml:"success_if"`
	OnFailure string `yaml:"on_failure"`
}
