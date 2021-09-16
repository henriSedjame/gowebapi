package web

type ServerProperties struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

type CorsProperties struct {
	AllowedOrigins string `yaml:"allowed_origins"`
	AllowedHeaders string `yaml:"allowed_headers"`
	AllowedMethods string `yaml:"allowed_methods"`
}