package grpc

import (
	"context"
	"recyclehub-service/model/web/request"
	"recyclehub-service/pb/recyclehub"
	"recyclehub-service/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecycleHubServer struct {
	recyclehub.UnimplementedRecycleHubServiceServer
	useCase usecase.IRecycleHubUsecase
}

func NewRecycleHubServer(useCase usecase.IRecycleHubUsecase) *RecycleHubServer {
	return &RecycleHubServer{useCase: useCase}
}

// CreateRecycleHub handles the gRPC request to create a new recycle hub.
func (s *RecycleHubServer) CreateRecycleHub(ctx context.Context, req *recyclehub.CreateRecycleHubRequest) (*recyclehub.RecycleHubResponse, error) {
	// Convert the gRPC request to a use case request
	createRequest := request.RecycleHubCreateRequest{
		Name:        req.Name,
		Phone:       req.Phone,
		AddressID:   req.AddressId,
		WasteTypeID: req.WasteTypeId,
	}

	// Call the use case with the converted request
	recycleHub, err := s.useCase.Create(ctx, createRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create recycle hub: %v", err)
	}

	// Convert the domain object to a gRPC response
	return &recyclehub.RecycleHubResponse{
		Id:          recycleHub.ID, // Use ID directly as it is already a string
		Name:        recycleHub.Name,
		Phone:       recycleHub.Phone,
		AddressId:   recycleHub.AddressID,
		WasteTypeId: recycleHub.WasteTypeID,
	}, nil
}

// GetAllRecycleHubs handles the gRPC request to retrieve all recycle hubs.
func (s *RecycleHubServer) GetAllRecycleHubs(ctx context.Context, req *recyclehub.Empty) (*recyclehub.RecycleHubList, error) {
	recycleHubs, err := s.useCase.FindAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve recycle hubs: %v", err)
	}

	var pbRecycleHubs []*recyclehub.RecycleHubResponse
	for _, hub := range recycleHubs {
		pbRecycleHubs = append(pbRecycleHubs, &recyclehub.RecycleHubResponse{
			Id:          hub.ID, // Use ID directly as it is already a string
			Name:        hub.Name,
			Phone:       hub.Phone,
			AddressId:   hub.AddressID,
			WasteTypeId: hub.WasteTypeID,
		})
	}

	return &recyclehub.RecycleHubList{RecycleHubs: pbRecycleHubs}, nil
}

// GetRecycleHubByID handles the gRPC request to retrieve a specific recycle hub by ID.
func (s *RecycleHubServer) GetRecycleHubByID(ctx context.Context, req *recyclehub.GetRecycleHubByIDRequest) (*recyclehub.RecycleHubResponse, error) {
	recycleHub, err := s.useCase.FindById(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Recycle hub not found: %v", err)
	}

	return &recyclehub.RecycleHubResponse{
		Id:          recycleHub.ID, // Use ID directly as it is already a string
		Name:        recycleHub.Name,
		Phone:       recycleHub.Phone,
		AddressId:   recycleHub.AddressID,
		WasteTypeId: recycleHub.WasteTypeID,
	}, nil
}

// UpdateRecycleHub handles the gRPC request to update an existing recycle hub.
func (s *RecycleHubServer) UpdateRecycleHub(ctx context.Context, req *recyclehub.UpdateRecycleHubRequest) (*recyclehub.RecycleHubResponse, error) {
	// Convert the gRPC request to a use case request
	updateRequest := request.RecycleHubUpdateRequest{
		Name:        req.Name,
		Phone:       req.Phone,
		AddressID:   req.AddressId,
		WasteTypeID: req.WasteTypeId, // Ensure this matches the type in the request
	}

	recycleHub, err := s.useCase.Update(ctx, updateRequest, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update recycle hub: %v", err)
	}

	return &recyclehub.RecycleHubResponse{
		Id:          recycleHub.ID, // Use ID directly as it is already a string
		Name:        recycleHub.Name,
		Phone:       recycleHub.Phone,
		AddressId:   recycleHub.AddressID,
		WasteTypeId: recycleHub.WasteTypeID,
	}, nil
}

// DeleteRecycleHub handles the gRPC request to delete a recycle hub by ID.
func (s *RecycleHubServer) DeleteRecycleHub(ctx context.Context, req *recyclehub.DeleteRecycleHubRequest) (*recyclehub.Empty, error) {
	err := s.useCase.Delete(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete recycle hub: %v", err)
	}
	return &recyclehub.Empty{}, nil
}
