package singletonpattern

type Single struct {
	single Single
}

func newSingle() *Single {
	return nil
}
