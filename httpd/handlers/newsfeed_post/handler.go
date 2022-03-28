package handlersNewsfeedPost

import (
	"encoding/json"
	"net/http"
	"newsfeeder/platform/newsfeed"
)

type IServices interface {
	NewsFeedService() newsfeed.INewsFeedService
}

func Handler(services IServices) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body RequestBody
		json.NewDecoder(r.Body).Decode(&body)

		item := &newsfeed.Item{Message: body.Message}
		services.NewsFeedService().Add(*item)

		w.WriteHeader(204)
		w.Write([]byte{})
	}
}
