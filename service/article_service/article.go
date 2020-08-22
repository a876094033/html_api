package article_service

import (
	"encoding/json"
	"html_api/models"
	"html_api/pkg/gredis"
	"html_api/pkg/logging"
	"html_api/service/cache_service"
)

type Article struct {
	ID             int
	ArticcleName   string
	ArticcleType   int
	CreateAt       string
	ArticleContent string
	PageNum        int
	PageSize       int
}

func (a *Article) GetArticles() ([]*models.Article, error) {
	var (
		article, cacheArticle []*models.Article
	)
	cache := cache_service.Borrow{
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetBorrowsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	//gredis.Set(key, articles, 3600)
	return article, nil
}
func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.getMaps())
}
func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	//if b.AmountLimit >= 0 {
	//	maps["amount_limit"] = b.AmountLimit
	//}
	return maps
}

func (a *Article) GetArticle() (*models.Article, error) {
	return models.GetArticle(a.ID)
}