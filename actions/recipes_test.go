package actions

import (
	"github.com/shanebarringer/meal_planner/models"
)

func (as *ActionSuite) Test_RecipesResource_List() {
	as.LoadFixture("load recipes")
	res := as.JSON("/recipes").Get()

	as.Equal(200, res.Code)
	body := res.Body.String()
	as.Contains(body, "Veggie Dogs")
	as.Contains(body, "Soba with Tofu")
	as.Contains(body, "Something Delicious")

}

func (as *ActionSuite) Test_RecipesResource_Show() {
	as.LoadFixture("load recipes")

	res := as.JSON("/recipes/3e9d4326-fc93-4946-845e-789790ae5116").Get()
	as.Equal(200, res.Code)

	body := res.Body.String()
	as.Contains(body, "Veggie Dogs")
}

func (as *ActionSuite) Test_RecipesResource_Create() {
	r := &models.Recipe{
		Name:     "Portabella Sandwich",
		PrepTime: 22,
	}
	res := as.JSON("/recipes").Post(r)

	as.Equal(201, res.Code)

	err := as.DB.First(r)
	as.NoError(err)
	as.NotZero(r.ID)
	as.Equal("Portabella Sandwich", r.Name)

}

func (as *ActionSuite) Test_RecipesResource_Update() {
	as.LoadFixture("load recipes")

	res := as.JSON("/recipes/3e9d4326-fc93-4946-845e-789790ae5116").Put(&models.Recipe{Name: "Beyond Burger"})

	as.Equal(200, res.Code)

	body := res.Body.String()
	as.Contains(body, "Beyond Burger")

}

func (as *ActionSuite) Test_RecipesResource_Destroy() {
	as.LoadFixture("load recipes")

	res := as.JSON("/recipes/3e9d4326-fc93-4946-845e-789790ae5116").Delete()

	as.Equal(200, res.Code)

	err := as.DB.Find(&models.Recipe{}, "3e9d4326-fc93-4946-845e-789790ae5116")

	as.NotNil(err)
}
