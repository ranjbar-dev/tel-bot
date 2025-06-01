package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

var bot *tele.Bot

func main() {

	// create channel to stop the bot
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {

		<-sigChan
		bot.Stop()
	}()

	// load env fron .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	// config telegram bot
	pref := tele.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	// create new bot
	bot, err = tele.NewBot(pref)
	if err != nil {

		panic(err)
	}

	// handle /start command
	bot.Handle("/start", func(c tele.Context) error {

		user, err := findUser(c.Sender().ID)
		if err != nil {

			if err != sql.ErrNoRows {

				return c.Send("Error while finding user")
			}

			user, err = insertUser(c.Sender().ID, c.Sender().Username)
			if err != nil {

				return c.Send("Error while inserting user")
			}
		}

		// register commands

		return c.Send(fmt.Sprintf("Hello, %d!", user.ChatID))
	})

	// start bot
	bot.Start()
}
