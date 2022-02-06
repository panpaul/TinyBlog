package main

import (
	"log"
	"server/model"
	"server/service"
)

func devDebug() error {
	model.MigrateDatabase()

	//var article model.Article
	//author1, _ := uuid.FromString("abbc188a-7b77-485a-80e1-bbcedb6de7e3")
	//author3, _ := uuid.FromString("a22dcbb2-7d61-4374-b2e5-926a285b523b")
	//for i := 0; i <= 105; i++ {
	//	article = model.Article{
	//		AuthorID:    author1,
	//		Title:       "111:title" + strconv.Itoa(i),
	//		Content:     "111:content" + strconv.Itoa(i),
	//		Description: "111:description" + strconv.Itoa(i),
	//		Level:       model.RoleGuest,
	//		Tags:        []string{},
	//	}
	//	article.UUID, _ = uuid.NewV4()
	//	service.ArticleApp.CreateArticle(&article)
	//
	//	article = model.Article{
	//		AuthorID:    author3,
	//		Title:       "333:title" + strconv.Itoa(i),
	//		Content:     "333:content" + strconv.Itoa(i),
	//		Description: "333:description" + strconv.Itoa(i),
	//		Level:       model.RoleGuest,
	//		Tags:        []string{},
	//	}
	//	article.UUID, _ = uuid.NewV4()
	//	service.ArticleApp.CreateArticle(&article)
	//}

	a, e := service.ArticleApp.GetArticlePages("dev1", "tag2")
	log.Println(a)
	log.Println(e)
	return nil
}
