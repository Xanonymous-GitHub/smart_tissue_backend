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

func GetToilet(id string) Toilet {
	return toilets[id]
}

func IsUndeployedToiletExist(toiletId string) bool {
	for _, currentToiletId := range undeployedToiletIdList {
		if currentToiletId == toiletId {
			return true
		}
	}
	return false
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

func RegisterToilet(toiletId string, restroomId string) {
	restroom := restrooms[restroomId]
	restroom.AddToiletId(toiletId)
	restrooms[restroomId] = restroom
	undeployedToiletIdList = RemoveIdFromUndeployedToiletId(toiletId)
}

func RemoveIdFromUndeployedToiletId(toiletId string) []string {
	for toiletIndex, currentToiletId := range undeployedToiletIdList {
		if currentToiletId == toiletId {
			undeployedToiletIdList[toiletIndex] = undeployedToiletIdList[len(undeployedToiletIdList)-1]
			return undeployedToiletIdList[:len(undeployedToiletIdList)-1]
		}
	}
	return undeployedToiletIdList
}

func RemoveToilet(toiletId string, restroomId string) {
	restroom := restrooms[restroomId]
	restroom.RemoveIdFromToiletIdList(toiletId)
	restrooms[restroomId] = restroom
	undeployedToiletIdList = append(undeployedToiletIdList, toiletId)
}

func UpdateToiletData(toilet Toilet) {
	toilets[toilet.GetId()] = toilet
}
