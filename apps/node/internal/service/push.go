package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "prometheus-manager/api/strategy/v1/push"
)

type (
	IPushLogic interface {
		Strategies(ctx context.Context, req *pb.StrategiesRequest) (*pb.StrategiesReply, error)
	}

	PushService struct {
		pb.UnimplementedPushServer

		logger *log.Helper
		logic  IPushLogic
	}
)

var _ pb.PushServer = (*PushService)(nil)

func NewPushService(logic IPushLogic, logger log.Logger) *PushService {
	return &PushService{logic: logic, logger: log.NewHelper(log.With(logger, "module", "service/Push"))}
}

func (l *PushService) Strategies(ctx context.Context, req *pb.StrategiesRequest) (*pb.StrategiesReply, error) {
	l.logger.Debugf("Strategies req: %v", req)
	return l.logic.Strategies(ctx, req)
}