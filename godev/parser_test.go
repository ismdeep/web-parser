package godev

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDownloadLinks(t *testing.T) {
	links, err := GetDownloadLinks()
	assert.NoError(t, err)
	assert.Greater(t, len(links), 0)
	for i, link := range links {
		t.Logf("links[%v] = %v", i, link)
	}
}

func TestGetStableVersions(t *testing.T) {
	versions, err := GetStableVersions()
	assert.NoError(t, err)
	assert.Greater(t, len(versions), 0)
	for i, version := range versions {
		t.Logf("versions[%v] = %v", i, version)
	}
}
