package main

import (
    "github.com/yuaoki/test-aoki/restAPI/httpd/handler"

    "github.com/yuaoki/test-aoki/restAPI/article"

    "github.com/gin-gonic/gin"
)

func main() {
    article := article.New()
    r := gin.Default()
    r.GET("/article", handler.ArticlesGet(article))
    r.POST("/article", handler.ArticlePost(article))

    r.Run() // listen and serve on 0.0.0.0:8080

}