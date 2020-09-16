package actions

import (
	"net/http"
	"string_calculator_app/models"

	"github.com/gobuffalo/buffalo"
)

// CalculatorsShow default implementation.
func CalculatorsShow(c buffalo.Context) error {
	theme := c.Params().Get("theme")
	if theme == "" {
		theme = "day"
	}

	c.Set("result", "")
	c.Set("theme", theme)
	return c.Render(http.StatusOK, r.HTML("calculators/show.html"))
}

// CalculatorsCalculate default implementation.
func CalculatorsCalculate(c buffalo.Context) error {
	theme := c.Params().Get("theme")
	if theme == "" {
		theme = "day"
	}

	calculator := models.Calculator{}
	if err := c.Bind(&calculator); err != nil {
		return err
	}
	result, _ := models.Calculate(calculator.EnteredValue)

	c.Set("result", result)
	c.Set("theme", theme)
	return c.Render(http.StatusOK, r.HTML("calculators/show.html"))
}
