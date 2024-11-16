package de

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gopkg.in/yaml.v2"
)

type Config struct {
    AppName  string `yaml:"app_name" env:"DOT_ENV_APP_NAME,required"`
    Database struct {
        Host     string `yaml:"host" env:"DOT_ENV_DB_HOST,required"`
        // Port     int    `yaml:"port" env:"DB_PORT,required"`
        // Username string `yaml:"username" env:"DB_USERNAME,required"`
        // Password string `yaml:"password" env:"DB_PASSWORD,required"`
    } `yaml:"database"`
}

func readConfigFromFile(filename string) (*Config, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("error reading YAML file: %w", err)
    }
    var config Config
    if err = yaml.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("error unmarshalling YAML data: %w", err)
    }
    return &config, nil
}

func readConfigFromEnv(config *Config) error {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    requiredVars := map[string]string{
        "APP_NAME":     os.Getenv("DOT_ENV_APP_NAME"),
        "DB_HOST":      os.Getenv("DOT_ENV_DB_HOST"),
    }

    for key, value := range requiredVars {
        if value == "" {
            return fmt.Errorf("required environment variable %s is missing", key)
        }
    }

    if appName := os.Getenv("DOT_ENV_APP_NAME"); appName != "" {
        config.AppName = appName
    }
    if dbHost := os.Getenv("DOT_ENV_DB_HOST"); dbHost != "" {
        config.Database.Host = dbHost
    }
    // if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
    //     fmt.Sscanf(dbPort, "%d", &config.Database.Port)
    // }
    // if dbUsername := os.Getenv("DB_USERNAME"); dbUsername != "" {
    //     config.Database.Username = dbUsername
    // }
    // if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
    //     config.Database.Password = dbPassword
    // }

    return nil
}

func LoadConfigDotEnv(path string) (*Config, error) {
    cfg, err := readConfigFromFile(path)
    if err != nil {
        log.Printf("Error reading config file: %v", err)
        return cfg, err
    }

    if err = readConfigFromEnv(cfg); err != nil {
        log.Printf("Error reading config from environment variables: %v", err)
    } else {   
        fmt.Printf("DOT ENV ---- Config: %+v\n\n", cfg)
    }

    return cfg, err
}
