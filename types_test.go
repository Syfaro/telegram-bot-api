package tgbotapi_test

import (
	"github.com/zhulik/telegram-bot-api"
	"testing"
	"time"
)

func TestUserStringWith(t *testing.T) {
	user := tgbotapi.User{0, "Test", "Test", ""}

	if user.String() != "Test Test" {
		t.Fail()
	}
}

func TestUserStringWithUserName(t *testing.T) {
	user := tgbotapi.User{0, "Test", "Test", "@test"}

	if user.String() != "@test" {
		t.Fail()
	}
}

func TestMessageIsGroup(t *testing.T) {
	from := tgbotapi.User{ID: 0}
	chat := tgbotapi.Chat{ID: 10}
	message := tgbotapi.Message{From: from, Chat: chat}

	if message.IsGroup() != true {
		t.Fail()
	}
}

func TestMessageTime(t *testing.T) {
	message := tgbotapi.Message{Date: 0}

	date := time.Unix(0, 0)
	if message.Time() != date {
		t.Fail()
	}
}

func TestChatIsPrivate(t *testing.T) {
	chat := tgbotapi.Chat{ID: 10, Type: "private"}

	if chat.IsPrivate() != true {
		t.Fail()
	}
}

func TestChatIsGroup(t *testing.T) {
	chat := tgbotapi.Chat{ID: 10, Type: "group"}

	if chat.IsGroup() != true {
		t.Fail()
	}
}

func TestChatIsChannel(t *testing.T) {
	chat := tgbotapi.Chat{ID: 10, Type: "channel"}

	if chat.IsChannel() != true {
		t.Fail()
	}
}

func TestFileLink(t *testing.T) {
	file := tgbotapi.File{FilePath: "test/test.txt"}

	if file.Link("token") != "https://api.telegram.org/file/bottoken/test/test.txt" {
		t.Fail()
	}
}