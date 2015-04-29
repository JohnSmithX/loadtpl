### What

`loadtpl` can load all templates from a path which has lots of subdirectories into cache.

### Use with [gin](https://github.com/gin-gonic/gin)

#### templates

`Directory Structure`

..views

....body

......body.html

....footer

......footer.html

....index.html


`index.html`

```html
<!DOCTYPE html>
<html>
<head>
	<title>example</title>
</head>
<body>
	{{template "body/body.html"}}
	{{template "footer/footer.html"}}
</body>
</html>
```

#### gin code

```go
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

```