package input_test

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/services/input"
)

var TwitchSourceRecord = database.Source {
	ID:     uuid.New(),
	Name:   "nintendo",
	Source: "Twitch",
}

var TwitchInvalidRecord = database.Source {
	ID:     uuid.New(),
	Name:   "EvilNintendo",
	Source: "Twitch",
}

func TestTwitchLogin(t *testing.T) {
	tc, err := input.NewTwitchClient()
	if err != nil {
		t.Error(err)
	}
	tc.ReplaceSourceRecord(TwitchSourceRecord)

	err = tc.Login()
	if err != nil {
		t.Error(err)
	}
}

// reach out and confirms that the API returns posts made by the user.
func TestTwitchReturnsUserPosts(t *testing.T) {
	tc, err := input.NewTwitchClient()
	if err != nil {
		t.Error(err)
	}
	tc.ReplaceSourceRecord(TwitchSourceRecord)

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

func TestTwitchReturnsNothingDueToInvalidUserName(t *testing.T) {
	tc, err := input.NewTwitchClient()
	if err != nil {
		t.Error(err)
	}
	tc.ReplaceSourceRecord(TwitchInvalidRecord)

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

func TestTwitchReturnsVideoAuthor(t *testing.T) {
	tc, err := input.NewTwitchClient()
	if err != nil {
		t.Error(err)
	}
	tc.ReplaceSourceRecord(TwitchSourceRecord)

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

func TestTwitchReturnsThumbnail(t *testing.T) {
	tc, err := input.NewTwitchClient()
	if err != nil {t.Error(err) }
	tc.ReplaceSourceRecord(TwitchSourceRecord)

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

func TestTwitchReturnsPubDate(t *testing.T) {
	tc, err := input.NewTwitchClient()
	if err != nil { t.Error(err) }
	tc.ReplaceSourceRecord(TwitchSourceRecord)

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

func TestTwitchReturnsDescription(t *testing.T) {
	tc, err := input.NewTwitchClient()
	if err != nil {
		t.Error(err)
	}
	tc.ReplaceSourceRecord(TwitchSourceRecord)

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

func TestTwitchReturnsAuthorImage(t *testing.T) {
	tc, err := input.NewTwitchClient()
	if err != nil {t.Error(err) }
	tc.ReplaceSourceRecord(TwitchSourceRecord)

	err = tc.Login()
	if err != nil { t.Error(err) }

	user, err := tc.GetUserDetails()
	if err != nil {t.Error(err) }
	
	_, err = tc.ExtractAuthorImage(user)
	if err != nil { t.Error(err) }
}

func TestTwitchReturnsTags(t *testing.T) {
	tc, err := input.NewTwitchClient()
	if err != nil {
		t.Error(err)
	}
	tc.ReplaceSourceRecord(TwitchSourceRecord)

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

func TestTwitchReturnsTitle(t *testing.T) {
	tc, err := input.NewTwitchClient()
	if err != nil {
		t.Error(err)
	}
	tc.ReplaceSourceRecord(TwitchSourceRecord)

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

func TestTwitchReturnsUrl(t *testing.T) {
	tc, err := input.NewTwitchClient()
	if err != nil { t.Error(err) }
	tc.ReplaceSourceRecord(TwitchSourceRecord)

	err = tc.Login()
	if err != nil { t.Error(err) }

	user, err := tc.GetUserDetails()
	if err != nil { t.Error(err) }

	posts, err := tc.GetPosts(user)
	if err != nil { t.Error(err) }

	res, err := tc.ExtractUrl(posts[0])
	if err != nil { t.Error(err) }
	if res == "" { t.Error("expected a filled string but got nil")}
}

func TestTwitchGetContent(t *testing.T) {
	tc, err := input.NewTwitchClient()
	if err != nil { t.Error(err) }
	tc.ReplaceSourceRecord(TwitchSourceRecord)

	err = tc.Login()
	if err != nil { t.Error(err) }
	
	posts, err := tc.GetContent()
	if err != nil {t.Error(err) }
	if len(posts) == 0 { t.Error("posts came back with 0 posts") }
	if len(posts) != 20 { t.Error("expected 20 posts") } 
}