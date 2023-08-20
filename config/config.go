package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Mail     Mail
	Logger   LoggerConfig
	Cookie   Cookie
	Limiter  Limiter
}

type ServerConfig struct {
	AppVersion               string
	AppName                  string
	AppHost                  string
	Port                     string
	Mode                     string
	JwtSecretKey             string
	PrivateKeyToken          string
	PublicKeyToken           string
	IDExpiration             time.Duration
	RefreshExpiratio         time.Duration
	ActivationTokenExpiratio time.Duration
	PepperSecreKey           string
	ReadTimeout              time.Duration
	WriteTimeout             time.Duration
	CtxDefaultTimeout        time.Duration
	CSRF                     bool
}

type LoggerConfig struct {
	Development bool
	Formate     string
	Level       string
}

type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	PgDriver           string
}
type RedisConfig struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultdb string
	MinIdleConns   int
	PoolSize       int
	PoolTimeout    time.Duration
	Password       string
	DB             int
}

type Mail struct {
	Host     string
	Port     int
	UserName string
	Password string
	Sender   string
	TimeOut  time.Duration
}

type Cookie struct {
	Name     string
	MaxAge   int
	Secure   bool
	HTTPOnly bool
}
type Limiter struct {
	Rps     float32
	Burst   int
	Enabled bool
}

func LoadConfig(fileName string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf(fmt.Sprintf("config file %s not found", fileName))
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
