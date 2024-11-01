package format

import (
	"yuki-image/internal/db"
	"yuki-image/internal/model"
)

func Insert(format model.Format) (uint64, error) {
	return db.InsertFormat(format.ToDBModel())
}
