package config

type Config struct {
	SECRET string
}

func GetConfig() *Config {
	return &Config{
		SECRET: "!ebbdg^8n25lfdREACH@#$%!)<vdyd;-d6cb",
	}
}