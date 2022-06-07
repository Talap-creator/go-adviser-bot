package telegram

import "go-adviser-bot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// offset
}

func New(client *telegram.Client, storage) {
	
}
