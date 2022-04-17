package services

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
	//"github.com/go-rod/rod"
	//"github.com/mmcdole/gofeed"
	//"github.com/jtom38/newsbot/collector/domain/model"
)

var (

)

const (
	FFXIV_FEED_URL string = "https://na.finalfantasyxiv.com/lodestone/"
)

type FFXIVClient struct {
	SourceID uint
	Url string
}

func NewFFXIVClient() FFXIVClient {
	return FFXIVClient{}
}

//func (fc *FFXIVClient) CheckSource() error {

//}

func (fc *FFXIVClient) GetParser(uri string) (*goquery.Document, error) {
	html, err := http.Get(uri)
	if err != nil { return nil, err }
	defer html.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(html.Body)
	if err != nil { return nil, err }
	return doc, nil
}

func (fc *FFXIVClient) PullFeed() (error) {
	_, err := fc.GetParser(FFXIV_FEED_URL)
	if err != nil { return err }

	//parser.find
	return nil
}