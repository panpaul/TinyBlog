package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"server/e"
	"server/global"
	"server/middleware"
	"server/model"
	"server/service"
)

type ArticleForm struct {
	Author string `json:"author"`
	Page   int    `json:"page"`
	Tag    string `json:"tag"`
	UUID   string `json:"uuid"`
}

type ModifyArticleForm struct {
	UUID        string   `json:"uuid"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	Tags        []string `json:"tags"`
}

// GetArticle godoc
// @Summary     get article by uuid
// @Description get details of an article by uuid
// @Tags        Article
// @Accept      json
// @Produce     json
// @Param       article body v1.ArticleForm true "article uuid"
// @Router      /article/content [post]
func GetArticle(c *gin.Context) {
	var a ArticleForm
	_ = c.ShouldBindJSON(&a)
	id, _ := uuid.FromString(a.UUID)
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

// ModifyArticle godoc
// @Summary     create or update an article
// @Description create(UUID must be empty string) or update(with UUID provided) an article
// @Tags        Article
// @Accept      json
// @Produce     json
// @Param       article body v1.ModifyArticleForm true "uuid, title, description, content and tags"
// @Router      /article/modify [post]
func ModifyArticle(c *gin.Context) {
	var a ModifyArticleForm
	_ = c.ShouldBindJSON(&a)

	id, err := uuid.FromString(a.UUID)
	if err != nil && a.UUID != "" {
		global.Pong(e.InvalidParameter, nil, c)
		return
	}

	claim, _ := c.Get("claim") // casbin requires it, so it must exists
	user := claim.(*model.Claims)

	article := model.Article{
		UUID:        id,
		AuthorID:    user.UUID,
		Title:       a.Title,
		Description: a.Description,
		Content:     a.Content,
		Level:       0,
		Tags:        a.Tags,
	}

	if a.UUID != "" {
		// update -> check permission
		old, err := service.ArticleApp.GetArticle(id)
		if err != e.Success || old.UUID.IsNil() {
			global.Pong(e.InvalidParameter, nil, c)
			return
		}

		if old.AuthorID != user.UUID && user.Role != model.RoleAdmin {
			global.Pong(e.InsufficientPermission, nil, c)
			return
		}

		article.AuthorID = old.AuthorID

		_, code := service.ArticleApp.UpdateArticle(&article)
		global.Pong(code, article.UUID, c)
	} else {
		// create article
		article.UUID, _ = uuid.NewV4()
		_, code := service.ArticleApp.CreateArticle(&article)
		global.Pong(code, article.UUID, c)
	}
}

func ArticleApi(c *gin.RouterGroup) {
	c.POST("/content", GetArticle)
	c.POST("/list", GetArticleList)
	c.POST("/page", GetArticlePage)
	c.POST(
		"/modify",
		middleware.JwtHandler(),
		middleware.CasbinHandler(),
		ModifyArticle,
	)
}

func ArticleSetup(base string) {
	global.LOG.Debug("api v1 /article setup")
	service.CasbinApp.AddRange([]model.CasbinRule{
		{
			Role:   model.RoleUser,
			Path:   base + "/modify",
			Method: "POST",
		},
		{
			Role:   model.RoleAdmin,
			Path:   base + "/modify",
			Method: "POST",
		},
	})
}
