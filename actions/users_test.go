package actions

import (
	"net/http"
	"net/url"
	"string_calculator_app/models"
)

func (as *ActionSuite) Test_Users_Create() {
	res := as.HTML("/users/create").Post(url.Values{
		"Name":     []string{"Edwin"},
		"LastName": []string{"Polo"},
		"Email":    []string{"edwin@polo.com"},
		"Age":      []string{"26"},
	})

	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/calculator/show", res.Location())

	user := models.User{}
	as.NoError(as.DB.Where("email = ?", "edwin@polo.com").First(&user))

}
