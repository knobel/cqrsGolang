package requests

import (
	"github.com/satori/go.uuid"
)

type IRequest interface{
	CommandId() uuid.UUID
}

type BaseIRequest struct {
	Id uuid.UUID `json:"id"`
}

func (c BaseIRequest) CommandId() uuid.UUID{
	return c.Id
}

type ReserveRoomRequest struct{
	BaseIRequest
	RoomType string `json:"roomtype"`
	HotelId uuid.UUID `json:"hotelid"`
}

func NewReserveRoomRequest(id uuid.UUID, hotelId uuid.UUID, roomType string) ReserveRoomRequest {
	s := ReserveRoomRequest{}
	s.Id = id
	s.HotelId = hotelId
	s.RoomType = roomType
	return s
}

type UnReserveRoomRequest struct{
	BaseIRequest
	HotelId uuid.UUID `json:"hotelid"`
}

func NewUnReserveRoomRequest(id uuid.UUID, hotelId uuid.UUID) UnReserveRoomRequest {
	s := UnReserveRoomRequest{}
	s.Id = id
	s.HotelId = hotelId
	return s
}

