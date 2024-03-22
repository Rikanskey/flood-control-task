package runner

import (
	"log"
	"task/internal/config"
	"task/internal/database/repository"
	"task/internal/database/repository/db"
	"task/internal/domain"
)

func Start(configDir string) {
	cfg := newConfig(configDir)
	ramDb := initDb(cfg)
	if ramDb == nil {
		return
	}
	application := newApplication(cfg.App, ramDb)
	startApp(application)
}

func newConfig(configDir string) *config.Config {
	cfg, err := config.New(configDir)
	if err != nil {
		log.Panicln(err)
	}

	return cfg
}

func initDb(cfg *config.Config) repository.RamRepository {
	redis, err := db.NewRedisRepository(cfg.Redis)
	if err != nil {
		log.Print(err)
		return nil
	}
	return redis
}

func newApplication(cfg config.AppConfig, db repository.RamRepository) domain.FloodControl {
	return domain.NewFloodController(cfg, db)
}

// Заглушка, моделирующая работу контроля флуда в контексте веб-приложения
func startApp(app domain.FloodControl) {
	for true {
	}
}
