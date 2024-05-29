package model

type JwtBlacklist struct {
	BASE_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
