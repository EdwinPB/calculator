package actions

import (
	"fmt"
	"net/http"
	"strconv"
	"string_calculator_app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

// CalculatorsShow default implementation.
func CalculatorsShow(c buffalo.Context) error {
	c.Set("result", "")
	c.Set("theme", c.Params().Get("theme"))

	userID := c.Session().Get("current_user_id")
	if userID == nil {
		return c.Render(http.StatusUnprocessableEntity, r.HTML("index.html"))
	}

	return c.Render(http.StatusOK, r.HTML("calculators/show.html"))
}

// CalculatorsCalculate default implementation.
func CalculatorsCalculate(c buffalo.Context) error {
	calculator := models.Calculator{}
	if err := c.Bind(&calculator); err != nil {
		return err
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errTransactionNoFound
	}

	result, err := models.Calculate(fmt.Sprintf("%s", calculator.EnteredValue))
	if err != nil {
		c.Set("calculateError", err)
		c.Set("theme", c.Params().Get("theme"))
		c.Set("result", "0")
		return c.Render(http.StatusUnprocessableEntity, r.HTML("calculators/show.html"))
	}

	user := models.User{}
	tx.Find(&user, c.Session().Get("current_user_id"))

	calculator.CalculatedValue = strconv.Itoa(result)
	calculator.UserID = user.ID

	verrs, err := calculator.Validate(tx)

	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("calculateError", verrs)
		c.Set("theme", c.Params().Get("theme"))
		c.Set("result", "0")
		return c.Render(http.StatusUnprocessableEntity, r.HTML("calculators/show.html"))
	}

	if err := tx.Create(&calculator); err != nil {
		return err
	}

	c.Set("result", result)
	c.Set("theme", c.Params().Get("theme"))
	return c.Render(http.StatusOK, r.HTML("calculators/show.html"))
}

// CalculatorsShowReport default implementation.
func CalculatorsShowReport(c buffalo.Context) error {

	return c.Render(http.StatusOK, r.HTML("calculators/report.html"))
}
