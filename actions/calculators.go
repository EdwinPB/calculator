package actions

import (
	"net/http"
	"string_calculator_app/models"

	"github.com/gobuffalo/buffalo"
)

// CalculatorsShow default implementation.
func CalculatorsShow(c buffalo.Context) error {
	c.Set("result", "")
	return c.Render(http.StatusOK, r.HTML("calculators/show.html"))
}

// CalculatorsCalculate default implementation.
func CalculatorsCalculate(c buffalo.Context) error {
	calculator := models.Calculator{}
	if err := c.Bind(&calculator); err != nil {
		return err
	}
	result, _ := models.Calculate(calculator.Numbers)
	c.Set("result", result)
	return c.Render(http.StatusOK, r.HTML("calculators/show.html"))
}