// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package server

import (
	"context"

	"go-travel/service/order/cmd/rpc/internal/logic"
	"go-travel/service/order/cmd/rpc/internal/svc"
	"go-travel/service/order/cmd/rpc/pb"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

// 民宿下订单
func (s *OrderServer) CreateHomestayOrder(ctx context.Context, in *pb.CreateHomestayOrderReq) (*pb.CreateHomestayOrderResp, error) {
	l := logic.NewCreateHomestayOrderLogic(ctx, s.svcCtx)
	return l.CreateHomestayOrder(in)
}

// 民宿订单详情
func (s *OrderServer) HomestayOrderDetail(ctx context.Context, in *pb.HomestayOrderDetailReq) (*pb.HomestayOrderDetailResp, error) {
	l := logic.NewHomestayOrderDetailLogic(ctx, s.svcCtx)
	return l.HomestayOrderDetail(in)
}

// 更新民宿订单状态
func (s *OrderServer) UpdateHomestayOrderTradeState(ctx context.Context, in *pb.UpdateHomestayOrderTradeStateReq) (*pb.UpdateHomestayOrderTradeStateResp, error) {
	l := logic.NewUpdateHomestayOrderTradeStateLogic(ctx, s.svcCtx)
	return l.UpdateHomestayOrderTradeState(in)
}

// 用户民宿订单
func (s *OrderServer) UserHomestayOrderList(ctx context.Context, in *pb.UserHomestayOrderListReq) (*pb.UserHomestayOrderListResp, error) {
	l := logic.NewUserHomestayOrderListLogic(ctx, s.svcCtx)
	return l.UserHomestayOrderList(in)
}

// provide to admin use
func (s *OrderServer) AddOrder(ctx context.Context, in *pb.AddOrderReq) (*pb.AddOrderResp, error) {
	l := logic.NewAddOrderLogic(ctx, s.svcCtx)
	return l.AddOrder(in)
}

func (s *OrderServer) DeleteOrder(ctx context.Context, in *pb.DeleteOrderReq) (*pb.DeleteOrderResp, error) {
	l := logic.NewDeleteOrderLogic(ctx, s.svcCtx)
	return l.DeleteOrder(in)
}

func (s *OrderServer) ListOrder(ctx context.Context, in *pb.ListOrderReq) (*pb.ListOrderResp, error) {
	l := logic.NewListOrderLogic(ctx, s.svcCtx)
	return l.ListOrder(in)
}

func (s *OrderServer) UpdateOrder(ctx context.Context, in *pb.UpdateOrderReq) (*pb.UpdateOrderResp, error) {
	l := logic.NewUpdateOrderLogic(ctx, s.svcCtx)
	return l.UpdateOrder(in)
}
