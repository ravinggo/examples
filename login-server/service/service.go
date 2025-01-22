package service

import (
	"reflect"
	"time"

	"github.com/ravinggo/game/common/berror"
	"github.com/ravinggo/game/common/ctx"
	"github.com/ravinggo/game/common/handler"
	"github.com/ravinggo/game/common/logger"
	"github.com/ravinggo/game/common/natsclient"
	"github.com/ravinggo/game/common/service"

	"github.com/ravinggo/examples/errmsg"
	"github.com/ravinggo/examples/pbmsg/game"
	"github.com/ravinggo/examples/pbmsg/login"
)

type LoginServer struct {
	*service.BaseService[ctx.IntTrace, *ctx.IntTrace]
}

func NewLoginServer(
	natsUrls []string,
	lockQueueThread bool,
	hashMode service.HashRunMode,
	taskMode service.TaskRunMode,
	rpcTimeout time.Duration,
) *LoginServer {
	return &LoginServer{
		BaseService: service.NewBaseService[ctx.IntTrace, *ctx.IntTrace](
			natsUrls, lockQueueThread, hashMode, taskMode, rpcTimeout,
		),
	}
}

func (s *LoginServer) Start() {
	s.Router()
	s.BaseService.Start(
		func(e any) {
			logger.Log.Warn().Str("eventName", reflect.TypeOf(e).String()).Msg("unknown event")
		},
	)
}

func (s *LoginServer) Router() {
	handler.RegisterRPCResp(s.GetHandler(), "login request", s.Login)
}

func (s *LoginServer) Stop() {
	s.BaseService.Stop()
}

func (s *LoginServer) Login(c *ctx.Int64TraceCtx, req *login.LoginReq, res *login.LoginRsp) *berror.ErrMsg {
	if req.Username == "" {
		return errmsg.NewLoginErrorAccountIsWrong()
	}
	if req.Password == "" {
		return errmsg.NewLoginErrorPasswordFormatError()
	}

	rpc := natsclient.NewClusterRequest[game.LoginSuccessReq, game.LoginSuccessRsp]()
	defer rpc.Free()
	rpc.Req.UserName = req.Username
	err := rpc.Request(s.GetNatsCluster(), c)
	if err != nil {
		return err
	}

	if rpc.Resp.UserName == req.Username {
		res.LoginResult = 1
	}
	return nil
}
