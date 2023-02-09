package config

import (
	"os"
	"time"
)

var JwtTokenExpiration time.Duration = 10
var JwtRefreshTokenExpiration time.Duration = 720

var CloudinaryCloudName = os.Getenv("cloudinary_cloud_name")
var CloudinaryApiKey = os.Getenv("cloudinary_api_key")
var CloudinaryApiSecret = os.Getenv("cloudinary_api_secret")
