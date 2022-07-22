package main

import (
	"accounting_teko/pkg/repository"
	"accounting_teko/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	viper.SetConfigFile("configs/config.yml")
	viper.ReadInConfig()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	db, err := repository.NewPostgresDB(repository.Config{
		Password: os.Getenv("DB_PASSWORD"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)

	telegramBot := telegram.NewBot(bot, repos)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
