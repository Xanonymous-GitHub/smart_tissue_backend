package model

type Restroom struct {
	Id           string
	Location     string
	ToiletIdList []string
}

func (restroom *Restroom) GetId() string {
	return restroom.Id
}

func (restroom *Restroom) AddToiletId(toiletId string) {
	restroom.ToiletIdList = append(restroom.ToiletIdList, toiletId)
}
