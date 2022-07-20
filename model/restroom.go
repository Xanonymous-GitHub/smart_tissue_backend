package model

type Restroom struct {
	Id           string
	Location     string `json:"location"`
	ToiletIdList []string
}
