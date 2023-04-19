package models

type Information struct {
	Vegetarian               bool   `json:"vegetarian"`
	Vegan                    bool   `json:"vegan"`
	GlutenFree               bool   `json:"glutenFree"`
	DairyFree                bool   `json:"dairyFree"`
	VeryHealthy              bool   `json:"veryHealthy"`
	Cheap                    bool   `json:"cheap"`
	VeryPopular              bool   `json:"veryPopular"`
	Sustainable              bool   `json:"sustainable"`
	WeightWatcherSmartPoints int    `json:"weightWatcherSmartPoints"`
	Gaps                     string `json:"gaps"`
	LowFodmap                bool   `json:"lowFodmap"`
	Ketogenic                bool   `json:"ketogenic"`
	Whole30                  bool   `json:"whole30"`
	Servings                 int    `json:"servings"`
	SourceURL                string `json:"sourceUrl"`
	SpoonacularSourceURL     string `json:"spoonacularSourceUrl"`
	AggregateLikes           int    `json:"aggregateLikes"`
	CreditText               string `json:"creditText"`
	SourceName               string `json:"sourceName"`
	ExtendedIngredients      []struct {
		ID              int           `json:"id"`
		Aisle           string        `json:"aisle"`
		Image           string        `json:"image"`
		Name            string        `json:"name"`
		Amount          float64       `json:"amount"`
		Unit            string        `json:"unit"`
		UnitShort       string        `json:"unitShort"`
		UnitLong        string        `json:"unitLong"`
		OriginalString  string        `json:"originalString"`
		MetaInformation []interface{} `json:"metaInformation"`
	} `json:"extendedIngredients"`
	ID             int    `json:"id"`
	Title          string `json:"title"`
	ReadyInMinutes int    `json:"readyInMinutes"`
	Image          string `json:"image"`
	ImageType      string `json:"imageType"`
	Instructions   string `json:"instructions"`
}