package actions

import (
	"net/http"
	"string_calculator_app/models"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	user := models.User{}
	c.Set("user", user)
	return c.Render(http.StatusOK, r.HTML("index.html"))
}
