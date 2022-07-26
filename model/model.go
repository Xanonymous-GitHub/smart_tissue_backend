package model

import "strconv"

var (
	restrooms              map[string]Restroom
	toilets                map[string]Toilet
	undeployedToiletIdList []string
	nextRestroomId         int
)

func Setup() {
	restrooms = map[string]Restroom{}
	toilets = map[string]Toilet{}
	undeployedToiletIdList = []string{}
	nextRestroomId = 1
}

func GetAllRestrooms() map[string]Restroom {
	return restrooms
}

func GetRestroom(id string) Restroom {
	return restrooms[id]
}

func RegisterRestroom(restroom Restroom) {
	restrooms[restroom.GetId()] = restroom
}

func GenerateNextRestroomId() string {
	defer func() { nextRestroomId += 1 }()
	return strconv.Itoa(nextRestroomId)
}

func IsToiletExist(toiletId string) bool {
	_, isExist := toilets[toiletId]
	return isExist
}

func IsRestroomExist(restrooomId string) bool {
	_, isExist := restrooms[restrooomId]
	return isExist
}

func RemoveToilet(toiletId string, restroomId string) {
	delete(toilets, toiletId)
	restroom := restrooms[restroomId]
	restroom.RemoveIdFromToiletIdList(toiletId)
	restrooms[restroomId] = restroom
	undeployedToiletIdList = append(undeployedToiletIdList, toiletId)
}
