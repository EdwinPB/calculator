package models

func (ms ModelSuite) Test_User() {
	user := User{
		Name:     "Edwin",
		LastName: "Polo",
		Email:    "edwin@polo.com",
		Age:      "33",
	}
	verrs, err := ms.DB.ValidateAndCreate(&user)
	ms.NoError(err)
	ms.Empty(verrs.Errors)
}
