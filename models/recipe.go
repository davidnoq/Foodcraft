package models

type Recipe struct {
	ID                    int    `bson:"id"`
	Title                 string `bson:"title"`
	Image                 string `bson:"image"`
	ImageType             string `bson:"imageType"`
	UsedIngredientCount   int    `bson:"usedIngredientCount"`
	MissedIngredientCount int    `bson:"missedIngredientCount"`

	MissedIngredients []struct {
		ID           int      `bson:"id"`
		Amount       float64  `bson:"amount"`
		Unit         string   `bson:"unit"`
		UnitLong     string   `bson:"unitLong"`
		UnitShort    string   `bson:"unitShort"`
		Aisle        string   `bson:"aisle"`
		Name         string   `bson:"name"`
		Original     string   `bson:"original"`
		OriginalName string   `bson:"originalName"`
		Meta         []string `bson:"meta"`
		Image        string   `bson:"image"`
	} `bson:"missedIngredients"`
	UsedIngredients []struct {
		ID           int      `bson:"id"`
		Amount       int      `bson:"amount"`
		Unit         string   `bson:"unit"`
		UnitLong     string   `bson:"unitLong"`
		UnitShort    string   `bson:"unitShort"`
		Aisle        string   `bson:"aisle"`
		Name         string   `bson:"name"`
		Original     string   `bson:"original"`
		OriginalName string   `bson:"originalName"`
		Meta         []string `bson:"meta"`
		Image        string   `bson:"image"`
	} `bson:"usedIngredients"`
	Likes int `bson:"likes"`
}
