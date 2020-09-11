package actions

import (
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
