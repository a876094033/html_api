package models

import "github.com/jinzhu/gorm"

type Article struct {
	ID             int    `json:"id"`
	ArticcleName   string `json:"articcle_name"`
	ArticcleType   int    `json:"articcle_type"`
	CreatedAt       string `json:"created_at"`
	ArticleContent string `json:"article_content"`
}

func GetArticleTotal(maps map[string]interface{}) (int, error) {
	var count int
	if err := db.Debug().Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var article []*Article
	err := db.Debug().Model(&Article{}).Where(maps).Offset(pageNum / 10 * pageSize).Limit(pageSize).Find(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return article, nil
}
func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Debug().Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}
