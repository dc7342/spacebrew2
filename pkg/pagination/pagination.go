package pagination

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type Keyboard struct {
	Before  []tgbotapi.InlineKeyboardMarkup
	Current tgbotapi.InlineKeyboardMarkup
	After   []tgbotapi.InlineKeyboardMarkup
}

type Symbol struct {
	First    string
	Previous string
	Next     string
	Last     string
	Current  string
}

var (
	symbol = Symbol{
		First:    "<< %s",
		Previous: "< %s",
		Next:     "%s >",
		Last:     "%s >>",
		Current:  "%s",
	}
)

type Paginator struct {
	keyboard Keyboard
	symbol   Symbol
	msg      *tgbotapi.MessageConfig
}

func NewPaginator(msg *tgbotapi.MessageConfig) Paginator {
	return Paginator{
		symbol: symbol,
		msg:    msg,
	}
}

func (p *Paginator) Keyboard(currentPage int, maxPage int) (*tgbotapi.MessageConfig, error) {
	var keys []tgbotapi.InlineKeyboardButton
	if currentPage > maxPage {
		return p.msg, fmt.Errorf("current page can't be grater than max page")
	}

	// Button for the first page.
	keys = append(keys, p.pageKey(1))

	if currentPage < 5 {
		for i := 2; i <= 5; i++ {
			keys = append(keys, p.pageKey(i))
		}

		p.msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(keys...))
		return p.msg, nil
	}

	// Button for the previous page.
	keys = append(keys, p.pageKey(currentPage-1))
	// Button for the current page.
	keys = append(keys, p.pageKey(currentPage))
	// Button for the next page.
	keys = append(keys, p.pageKey(currentPage+1))
	// Button for the last page.
	keys = append(keys, p.pageKey(maxPage))

	p.msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(keys...))

	return p.msg, nil
}

func (p *Paginator) pageKey(page int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(strconv.FormatInt(int64(page), 10),
		fmt.Sprintf("page_%d", page))
}
