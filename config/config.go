package config

import (
    "log"
    "github.com/spf13/viper"
)

type Config struct {
    Server struct {
        Port string
    }
    Database struct {
        Host     string
        Port     int
        User     string
        Password string
        DBName   string
    }
    Kafka struct {
        Brokers []string
        Topic   string
    }
}

var AppConfig Config

func LoadConfig() {
    viper.SetConfigFile("config/config.yaml")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }
    if err := viper.Unmarshal(&AppConfig); err != nil {
        log.Fatalf("Error unmarshaling config: %v", err)
    }
}
