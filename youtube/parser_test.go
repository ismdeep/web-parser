package youtube

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChannelInfoByHomePageURL(t *testing.T) {
	info, err := GetChannelInfoByHomePageURL("https://www.youtube.com/@helloworld-cn")
	assert.NoError(t, err)
	t.Logf("ChannelID = %v", info.ID)
	t.Logf("Title = %v", info.Title)
	t.Logf("ChannelURL = %v", info.ChannelURL)
	t.Logf("RSSLink = %v", info.RSSLink)
}
