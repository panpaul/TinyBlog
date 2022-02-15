package service

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"server/e"
	"server/global"
	"server/model"
	"sync"
)

type ArticleService struct {
	tableName string
	pageLimit int
	once      sync.Once
}

var ArticleApp = new(ArticleService)

func (a *ArticleService) Setup() {
	a.once.Do(func() {
		a.pageLimit = 10
		stmt := &gorm.Statement{DB: global.DB}
		err := stmt.Parse(&model.Article{})
		if err != nil {
			global.LOG.Panic("Failed to get Article table name", zap.Error(err))
			return
		}
		a.tableName = stmt.Schema.Table
	})
}

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

func (a *ArticleService) GetArticleList(author string, tag string, page int) ([]*model.ArticleDescription, e.Err) {
	var articles []*model.ArticleDescription

	qry := global.DB.
		Model(&model.Article{}).
		Joins("Author")
	if author != "" {
		qry = qry.Where("\"Author\".\"user_name\" = ?", author)
	}
	if tag != "" {
		qry = qry.Where("? = ANY(\"tags\")", tag)
	}
	qry = qry.
		Select([]string{
			fmt.Sprintf("\"%s\".\"uuid\"", a.tableName),
			fmt.Sprintf("\"%s\".\"created_at\"", a.tableName),
			fmt.Sprintf("\"%s\".\"title\"", a.tableName),
			fmt.Sprintf("\"%s\".\"description\"", a.tableName),
		}).
		Limit(a.pageLimit).
		Offset(a.pageLimit * page).
		Order(fmt.Sprintf("\"%s\".\"created_at\" desc", a.tableName))

	if errors.Is(
		qry.Scan(&articles).Error,
		gorm.ErrRecordNotFound,
	) {
		return []*model.ArticleDescription{}, e.NotFound
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

	qry := global.DB.
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

	cnt := count / int64(a.pageLimit)
	if count%int64(a.pageLimit) != 0 {
		cnt++
	}

	return cnt, e.Success
}

func (a *ArticleService) UpdateArticle(article *model.Article) (*model.Article, e.Err) {
	if err := global.DB.Model(article).
		Select("Title", "Description", "Content", "EnableMath", "Tags").
		Updates(article).Error; err != nil {
		return &model.Article{}, e.DBUpdateError
	}

	return article, e.Success
}
