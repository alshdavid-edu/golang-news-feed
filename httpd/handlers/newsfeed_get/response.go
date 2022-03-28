package handlersNewsfeedGet

type ResponseItem struct {
	Message string `json:"message"`
}

type Response struct {
	Items []*ResponseItem `json:"items"`
}
