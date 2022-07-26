package rowdata

import "time"

type RowData struct {
	Uuid     string    `json:"uuid"     bson:"uuid"    `
	Parentid string    `json:"parentid" bson:"parentid" validate:"required"`
	Comment  string    `json:"comment"  bson:"comment"  validate:"required"`
	Author   string    `json:"author"   bson:"author"   validate:"required"`
	Update   time.Time `json:"update"   bson:"update"   validate:"required"`
	Favorite bool      `json:"favorite" bson:"favorite"`
}
