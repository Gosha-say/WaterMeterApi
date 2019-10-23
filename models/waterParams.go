package models

type WaterParams struct {
	Id     string
	WCold1 int64
	WCold2 int64
	WHot1  int64
	WHot2  int64
	Power  int64
	Date   int32
}

func (params WaterParams) Validate() (bool, string) {
	return true, ""
}
