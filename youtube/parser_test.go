package youtube

import (
	"testing"
	"time"

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

func TestGetVideoInfo(t *testing.T) {
	info, err := GetVideoInfo("NP08uUsHFkA")
	assert.NoError(t, err)
	if info != nil {
		t.Logf("ID = %v", info.ID)
		t.Logf("URL = %v", info.URL)
		t.Logf("Title = %v", info.Title)
		t.Logf("Description = %v", info.Description)
		loc, err := time.LoadLocation("Asia/Shanghai")
		assert.NoError(t, err)
		t.Logf("DatePublished = %v", info.DatePublished.In(loc).Format(time.RFC3339))
		t.Logf("ChannelID = %v", info.ChannelID)
		t.Logf("LengthSeconds = %v", info.LengthSeconds)
	}
}

func TestGetVideoIDListByHomePageURL(t *testing.T) {
	videoIDLst, err := GetVideoIDListByHomePageURL("https://www.youtube.com/@helloworld-cn")
	assert.NoError(t, err)
	assert.Greater(t, len(videoIDLst), 0)
	t.Logf("video lst len = %v", len(videoIDLst))
}
