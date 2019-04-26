package thirdparty

type recipe struct {
	URI         string  `json:"uri"`
	Label       string  `json:"label"`
	Image       string  `json:"image"`
	URL         string  `json:"url"`
	Yield       float64 `json:"yield"`
	Calories    float64 `json:"calories"`
	TotalWeight float64 `json:"totalWeight"`
	TotalTime   float64 `json:"totalTime"`
}

type recipeWrapper struct {
	Recipe *recipe `json:"recipe"`
}

type result struct {
	Hits []*recipeWrapper `json:"hits"`
}
