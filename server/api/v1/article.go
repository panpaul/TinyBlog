package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"server/global"
	"server/service"
)

type ArticleForm struct {
	Author string `json:"author"`
	Page   int    `json:"page"`
	Tag    string `json:"tag"`
}

// GetArticle godoc
// @Summary     get article by uuid
// @Description get details of an article by uuid
// @Tags        Article
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Param       uuid path string true "article uuid"
// @Router      /article/content/{uuid} [get]
func GetArticle(c *gin.Context) {
	id, _ := uuid.FromString(c.Param("uuid"))
	article, err := service.ArticleApp.GetArticle(id)
	global.Pong(err, article, c)
}

// GetArticleList godoc
// @Summary     get article list
// @Description gets a page of article list
// @Tags        Article
// @Accept      json
// @Produce     json
// @Param       article body v1.ArticleForm true "filter by author or tag or page"
// @Router      /article/list [post]
func GetArticleList(c *gin.Context) {
	var a ArticleForm
	_ = c.ShouldBindJSON(&a)
	articles, err := service.ArticleApp.GetArticleList(a.Author, a.Tag, a.Page)
	global.Pong(err, articles, c)
}

// GetArticlePage godoc
// @Summary     get total pages of article list
// @Description gets pages of article list
// @Tags        Article
// @Accept      json
// @Produce     json
// @Param       article body v1.ArticleForm true "filter by author or tag or page"
// @Router      /article/page [post]
func GetArticlePage(c *gin.Context) {
	var a ArticleForm
	_ = c.ShouldBindJSON(&a)
	page, err := service.ArticleApp.GetArticlePages(a.Author, a.Tag)
	global.Pong(err, page, c)
}

func ArticleApi(c *gin.RouterGroup) {
	c.GET("/content/:uuid", GetArticle)
	c.POST("/list", GetArticleList)
	c.POST("/page", GetArticlePage)
}

func ArticleSetup(base string) {
	global.LOG.Debug("api v1 /article setup")
}
