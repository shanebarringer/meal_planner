package actions

func (as *ActionSuite) Test_IngredientsResource_List() {
	as.LoadFixture("load recipes")

	res := as.JSON("/ingredients").Get()

	as.Equal(200, res.Code)

	body := res.Body.String()
	as.Contains(body, "Hot Dog Buns")
}

func (as *ActionSuite) Test_IngredientsResource_Show() {
	// as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_IngredientsResource_Create() {
	// as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_IngredientsResource_Update() {
	// as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_IngredientsResource_Destroy() {
	// as.Fail("Not Implemented!")
}
