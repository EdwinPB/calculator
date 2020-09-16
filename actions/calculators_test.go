package actions

import (
	"fmt"
	"net/http"
	"net/url"
)

func (as *ActionSuite) Test_Calculators_Show() {
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
	}
	for _, tcase := range tcases {
		res := as.HTML("/calculators/calculate").Post(url.Values{"Numbers": []string{tcase.sNumbers}})
		as.Equal(http.StatusOK, res.Code)
		as.Contains(res.Body.String(), tcase.resultado)
	}

}

func (as *ActionSuite) Test_Calculators_Show_Theme() {
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
