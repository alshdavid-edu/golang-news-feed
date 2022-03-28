package servicecontainer

import (
	"newsfeeder/platform/newsfeed"
)

type ServiceContainer struct {
	newsFeedService newsfeed.INewsFeedService
}

func (s *ServiceContainer) NewsFeedService() newsfeed.INewsFeedService {
	return s.newsFeedService
}

func New() *ServiceContainer {
	newsFeedService := newsfeed.NewNewsFeedService()

	return &ServiceContainer{
		newsFeedService,
	}
}
