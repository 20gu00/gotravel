// Code generated by goctl. DO NOT EDIT!
// Source: payment.proto

package server

import (
	"context"

	"go-travel/service/payment/cmd/rpc/internal/logic"
	"go-travel/service/payment/cmd/rpc/internal/svc"
	"go-travel/service/payment/cmd/rpc/pb"
)

type PaymentServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedPaymentServer
}

func NewPaymentServer(svcCtx *svc.ServiceContext) *PaymentServer {
	return &PaymentServer{
		svcCtx: svcCtx,
	}
}

// 创建微信支付预处理订单
func (s *PaymentServer) CreatePayment(ctx context.Context, in *pb.CreatePaymentReq) (*pb.CreatePaymentResp, error) {
	l := logic.NewCreatePaymentLogic(ctx, s.svcCtx)
	return l.CreatePayment(in)
}

// 根据sn查询流水记录
func (s *PaymentServer) GetPaymentBySn(ctx context.Context, in *pb.GetPaymentBySnReq) (*pb.GetPaymentBySnResp, error) {
	l := logic.NewGetPaymentBySnLogic(ctx, s.svcCtx)
	return l.GetPaymentBySn(in)
}

// 更新交易状态
func (s *PaymentServer) UpdateTradeState(ctx context.Context, in *pb.UpdateTradeStateReq) (*pb.UpdateTradeStateResp, error) {
	l := logic.NewUpdateTradeStateLogic(ctx, s.svcCtx)
	return l.UpdateTradeState(in)
}

// 根据订单sn查询流水记录
func (s *PaymentServer) GetPaymentSuccessRefundByOrderSn(ctx context.Context, in *pb.GetPaymentSuccessRefundByOrderSnReq) (*pb.GetPaymentSuccessRefundByOrderSnResp, error) {
	l := logic.NewGetPaymentSuccessRefundByOrderSnLogic(ctx, s.svcCtx)
	return l.GetPaymentSuccessRefundByOrderSn(in)
}

// provide to admin use
func (s *PaymentServer) ListPayment(ctx context.Context, in *pb.ListPaymentReq) (*pb.ListPaymentResp, error) {
	l := logic.NewListPaymentLogic(ctx, s.svcCtx)
	return l.ListPayment(in)
}
