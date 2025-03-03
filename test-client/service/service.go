package service

import (
	"time"

	"github.com/ravinggo/game/common/berror"
	"github.com/ravinggo/game/common/ctx"
	"github.com/ravinggo/game/common/natsclient"
	"github.com/ravinggo/objectpool"

	"github.com/ravinggo/examples/pbmsg/game"
	"github.com/ravinggo/examples/pbmsg/login"
	"github.com/ravinggo/examples/test-client/stat"
)

type Client struct {
	cnc   *natsclient.ClusterClient
	stats *stat.Stat
}

func NewClient(cnc *natsclient.ClusterClient, stats *stat.Stat) *Client {
	return &Client{
		// cnc: natsclient.NewClusterClient(env.GetConfig().ServerType, urls, time.Second*10),
		cnc:   cnc,
		stats: stats,
	}
}

func (client *Client) LoginReq() *berror.ErrMsg {
	now := time.Now()
	defer func() {
		client.stats.Add(time.Since(now))
	}()
	c := objectpool.Get[ctx.Int64TraceCtx]()
	defer objectpool.Put[ctx.Int64TraceCtx](c)
	c.TD.TraceId = "123XASDASDASD"
	rpc := natsclient.NewClusterRequest[login.LoginReq, login.LoginRsp]()
	defer rpc.Free()
	rpc.Req.Username = "ravinggo"
	rpc.Req.Password = "123456"
	err := rpc.Request(client.cnc, c)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) Echo() *berror.ErrMsg {
	now := time.Now()
	defer func() {
		client.stats.Add(time.Since(now))
	}()
	c := objectpool.Get[ctx.Int64TraceCtx]()
	defer objectpool.Put[ctx.Int64TraceCtx](c)
	c.TD.TraceId = "123XASDASDASD"
	rpc := natsclient.NewClusterRequest[game.EchoReq, game.EchoRsp]()
	defer rpc.Free()
	const msg = "asdasdasda"
	rpc.Req.Msg = msg
	err := rpc.Request(client.cnc, c)
	if err != nil {
		return err
	}
	if rpc.Resp.Msg != msg {
		return berror.NewProtocolStr("rpc.Resp.Msg != " + msg)
	}
	return nil
}
