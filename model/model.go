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

func GetUndeployedToiletIdList() []string {
	return undeployedToiletIdList
}

func RegisterRestroom(restroom Restroom) {
	restrooms[restroom.GetId()] = restroom
}

func GenerateNextRestroomId() string {
	defer func() { nextRestroomId += 1 }()
	return strconv.Itoa(nextRestroomId)
}
