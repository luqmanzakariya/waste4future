package handler

import (
	"address-service/model"
	pb "address-service/pb/address"
	"address-service/usecase"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AddressGRPCHandler struct {
	pb.UnimplementedAddressServiceServer
	AddressUsecase usecase.IAddressUsecase
}

func NewAddressGRPCHandler(addressUsecase usecase.IAddressUsecase) *AddressGRPCHandler {
	return &AddressGRPCHandler{
		AddressUsecase: addressUsecase,
	}
}

func (h *AddressGRPCHandler) CreateAddress(ctx context.Context, req *pb.CreateAddressRequest) (*pb.AddressResponse, error) {
	payload := model.PayloadCreateAddress{
		Name:      req.GetName(),
		Latitude:  req.GetLatitude(),
		Longitude: req.GetLongitude(),
	}

	created, err := h.AddressUsecase.Create(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create address: %v", err)
	}

	return &pb.AddressResponse{
		Id:        created.ID,
		Name:      created.Name,
		Latitude:  created.Latitude,
		Longitude: created.Longitude,
		CreatedAt: created.CreatedAt.String(),
		UpdatedAt: created.UpdatedAt.String(),
	}, nil
}

func (h *AddressGRPCHandler) GetAddressByID(ctx context.Context, req *pb.GetAddressByIDRequest) (*pb.AddressResponse, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "address ID is required")
	}

	data, err := h.AddressUsecase.FindByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get address: %v", err)
	}

	return &pb.AddressResponse{
		Id:        data.ID,
		Name:      data.Name,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		CreatedAt: data.CreatedAt.String(),
		UpdatedAt: data.UpdatedAt.String(),
	}, nil
}

func (h *AddressGRPCHandler) GetAllAddresses(ctx context.Context, req *pb.GetAddressesRequest) (*pb.GetAddressesResponse, error) {
	addresses, err := h.AddressUsecase.FindAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all addresses: %v", err)
	}

	var pbAddresses []*pb.AddressResponse
	for _, addr := range addresses {
		pbAddresses = append(pbAddresses, &pb.AddressResponse{
			Id:        addr.ID.Hex(),
			Name:      addr.Name,
			Latitude:  addr.Latitude,
			Longitude: addr.Longitude,
			CreatedAt: addr.CreatedAt.String(),
			UpdatedAt: addr.UpdatedAt.String(),
		})
	}

	return &pb.GetAddressesResponse{
		Addresses: pbAddresses,
	}, nil
}

func (h *AddressGRPCHandler) UpdateAddress(ctx context.Context, req *pb.UpdateAddressRequest) (*pb.AddressResponse, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "address ID is required")
	}

	payload := model.PayloadUpdateAddress{
		Name:      req.GetName(),
		Latitude:  req.GetLatitude(),
		Longitude: req.GetLongitude(),
	}

	updated, err := h.AddressUsecase.Update(ctx, id, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update address: %v", err)
	}

	return &pb.AddressResponse{
		Id:        updated.ID,
		Name:      updated.Name,
		Latitude:  updated.Latitude,
		Longitude: updated.Longitude,
		CreatedAt: updated.CreatedAt.String(),
		UpdatedAt: updated.UpdatedAt.String(),
	}, nil
}

func (h *AddressGRPCHandler) DeleteAddress(ctx context.Context, req *pb.DeleteAddressRequest) (*pb.DeleteAddressResponse, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "address ID is required")
	}

	err := h.AddressUsecase.Delete(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete address: %v", err)
	}

	return &pb.DeleteAddressResponse{
		Success: true,
		Message: "Address deleted successfully",
	}, nil
}
