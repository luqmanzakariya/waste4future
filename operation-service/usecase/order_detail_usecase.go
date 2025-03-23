// order_detail_usecase.go
package usecase

import (
	"context"
	"errors"
	"math"
	"operation-service/model"
	pbAddress "operation-service/pb/address"    // Address service proto
	pbRecycle "operation-service/pb/recyclehub" // Recyclehub service proto (assumed)
	pbUser "operation-service/pb/user"          // User service proto (assumed)
	"operation-service/repository"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IOrderDetailUsecase interface {
	Create(ctx context.Context, payload model.PayloadCreateOrderDetail) (model.ResponseOrderDetail, error)
	FindAll(ctx context.Context) ([]model.OrderDetail, error)
	FindByID(ctx context.Context, id string) (model.ResponseOrderDetail, error)
	Update(ctx context.Context, id string, payload model.PayloadUpdateOrderDetail) (model.ResponseOrderDetail, error)
	Delete(ctx context.Context, id string) error
}

type orderDetailUsecase struct {
	OrderDetailRepo  repository.IOrderDetailRepository
	AddressClient    pbAddress.AddressServiceClient
	RecycleHubClient pbRecycle.RecycleHubServiceClient // Assumed client
	UserClient       pbUser.UserServiceClient          // Assumed client
	Validate         *validator.Validate
}

func NewOrderDetailUsecase(
	orderDetailRepo repository.IOrderDetailRepository,
	addressConn *grpc.ClientConn,
	recycleHubConn *grpc.ClientConn,
	userConn *grpc.ClientConn,
	validate *validator.Validate,
) IOrderDetailUsecase {
	return &orderDetailUsecase{
		OrderDetailRepo:  orderDetailRepo,
		AddressClient:    pbAddress.NewAddressServiceClient(addressConn),
		RecycleHubClient: pbRecycle.NewRecycleHubServiceClient(recycleHubConn),
		UserClient:       pbUser.NewUserServiceClient(userConn),
		Validate:         validate,
	}
}

// Haversine formula for distance calculation
func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Earth radius in kilometers

	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func (u *orderDetailUsecase) Create(ctx context.Context, payload model.PayloadCreateOrderDetail) (model.ResponseOrderDetail, error) {
	if err := u.Validate.Struct(payload); err != nil {
		return model.ResponseOrderDetail{}, err
	}

	// 1. Get origin address from address service
	originResp, err := u.AddressClient.GetAddressByID(ctx, &pbAddress.GetAddressByIDRequest{
		Id: payload.OriginAddressID,
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return model.ResponseOrderDetail{}, errors.New("origin address not found")
		}
		return model.ResponseOrderDetail{}, errors.New("failed to fetch origin address: " + err.Error())
	}

	// 2. Get destination address from address service
	destResp, err := u.AddressClient.GetAddressByID(ctx, &pbAddress.GetAddressByIDRequest{
		Id: payload.DestinationAddressID,
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return model.ResponseOrderDetail{}, errors.New("destination address not found")
		}
		return model.ResponseOrderDetail{}, errors.New("failed to fetch destination address: " + err.Error())
	}

	// 3. Get recycle hub details and waste price
	recycleResp, err := u.RecycleHubClient.GetRecycleByID(ctx, &pbRecycle.GetRecycleByIDRequest{
		Id: payload.RecycleHubID,
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return model.ResponseOrderDetail{}, errors.New("recycle hub not found")
		}
		return model.ResponseOrderDetail{}, errors.New("failed to fetch recycle hub: " + err.Error())
	}

	// 4. Get waste price from recyclehub service
	wasteResp, err := u.RecycleHubClient.GetWasteByID(ctx, &pbRecycle.GetWasteByIDRequest{
		Id: recycleResp.WasteTypeId,
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return model.ResponseOrderDetail{}, errors.New("waste type not found")
		}
		return model.ResponseOrderDetail{}, errors.New("failed to fetch waste price: " + err.Error())
	}

	// Convert coordinates to float64
	originLat, err := strconv.ParseFloat(originResp.Latitude, 64)
	if err != nil {
		return model.ResponseOrderDetail{}, errors.New("invalid origin latitude")
	}
	originLon, err := strconv.ParseFloat(originResp.Longitude, 64)
	if err != nil {
		return model.ResponseOrderDetail{}, errors.New("invalid origin longitude")
	}
	destLat, err := strconv.ParseFloat(destResp.Latitude, 64)
	if err != nil {
		return model.ResponseOrderDetail{}, errors.New("invalid destination latitude")
	}
	destLon, err := strconv.ParseFloat(destResp.Longitude, 64)
	if err != nil {
		return model.ResponseOrderDetail{}, errors.New("invalid destination longitude")
	}

	// Calculate delivery price
	distance := calculateDistance(originLat, originLon, destLat, destLon)
	deliveryPrice := distance * 15000 // 15,000 per km

	// Calculate subtotal
	subTotal := payload.WasteWeight * wasteResp.Price

	orderDetail := model.OrderDetail{
		UserID:               payload.UserID,
		RecycleHubID:         payload.RecycleHubID,
		WasteWeight:          payload.WasteWeight,
		SubTotal:             subTotal,
		OriginAddressID:      payload.OriginAddressID,
		DestinationAddressID: payload.DestinationAddressID,
		DeliveryPrice:        deliveryPrice,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	res, err := u.OrderDetailRepo.Create(ctx, orderDetail)
	if err != nil {
		return model.ResponseOrderDetail{}, errors.New("failed to create order detail: " + err.Error())
	}

	return model.ResponseOrderDetail{
		ID:                   res.ID.Hex(),
		UserID:               res.UserID,
		RecycleHubID:         res.RecycleHubID,
		WasteWeight:          res.WasteWeight,
		SubTotal:             res.SubTotal,
		OriginAddressID:      res.OriginAddressID,
		DestinationAddressID: res.DestinationAddressID,
		DeliveryPrice:        res.DeliveryPrice,
		CreatedAt:            res.CreatedAt,
		UpdatedAt:            res.UpdatedAt,
	}, nil
}

func (u *orderDetailUsecase) Update(ctx context.Context, id string, payload model.PayloadUpdateOrderDetail) (model.ResponseOrderDetail, error) {
	if err := u.Validate.Struct(payload); err != nil {
		return model.ResponseOrderDetail{}, err
	}

	existing, err := u.OrderDetailRepo.ReadByID(ctx, id)
	if err != nil {
		return model.ResponseOrderDetail{}, errors.New("order detail not found")
	}

	// Get addresses and recalculate delivery price
	originResp, err := u.AddressClient.GetAddressByID(ctx, &pbAddress.GetAddressByIDRequest{
		Id: payload.OriginAddressID,
	})
	if err != nil {
		return model.ResponseOrderDetail{}, errors.New("failed to fetch origin address: " + err.Error())
	}

	destResp, err := u.AddressClient.GetAddressByID(ctx, &pbAddress.GetAddressByIDRequest{
		Id: payload.DestinationAddressID,
	})
	if err != nil {
		return model.ResponseOrderDetail{}, errors.New("failed to fetch destination address: " + err.Error())
	}

	// Get waste price
	recycleResp, err := u.RecycleHubClient.GetRecycleByID(ctx, &pbRecycle.GetRecycleByIDRequest{
		Id: existing.RecycleHubID,
	})
	if err != nil {
		return model.ResponseOrderDetail{}, errors.New("failed to fetch recycle hub: " + err.Error())
	}

	wasteResp, err := u.RecycleHubClient.GetWasteByID(ctx, &pbRecycle.GetWasteByIDRequest{
		Id: recycleResp.WasteTypeId,
	})
	if err != nil {
		return model.ResponseOrderDetail{}, errors.New("failed to fetch waste price: " + err.Error())
	}

	// Convert coordinates and calculate distance
	originLat, _ := strconv.ParseFloat(originResp.Latitude, 64)
	originLon, _ := strconv.ParseFloat(originResp.Longitude, 64)
	destLat, _ := strconv.ParseFloat(destResp.Latitude, 64)
	destLon, _ := strconv.ParseFloat(destResp.Longitude, 64)
	distance := calculateDistance(originLat, originLon, destLat, destLon)
	deliveryPrice := distance * 15000

	// Calculate subtotal
	subTotal := payload.WasteWeight * wasteResp.Price

	updatedOrder := model.OrderDetail{
		ID:                   existing.ID,
		UserID:               existing.UserID,
		RecycleHubID:         existing.RecycleHubID,
		WasteWeight:          payload.WasteWeight,
		SubTotal:             subTotal,
		OriginAddressID:      payload.OriginAddressID,
		DestinationAddressID: payload.DestinationAddressID,
		DeliveryPrice:        deliveryPrice,
		CreatedAt:            existing.CreatedAt,
		UpdatedAt:            time.Now(),
	}

	res, err := u.OrderDetailRepo.Update(ctx, id, updatedOrder)
	if err != nil {
		return model.ResponseOrderDetail{}, errors.New("failed to update order detail: " + err.Error())
	}

	return model.ResponseOrderDetail{
		ID:                   res.ID.Hex(),
		UserID:               res.UserID,
		RecycleHubID:         res.RecycleHubID,
		WasteWeight:          res.WasteWeight,
		SubTotal:             res.SubTotal,
		OriginAddressID:      res.OriginAddressID,
		DestinationAddressID: res.DestinationAddressID,
		DeliveryPrice:        res.DeliveryPrice,
		CreatedAt:            res.CreatedAt,
		UpdatedAt:            res.UpdatedAt,
	}, nil
}

func (u *orderDetailUsecase) FindAll(ctx context.Context) ([]model.OrderDetail, error) {
	return u.OrderDetailRepo.ReadAll(ctx)
}

func (u *orderDetailUsecase) FindByID(ctx context.Context, id string) (model.ResponseOrderDetail, error) {
	res, err := u.OrderDetailRepo.ReadByID(ctx, id)
	if err != nil {
		return model.ResponseOrderDetail{}, errors.New("order detail not found")
	}

	return model.ResponseOrderDetail{
		ID:                   res.ID.Hex(),
		UserID:               res.UserID,
		RecycleHubID:         res.RecycleHubID,
		WasteWeight:          res.WasteWeight,
		SubTotal:             res.SubTotal,
		OriginAddressID:      res.OriginAddressID,
		DestinationAddressID: res.DestinationAddressID,
		DeliveryPrice:        res.DeliveryPrice,
		CreatedAt:            res.CreatedAt,
		UpdatedAt:            res.UpdatedAt,
	}, nil
}

func (u *orderDetailUsecase) Delete(ctx context.Context, id string) error {
	return u.OrderDetailRepo.Delete(ctx, id)
}
