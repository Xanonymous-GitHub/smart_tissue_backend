package model

type Restroom struct {
	Id           string
	Location     string
	ToiletIdList []string
}

func (restroom *Restroom) GetId() string {
	return restroom.Id
}

func (restroom *Restroom) RemoveIdFromToiletIdList(toiletId string) []string {
	for i := 0; i < len(restroom.ToiletIdList); i++ {
		if restroom.ToiletIdList[i] == toiletId {
			restroom.ToiletIdList[i] = restroom.ToiletIdList[len(restroom.ToiletIdList) - 1]
			return restroom.ToiletIdList[:len(restroom.ToiletIdList) - 1]
		}
	}
	return restroom.ToiletIdList
}

func (restroom *Restroom) IsToiletIdInList(toiletId string) bool {
	for i := 0; i < len(restroom.ToiletIdList); i++ {
		if restroom.ToiletIdList[i] == toiletId {
			return true
		}
	}
	return false
}