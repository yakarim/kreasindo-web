package controller

import (
	"fmt"
	"log"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/model"
	"github.com/yakarim/kreasindo-web/protobuf"
)

// Person ...
type Person struct {
	config.Config
	model.Model
}

// View person.
func (c *Person) View(ctx *atreugo.RequestCtx) error {
	data, _ := c.Contact.Query()
	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "person/view", config.H{
		"title":    "Contact",
		"username": string(u.Username),
		"signIn":   signIn,
		"data":     data,
	})
}

// Create person.
func (c *Person) Create(ctx *atreugo.RequestCtx) error {
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Fatal("uid error", err)
	}
	pr := &protobuf.Person{
		Name:     "walid",
		Email:    "string",
		Password: "string",
		Base: &protobuf.Base{
			Uid: uid.String(),
			//CreatedAt: time.Now(),
		},
	}

	db.Save(&pr)
	data, err := proto.Marshal(pr)
	if err != nil {
		log.Fatal("marsaling err", err)
	}
	fmt.Println(data)

	n := int64(len(data))
	Length := strconv.FormatInt(n, 16)
	ctx.Response.Header.Set("Content-Type", "application/x-protobuf")
	ctx.Response.Header.Set("Content-Length", Length)
	_, err = ctx.Write(data)
	return err
}
