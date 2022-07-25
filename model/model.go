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
	for i := 0; i < len(undeployedToiletIdList); i++ {
		if undeployedToiletIdList[i] == toiletId {
			return true
		}
	}
	return false
}

func IsRestroomExist(restrooomId string) bool {
	_, isExist := restrooms[restrooomId]
	return isExist
}

func RegisterToilet(toilet Toilet, restroomId string) {
	toilets[toilet.GetId()] = toilet
	restroom := restrooms[restroomId]
	restroom.AddToiletId(toilet.GetId())
	restrooms[restroom.GetId()] = restroom
	undeployedToiletIdList = RemoveIdFromUndeployedToiletId(toilet.GetId())
}

func RemoveIdFromUndeployedToiletId(toiletId string) []string {
	for i := 0; i < len(undeployedToiletIdList); i++ {
		if undeployedToiletIdList[i] == toiletId {
			undeployedToiletIdList[i] = undeployedToiletIdList[len(undeployedToiletIdList) - 1]
			return undeployedToiletIdList[:len(undeployedToiletIdList) - 1]
		}
	}
	return undeployedToiletIdList
}
