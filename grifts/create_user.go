package grifts

import (
	"string_calculator_app/models"

	. "github.com/markbates/grift/grift"
)

var _ = Desc("create_user", "Task Description")
var _ = Add("create_user", func(c *Context) error {
	tx := models.DB

	users := models.Users{
		models.User{
			Name:     "Edwin",
			LastName: "Polo",
			Email:    "edwin@polo.com",
			Age:      "25",
		},
		models.User{
			Name:     "Larry",
			LastName: "M",
			Email:    "larry@polo.com",
			Age:      "21",
		},
		models.User{
			Name:     "Rodo",
			LastName: "M",
			Email:    "rodo@polo.com",
			Age:      "21",
		},
	}

	if err := tx.Create(&users); err != nil {
		return err
	}
	return nil
})
