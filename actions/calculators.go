package actions

import (
	"fmt"
	"net/http"
	"string_calculator_app/models"

	"github.com/gobuffalo/buffalo"
)

// CalculatorsShow default implementation.
func CalculatorsShow(c buffalo.Context) error {
	c.Set("result", "")
	c.Set("theme", c.Params().Get("theme"))
	return c.Render(http.StatusOK, r.HTML("calculators/show.html"))
}

// CalculatorsCalculate default implementation.
func CalculatorsCalculate(c buffalo.Context) error {
	calculator := models.Calculator{}
	if err := c.Bind(&calculator); err != nil {
		return err
	}

	result, err := models.Calculate(fmt.Sprintf("%s", calculator.EnteredValue))
	if err != nil {
		// show err
		c.Set("calculateError", err)
	}

	c.Set("result", result)
	c.Set("theme", c.Params().Get("theme"))
	return c.Render(http.StatusOK, r.HTML("calculators/show.html"))
}
