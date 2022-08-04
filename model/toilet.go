package model

type Toilet struct {
	Id          string      `json:"id"`
	Percentage  float64     `json:"percentage"`
	Distance    float64     `json:"distance"`
	MaxDistance float64     `json:"maxDistance"`
	Location    string      `json:"location"`
	State       ToiletState `json:"state"`
}
