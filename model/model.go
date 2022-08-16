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

func GetRestroom(restroomId string) Restroom {
	return restrooms[restroomId]
}

func GetToiletsFromRestroom(restroomId string) []Toilet {
	restroomToilets := []Toilet{}
	restroom := restrooms[restroomId]
	toiletIdList := restroom.ToiletIdList

	for _, toiletId := range toiletIdList {
		restroomToilets = append(restroomToilets, toilets[toiletId])
	}

	return restroomToilets
}

func GetSingleToilet(toiletId string) Toilet {
	return toilets[toiletId]
}

func GetUndeployedToiletIdList() []string {
	return undeployedToiletIdList
}

func IsRestroomExists(restroomId string) bool {
	_, isExist := restrooms[restroomId]
	return isExist
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

func IsToiletIdInRestroom(toiletId string, restrooomId string) bool {
	restroom := restrooms[restrooomId]
	for _, currentToiletId := range restroom.ToiletIdList {
		if currentToiletId == toiletId {
			return true
		}
	}
	return false
}

func GenerateNextRestroomId() string {
	defer func() { nextRestroomId += 1 }()
	return strconv.Itoa(nextRestroomId)
}

func RegisterRestroom(restroom Restroom) {
	restrooms[restroom.Id] = restroom
}

func RegisterToilet(toiletId string, restroomId string, location string) {
	restroom := restrooms[restroomId]
	restroom.ToiletIdList = append(restroom.ToiletIdList, toiletId)
	restrooms[restroomId] = restroom
	RemoveIdFromUndeployedToiletIds(toiletId)

	toilet := toilets[toiletId]
	toilet.Location = location
	toilets[toilet.Id] = toilet
}

func UpdateRestroomLocation(restroomId string, location string) {
	restroom := restrooms[restroomId]
	restroom.Location = location
	restrooms[restroomId] = restroom
}

func UpdateToiletData(toilet Toilet) {
	currentToilet := toilets[toilet.Id]
	toilet.Distance = currentToilet.Distance
	toilet.Percentage = math.Round(toilet.Distance/toilet.MaxDistance*10000) / 10000

	toilets[toilet.Id] = toilet
}

func UploadTissueBoxData(toilet Toilet) {
	currentToilet, isExist := toilets[toilet.Id]

	if isExist {
		toilet.Location = currentToilet.Location
	} else {
		undeployedToiletIdList = append(undeployedToiletIdList, toilet.Id)
	}

	toilet.Percentage = math.Round(toilet.Distance/toilet.MaxDistance*10000) / 10000

	toilets[toilet.Id] = toilet
}

func DeleteRestroom(restroomId string) {
	undeployedToiletIdList = append(undeployedToiletIdList, restrooms[restroomId].ToiletIdList...)
	delete(restrooms, restroomId)
}

func RemoveIdFromUndeployedToiletIds(toiletId string) {
	for toiletIndex, currentToiletId := range undeployedToiletIdList {
		if currentToiletId == toiletId {
			undeployedToiletIdList[toiletIndex] = undeployedToiletIdList[len(undeployedToiletIdList)-1]
			undeployedToiletIdList = undeployedToiletIdList[:len(undeployedToiletIdList)-1]
		}
	}
}

func RemoveToilet(toiletId string, restroomId string) {
	restroom := restrooms[restroomId]
	for toiletIndex, currentToiletId := range restroom.ToiletIdList {
		if currentToiletId == toiletId {
			restroom.ToiletIdList[toiletIndex] = restroom.ToiletIdList[len(restroom.ToiletIdList)-1]
			restroom.ToiletIdList = restroom.ToiletIdList[:len(restroom.ToiletIdList)-1]
		}
	}
	restrooms[restroomId] = restroom
	undeployedToiletIdList = append(undeployedToiletIdList, toiletId)
}
