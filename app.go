package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	hd "github.com/olondar/bot/src/handled"
	"log"
	"os"
	"os/signal"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(hd.DefaultHandler),
	}

	var token = os.Getenv("TELEGRAM_BOT_TOKEN")

	b, err := bot.New(token, opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	hd.Exec(b)

	b.Start(ctx)
}
