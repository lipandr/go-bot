package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("incorrect argument")
		return
	}
	product, err := c.productService.GetProduct(arg)
	var text string
	if err != nil {
		text = err.Error()
	} else {
		text = fmt.Sprintf("Product idx %v Title: %v", arg, product.Title)
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	_, _ = c.bot.Send(msg)
}
