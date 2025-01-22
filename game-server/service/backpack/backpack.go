package backpack

import (
	"github.com/ravinggo/game/common/berror"
	"github.com/ravinggo/game/common/ctx"
	"github.com/ravinggo/game/common/handler"
	"github.com/ravinggo/game/common/localevent"
	"github.com/ravinggo/game/common/service"

	"github.com/ravinggo/examples/errmsg"
	"github.com/ravinggo/examples/game-server/event"
	"github.com/ravinggo/examples/pbmsg/game"
)

type BackPack struct {
	svc *service.BaseService[ctx.IntTrace, *ctx.IntTrace]
}

func NewBackpack(
	svc *service.BaseService[ctx.IntTrace, *ctx.IntTrace],
) *BackPack {
	return &BackPack{
		svc: svc,
	}
}

func (bp *BackPack) Router() {
	handler.RegisterRPCResp(bp.svc.GetHandler(), "GetBackpack", bp.GetBackpack)
	handler.RegisterRPCResp(bp.svc.GetHandler(), "AddItem", bp.AddItem)
	handler.RegisterRPCResp(bp.svc.GetHandler(), "SubItem", bp.SubItem)
}

func (bp *BackPack) GetBackpack(c *ctx.Int64TraceCtx, req *game.GetBackpackReq, resp *game.GetBackpackRsp) *berror.ErrMsg {
	items := make([]*game.Item, 0, 4)
	items = append(
		items,
		&game.Item{
			Id:    1,
			Cid:   1,
			Value: 9999,
		},
		&game.Item{
			Id:    2,
			Cid:   2,
			Value: 9999,
		},
		&game.Item{
			Id:    3,
			Cid:   3,
			Value: 9999,
		},
		&game.Item{
			Id:    4,
			Cid:   3,
			Value: 9999,
		},
	)
	resp.Items = items
	return nil
}

func (bp *BackPack) AddItem(c *ctx.Int64TraceCtx, req *game.AddItemReq, resp *game.AddItemRsp) *berror.ErrMsg {
	before := int64(9999)
	after := 9999 + req.Value
	resp.Item = &game.Item{
		Id:    req.Cid,
		Cid:   req.Cid,
		Value: after,
	}
	return localevent.Call(
		c, event.ItemChangeEvent{
			Id:           req.Cid,
			Cid:          req.Cid,
			ChangeBefore: before,
			ChangeAfter:  after,
		},
	)
}

func (bp *BackPack) SubItem(c *ctx.Int64TraceCtx, req *game.SubItemReq, resp *game.SubItemRsp) *berror.ErrMsg {
	if req.Value < -9999 {
		return errmsg.NewBackpackErrorItemNotEnough()
	}
	before := int64(9999)
	after := 9999 - req.Value
	resp.Item = &game.Item{
		Id:    req.Id,
		Cid:   req.Id,
		Value: after,
	}
	return localevent.Call(
		c, event.ItemChangeEvent{
			Id:           req.Id,
			Cid:          req.Id,
			ChangeBefore: before,
			ChangeAfter:  after,
		},
	)
}
