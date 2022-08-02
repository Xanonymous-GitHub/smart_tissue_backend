package model

type Toilet struct {
	Id         string
	Percentage float32
	Location   string
	State      ToiletState
}

func (toilet *Toilet) GetId() string {
	return toilet.Id
}
