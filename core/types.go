package core

type Ingredient struct {
	Name      string  `yaml:"name"`
	Subname   string  `yaml:"subname"`
	Calories  float32 `yaml:"calories"`
	Serving   float32 `yaml:"serving"`
	Carbs     float32 `yaml:"carbs"`
	Fiber     float32 `yaml:"fiber"`
	Protein   float32 `yaml:"protein"`
	Fat       float32 `yaml:"fat"`
	Sodium    float32 `yaml:"sodium"`
	Magnesium float32 `yaml:"magnesium"`
	Potassium float32 `yaml:"potassium"`
}

type Amount struct {
	Value float32 `yaml:"value"`
	Unit  string  `yaml:"unit"`
}

type Item struct {
	Name     string `yaml:"name"`
	Servings Amount `yaml:"servings"`
	Contains []Item `yaml:"contains"`
}
