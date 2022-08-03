package model

type Restroom struct {
	Id           string
	Location     string
	ToiletIdList []string
}

func (restroom *Restroom) AddToiletId(toiletId string) {
	restroom.ToiletIdList = append(restroom.ToiletIdList, toiletId)
}


func (restroom *Restroom) RemoveIdFromToiletIdList(toiletId string) []string {
	for toiletIndex, currentToiletId := range restroom.ToiletIdList{
		if currentToiletId == toiletId {
			restroom.ToiletIdList[toiletIndex] = restroom.ToiletIdList[len(restroom.ToiletIdList) - 1]
			return restroom.ToiletIdList[:len(restroom.ToiletIdList) - 1]
		}
	}
	return restroom.ToiletIdList
}

func (restroom *Restroom) IsToiletIdInList(toiletId string) bool {
	for _, currentToiletId := range restroom.ToiletIdList{
		if currentToiletId == toiletId {
			return true
		}
	}
	return false
}
