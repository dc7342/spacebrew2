package telegram

import (
	tm "github.com/and3rson/telemux/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/je09/spacebrew2/internal/entity"
	"strconv"
	"strings"
)

const (
	// States
	// A user is being asked about the name of the task
	stateNewName = "new_name"
	// A user is being asked about task description for the task
	stateNewDescription = "new_description"
	// A user is being asked of which part of the task he wants to edit
	stateEditChoice = "edit_choice"
	// A user wants to edit task name
	stateEditName = "edit_name"
	// A user wants to edit task description
	stateEditDescription = "edit_description"
	// A user wants to see all tasks
	stateShow = "show_task"
	// A user wants to close the task
	stateClose = "close_task"
	// A user is being asked for confirmation
	stateConfirm = "confirm"
	// Default menu state
	stateDefault = "menu"

	// Context data
	taskID        = "task_id"
	confirmAction = "confirm_action"
	closeAction   = "close_action"
)

// HandlerFunction for /start command
func (t *Telegram) startCmd(u *tm.Update) bool {
	answer := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.WelcomeMessage)
	_, err := t.bot.Send(answer)
	if err != nil {
		return false
	}

	return true
}

// handles new task action by inline button
func (t *Telegram) newTask(u *tm.Update) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.NewTaskName)
	t.bot.Send(msg)
	u.PersistenceContext.SetState(stateNewName)
}

// handles a naming stage in process of creating a new task
func (t *Telegram) newName(u *tm.Update) {
	// Typing status
	t.bot.Send(tgbotapi.NewChatAction(u.Message.Chat.ID, tgbotapi.ChatTyping))
	// Get data from previous step
	data := u.PersistenceContext.GetData()
	// Send name to the context
	data[stateNewName] = u.Message.Text
	u.PersistenceContext.SetData(data)

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.NewDescriptionName)
	t.bot.Send(msg)
	u.PersistenceContext.SetState(stateNewDescription)
}

// handles a descriptive stage in process of creating a new task
func (t *Telegram) newDescription(u *tm.Update) {
	// Typing status
	t.bot.Send(tgbotapi.NewChatAction(u.Message.Chat.ID, tgbotapi.ChatTyping))
	// Get data from previous step
	data := u.PersistenceContext.GetData()

	task := entity.Task{
		Title:       data[stateNewName].(string),
		Description: u.Message.Text,
		Open:        true,
		AuthorID:    u.Message.Chat.ID,
	}
	t.NewPost(task)

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.NewTaskDone)
	t.bot.Send(msg)
	// remembers description ads it to db
	u.PersistenceContext.SetState(stateDefault)
}

// asks user what is they want to edit
func (t *Telegram) editTask(u *tm.Update) {
	// Get data from previous step
	data := u.PersistenceContext.GetData()
	// Send name to the context
	idStr := strings.Split(u.Message.Text, "_")[1]
	id64, _ := strconv.ParseInt(idStr, 10, 64)
	data[taskID] = id64
	u.PersistenceContext.SetData(data)

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.EditTaskChoice)
	t.bot.Send(msg)
	u.PersistenceContext.SetState(stateEditChoice)
}

// handles what user wants to edit
func (t *Telegram) editChoice(u *tm.Update) {
	t.bot.Send(tgbotapi.NewChatAction(u.Message.Chat.ID, tgbotapi.ChatTyping))
	var msg tgbotapi.MessageConfig
	switch u.Message.Text {
	case t.conf.Button.EditTitle:
		u.PersistenceContext.SetState(stateEditName)
		msg = tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.EditName)
	case t.conf.Button.EditDescription:
		u.PersistenceContext.SetState(stateEditDescription)
		msg = tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.EditDescription)
	default:
		u.PersistenceContext.SetState(stateDefault)
		msg = tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.UnknownMessage)
	}

	t.bot.Send(msg)
}

func (t *Telegram) editTitle(u *tm.Update) {
	t.bot.Send(tgbotapi.NewChatAction(u.Message.Chat.ID, tgbotapi.ChatTyping))
	// Get data from previous step
	data := u.PersistenceContext.GetData()
	u.PersistenceContext.SetState(stateDefault)

	t.services.EditTitle(data[taskID].(int64), u.Message.Text)
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.Done)
	t.bot.Send(msg)
}

func (t *Telegram) editDescription(u *tm.Update) {
	t.bot.Send(tgbotapi.NewChatAction(u.Message.Chat.ID, tgbotapi.ChatTyping))
	// Get data from previous step
	data := u.PersistenceContext.GetData()
	u.PersistenceContext.SetState(stateDefault)

	t.services.EditDescription(data[taskID].(int64), u.Message.Text)
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.Done)
	t.bot.Send(msg)
}

func (t *Telegram) showTasks(u *tm.Update) {
	c, _ := t.services.Post.Pages(t.conf.Page.TasksPerPage)

}

func (t *Telegram) closeTask(u *tm.Update) {
	t.bot.Send(tgbotapi.NewChatAction(u.Message.Chat.ID, tgbotapi.ChatTyping))
	data := u.PersistenceContext.GetData()
	data[confirmAction] = closeAction

	idStr := strings.Split(u.Message.Text, "_")[1]
	id64, _ := strconv.ParseInt(idStr, 10, 64)
	data[taskID] = id64

	u.PersistenceContext.SetData(data)

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.Confirmation)
	t.bot.Send(msg)
}

func (t *Telegram) confirmChoice(u *tm.Update) {
	t.bot.Send(tgbotapi.NewChatAction(u.Message.Chat.ID, tgbotapi.ChatTyping))
	// Get data from previous step
	data := u.PersistenceContext.GetData()

	if u.Message.Text == t.conf.Button.No {
		u.PersistenceContext.SetState(stateDefault)
		msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.ConfirmationNegative)
		t.bot.Send(msg)
		return
	}

	if u.Message.Text != t.conf.Button.Yes {
		u.PersistenceContext.SetState(stateDefault)
		msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.UnknownMessage)
		t.bot.Send(msg)
		return
	}

	// There may be more than one condition
	switch data[confirmAction].(string) {
	case closeAction:
		t.services.Close(data[taskID].(int64))
	}

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.Done)
	t.bot.Send(msg)
}

func (t *Telegram) unknownCmd(u *tm.Update) {
	answer := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.UnknownMessage)
	t.bot.Send(answer)
}
