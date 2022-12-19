package twitter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func (c client) GetUser(username string) (u User, err error) {
	u.client = &c
	endpoint := "https://twitter.com/i/api/graphql/mCbpQvZAw6zu_4PvuAUVVQ/UserByScreenName"
	variables := struct {
		Username string `json:"screen_name"`
		A        bool   `json:"withSafetyModeUserFields"`
		B        bool   `json:"withSuperFollowsUserFields"`
	}{username, false, false}
	j, _ := json.Marshal(variables)
	v := url.Values{}
	v.Add("variables", string(j))

	req, _ := http.NewRequest("GET", endpoint+"?"+v.Encode(), nil)
	req.Header = headers(c.token)
	resp, err := http.DefaultClient.Do(req)
	err = json.NewDecoder(resp.Body).Decode(&u)
	return u, err
}

type User struct {
	restId          string
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	Username        string    `json:"username"`
	CreatedAt       time.Time `json:"created_at"`
	Description     string    `json:"description"`
	Entities        Entities  `json:"entities"`
	Location        string    `json:"location"`
	PinnedTweetIds  []string  `json:"pinned_tweet_id"`
	ProfileImageUrl string    `json:"profile_image_url"`
	Protected       bool      `json:"protected"`
	PublicMetrics   Metrics   `json:"public_metrics"`
	Url             string    `json:"url"`
	Verified        bool      `json:"verified"`
	client          *client
}

type Entities struct {
	Url         UrlEntities         `json:"url"`
	Description DescriptionEntities `json:"description"`
}

type DescriptionEntities struct {
	Urls     []UrlEntity `json:"urls"`
	Hashtags []TagEntity `json:"hashtags"`
	Mentions []TagEntity `json:"mentions"`
	Cashtags []TagEntity `json:"cashtags"`
}

type UrlEntities struct {
	Urls []UrlEntity `json:"urls"`
}

type UrlEntity struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	Url         string `json:"url"`
	ExpandedUrl string `json:"expanded_url"`
	DisplayUrl  string `json:"display_url"`
}

type TagEntity struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type Metrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}

var twittertime = "Mon Jan 2 15:04:05 -0700 2006"

func (u *User) UnmarshalJSON(data []byte) error {
	data, err := JsonGet(data, "data", "user", "result")
	if err != nil {
		return err
	}
	type ApiResp struct {
		Id     string `json:"id"`
		RestId string `json:"rest_id"`
		Legacy struct {
			CreatedAt            string   `json:"created_at"`
			Description          string   `json:"description"`
			Entities             Entities `json:"entities"`
			FollowersCount       int      `json:"followers_count"`
			FriendsCount         int      `json:"friends_count"`
			Name                 string   `json:"name"`
			NormalFollowersCount int      `json:"normal_followers_count"`
			PinnedTweetIdsStr    []string `json:"pinned_tweet_ids_str"`
			ProfileImageUrlHttps string   `json:"profile_image_url_https"`
			Protected            bool     `json:"protected"`
			ScreenName           string   `json:"screen_name"`
			Verified             bool     `json:"verified"`
			Location             string   `json:"location"`
			StatusesCount        int      `json:"statuses_count"`
			ListedCount          int      `json:"listed_count"`
			Url                  string   `json:"url"`
		} `json:"legacy"`
	}
	var resp ApiResp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}
	u.restId = resp.RestId
	u.Id = resp.Id
	u.Name = resp.Legacy.Name
	u.Username = resp.Legacy.ScreenName
	u.CreatedAt, err = time.Parse(twittertime, resp.Legacy.CreatedAt)
	if err != nil {
		return err
	}
	u.Description = strings.Trim(resp.Legacy.Description, " \n\r\t")
	u.Entities = resp.Legacy.Entities
	u.Location = resp.Legacy.Location
	u.PinnedTweetIds = resp.Legacy.PinnedTweetIdsStr
	u.ProfileImageUrl = resp.Legacy.ProfileImageUrlHttps
	u.Protected = resp.Legacy.Protected
	u.PublicMetrics = Metrics{
		FollowersCount: resp.Legacy.FollowersCount,
		FollowingCount: resp.Legacy.FriendsCount,
		TweetCount:     resp.Legacy.StatusesCount,
		ListedCount:    resp.Legacy.ListedCount,
	}
	u.Url = resp.Legacy.Url
	u.Verified = resp.Legacy.Verified
	return nil
}

func Itoa(i int) string {
	if i > 1000000 {
		return fmt.Sprintf("%.1fM", float32(i)/1000000)
	}
	if i > 1000 {
		return fmt.Sprintf("%.1fK", float32(i)/1000)
	}
	return fmt.Sprintf("%d", i)
}

var nbsp = string(rune(0xA0))

func (u User) String() (s string) {
	bold := lipgloss.NewStyle().Bold(true).Render
	dimmed := lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render
	verified := lipgloss.NewStyle().Background(lipgloss.Color("21")).Foreground(lipgloss.Color("15")).Render
	full := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Align(lipgloss.Left).Width(60).Render
	s += bold(u.Name) + " "
	if u.Verified {
		s += verified("☑ ") + " "
	}
	s += dimmed(" @"+u.Username) + "\n"
	s += u.Description + "\n"
	if u.Location != "" {
		s += dimmed("📍" + nbsp + u.Location + " ")
	}
	if u.Url != "" {
		s += dimmed("🔗" + nbsp + u.Entities.Url.Urls[0].DisplayUrl + " ")
	}
	s += dimmed("📅"+nbsp+u.CreatedAt.Format("February 2006")) + "\n"
	s += bold(Itoa(u.PublicMetrics.FollowingCount)) + dimmed(" Following  ") + bold(Itoa(u.PublicMetrics.FollowersCount)) + dimmed(" Followers")
	return full(s)
}
