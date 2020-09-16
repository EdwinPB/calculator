package models

import "errors"

func (ms ModelSuite) Test_Calculator_Calculate() {
	tcases := []struct {
		sNumbers string
		result   int
		err      error
	}{
		{"", 0, nil},
		{"1", 1, nil},
		{"2", 2, nil},
		{"3", 3, nil},
		{"1,2", 3, nil},
		{"13,2", 15, nil},
		{"1,2,3,5", 11, nil},
		{"2\n3,4", 9, nil},
		{"//;\n1;2", 3, nil},
		{"//:\n1:2", 3, nil},
		{"-1", 0, errors.New("Invalid Input")},
		{"1,-2", 0, errors.New("Invalid Input")},
		{"2\n3,-4", 0, errors.New("Invalid Input")},
	}

	for _, tcase := range tcases {
		result, err := Calculate(tcase.sNumbers)
		ms.Equal(tcase.result, result)

		if err != nil {
			ms.Equal(tcase.err.Error(), err.Error())
		}
	}
}

func (ms ModelSuite) Test_Calculator_Storage_Calculator() {

}
