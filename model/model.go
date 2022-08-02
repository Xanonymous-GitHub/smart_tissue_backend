package model

import (
	"math"
	"strconv"
)

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

func GetSingleToilet(toiletId string) Toilet {
	return toilets[toiletId]
}

func GetToiletsFromRestroom(id string) []Toilet {
	restroomToilets := []Toilet{}
	restroom := restrooms[id]
	toiletIdList := restroom.ToiletIdList

	for _, toiletId := range toiletIdList {
		restroomToilets = append(restroomToilets, toilets[toiletId])
	}

	return restroomToilets
}

func GetUndeployedToiletIdList() []string {
	return undeployedToiletIdList
}

func RegisterRestroom(restroom Restroom) {
	restrooms[restroom.Id] = restroom
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
	restroom.Location = location
	restrooms[id] = restroom
}

func DeleteRestroom(id string) {
	delete(restrooms, id)
}

func UploadTissueBoxData(toilet Toilet) {
	currentToilet, exists := toilets[toilet.Id]

	if exists {
		toilet.Location = currentToilet.Location
	} else {
		undeployedToiletIdList = append(undeployedToiletIdList, toilet.Id)
	}

	toilet.Percentage = math.Round(toilet.Distance/toilet.MaxDistance*10000) / 10000

	toilets[toilet.Id] = toilet
}
