package model

type Restroom struct {
	Id           string   `json:"id"`
	Location     string   `json:"location"`
	ToiletIdList []string `json:"toiletIdList"`
}
