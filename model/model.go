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

func IsToiletExists(toiletId string) bool {
	_, isExist := toilets[toiletId]
	return isExist
}

func IsRestroomExists(restrooomId string) bool {
	_, isExist := restrooms[restrooomId]
	return isExist
}

func IsToiletIdInRestroom(toiletId string, restrooomId string) bool {
	restroom := restrooms[restrooomId]
	return restroom.IsToiletIdInList(toiletId)
}

func RemoveToilet(toiletId string, restroomId string) {
	restroom := restrooms[restroomId]
	restroom.RemoveIdFromToiletIdList(toiletId)
	restrooms[restroomId] = restroom
	undeployedToiletIdList = append(undeployedToiletIdList, toiletId)
}
