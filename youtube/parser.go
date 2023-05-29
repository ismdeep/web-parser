package youtube

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"

	"github.com/ismdeep/web-parser/pkg/httpdoc"
)

// ChannelInfo YouTube channel info struct
type ChannelInfo struct {
	ID         string
	RSSLink    string
	Title      string
	ChannelURL string // home page url with channel id
}

// GetChannelInfoByHomePageURL get channel info by home page url, e.g. https://www.youtube.com/@helloworld-cn
func GetChannelInfoByHomePageURL(url string) (*ChannelInfo, error) {
	doc, err := httpdoc.GetHTMLNode(url)
	if err != nil {
		return nil, err
	}

	var info ChannelInfo

	title := htmlquery.FindOne(doc, `//meta[@property="og:title"]/@content`)
	if title == nil {
		return nil, errors.New("og:title not found")
	}
	info.Title = htmlquery.InnerText(title)

	ogURL := htmlquery.FindOne(doc, `//meta[@property="og:url"]/@content`)
	if ogURL == nil {
		return nil, errors.New("og:url not found")
	}
	info.ChannelURL = htmlquery.InnerText(ogURL)
	info.ID = strings.ReplaceAll(info.ChannelURL, "https://www.youtube.com/channel/", "")

	rss := htmlquery.FindOne(doc, `//link[@title="RSS"]/@href`)
	if rss == nil {
		return nil, errors.New("rss not found")
	}
	info.RSSLink = htmlquery.InnerText(rss)

	return &info, nil
}

// VideoInfo YouTube video info struct
type VideoInfo struct {
	ID            string
	URL           string
	Title         string
	Description   string
	DatePublished string
	ChannelID     string
	LengthSeconds int
}

// GetVideoInfo get video info
func GetVideoInfo(videoID string) (*VideoInfo, error) {
	doc, err := httpdoc.GetHTMLNode(fmt.Sprintf("https://www.youtube.com/watch?v=%v", videoID))
	if err != nil {
		return nil, err
	}

	var info VideoInfo
	info.ID = videoID

	ogURL := htmlquery.FindOne(doc, `//meta[@property="og:url"]/@content`)
	if ogURL == nil {
		return nil, errors.New("og:url not found")
	}
	info.URL = htmlquery.InnerText(ogURL)

	title := htmlquery.FindOne(doc, `//meta[@name="title"]/@content`)
	if title == nil {
		return nil, errors.New("title not found")
	}
	info.Title = htmlquery.InnerText(title)

	description := htmlquery.FindOne(doc, `//meta[@property="og:description"]/@content`)
	if description == nil {
		return nil, errors.New("og:description not found")
	}
	info.Description = htmlquery.InnerText(description)

	datePublished := htmlquery.FindOne(doc, `//meta[@itemprop="datePublished"]/@content`)
	if datePublished == nil {
		return nil, errors.New("datePublished not found")
	}
	info.DatePublished = htmlquery.InnerText(datePublished)

	content := htmlquery.InnerText(doc)
	index := strings.Index(content, `,"subscribeEndpoint":{"channelIds":["`)
	if index < 0 {
		return nil, errors.New("channel id not found")
	}
	channelID := content[index+len(`,"subscribeEndpoint":{"channelIds":["`):]
	info.ChannelID = channelID[:strings.Index(channelID, `"`)]

	lengthSeconds := content[strings.Index(content, `},"lengthSeconds":"`)+len(`},"lengthSeconds":"`):]
	lengthSeconds = lengthSeconds[:strings.Index(lengthSeconds, `"`)]
	lengthSecondsInt, err := strconv.ParseInt(lengthSeconds, 10, 64)
	if err != nil {
		return nil, err
	}
	info.LengthSeconds = int(lengthSecondsInt)

	return &info, nil
}

// GetVideoIDListByHomePageURL get video id list by home page url
func GetVideoIDListByHomePageURL(url string) ([]string, error) {
	doc, err := httpdoc.GetHTMLNode(url)
	if err != nil {
		return nil, err
	}

	exists := make(map[string]bool)

	var lst []string
	content := htmlquery.InnerText(doc)
	for {
		index := strings.Index(content, `"videoId":"`)
		if index < 0 {
			break
		}

		content = content[index+len(`"videoId":"`):]
		videoID := content[:strings.Index(content, `"`)]
		if _, ok := exists[videoID]; !ok {
			lst = append(lst, videoID)
			exists[videoID] = true
		}

	}

	return lst, nil
}
