package twitter

var (
	auth = "Bearer AAAAAAAAAAAAAAAAAAAAAPYXBAAAAAAACLXUNDekMxqa8h%2F40K4moUkGsoc%3DTYfbDKbT3jJPCEVnMYqilB28NHfOPqkca3qaAxGfsyKCs0wRbw"

	api      = "https://api.twitter.com"
	activate = api + "/1.1/guest/activate.json"

	userShow  = api + "/1.1/users/show.json"
	photoRail = api + "/1.1/statuses/media_timeline.json"
	status    = api + "/1.1/statuses/show"
	search    = api + "/2/search/adaptive.json"

	timelineApi   = api + "/2/timeline"
	timeline      = timelineApi + "/profile"
	mediaTimeline = timelineApi + "/media"
	listTimeline  = timelineApi + "/list.json"
	tweet         = timelineApi + "/conversation"

	graphql          = api + "/graphql"
	graphUser        = graphql + "/I5nvpI91ljifos1Y3Lltyg/UserByRestId"
	graphList        = graphql + "/JADTh6cjebfgetzvF3tQvQ/List"
	graphListBySlug  = graphql + "/ErWsz9cObLel1BF-HjuBlA/ListBySlug"
	graphListMembers = graphql + "/Ke6urWMeCV2UlKXGRy4sow/ListMembers"
)
