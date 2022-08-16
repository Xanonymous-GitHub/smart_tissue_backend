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
	nextRestroomId = 4
	restrooms["1"] = Restroom{Id: "1", Location: "科研12樓", ToiletIdList: []string{"e8e3-45c2-b9b3", "3597-47c8-91d1", "0fc6-49f6-b55a"}}
	toilets["e8e3-45c2-b9b3"] = Toilet{Id: "e8e3-45c2-b9b3", Percentage: 20, Distance: 8, MaxDistance: 10, Location: "科研12樓-第1間", State: "sufficient"}
	toilets["3597-47c8-91d1"] = Toilet{Id: "3597-47c8-91d1", Percentage: 60, Distance: 8, MaxDistance: 20, Location: "科研12樓-第2間", State: "sufficient"}
	toilets["0fc6-49f6-b55a"] = Toilet{Id: "0fc6-49f6-b55a", Percentage: 5, Distance: 9.5, MaxDistance: 10, Location: "科研12樓-第3間", State: "insufficient"}
	restrooms["2"] = Restroom{Id: "2", Location: "科研13樓", ToiletIdList: []string{"527f-427b-894e", "8abd-4680-9214", "bbab-4e7b-92ce"}}
	toilets["527f-427b-894e"] = Toilet{Id: "527f-427b-894e", Percentage: 50, Distance: 10, MaxDistance: 20, Location: "科研13樓-第1間", State: "sufficient"}
	toilets["8abd-4680-9214"] = Toilet{Id: "8abd-4680-9214", Percentage: 100, Distance: 10, MaxDistance: 10, Location: "科研13樓-第2間", State: "sufficient"}
	toilets["bbab-4e7b-92ce"] = Toilet{Id: "bbab-4e7b-92ce", Percentage: 0, Distance: 0, MaxDistance: 10, Location: "科研13樓-第3間", State: "insufficient"}
	restrooms["3"] = Restroom{Id: "3", Location: "科研14樓", ToiletIdList: []string{"5e44-41e4-aa4e", "abc0-4ee0-b2b2", "a693-4a8f-9098"}}
	toilets["5e44-41e4-aa4e"] = Toilet{Id: "5e44-41e4-aa4e", Percentage: 25, Distance: 7.5, MaxDistance: 10, Location: "科研14樓-第1間", State: "sufficient"}
	toilets["abc0-4ee0-b2b2"] = Toilet{Id: "abc0-4ee0-b2b2", Percentage: 75, Distance: 2.5, MaxDistance: 10, Location: "科研14樓-第2間", State: "sufficient"}
	toilets["a693-4a8f-9098"] = Toilet{Id: "a693-4a8f-9098", Percentage: 20, Distance: 16, MaxDistance: 20, Location: "科研14樓-第3間", State: "sufficient"}
	undeployedToiletIdList = append(undeployedToiletIdList, "2f8d-4ae8-8b8a", "90d3-4b14-919f", "221e-4b68-a054")
	toilets["2f8d-4ae8-8b8a"] = Toilet{Id: "2f8d-4ae8-8b8a", Percentage: 90, Distance: 1, MaxDistance: 10, State: "sufficient"}
	toilets["90d3-4b14-919f"] = Toilet{Id: "90d3-4b14-919f", Percentage: 15, Distance: 17, MaxDistance: 20, State: "sufficient"}
	toilets["221e-4b68-a054"] = Toilet{Id: "221e-4b68-a054", Percentage: 5, Distance: 19, MaxDistance: 20, State: "insufficient"}
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
