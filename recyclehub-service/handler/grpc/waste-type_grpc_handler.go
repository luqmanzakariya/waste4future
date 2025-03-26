package handler

import (
	"context"
	pb "reyclehub-service/pb/recycle_hub"
	"reyclehub-service/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WasteTypeGRPCHandler struct {
	pb.UnimplementedWasteTypeServiceServer
	WasteTypeUsecase usecase.IWasteTypeUsecase
}

func NewWasteTypeGRPCHandler(wasteTypeUsecase usecase.IWasteTypeUsecase) *WasteTypeGRPCHandler {
	return &WasteTypeGRPCHandler{
		WasteTypeUsecase: wasteTypeUsecase,
	}
}

func (h *WasteTypeGRPCHandler) GetWasteTypeByID(ctx context.Context, req *pb.GetWasteTypeByIDRequest) (*pb.WasteTypeResponse, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "waste type ID is required")
	}

	data, err := h.WasteTypeUsecase.FindByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get waste type: %v", err)
	}

	return &pb.WasteTypeResponse{
		Id:        data.ID.Hex(),
		Name:      data.Name,
		Price:     data.Price,
		CreatedAt: data.CreatedAt.String(),
		UpdatedAt: data.UpdatedAt.String(),
	}, nil
}

func (h *WasteTypeGRPCHandler) GetAllWasteTypes(ctx context.Context, req *pb.GetWasteTypesRequest) (*pb.GetWasteTypesResponse, error) {
	wasteTypes, err := h.WasteTypeUsecase.FindAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all waste types: %v", err)
	}

	var pbWasteTypes []*pb.WasteTypeResponse
	for _, wt := range wasteTypes {
		pbWasteTypes = append(pbWasteTypes, &pb.WasteTypeResponse{
			Id:        wt.ID.Hex(),
			Name:      wt.Name,
			Price:     wt.Price,
			CreatedAt: wt.CreatedAt.String(),
			UpdatedAt: wt.UpdatedAt.String(),
		})
	}

	return &pb.GetWasteTypesResponse{
		WasteTypes: pbWasteTypes,
	}, nil
}
