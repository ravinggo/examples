package service

import (
	"reflect"

	"github.com/ravinggo/game/common/berror"
	"github.com/ravinggo/game/common/ctx"
	"github.com/ravinggo/game/common/handler"
	"github.com/ravinggo/game/common/localevent"
	"github.com/ravinggo/game/common/logger"
	"github.com/ravinggo/game/common/service"

	"github.com/ravinggo/examples/game-server/event"
	"github.com/ravinggo/examples/game-server/service/backpack"
	"github.com/ravinggo/examples/pbmsg/game"
)

type GameServer struct {
	*service.BaseService[ctx.IntTrace, *ctx.IntTrace]
}

func NewGameServer(
	natsUrls []string,
) *GameServer {
	return &GameServer{
		BaseService: service.NewBaseService[ctx.IntTrace](
			natsUrls,
		),
	}
}

func (s *GameServer) Start() {
	s.Router()
	s.BaseService.Start(
		func(e any) {
			logger.Log.Warn().Str("eventName", reflect.TypeOf(e).String()).Msg("unknown event")
		},
	)
}

func (s *GameServer) Router() {
	handler.RegisterRPC(s.GetHandler(), "login server login success request", s.LoginSuccess)
	handler.RegisterRPCResp(s.GetHandler(), "game server echo test", s.Echo)
	localevent.Register("GameServer ItemChange", s.ItemChange)

	backpack.NewBackpack(s.BaseService).Router()
}

func (s *GameServer) Stop() {
	s.BaseService.Stop()
}

func (s *GameServer) LoginSuccess(c *ctx.Int64TraceCtx, req *game.LoginSuccessReq) (*game.LoginSuccessRsp, *berror.ErrMsg) {
	return &game.LoginSuccessRsp{UserName: req.UserName}, nil
}

func (s *GameServer) Echo(c *ctx.Int64TraceCtx, req *game.EchoReq, resp *game.EchoRsp) *berror.ErrMsg {
	resp.Msg = req.Msg
	return nil
}

func (s *GameServer) ItemChange(c *ctx.Int64TraceCtx, e event.ItemChangeEvent) *berror.ErrMsg {
	logger.Log.Info().Int64("beforeCount", e.ChangeBefore).Int64("afterCount", e.ChangeAfter).Msg("GameServer listen item change")
	return nil
}
