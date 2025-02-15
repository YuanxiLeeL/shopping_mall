package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	RedisDB  *redis.Client
	Enforcer *casbin.Enforcer
	Validate *validator.Validate
)
