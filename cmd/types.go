package cmd

type Food struct {
	Name   string
	Amount float32
	Unit   string
}

type Day struct {
	Foods []Food
}
