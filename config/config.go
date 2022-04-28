package config

import "github.com/Viva-con-Agua/vcago"

type Config struct {
	// Which origins are allowed in CORS requests
	AllowOrigins []string

	// On which port this app will listen to
	AppPort int
}

func LoadFromEnv() Config {
	return Config{
		AllowOrigins: vcago.Config.GetEnvStringList("APP_ALLOW_ORIGINS", "e", []string{"*"}),
		AppPort:      vcago.Config.GetEnvInt("APP_PORT", "n", 8080),
	}
}
