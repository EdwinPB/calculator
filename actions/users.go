package actions

import (
	"fmt"
	"net/http"
	"string_calculator_app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

// UsersCreate default implementation.
func UsersCreate(c buffalo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	_, err := tx.ValidateAndCreate(&user)

	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/calculator/show")
}
