package app

import (
	"fmt"
	"ginBlog/core/pkg/logger"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

func BindForm(c *gin.Context, form interface{}) []string { return Bind("form", c, form) }

func BindJson(c *gin.Context, json interface{}) []string {
	return Bind("json", c, json)
}

func BindQuery(c *gin.Context, query interface{}) []string {
	return Bind("query", c, query)
}

func BindHeader(c *gin.Context, query interface{}) []string {
	return Bind("header", c, query)
}

func Bind() {
	
}
