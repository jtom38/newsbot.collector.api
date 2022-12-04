package config_test

import (
	"os"
	"testing"

	"github.com/jtom38/newsbot/collector/services/config"
)

func TestNewClient(t *testing.T) {
	config.New()
}

func TestGetConfigExpectNull(t *testing.T) {
	cc := config.New()
	os.Setenv(config.REDDIT_PULL_HOT, "")
	res := cc.GetConfig(config.REDDIT_PULL_HOT)
	if res != "" {
		panic("expected blank")
	}

}
