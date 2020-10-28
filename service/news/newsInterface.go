package news

import "news.com/models"

type Collector interface {
	GetTitleData() []models.Title
}

