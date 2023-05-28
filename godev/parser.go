package godev

import (
	"fmt"

	"github.com/antchfx/htmlquery"

	"github.com/ismdeep/web-parser/pkg/httpdoc"
)

// GetDownloadLinks get download links
func GetDownloadLinks() ([]string, error) {
	doc, err := httpdoc.GetHTMLNode("https://go.dev/dl/")
	if err != nil {
		return nil, err
	}

	var results []string
	lst := htmlquery.Find(doc, `//table[@class="downloadtable"]//tr/td/a[@class="download"]/@href`)
	for _, node := range lst {
		link := htmlquery.InnerText(node)
		results = append(results, fmt.Sprintf("https://go.dev%v", link))
	}

	return results, nil
}

// GetStableVersions get stable versions
func GetStableVersions() ([]string, error) {
	doc, err := httpdoc.GetHTMLNode("https://go.dev/dl/")
	if err != nil {
		return nil, err
	}

	var results []string
	lst := htmlquery.Find(doc, `//div[@class="toggleVisible"]/@id`)
	for _, node := range lst {
		results = append(results,
			htmlquery.InnerText(node))
	}

	return results, nil
}
