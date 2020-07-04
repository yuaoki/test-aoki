package handler

import (
    "net/http"

    "github.com/yuaoki/test-aoki/tree/master/restAPI/article"

    "github.com/gin-gonic/gin"
)

func ArticlesGet(articles *article.Articles) gin.HandlerFunc {
    return func(c *gin.Context) {
        result := articles.GetAll()
        c.JSON(http.StatusOK, result)
    }
}

type ArticlePostRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}

func ArticlePost(post *article.Articles) gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := ArticlePostRequest{}
        c.Bind(&requestBody)

        item := article.Item{
            Title:       requestBody.Title,
            Description: requestBody.Description,
        }
        post.Add(item)

        c.Status(http.StatusNoContent)
    }
}