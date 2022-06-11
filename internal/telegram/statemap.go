package telegram

import (
	tm "github.com/and3rson/telemux/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) initMux() {
	stateMap := tm.StateMap{
		stateDefault: {
			tm.NewHandler(tm.IsCommandMessage(stateNewName), t.newTask),
			tm.NewHandler(isEditCommand(), t.editTask),
			tm.NewHandler(isCloseCommand(), t.closeTask),
		},
		stateNewName: {
			tm.NewHandler(tm.HasText(), t.newName),
		},
		stateNewDescription: {
			tm.NewHandler(tm.HasText(), t.newDescription),
		},
		stateEditChoice: {
			tm.NewHandler(tm.HasText(), t.editChoice),
		},
		stateEditName: {
			tm.NewHandler(tm.HasText(), t.editTitle),
		},
		stateEditDescription: {
			tm.NewHandler(tm.HasText(), t.editDescription),
		},
		stateShow: {
			tm.NewHandler(tm.IsCallbackQuery(), t.showTasks),
		},
		stateClose: {
			tm.NewHandler(tm.IsCallbackQuery(), t.closeTask),
		},
		stateConfirm: {
			tm.NewHandler(tm.IsCallbackQuery(), t.confirmChoice),
		},
	}

	def := []*tm.Handler{
		tm.NewHandler(tm.IsCommandMessage("cancel"), func(u *tm.Update) {
			u.PersistenceContext.ClearData()
			u.PersistenceContext.SetState(stateDefault)
			t.bot.Send(tgbotapi.NewMessage(u.Message.Chat.ID, "Cancelled."))
		}),
	}

	cmdHandler := []*tm.Handler{
		tm.NewHandler(tm.And(tm.IsPrivate(), tm.IsCommandMessage("start"), t.startCmd)),
	}

	t.cnvs = tm.NewMux().AddHandler(tm.NewConversationHandler("menu", tm.NewLocalPersistence(), stateMap, def))
	t.cmds = tm.NewMux().AddHandler(cmdHandler...)
}
