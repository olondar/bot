package handled

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

func Exec(b *bot.Bot) {
	haldlerList := map[string]bot.HandlerFunc{
		"/hello": helloHandler,
		"/phone": phoneHandler,
	}

	for key, function := range haldlerList {
		b.RegisterHandler(bot.HandlerTypeMessageText, key, bot.MatchTypeExact, function)
	}

	//b.RegisterHandler(bot.HandlerTypeMessageText, "/hello", bot.MatchTypeExact, helloHandler)
	//b.RegisterHandler(bot.HandlerTypeMessageText, "/phone", bot.MatchTypeExact, phoneHandler)
}

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	text := update.Message.Text
	phone := update.Message.Contact.PhoneNumber

	if text != "" {
		log.Println(text)
	} else if phone != "" {
		log.Println(phone)
	}
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "/hello - sey\n" +
			"/phone - send phone",
	})

	if err != nil {
		return
	}
}
func helloHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Printf("%s %s\n", update.Message.From.FirstName, update.Message.From.LastName)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Hello, *" + bot.EscapeMarkdown(update.Message.From.FirstName) + "*",
		ParseMode: models.ParseModeMarkdown,
	})

	if err != nil {
		return
	}
}

func phoneHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Println(update.Message.Text)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Для идентификации вас в нашей базе нам нужен вашь номер телефона.",
		ReplyMarkup: &models.ReplyKeyboardMarkup{
			Keyboard: [][]models.KeyboardButton{
				{
					{
						Text:           "ОТПРАВИТЬ ТЕЛЕФОН",
						RequestContact: true,
					},
				},
			},
			ResizeKeyboard:  true,
			OneTimeKeyboard: true,
			//Selective:       false,
		},
	})

	if err != nil {
		return
	}
}
