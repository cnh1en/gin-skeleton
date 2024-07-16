package global

import (
	"github.com/cnh1en/lets_go/pkg/setting"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *zap.Logger
	DB     *gorm.DB
	Rdb    *redis.Client
	Server setting.SeverSetting
)
