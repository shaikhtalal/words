package router

import (
	"words/pkg/word"

	"github.com/gin-gonic/gin"
)

func HandleHTTP(e *gin.Engine) {
	e.POST("/words", word.CountWords)
}
