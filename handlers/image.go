package handlers

func UploadImage(filepath string, album uint64) error {

	return nil
	// 保存文件到临时位置
	// dst := fmt.Sprintf("tmp/%s", filepath)

	// err := image.ManipulateJPEG("tmp/"+file.Filename, conf.Conf.Server.Path+"/"+file.Filename, 1080, 1080)
	// if err != nil {
	// 	ctx.JSON(400, gin.H{"error": err})
	// }
	// ctx.JSON(200, gin.H{"url": "http://localhost:" + conf.Conf.Server.Port + "/" + "i/" + file.Filename})
}
