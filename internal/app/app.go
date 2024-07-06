package app

import "finance_manager/internal/config"

func Run(configPath string) {

	cfg, err := config.NewConfig(configPath)
}
