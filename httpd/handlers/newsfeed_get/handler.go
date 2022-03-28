package handlersNewsfeedGet

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
		response := &Response{Items: []*ResponseItem{}}
		items := services.NewsFeedService().GetAll()

		for _, value := range items {
			responseItem := &ResponseItem{Message: value.Message}
			response.Items = append(response.Items, responseItem)
		}

		w.Header().Add("Content-Type", "application/json")
		bytes, _ := json.Marshal(response)
		w.WriteHeader(200)
		w.Write(bytes)
	}
}
