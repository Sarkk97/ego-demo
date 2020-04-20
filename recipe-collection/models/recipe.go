package models

//RecipeIngredient is a recipe ingredient
type RecipeIngredient struct {
	ID        string `json:"id" gorm:"PRIMARY_KEY;varchar(120)"`
	RecipeID  string `json:"recipeId" gorm:"int;not null"`
	Name      string `json:"name" gorm:"varchar(120);not null"`
	UOM       string `json:"uom" gorm:"varchar(120);not  null"`
	Quantity  uint   `json:"quantity" gorm:"int;not null"`
	ImageURL  string `json:"image" gorm:"string"`
	CreatedAt string `json:"-" gorm:"DATETIME"`
	UpdatedAt string `json:"-" gorm:"TIMESTAMP"`
}

//Recipe is a recipe
type Recipe struct {
	ID          string              `json:"id" gorm:"PRIMARY_KEY;varchar(120)"`
	Name        string              `json:"name" gorm:"varchar(120);not null"`
	PrepTime    uint                `json:"prepTime" gorm:"int;not null"`
	Difficulty  uint                `json:"difficulty" gorm:"int;not null"`
	Ingredients []*RecipeIngredient `json:"ingredients"`
	CreatedAt   string              `json:"-" gorm:"DATETIME"`
	UpdatedAt   string              `json:"-" gorm:"TIMESTAMP"`
}
