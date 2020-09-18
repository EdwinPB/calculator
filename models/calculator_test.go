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
		{"//;\n1;2;6;7", 16, nil},
		{"//:\n1:2", 3, nil},
		{"-1", 0, errors.New("invalid input")},
		{"1,-2", 0, errors.New("invalid input")},
		{"2\n3,-4", 0, errors.New("invalid input")},
	}

	for _, tcase := range tcases {
		result, err := Calculate(tcase.sNumbers)
		ms.Equal(tcase.result, result)

		if err != nil {
			ms.Equal(tcase.err.Error(), err.Error())
		}
	}
}

func (ms ModelSuite) Test_Calculator_Average_Calculator() {
	calculators := Calculators{}
	ms.Equal(float64(0), calculators.AverageCalculatedValues())

	calculators = Calculators{
		Calculator{
			CalculatedValue: "6",
		},
		Calculator{
			CalculatedValue: "6",
		},
		Calculator{
			CalculatedValue: "9",
		},
	}

	ms.Equal(float64(7), calculators.AverageCalculatedValues())

	calculators = Calculators{
		Calculator{
			CalculatedValue: "6",
		},
		Calculator{
			CalculatedValue: "13",
		},
	}

	ms.Equal(float64(9.5), calculators.AverageCalculatedValues())

	calculators = Calculators{
		Calculator{
			CalculatedValue: "0",
		},
		Calculator{
			CalculatedValue: "0",
		},
	}

	ms.Equal(float64(0), calculators.AverageCalculatedValues())

}
