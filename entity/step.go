package entity

type Step struct {
	Row    int
	Column int
}

func NewStep(step int) *Step {
	return &Step{
		Row:    step / 10,
		Column: step % 10,
	}
}
