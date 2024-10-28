package model

type Album struct {
	Id            uint64    `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	MaxHeight     uint64    `json:"max_height,omitempty"`
	MaxWidth      uint64    `json:"max_width,omitempty"`
	FormatSupport []Format  `json:"format_support,omitempty"`
	UpdateTime    string    `json:"update_time,omitempty"`
	CreateTime    string    `json:"create_time,omitempty"`
	Image         ImageList `json:"image,omitempty"`
}

type ImageList struct {
	Image []Image `json:"image,omitempty"`
	Page  uint64  `json:"page,omitempty"`
	Size  uint64  `json:"size,omitempty"`
	Total uint64  `json:"total"`
}
