package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (t *Telegram) initInlineKeyboard() {
	t.menu = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(t.conf.Button.AllTasks, stateShow),
			tgbotapi.NewInlineKeyboardButtonData(t.conf.Button.AddTask, stateNewName),
		),
	)

	t.edit = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(t.conf.Button.EditTitle, stateEditName),
			tgbotapi.NewInlineKeyboardButtonData(t.conf.Button.EditDescription, stateEditDescription),
		),
	)

	t.confirm = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(t.conf.Button.Yes, t.conf.Button.Yes),
			tgbotapi.NewInlineKeyboardButtonData(t.conf.Button.No, t.conf.Button.No),
		),
	)
}

func (t *Telegram) pagination(current int, maxPages int) {
	var k []tgbotapi.InlineKeyboardButton
	if current > 1 {
		k = append(k, tgbotapi.NewInlineKeyboardButtonData("←1", "1"))
	}
	if current > 2 {

	}
}
