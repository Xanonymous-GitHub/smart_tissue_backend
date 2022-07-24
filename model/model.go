package model

import "strconv"

var (
	restrooms              map[string]Restroom
	toilets                map[string]Toilet
	undeployedToiletIdList []string
	nextRestroomId         int
	nextToiletId           int
)

func Setup() {
	restrooms = map[string]Restroom{}
	toilets = map[string]Toilet{}
	undeployedToiletIdList = []string{}
	nextRestroomId = 1
	nextToiletId = 1
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

func RegisterToilet(toilet Toilet) {
	toilets[toilet.GetId()] = toilet
	undeployedToiletIdList = append(undeployedToiletIdList, toilet.GetId())
}

func GenerateNextToiletId() string {
	defer func() { nextToiletId += 1 }()
	return strconv.Itoa(nextToiletId)
}
