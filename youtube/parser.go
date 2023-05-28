package youtube

import (
	"errors"
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
