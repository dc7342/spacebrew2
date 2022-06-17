package pagination

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Update struct {
	tgbotapi.Update
}

type Processor interface {
	Process(u *Update) bool
}

func (p *Paginator) Process(u *Update) bool {
	// Check if it has keyboard event

}
