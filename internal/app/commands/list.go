package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsg := "Here is all the products\n\n"
	products := c.productService.List()
	for _, p := range products {
		outputMsg += p.Title + "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
	_, _ = c.bot.Send(msg)
}
