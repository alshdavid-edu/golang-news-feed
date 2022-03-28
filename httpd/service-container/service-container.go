package servicecontainer

import (
	"newsfeeder/platform/newsfeed"
)

type ServiceContainer struct {
	FeedService *newsfeed.NewsFeedService
}

func (s *ServiceContainer) NewsFeedService() newsfeed.INewsFeedService {
	return s.FeedService
}

func New() *ServiceContainer {
	newsFeedService := newsfeed.NewNewsFeedService()
	return &ServiceContainer{
		FeedService: newsFeedService,
	}
}
