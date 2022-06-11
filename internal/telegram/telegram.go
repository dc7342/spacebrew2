package telegram

import (
	tm "github.com/and3rson/telemux/v2"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/je09/spacebrew2/internal/service"
	"github.com/je09/spacebrew2/pkg/config"
)

type Telegram struct {
	bot      *tgbotapi.BotAPI
	services *service.Service
	conf     config.Config

	cmds *tm.Mux
	cnvs *tm.Mux

	menu    tgbotapi.InlineKeyboardMarkup
	edit    tgbotapi.InlineKeyboardMarkup
	confirm tgbotapi.InlineKeyboardMarkup
}

func NewTelegram(serv *service.Service, conf config.Config) *Telegram {
	return &Telegram{
		services: serv,
		conf:     conf,
	}
}

func (t *Telegram) Start() error {
	bot, err := tgbotapi.NewBotAPI(t.conf.Telegram.Token)
	if err != nil {
		return err
	}
	t.bot = bot

	u := tgbotapi.NewUpdate(0)
	u.Timeout = t.conf.Telegram.Timeout
	updates := t.bot.GetUpdatesChan(u)

	t.initMux()
	t.initInlineKeyboard()

	for update := range updates {
		t.cmds.Dispatch(bot, update)
		t.cnvs.Dispatch(bot, update)
	}

	return nil
}
