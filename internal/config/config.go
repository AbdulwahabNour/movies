package config

import "log"

type Config struct{
    Port int
    Env string
    Version string
}

type  App struct{
    Config *Config
    InfoLog *log.Logger
    ErrorLog *log.Logger
 
}