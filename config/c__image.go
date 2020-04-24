package config

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/nfnt/resize"
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/database"
)

// Image ...
func (c *Config) Image(file *multipart.FileHeader, width, height uint) (string, []byte, string, int64) {

	f, err := file.Open()
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	name := file.Filename
	size1 := file.Size
	buffer := make([]byte, size1)

	f.Read(buffer)

	types := http.DetectContentType(buffer)
	img, size := size(buffer, width, height, types)
	return name, img, types, size
}

// ShowImage ...
func (c *Config) ShowImage(ctx *atreugo.RequestCtx) error {
	var b database.Images
	id := ctx.UserValue("id")

	if err := db.Where("UID = ?", id.(string)).First(&b).Error; err != nil {
		return err
	}
	n := int64(b.Size)
	Length := strconv.FormatInt(n, 16)
	ctx.Response.Header.Set("Content-Type", b.Type)
	ctx.Response.Header.Set("Content-Length", Length)
	ctx.Write(b.Body)
	return nil
}

func size(imgByte []byte, width, height uint, types string) ([]byte, int64) {
	img, _, _ := image.Decode(bytes.NewReader(imgByte))
	m := resize.Resize(width, height, img, resize.Lanczos3)
	buf := new(bytes.Buffer)
	if types == "image/jpeg" {
		jpeg.Encode(buf, m, nil)
	}
	if types == "image/png" {
		png.Encode(buf, m)
	}
	if types == "image/gif" {
		gif.Encode(buf, m, nil)
	}
	size := int64(len(buf.Bytes()))

	return buf.Bytes(), size
}
