package album

import (
	"yuki-image/internal/db"
	"yuki-image/internal/model"
)

func SelectImage(id uint64, page uint64, size uint64) (model.ImageList, error) {
	dbiamges, err := db.SelectImageFromAlbum(id, page, size)
	if err != nil {
		return model.ImageList{}, err
	}
	var images []model.Image
	for _, v := range dbiamges {
		var image model.Image
		image.FromDBModel(v)
		images = append(images, image)
	}
	total, err := db.GetAlbumImageTotal(id)
	if err != nil {
		return model.ImageList{}, err
	}
	var imageList model.ImageList
	imageList.Image = images
	imageList.Total = total
	imageList.Page = page
	imageList.Size = size
	return imageList, nil
}

func SelectImageFromName(name string, page uint64, size uint64) (model.ImageList, error) {
	id, err := db.SelectAlbumIdFromName(name)
	if err != nil {
		return model.ImageList{}, err
	}
	return SelectImage(id, page, size)
}

func GetImageTotal(id uint64) (uint64, error) {
	return db.GetAlbumImageTotal(id)
}
