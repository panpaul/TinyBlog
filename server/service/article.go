package service

import (
	"errors"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"server/e"
	"server/global"
	"server/model"
)

type ArticleService struct{}

const PageLimit = 20

var ArticleApp = new(ArticleService)

func (a *ArticleService) GetArticle(uuid uuid.UUID) (*model.Article, e.Err) {
	article := &model.Article{}

	if errors.Is(
		global.DB.Model(&model.Article{}).
			Preload(clause.Associations).
			Where("uuid = ?", uuid).
			First(&article).Error,
		gorm.ErrRecordNotFound,
	) {
		return &model.Article{}, e.NotFound
	}

	return article, e.Success
}

func (a *ArticleService) GetArticleList(author string, tag string, page int) ([]*model.Article, e.Err) {
	var articles []*model.Article

	qry := global.DB.Debug().
		Model(&model.Article{}).
		Joins("Author")
	if author != "" {
		qry = qry.Where("\"Author\".\"user_name\" = ?", author)
	}
	if tag != "" {
		qry = qry.Where("? = ANY(\"tags\")", tag)
	}
	qry = qry.
		Limit(PageLimit).
		Offset(PageLimit * page).
		Order("created_at desc")

	if errors.Is(
		qry.Find(&articles).Error,
		gorm.ErrRecordNotFound,
	) {
		return []*model.Article{}, e.NotFound
	}

	return articles, e.Success
}

func (a *ArticleService) CreateArticle(article *model.Article) (*model.Article, e.Err) {
	if err := global.DB.Create(article).Error; err != nil {
		return &model.Article{}, e.DBCreateError
	}

	return article, e.Success
}

func (a *ArticleService) GetArticlePages(author string, tag string) (int64, e.Err) {
	var count int64

	qry := global.DB.Debug().
		Model(&model.Article{}).
		Joins("Author")
	if author != "" {
		qry = qry.Where("\"Author\".\"user_name\" = ?", author)
	}
	if tag != "" {
		qry = qry.Where("? = ANY(\"tags\")", tag)
	}

	err := qry.Count(&count).Error
	if err != nil {
		return 0, e.DBQueryError
	}

	cnt := count / PageLimit
	if count%PageLimit != 0 {
		cnt++
	}

	return cnt, e.Success
}

func (a *ArticleService) UpdateArticle(article *model.Article) (*model.Article, e.Err) {
	if err := global.DB.Model(article).Updates(article).Error; err != nil {
		return &model.Article{}, e.DBUpdateError
	}

	return article, e.Success
}
