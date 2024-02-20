package grpc

import (
	"context"

	"github.com/novriyantoAli/product-service/domain"
	"github.com/novriyantoAli/product-service/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type radcheckServer struct {
	model.UnimplementedRadchecksServer
	UCase domain.RadcheckUsecase
}

func NewHandler(srv *grpc.Server, ucase domain.RadcheckUsecase) {
	handler := &radcheckServer{UCase: ucase}

	model.RegisterRadchecksServer(srv, handler)
}

func (s *radcheckServer) CreateBatch(_ context.Context, param *model.RadcheckList) (res *emptypb.Empty, err error) {
	err = s.UCase.CreateBatch(param)

	return
}
