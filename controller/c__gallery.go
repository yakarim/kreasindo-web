package controller

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/model"
	"github.com/yakarim/kreasindo-web/protobuf"
	"google.golang.org/protobuf/proto"
)

// Gallery ...
type Gallery struct {
	config.Config
	model.Model
}

// GalleryDepan ...
func (c *Gallery) GalleryDepan(ctx *atreugo.RequestCtx) error {
	data, err := c.Gallery.Query()
	if err != nil {
		log.Fatal("gallery err", err)
	}
	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "pages/gallery", config.H{
		"title":    "Gallery",
		"username": string(u.Username),
		"signIn":   signIn,
		"data":     data,
	})
}

// View person.
func (c *Gallery) View(ctx *atreugo.RequestCtx) error {

	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "gallery/view", config.H{
		"title":    "Gallery",
		"username": string(u.Username),
		"signIn":   signIn,
	})
}

// JSON gallery.
func (c *Gallery) JSON(ctx *atreugo.RequestCtx) error {

	data, err := c.Gallery.Query()
	if err != nil {
		log.Fatal("JSON gallery error", err)
	}

	b, _ := proto.Marshal(data)
	fmt.Println(b)

	m := jsonpb.Marshaler{}

	return m.Marshal(ctx, data)

}

// Create person.
func (c *Gallery) Create(ctx *atreugo.RequestCtx) error {
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Fatal("uid error", err)
	}
	pr := &protobuf.Gallery{
		Uid:  uid.String(),
		Name: "walid",
		Desc: "string",
	}

	db.Save(&pr)

	return nil
}
