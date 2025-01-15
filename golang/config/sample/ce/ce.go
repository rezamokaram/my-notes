package ce

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App APP `env-required:"true" json:"app"`
		// Postgres POSTGRES `env-required:"true" json:"postgres"`
	}

	APP struct {
		Name    string `env-required:"false" json:"name" env:"CLEANENV_APP_NAME"`
		Version string `env-required:"false" json:"version" env:"CLEANENV_APP_VERSION"`
		// Host    string `env-required:"true" json:"host" env:"INSPECTION_APP_HOST"`
		// Port    string `env-required:"true" json:"port" env:"INSPECTION_APP_PORT"`
	}

	// POSTGRES struct {
	// 	DB       string `env-required:"true" json:"db" env:"POSTGRES_DB"`
	// 	User     string `env-required:"true" json:"user" env:"POSTGRES_USER"`
	// 	Password string `env-required:"true" json:"password" env:"POSTGRES_PASSWORD"`
	// 	Host     string `env-required:"true" json:"host" env:"POSTGRES_HOST"`
	// 	Port     string `env-required:"true" json:"port" env:"POSTGRES_PORT"`
	// 	SSLMode  string `env-required:"true" json:"ssl_mode" env:"POSTGRES_SSLMODE"`
	// 	Timezone string `env-required:"true" json:"timezone" env:"POSTGRES_TIMEZONE"`
	// }
)

func LoadConfig(path string) (*Config, error) {
	config := new(Config)
	err := cleanenv.ReadConfig(path, config)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func LoadConfigCleanEnv(path string) (*Config, error) {
	cfg, err := LoadConfig(path)
	if err != nil {
		log.Printf("Clean env failed: %v", err)
	} else {
		log.Printf("clean env: %v", cfg)
	}
	return cfg, err
}
