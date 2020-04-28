package config

import (
	"log"
	"strconv"

	"github.com/savsgio/atreugo/v10"
	"google.golang.org/protobuf/proto"
)

// PBdata person.
func (c *Config) PBdata(ctx *atreugo.RequestCtx, code int, pr proto.Message) error {
	data, err := proto.Marshal(pr)
	if err != nil {
		log.Fatal("marsaling err", err)
	}

	n := int64(len(data))
	Length := strconv.FormatInt(n, 16)
	ctx.SetStatusCode(code)
	ctx.Response.Header.Set("Content-Type", "application/protobuf")
	ctx.Response.Header.Set("Content-Length", Length)
	_, err = ctx.Write(data)

	return err
}
