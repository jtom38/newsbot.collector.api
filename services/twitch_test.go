package services_test

import (
	"log"
	"testing"

	"github.com/jtom38/newsbot/collector/domain/model"
	"github.com/jtom38/newsbot/collector/services"
)

var sourceRecord = model.Sources{
	ID:     1,
	Name:   "nintendo",
	Source: "Twitch",
}

var invalidRecord = model.Sources{
	ID:     1,
	Name:   "EvilNintendo",
	Source: "Twitch",
}

func TestTwitchLogin(t *testing.T) {
	tc, err := services.NewTwitchClient(sourceRecord)
	if err != nil {
		t.Error(err)
	}

	err = tc.Login()
	if err != nil {
		t.Error(err)
	}
}

// reach out and confirms that the API returns posts made by the user.
func TestReturnsUserPosts(t *testing.T) {
	tc, err := services.NewTwitchClient(sourceRecord)
	if err != nil {
		t.Error(err)
	}

	err = tc.Login()
	if err != nil {
		t.Error(err)
	}

	user, err := tc.GetUserDetails()
	if err != nil {
		t.Error(err)
	}

	posts, err := tc.GetPosts(user)
	if err != nil {
		t.Error(err)
	}
	if len(posts) == 0 {
		t.Error("expected videos but got none")
	}
}

func TestReturnsNothingDueToInvalidUserName(t *testing.T) {
	tc, err := services.NewTwitchClient(invalidRecord)
	if err != nil {
		t.Error(err)
	}

	err = tc.Login()
	if err != nil {
		t.Error(err)
	}

	user, err := tc.GetUserDetails()
	if err != nil {
		t.Error(err)
	}

	posts, err := tc.GetPosts(user)
	if err != nil {
		t.Error(err)
	}
	if len(posts) != 0 {
		t.Error("expected videos but got none")
	}
}

func TestReturnsVideoAuthor(t *testing.T) {
	tc, err := services.NewTwitchClient(sourceRecord)
	if err != nil {
		t.Error(err)
	}

	err = tc.Login()
	if err != nil {
		t.Error(err)
	}

	user, err := tc.GetUserDetails()
	if err != nil {
		t.Error(err)
	}

	posts, err := tc.GetPosts(user)
	if err != nil {
		t.Error(err)
	}
	if posts[0].UserName == "" {
		t.Error("uable to parse username")
	}
}

func TestReturnsThumbnail(t *testing.T) {
	tc, err := services.NewTwitchClient(sourceRecord)
	if err != nil {t.Error(err) }

	err = tc.Login()
	if err != nil { t.Error(err) }

	user, err := tc.GetUserDetails()
	if err != nil { t.Error(err) }

	posts, err := tc.GetPosts(user)
	if err != nil { t.Error(err) }

	value, err := tc.ExtractThumbnail(posts[0])
	if err != nil { t.Error(err) }
	if value == "" { t.Error("uable to parse username") }
}

func TestReturnsPubDate(t *testing.T) {
	tc, err := services.NewTwitchClient(sourceRecord)
	if err != nil { t.Error(err) }

	err = tc.Login()
	if err != nil { t.Error(err) }

	user, err := tc.GetUserDetails()
	if err != nil { t.Error(err) }

	posts, err := tc.GetPosts(user)
	if err != nil { t.Error(err) }

	date, err := tc.ExtractPubDate(posts[0])
	log.Println(date)
	if err != nil { t.Error(err) }
}

func TestReturnsDescription(t *testing.T) {
	tc, err := services.NewTwitchClient(sourceRecord)
	if err != nil {
		t.Error(err)
	}

	err = tc.Login()
	if err != nil {
		t.Error(err)
	}

	user, err := tc.GetUserDetails()
	if err != nil {
		t.Error(err)
	}

	posts, err := tc.GetPosts(user)
	if err != nil {
		t.Error(err)
	}

	_, err = tc.ExtractDescription(posts[0])
	if err != nil {
		t.Error(err)
	}
}

func TestReturnsAuthorImage(t *testing.T) {
	tc, err := services.NewTwitchClient(sourceRecord)
	if err != nil {t.Error(err) }

	err = tc.Login()
	if err != nil { t.Error(err) }

	user, err := tc.GetUserDetails()
	if err != nil {t.Error(err) }
	
	_, err = tc.ExtractAuthorImage(user)
	if err != nil { t.Error(err) }
}

func TestReturnsTags(t *testing.T) {

	tc, err := services.NewTwitchClient(sourceRecord)
	if err != nil {
		t.Error(err)
	}

	err = tc.Login()
	if err != nil {
		t.Error(err)
	}

	user, err := tc.GetUserDetails()
	if err != nil {
		t.Error(err)
	}

	posts, err := tc.GetPosts(user)
	if err != nil { t.Error(err) }

	_, err = tc.ExtractTags(posts[0], user)
	if err != nil { t.Error(err) }
}

func TestReturnsTitle(t *testing.T) {
	tc, err := services.NewTwitchClient(sourceRecord)
	if err != nil {
		t.Error(err)
	}

	err = tc.Login()
	if err != nil {
		t.Error(err)
	}

	user, err := tc.GetUserDetails()
	if err != nil {
		t.Error(err)
	}

	posts, err := tc.GetPosts(user)
	if err != nil { t.Error(err) }

	res, err := tc.ExtractTitle(posts[0])
	if err != nil { t.Error(err) }
	if res == "" { t.Error("expected a filled string but got nil")}
}
