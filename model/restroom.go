package model

type Restroom struct {
	Id           string
	Location     string
	ToiletIdList []Toilet
}
