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

func GetMultipleToilets(id string) []Toilet {
	multipleToilets := []Toilet{}
	restroom := restrooms[id]
	toiletIdList := restroom.GetToiletIdList()

	for _, toiletId := range toiletIdList {
		multipleToilets = append(multipleToilets, toilets[toiletId])
	}

	return multipleToilets
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

func IsRestroomExists(id string) bool {
	_, exists := restrooms[id]
	return exists
}

func UpdateRestroomLocation(id string, location string) {
	restroom := restrooms[id]
	restroom.SetLocation(location)
	restrooms[id] = restroom
}

func DeleteRestroom(id string) {
	delete(restrooms, id)
}
