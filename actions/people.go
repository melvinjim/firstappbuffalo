package actions

import (
	"net/http"
	"project_buffalo/models"

	"github.com/gobuffalo/buffalo"
)

func PeopleList(c buffalo.Context) error {
	person := models.Person{
		FirstName: "Melvin",
		LastName:  "Jimenez",
	}

	//esto se usa cada vez que yo quiero enviar algo al html
	c.Set("people", person)

	return c.Render(http.StatusOK, r.HTML("people/index.plush.html"))
}
