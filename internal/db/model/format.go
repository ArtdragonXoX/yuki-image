package model

type Format struct {
	Id            uint64   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string   `gorm:"type:varchar(20);not null;unique" json:"name"`
	FormatSupport []*Album `gorm:"many2many:format_supports;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;association_jointable_foreignkey:album_id;jointable_foreignkey:format_id"`
}
