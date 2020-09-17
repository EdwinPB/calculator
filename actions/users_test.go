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

	tcases := []struct {
		values     url.Values
		errMessaje string
	}{
		{
			values: url.Values{
				"Name":     []string{""},
				"LastName": []string{"Polo"},
				"Email":    []string{"edwin@polo.com"},
				"Age":      []string{"26"},
			},
			errMessaje: "Name cannot be empty",
		},
		{
			values: url.Values{
				"Name":     []string{"Edwin"},
				"LastName": []string{""},
				"Email":    []string{"edwin@polo.com"},
				"Age":      []string{"26"},
			},
			errMessaje: "Last Name cannot be empty",
		},
		{
			values: url.Values{
				"Name":     []string{"Edwin"},
				"LastName": []string{"Polo"},
				"Email":    []string{""},
				"Age":      []string{"26"},
			},
			errMessaje: "Email cannot be empty",
		},
		{
			values: url.Values{
				"Name":     []string{"Edwin"},
				"LastName": []string{"Polo"},
				"Email":    []string{"edwin@polo.com"},
				"Age":      []string{""},
			},
			errMessaje: "Age cannot be empty",
		},
	}

	for _, tcase := range tcases {
		res = as.HTML("/users/create").Post(tcase.values)

		as.Equal(http.StatusUnprocessableEntity, res.Code)
		as.Contains(res.Body.String(), tcase.errMessaje)
	}

}
