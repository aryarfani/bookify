package helpers

import "github.com/gin-gonic/gin"

func StoreImage(c *gin.Context, imageName string) string {
	var imagePath string

	form, _ := c.MultipartForm()
	if form != nil && len(form.File) > 0 {
		image := form.File[imageName][0]
		if image != nil {
			c.SaveUploadedFile(image, "./images/"+image.Filename)
			imagePath = "images/" + image.Filename
		}
	}

	return imagePath
}
