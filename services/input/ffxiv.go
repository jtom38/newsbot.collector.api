package input

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/google/uuid"

	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/services/cache"
)

const (
	FFXIV_NA_FEED_URL string = "https://na.finalfantasyxiv.com/lodestone/"
	FFXIV_JP_FEED_URL string = "https://jp.finalfantasyxiv.com/lodestone/"

	FFXIV_TIME_FORMAT string = "1/2/2006 3:4 PM"
)

type FFXIVClient struct {
	record database.Source
	//SourceID uint
	//Url string
	//Region string

	cacheGroup string
}

func NewFFXIVClient(Record database.Source) FFXIVClient {
	return FFXIVClient{
		record: Record,
		cacheGroup: "ffxiv",
	}
}

func (fc *FFXIVClient) CheckSource() ([]database.Article, error) {
	var articles []database.Article

	parser := fc.GetBrowser()
	defer parser.Close()

	links, err := fc.PullFeed(parser)
	if err != nil { return articles, err }

	cache := cache.NewCacheClient(fc.cacheGroup)

	for _, link := range links {
		// Check cache/db if this link has been seen already, skip
		_, err := cache.FindByValue(link)
		if err == nil { continue }
		

		page := fc.GetPage(parser, link)

		title, err := fc.ExtractTitle(page)
		if err != nil { return articles, err }

		thumb, err := fc.ExtractThumbnail(page)
		if err != nil { return articles, err }

		pubDate, err := fc.ExtractPubDate(page)
		if err != nil { return articles, err }

		description, err := fc.ExtractDescription(page)
		if err != nil { return articles, err }

		authorName, err := fc.ExtractAuthor(page)
		if err != nil { return articles, err }

		authorImage, err := fc.ExtractAuthorImage(page)
		if err != nil { return articles, err }

		tags, err := fc.ExtractTags(page)
		if err != nil { return articles, err } 

		article := database.Article{
			Sourceid: fc.record.ID,
			Tags: tags,
			Title: title,
			Url: link,
			Pubdate: pubDate,
			Videoheight: 0,
			Videowidth: 0,
			Thumbnail: thumb,
			Description: description,
			Authorname: sql.NullString{String: authorName},
			Authorimage: sql.NullString{String: authorImage},
		}
		log.Printf("Collected '%v' from '%v'", article.Title, article.Url)

		cache.Insert(uuid.New().String(), link)

		articles = append(articles, article)
	}

	return articles, nil
}

func (fc *FFXIVClient) GetParser() (*goquery.Document, error) {
	html, err := http.Get(fc.record.Url)
	if err != nil { return nil, err }
	defer html.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(html.Body)
	if err != nil { return nil, err }
	return doc, nil
}

func (fc *FFXIVClient) GetBrowser() (*rod.Browser) {
	var browser *rod.Browser
	if path, exists := launcher.LookPath(); exists {
		u := launcher.New().Bin(path).MustLaunch()
		browser = rod.New().ControlURL(u).MustConnect()
	}
	return browser
}

func (fc *FFXIVClient) PullFeed(parser *rod.Browser) ([]string, error) {
	var links []string

	page := parser.MustPage(fc.record.Url)
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

func (fc *FFXIVClient) ExtractDescription(page *rod.Page) (string, error) {
	res := page.MustElement(".news__detail__wrapper").MustText()
	if res == "" { return "", errors.New("unable to locate the description on the post")}

	return res, nil
}

func (fc *FFXIVClient) ExtractAuthor(page *rod.Page) (string, error) {
	meta := page.MustElements("head > meta")
	for _, item := range meta {
		name, err := item.Property("name")
		if err != nil { return "", err }

		if name.String() != "author" { continue }
		content, err := item.Property("content")
		if err != nil { return "", err }
		
		return content.String(), nil
	}
	//log.Println(meta)
	return "", errors.New("unable to find the author on the page")
}

func (fc *FFXIVClient) ExtractTags(page *rod.Page) (string, error) {
	meta := page.MustElements("head > meta")
	for _, item := range meta {
		name, err := item.Property("name")
		if err != nil { return "", err }

		if name.String() != "keywords" { continue }
		content, err := item.Property("content")
		if err != nil { return "", err }
		
		return content.String(), nil
	}
	//log.Println(meta)
	return "", errors.New("unable to find the author on the page")
}

func (fc *FFXIVClient) ExtractTitle(page *rod.Page) (string, error) {
	title, err := page.MustElement("head > title").Text()
	if err != nil { return "", err }

	if !strings.Contains(title, "|") { return "", errors.New("unable to split the title, missing | in the string")}

	res := strings.Split(title, "|")
	if title != "" { return res[0], nil }
	
	//log.Println(meta)
	return "", errors.New("unable to find the author on the page")
}

func (fc *FFXIVClient) ExtractAuthorImage(page *rod.Page) (string, error) {
	meta := page.MustElements("head > link")
	for _, item := range meta {
		name, err := item.Property("rel")
		if err != nil { return "", err }

		if name.String() != "apple-touch-icon-precomposed" { continue }
		content, err := item.Property("href")
		if err != nil { return "", err }
		
		return content.String(), nil
	}
	//log.Println(meta)
	return "", errors.New("unable to find the author image on the page")
}

