package xg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinHTML func(c *gin.Context) (code int, name string, obj interface{})
type GinObject func(c *gin.Context) (code int, obj interface{})
type GinString func(c *gin.Context) (code int, str string)
type GinData func(c *gin.Context) (code int, contentType string, data []byte)
type GinFile func(c *gin.Context) (filepath string)

func HandleHTML(fn GinHTML) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, name, obj := fn(c)
		c.HTML(code, name, obj)
	}
}

func HandleJSON(fn GinObject) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, obj := fn(c)
		c.JSON(code, obj)
	}
}

func HandleXML(fn GinObject) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, obj := fn(c)
		c.XML(code, obj)
	}
}

func HandleYAML(fn GinObject) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, obj := fn(c)
		c.YAML(code, obj)
	}
}

func HandleString(fn GinString) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, obj := fn(c)
		c.String(code, obj)
	}
}

func HandleRedirect(fn GinString) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, obj := fn(c)
		c.Redirect(code, obj)
	}
}

func HandleData(fn GinData) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, contentType, data := fn(c)
		c.Data(code, contentType, data)
	}
}

func HandleFile(fn GinFile) gin.HandlerFunc {
	return func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, fn(c))
	}
}
