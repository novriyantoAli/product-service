package grpc

import (
	"context"

	"github.com/novriyantoAli/product-service/domain"
	"github.com/novriyantoAli/product-service/model"
	"google.golang.org/grpc"
)

type productServer struct {
	model.UnimplementedProductsServer
	UCase domain.ProductUsecase
}

func NewHandler(srv *grpc.Server, ucase domain.ProductUsecase) {
	handler := &productServer{UCase: ucase}

	model.RegisterProductsServer(srv, handler)
}

func (h *productServer) Find(_ context.Context, param *model.Product) (*model.ProductList, error) {
	// validation input
	res, err := h.UCase.Find(param)
	return res, err
}
