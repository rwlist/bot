package conf

import (
	"github.com/caarlos0/env/v6"
)

type App struct {
	PrometheusBind string `env:"PROMETHEUS_BIND" envDefault:":2112"`

	// PostgresDSN is a DSN for the postgres.
	PostgresDSN string `env:"POSTGRES_DSN,required"`

	// Telegram bot token.
	TelegramToken string `env:"TELEGRAM_TOKEN,required"`
}

func ParseEnv() (*App, error) {
	cfg := App{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
