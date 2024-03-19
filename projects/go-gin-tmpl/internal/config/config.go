package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Env           string
	ListenAddress string
	ListenPort    uint
	DB            struct {
		User         string
		Password     string
		Address      string
		Port         uint
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
	Limiter struct {
		RPS     float64
		Burst   int
		Enabled bool
	}
	SMTP struct {
		Host     string
		Port     int
		Username string
		Password string
		Sender   string
	}
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) Address() string {
	return fmt.Sprintf("%s:%d", cfg.ListenAddress, cfg.ListenPort)
}

func (cfg *Config) DBURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.DB.User, cfg.DB.Password, cfg.DB.Address, cfg.DB.Port)
}

func (cfg *Config) ParseFlags() error {
	flag.StringVar(&cfg.Env, "env", os.Getenv("ENV"), "Env type (production,development)")

	flag.StringVar(&cfg.ListenAddress, "address", os.Getenv("LISTEN_ADDRESS"), "API server address")
	listenPort, err := strconv.Atoi(os.Getenv("LISTEN_PORT"))
	if err != nil {
		return err
	}
	flag.UintVar(&cfg.ListenPort, "listen-port", uint(listenPort), "API server port")

	flag.StringVar(&cfg.DB.User, "db-user", os.Getenv("DB_USER"), "MongoDB User")
	flag.StringVar(&cfg.DB.Password, "db-password", os.Getenv("DB_PASSWORD"), "MongoDB Password")
	flag.StringVar(&cfg.DB.Address, "db-address", os.Getenv("DB_ADDRESS"), "MongoDB Address")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return err
	}
	flag.UintVar(&cfg.DB.Port, "db-port", uint(dbPort), "MongoDB Port")

	flag.IntVar(&cfg.DB.MaxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.DB.MaxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Float64Var(&cfg.Limiter.RPS, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.Limiter.Burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.Limiter.Enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.StringVar(&cfg.SMTP.Host, "smtp-host", os.Getenv("SMTP_HOST"), "SMTP host")
	flag.IntVar(&cfg.SMTP.Port, "smtp-port", 25, os.Getenv("SMTP_PORT"))
	flag.StringVar(&cfg.SMTP.Username, "smtp-username", os.Getenv("SMTP_USERNAME"), "SMTP username")
	flag.StringVar(&cfg.SMTP.Password, "smtp-password", os.Getenv("SMTP_PASSWORD"), "SMTP password")
	flag.StringVar(&cfg.SMTP.Sender, "smtp-sender", "openMovie <no-reply@test.user.net>", "SMTP sender")

	return nil
}
