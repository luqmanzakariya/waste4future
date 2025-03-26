package grpc

import (
	"context"
	"recyclehub-service/pb/recyclehub"
	"recyclehub-service/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WasteTypeServer struct {
	recyclehub.UnimplementedWasteTypeServiceServer
	useCase usecase.IWasteTypeUsecase
}

func NewWasteTypeServer(useCase usecase.IWasteTypeUsecase) *WasteTypeServer {
	return &WasteTypeServer{useCase: useCase}
}

// GetAllWasteTypes handles the gRPC request to retrieve all waste types.
func (s *WasteTypeServer) GetAllWasteTypes(ctx context.Context, req *recyclehub.Empty) (*recyclehub.WasteTypeList, error) {
	wasteTypes, err := s.useCase.FindAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve waste types: %v", err)
	}

	var pbWasteTypes []*recyclehub.WasteType
	for _, wt := range wasteTypes {
		pbWasteTypes = append(pbWasteTypes, &recyclehub.WasteType{
			Id:    wt.ID.Hex(), // Use the string ID directly
			Name:  wt.Name,
			Price: wt.Price,
		})
	}

	return &recyclehub.WasteTypeList{WasteTypes: pbWasteTypes}, nil
}

// GetWasteTypeByID handles the gRPC request to retrieve a specific waste type by ID.
func (s *WasteTypeServer) GetWasteTypeByID(ctx context.Context, req *recyclehub.GetWasteTypeByIDRequest) (*recyclehub.WasteType, error) {
	wasteType, err := s.useCase.FindById(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Waste type not found: %v", err)
	}

	return &recyclehub.WasteType{
		Id:    wasteType.ID.Hex(), // Use the string ID directly
		Name:  wasteType.Name,
		Price: wasteType.Price,
	}, nil
}
