package output_test

import (
	"os"
	"strings"
	"testing"
	//"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/services/output"
)

var (
	article database.Article = database.Article{
		ID: uuid.New(),
		Sourceid: uuid.New(),
		Tags: "unit, testing",
		Title: "Demo",
		Url: "https://github.com/jtom38/newsbot.collector.api",
		//Pubdate: time.Now(),
		Videoheight: 0,
		Videowidth: 0,
		Description: "Hello World",
	}
	blank string = ""
)

func TestDiscordMessageContainsTitle(t *testing.T) {
	d := output.NewDiscordWebHookMessage(article)
	msg, err := d.GeneratePayload()
	if err != nil {
		t.Error(err)
	}
	
	for _, i := range *msg.Embeds {
		if i.Title == &blank {
			t.Error("title missing")
		}
	}
}

func TestDiscordMessageContainsDescription(t *testing.T) {
	d := output.NewDiscordWebHookMessage(article)
	msg, err := d.GeneratePayload()
	if err != nil {
		t.Error(err)
	}
	
	for _, i := range *msg.Embeds {
		if i.Description == &blank {
			t.Error("description missing")
		}
	}
}

func TestDiscordMessageFooter(t *testing.T) {
	d := output.NewDiscordWebHookMessage(article)
	msg, err := d.GeneratePayload()
	if err != nil {
		t.Error(err)
	}
	for _, i := range *msg.Embeds {
		blank := ""
		if i.Footer.Value == &blank {
			t.Error("missing footer vlue")
		}
		if i.Footer.IconUrl == &blank {
			t.Error("missing footer url")
		} 
	}
}

func TestDiscordMessageFields(t *testing.T) {
	header := "Link"
	d := output.NewDiscordWebHookMessage(article)
	msg, err := d.GeneratePayload()
	if err != nil {
		t.Error(err)
	}
	for _, embed := range *msg.Embeds {
		for _, field := range embed.Fields {
			var fName string 
			if field.Name != nil {
				fName = *field.Name
			} else {
				t.Error("missing link field value")
			}

			if fName != header {
				t.Error("missing link field key")
			}

			var fValue string
			if field.Value != nil {
				fValue = *field.Value
			}

			if fValue == blank {
				t.Error("missing link field value")
			}
		}
	}
}

// This test requires a env value to be present to work
func TestDiscordMessagePost(t *testing.T) {
	_, err := os.Open(".env")
	if err != nil {
		t.Error(err)
	}
		
	err = godotenv.Load()
	if err != nil {
		t.Error(err)
	}

	res := os.Getenv("TESTS_DISCORD_WEBHOOK")
	if res == "" {
		t.Error("TESTS_DISCORD_WEBHOOK is missing")
	}
	endpoints := strings.Split(res, " ")
	if err != nil {
		t.Error(err)
	}

	d := output.NewDiscordWebHookMessage(article)
	msg, err := d.GeneratePayload()
	if err != nil {
		t.Error(err)
	}

 	err = d.SendPayload(msg, endpoints[0])
	if err != nil {
		t.Error(err)
	}
}