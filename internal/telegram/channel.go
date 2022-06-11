package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/je09/spacebrew2/internal/entity"
)

func (t *Telegram) NewPost(task entity.Task) {
	msg := tgbotapi.NewMessageToChannel(t.conf.Telegram.ChannelUsername, t.services.Text(task))
	s, _ := t.bot.Send(msg)
	task.ID = int64(s.MessageID)
	if err := t.services.Post.New(task); err != nil {
		panic(err)
	}
}

func (t *Telegram) EditPost(task entity.Task) {
	// There is no method for editing channel messages like NewMessageChannel.
	// Temporary solution from #465 issue.
	// https://github.com/go-telegram-bot-api/telegram-bot-api/issues/465
	msg := tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChannelUsername: t.conf.Telegram.ChannelUsername,
			MessageID:       int(task.ID),
		},
		Text: t.services.Post.Text(task),
	}
	if err := t.services.Post.Edit(task); err != nil {
		panic(err)
	}

	t.bot.Send(msg)
}
