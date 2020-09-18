package actions

import (
	"fmt"
	"net/http"
	"net/url"
	"string_calculator_app/models"
)

func (as *ActionSuite) Test_Calculators_Show() {
	as.Session.Set("current_user_id", "5cff7a63-a9d5-4ba3-9316-682c56df6866")
	res := as.HTML("/calculators/show").Get()
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_Calculators_Calculate() {
	tcases := []struct {
		sNumbers  string
		resultado string
	}{
		{"", "0"},
		{"1", "1"},
		{"1,2,3", "6"},
		{"//;\n1;2;3", "6"},
	}

	user := models.User{
		Name:     "Edwin",
		LastName: "Polo",
		Email:    "edwin@edwi.com",
		Age:      "26",
	}
	as.NoError(as.DB.Create(&user))
	as.Session.Set("current_user_id", user.ID)

	for _, tcase := range tcases {
		res := as.HTML("/calculators/calculate").Post(url.Values{"Numbers": []string{tcase.sNumbers}})
		as.Equal(http.StatusOK, res.Code)
		as.Contains(res.Body.String(), tcase.resultado)
	}

}

func (as *ActionSuite) Test_Calculators_Show_Theme() {
	as.Session.Set("current_user_id", "5cff7a63-a9d5-4ba3-9316-682c56df6866")
	res := as.HTML("/calculators/show").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), `<div class="center day">`)
	as.Contains(res.Body.String(), `<input id="day" name="day" type="radio" value="true" checked>`)

	res = as.HTML("/calculators/show/?theme=day").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), `<div class="center day">`)
	as.Contains(res.Body.String(), `<input id="day" name="day" type="radio" value="true" checked>`)

	tcases := []struct {
		theme  string
		rTheme string
	}{
		{"", "day"},
		{"day", "day"},
		{"night", "night"},
	}

	for _, tcase := range tcases {
		res := as.HTML(fmt.Sprintf("/calculators/show/?theme=%v", tcase.theme)).Get()

		as.Equal(http.StatusOK, res.Code)
		as.Contains(res.Body.String(), fmt.Sprintf(`<div class="center %v">`, tcase.rTheme))
		as.Contains(res.Body.String(), fmt.Sprintf(`<input id="%v" name="%v" type="radio" value="true" checked>`, tcase.rTheme, tcase.rTheme))
	}

}

func (as *ActionSuite) Test_Calculators_Show_Report() {
	res := as.HTML("/calculators/show/report").Get()
	as.Equal(http.StatusOK, res.Code)

	as.Contains(res.Body.String(), "<th>Name</th>")
	as.Contains(res.Body.String(), "<th>Last Name</th>")
	as.Contains(res.Body.String(), "<th>Number of calculations</th>")
	as.Contains(res.Body.String(), "<th>Max Calculated Value</th>")
	as.Contains(res.Body.String(), "<th>Min Calculated Value</th>")
	as.Contains(res.Body.String(), "<th>Average Of Calculated Values</th>")

}
