package initializer

import (
	"buyleaf/config"
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"log"
)

var Cld *cloudinary.Cloudinary
var CldCtx context.Context

func InitCLD() {
	var err error
	Cld, err = cloudinary.NewFromParams(config.CloudinaryCloudName, config.CloudinaryApiKey, config.CloudinaryApiSecret)
	CldCtx = context.Background()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
