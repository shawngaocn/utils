package xg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandleHTML func(c *gin.Context) (code int, name string, obj interface{})
type HandleObject func(c *gin.Context) (code int, obj interface{})
type HandleString func(c *gin.Context) (code int, str string)
type HandleData func(c *gin.Context) (code int, contentType string, data []byte)
type HandleFile func(c *gin.Context) (filepath string)

func HTML(fn GinHTML) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, name, obj := fn(c)
		c.HTML(code, name, obj)
	}
}

func JSON(fn GinObject) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, obj := fn(c)
		c.JSON(code, obj)
	}
}

func XML(fn GinObject) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, obj := fn(c)
		c.XML(code, obj)
	}
}

func YAML(fn GinObject) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, obj := fn(c)
		c.YAML(code, obj)
	}
}

func String(fn GinString) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, obj := fn(c)
		c.String(code, obj)
	}
}

func Redirect(fn GinString) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, obj := fn(c)
		c.Redirect(code, obj)
	}
}

func Data(fn GinData) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, contentType, data := fn(c)
		c.Data(code, contentType, data)
	}
}

func File(fn GinFile) gin.HandlerFunc {
	return func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, fn(c))
	}
}
