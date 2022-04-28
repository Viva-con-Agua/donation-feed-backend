package config

import "github.com/Viva-con-Agua/vcago"

type Config struct {
	// On which port this app will listen to
	AppPort int
}

func LoadFromEnv() Config {
	return Config{
		AppPort: vcago.Config.GetEnvInt("APP_PORT", "n", 8080),
	}
}
