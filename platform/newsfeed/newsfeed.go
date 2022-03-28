package newsfeed

import "newsfeeder/platform/syncutils"

type Item struct {
	Message string
}

type INewsFeedService interface {
	Add(item ...Item)
	GetAll() []Item
}

type NewsFeedService struct {
	items *syncutils.Slice[Item]
}

func NewNewsFeedService() *NewsFeedService {
	return &NewsFeedService{
		items: syncutils.NewConcurrentSlice[Item](),
	}
}

func (n *NewsFeedService) Add(items ...Item) {
	n.items.Append(items...)
}

func (n *NewsFeedService) GetAll() []Item {
	return n.items.Items()
}
