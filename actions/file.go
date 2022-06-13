package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func File(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("files/index.plush.html"))
}

func ReceiveFiles(c buffalo.Context) error {
	_, fileInfo, err := c.Request().FormFile("document")
	fmt.Println(err)
	if err != nil {
		panic(err)
	}

	fmt.Println(fileInfo.Filename)

	// fmt.Fprintf(c.Response(), fileInfo.Filename)

	return c.Redirect(http.StatusSeeOther, "/name")
	// return c.Render(http.StatusOK, r.HTML("files/name.plush.html"))
}

func Name(c buffalo.Context) error {

	c.Set("file_name", "melvin.jpg")
	return c.Render(http.StatusOK, r.HTML("files/name.plush.html"))
}
