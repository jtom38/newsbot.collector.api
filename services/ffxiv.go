package services

import (
	//"fmt"
	"errors"
	"log"
	"time"

	//"log"
	"net/http"

	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/jtom38/newsbot/collector/domain/model"
)

var (

)

const (
	FFXIV_NA_FEED_URL string = "https://na.finalfantasyxiv.com/lodestone/"
	FFXIV_JP_FEED_URL string = "https://jp.finalfantasyxiv.com/lodestone/"

	FFXIV_TIME_FORMAT string = "1/2/2006 3:4 PM"
)

type FFXIVClient struct {
	SourceID uint
	Url string
	Region string
}

func NewFFXIVClient(region string) FFXIVClient {
	var url string

	switch region {
		case "na": 
		url = FFXIV_NA_FEED_URL	
		case "jp":
		url = FFXIV_JP_FEED_URL
	}

	return FFXIVClient{
		Region: region,
		Url: url,
	}
}

func (fc *FFXIVClient) CheckSource() error {
	parser := fc.GetBrowser()
	defer parser.Close()
	//if err != nil { return err }

	links, err := fc.PullFeed(parser)
	if err != nil { return err }

	for _, link := range links {
		var article model.Articles

		page := fc.GetPage(parser, link)

		thumb, err := fc.ExtractThumbnail(page)
		if err != nil { return err }


		pubDate, err := fc.ExtractPubDate(page)
		if err != nil { return err }

		article.Thumbnail = thumb
		article.PubDate = pubDate

	}

	return nil
}

func (fc *FFXIVClient) GetParser() (*goquery.Document, error) {
	html, err := http.Get(fc.Url)
	if err != nil { return nil, err }
	defer html.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(html.Body)
	if err != nil { return nil, err }
	return doc, nil
}

func (fc *FFXIVClient) GetBrowser() (*rod.Browser) {
	browser := rod.New().MustConnect()
	return browser
}

func (fc *FFXIVClient) PullFeed(parser *rod.Browser) ([]string, error) {
	var links []string

	page := parser.MustPage(fc.Url)
	defer page.Close()

	// find the list by xpath
	res := page.MustElementX("/html/body/div[3]/div/div/div[1]/div[2]/div[1]/div[2]/ul")

	// find all the li items
	items := res.MustElements("li")
	
	for _, item := range items {
		// in each li, find the a items
		a, err := item.Element("a")
		if err != nil { 
			log.Println("Unable to find the a item, skipping")
			continue 
		}

		// find the href behind the a 
		url, err := a.Property("href")
		if err != nil { 
			log.Println("Unable to find a href link, skipping")
			continue 
		}

		urlString := url.String()
		isTopic := strings.Contains(urlString, "topics")
		if isTopic {
			links = append(links, urlString)	
		}
	}

	return links, nil
}

func (rc *FFXIVClient) GetPage(parser *rod.Browser, url string) *rod.Page {
	page := parser.MustPage(url)
	return page
}


func (fc *FFXIVClient) ExtractThumbnail(page *rod.Page) (string, error) {
	thumbnail := page.MustElementX("/html/body/div[3]/div[2]/div[1]/article/div[1]/img").MustProperty("src").String()
	if thumbnail == "" { return "", errors.New("unable to find thumbnail")}
	
	title := page.MustElement(".news__header > h1:nth-child(2)").MustText()
	log.Println(title)
	return thumbnail, nil
}

func (fc *FFXIVClient) ExtractPubDate(page *rod.Page) (time.Time, error) {
	stringDate := page.MustElement(".news__ic--topics").MustText()
	if stringDate == "" { return time.Now(), errors.New("unable to locate the publish date on the post")}

	PubDate, err := time.Parse(FFXIV_TIME_FORMAT, stringDate)
	if err != nil { return time.Now(), err }

	return PubDate, nil
}