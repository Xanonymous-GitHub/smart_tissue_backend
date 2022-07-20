package model

var (
	Restrooms              map[string]Restroom
	Toilets                map[string]Toilet
	UndeployedToiletIdList []string
	NextRestroomId         int
)

func Setup() {
	Restrooms = map[string]Restroom{}
	Toilets = map[string]Toilet{}
	UndeployedToiletIdList = []string{}
	NextRestroomId = 1
}
