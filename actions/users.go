package actions

import (
	"fmt"
	"net/http"
	"string_calculator_app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

var errTransactionNoFound = fmt.Errorf("no transaction found")

// UsersCreate default implementation.
func UsersCreate(c buffalo.Context) error {
	c.Session().Clear()
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errTransactionNoFound
	}

	verrs, err := user.Validate(tx)

	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("user", user)
		c.Set("errors", verrs)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("index.html"))
	}

	if err := tx.Create(&user); err != nil {
		return err
	}

	c.Session().Set("current_user_id", user.ID)

	return c.Redirect(http.StatusSeeOther, "/calculator/show")
}
