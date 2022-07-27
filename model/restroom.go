package model

type Restroom struct {
	Id           string
	Location     string
	ToiletIdList []string
}

func (restroom *Restroom) GetId() string {
	return restroom.Id
}

func (restroom *Restroom) GetToiletIdList() []string {
	return restroom.ToiletIdList
}
  
func (restroom *Restroom) GetLocation() string {
	return restroom.Location
}

func (restroom *Restroom) SetLocation(location string) {
	restroom.Location = location
}
