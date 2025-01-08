package vp

import (
	"fmt"
	"log"
	"strings"
	// "os"

	"github.com/spf13/viper"
)

type Config struct {
    AppName  string `mapstructure:"app_name"`
    Database struct {
        Host     string `mapstructure:"host"`
    } `mapstructure:"database"`
}

func loadConfig(path string) (*Config, error) {
    viper.SetConfigFile(path)

	viper.SetEnvPrefix("VIPER")
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    viper.AutomaticEnv()

    // requiredEnvVars := []string{
    //     "VIPER_APP_NAME",
    //     "VIPER_DATABASE_HOST",
    // }

    // for _, envVar := range requiredEnvVars {
    //     viper.BindEnv(envVar)
	// 	fmt.Println("---->", os.Getenv(envVar))
    //     if !viper.IsSet(envVar) {
    //         return nil, fmt.Errorf("environment variable %s is required but not set", envVar)
    //     }
    // }

    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("error reading config file: %w", err)
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, fmt.Errorf("unable to decode into struct: %w", err)
    }

    return &config, nil
}

func LoadConfigViper(path string) {
    config, err := loadConfig(path)
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    fmt.Printf("Config: %+v\n", config)
}
