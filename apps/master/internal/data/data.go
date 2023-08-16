package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"prometheus-manager/pkg/conn"

	"prometheus-manager/apps/master/internal/biz"
	"prometheus-manager/apps/master/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewPingRepo,
	wire.Bind(new(biz.IPingRepo), new(*PingRepo)),
	NewPushRepo,
	wire.Bind(new(biz.IPushRepo), new(*PushRepo)),
	NewPromV1Repo,
	wire.Bind(new(biz.IPromV1Repo), new(*PromV1Repo)),
	NewAlarmPageV1Repo,
	wire.Bind(new(biz.IAlarmPageV1Repo), new(*AlarmPageV1Repo)),
	NewDictRepo,
	wire.Bind(new(biz.IDictV1Repo), new(*DictV1Repo)),
)

type TransactionFunc func(tx *gorm.DB) error

type ITransaction interface {
	Transaction(fn TransactionFunc) error
}

// Data .
type Data struct {
	db    *gorm.DB
	cache *redis.Client
}

const (
	promModuleName      = "data/Prom"
	alarmPageModuleName = "data/AlarmPage"
	dictModuleName      = "data/Dict"
	pushModuleName      = "data/Push"
)

func (l *Data) DB() *gorm.DB {
	return l.db
}

func (l *Data) Transaction(fn TransactionFunc) error {
	return fn(l.db)
}

type ICrud interface {
	ITransaction
	DB() *gorm.DB
}

var _ ITransaction = (*Data)(nil)
var _ ICrud = (*Data)(nil)

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := conn.NewMysqlDB(c.GetDatabase(), logger)
	if err != nil {
		log.NewHelper(logger).Errorf("failed opening connection to mysql: %v", err)
		return nil, nil, err
	}

	cache := conn.NewRedisClient(c.GetRedis())

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		// 关闭数据库连接
		s, err := db.DB()
		if err != nil {
			log.NewHelper(logger).Errorf("failed to close data resources: %v", err)
			return
		}

		if err := s.Close(); err != nil {
			log.NewHelper(logger).Errorf("close mysql error: %v", err)
		}

		if err := cache.Close(); err != nil {
			log.NewHelper(logger).Errorf("failed to close data resources: %v", err)
		}
	}
	return &Data{
		db:    db,
		cache: cache,
	}, cleanup, nil
}
