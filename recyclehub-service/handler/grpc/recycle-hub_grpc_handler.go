package handler

import (
	"context"
	"reyclehub-service/model"
	pb "reyclehub-service/pb/recycle_hub"
	"reyclehub-service/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecycleHubGRPCHandler struct {
	pb.UnimplementedRecycleHubServiceServer
	RecycleHubUsecase usecase.IRecycleHubUsecase
}

func NewRecycleHubGRPCHandler(recycleHubUsecase usecase.IRecycleHubUsecase) *RecycleHubGRPCHandler {
	return &RecycleHubGRPCHandler{
		RecycleHubUsecase: recycleHubUsecase,
	}
}

func (h *RecycleHubGRPCHandler) CreateRecycleHub(ctx context.Context, req *pb.CreateRecycleHubRequest) (*pb.RecycleHubResponse, error) {
	payload := model.PayloadCreateRecycleHub{
		Name:        req.GetName(),
		Phone:       req.GetPhone(),
		AddressID:   req.GetAddressId(),
		WasteTypeID: req.GetWasteTypeId(),
	}

	created, err := h.RecycleHubUsecase.Create(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create recycle hub: %v", err)
	}

	return &pb.RecycleHubResponse{
		Id:          created.ID,
		Name:        created.Name,
		Phone:       created.Phone,
		AddressId:   created.AddressID,
		WasteTypeId: created.WasteTypeID,
		CreatedAt:   created.CreatedAt.String(),
		UpdatedAt:   created.UpdatedAt.String(),
	}, nil
}

func (h *RecycleHubGRPCHandler) GetRecycleHubByID(ctx context.Context, req *pb.GetRecycleHubByIDRequest) (*pb.RecycleHubResponse, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "recycle hub ID is required")
	}

	data, err := h.RecycleHubUsecase.FindByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get recycle hub: %v", err)
	}

	return &pb.RecycleHubResponse{
		Id:          data.ID,
		Name:        data.Name,
		Phone:       data.Phone,
		AddressId:   data.AddressID,
		WasteTypeId: data.WasteTypeID,
		CreatedAt:   data.CreatedAt.String(),
		UpdatedAt:   data.UpdatedAt.String(),
	}, nil
}

func (h *RecycleHubGRPCHandler) GetAllRecycleHubs(ctx context.Context, req *pb.GetRecycleHubsRequest) (*pb.GetRecycleHubsResponse, error) {
	recycleHubs, err := h.RecycleHubUsecase.FindAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all recycle hubs: %v", err)
	}

	var pbRecycleHubs []*pb.RecycleHubResponse
	for _, hub := range recycleHubs {
		pbRecycleHubs = append(pbRecycleHubs, &pb.RecycleHubResponse{
			Id:          hub.ID.Hex(),
			Name:        hub.Name,
			Phone:       hub.Phone,
			AddressId:   hub.AddressID,
			WasteTypeId: hub.WasteTypeID,
			CreatedAt:   hub.CreatedAt.String(),
			UpdatedAt:   hub.UpdatedAt.String(),
		})
	}

	return &pb.GetRecycleHubsResponse{
		RecycleHubs: pbRecycleHubs,
	}, nil
}

func (h *RecycleHubGRPCHandler) UpdateRecycleHub(ctx context.Context, req *pb.UpdateRecycleHubRequest) (*pb.RecycleHubResponse, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "recycle hub ID is required")
	}

	payload := model.PayloadUpdateRecycleHub{
		Name:        req.GetName(),
		Phone:       req.GetPhone(),
		AddressID:   req.GetAddressId(),
		WasteTypeID: req.GetWasteTypeId(),
	}

	updated, err := h.RecycleHubUsecase.Update(ctx, id, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update recycle hub: %v", err)
	}

	return &pb.RecycleHubResponse{
		Id:          updated.ID,
		Name:        updated.Name,
		Phone:       updated.Phone,
		AddressId:   updated.AddressID,
		WasteTypeId: updated.WasteTypeID,
		CreatedAt:   updated.CreatedAt.String(),
		UpdatedAt:   updated.UpdatedAt.String(),
	}, nil
}

func (h *RecycleHubGRPCHandler) DeleteRecycleHub(ctx context.Context, req *pb.DeleteRecycleHubRequest) (*pb.DeleteRecycleHubResponse, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "recycle hub ID is required")
	}

	err := h.RecycleHubUsecase.Delete(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete recycle hub: %v", err)
	}

	return &pb.DeleteRecycleHubResponse{
		Success: true,
		Message: "Recycle hub deleted successfully",
	}, nil
}
