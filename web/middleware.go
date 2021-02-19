package web

import (
	"github.com/gin-gonic/gin"
)

//Middleware
func Middleware(context *gin.Context)  {
	authorization(context)
}
