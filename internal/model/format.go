package model

import dbModel "yuki-image/internal/db/model"

type Format struct {
	Id   uint64 `json:"id"`
	Name string `json:"name,omitempty"`
}

func (f *Format) ToDBModel() dbModel.Format {
	return dbModel.Format{Id: f.Id, Name: f.Name}
}

func (f *Format) FromDBModel(model dbModel.Format) {
	f.Id = model.Id
	f.Name = model.Name
}
