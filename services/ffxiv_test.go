package services_test

import (
	"testing"

	"github.com/jtom38/newsbot/collector/services"
)

func TestGetParser(t *testing.T) {
	fc := services.NewFFXIVClient()
	fc.GetParser(services.FFXIV_FEED_URL)
}
