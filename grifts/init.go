package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/shanebarringer/meal_planner/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
