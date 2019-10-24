package models

type WaterParams struct {
	Id     string `db:"MeterId"`
	WCold1 int64  `db:"WCold1"`
	WCold2 int64  `db:"WCold2"`
	WHot1  int64  `db:"WHot1"`
	WHot2  int64  `db:"WHot2"`
	Power  int64  `db:"Power"`
	Date   int32  `db:"Date"`
}

func (params WaterParams) Validate() (bool, string) {
	return true, ""
}
