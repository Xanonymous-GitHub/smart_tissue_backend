package model

type Toilet struct {
	Id          string
	Percentage  float64
	Distance    float64
	MaxDistance float64
	Location    string
	State       ToiletState
}

func (toilet *Toilet) GetId() string {
	return toilet.Id
}
