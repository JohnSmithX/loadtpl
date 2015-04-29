package main

import (
	"log"
	"net/http"

	"github.com/JohnSmithX/loadtpl"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	//load templates form `views/`
	t, err := loadtpl.LoadTemplates("views/")
	if err != nil {
		log.Fatal(err)
	}
	r.SetHTMLTemplate(t)
	//router
	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", nil)
	})

	//listen
	r.Run(":8081")
}
