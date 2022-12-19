package twitter

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	assert := assert.New(t)

	client, err := NewClient()
	assert.Nil(err)
	twitter, err := client.GetUser("twitter")
	assert.Nil(err)
	assert.Equal("VXNlcjo3ODMyMTQ=", twitter.Id)
	assert.Equal("Twitter", twitter.Name)
	assert.Equal("Twitter", twitter.Username)
	assert.Equal(time.Date(2007, time.February, 20, 14, 35, 54, 0, time.FixedZone("", 0)), twitter.CreatedAt)
	assert.Equal("What's happening?!", twitter.Description)
	assert.Equal("about.twitter.com", twitter.Entities.Url.Urls[0].DisplayUrl)
	assert.Equal("everywhere", twitter.Location)
	assert.Equal("https://pbs.twimg.com/profile_images/1488548719062654976/u6qfBBkF_normal.jpg", twitter.ProfileImageUrl)
	assert.False(twitter.Protected)
	assert.GreaterOrEqual(15006, twitter.PublicMetrics.TweetCount)
	assert.True(twitter.Verified)

	twitter.GetTimeline(false)

	fmt.Println(twitter)
	mastodon, err := client.GetUser("joinmastodon")
	fmt.Println()
	fmt.Println(mastodon)

}
