package model

var (
	RestroomList   map[string]Restroom
	NextRestroomId int
)

func Setup() {
	RestroomList = map[string]Restroom{}
	NextRestroomId = 1
}
