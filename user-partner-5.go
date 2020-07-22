package main

import (
	"context"
	"time"

	"github.com/KeThichDua/ex5go/db"
	"github.com/KeThichDua/ex5go/rpc"
	"github.com/rs/xid"
)

// UserPartner5 la trinh trien khai cua UserPartnerService5
type UserPartner5 struct {
	Db db.Database
}

// Update cap nhat UpdatedAt cua User theo id
func (s *UserPartner5) Update(ctx context.Context, in *rpc.UpdateUserRequest) (*rpc.UpdateUserResponse, error) {
	userPartner := rpc.UserPartner{UpdatedAt: time.Now().Unix()}
	condition := rpc.UserPartner{UserId: in.UserId}
	err := s.Db.UpdateUser(&userPartner, &condition)
	if err != nil {
		return nil, err
	}
	return &rpc.UpdateUserResponse{}, nil
}

// Create tao user bat ki
func (s *UserPartner5) Create(ctx context.Context, in *rpc.CreateUserRequest) (*rpc.CreateUserResponse, error) {
	guid := xid.New()
	guid1 := xid.New()
	userPartner := rpc.UserPartner{Id: guid.String(), UserId: guid.String(), Phone: guid1.String()}
	err := s.Db.InsertUser(&userPartner)
	if err != nil {
		return nil, err
	}
	return &rpc.CreateUserResponse{}, nil
}

// GetList lay tat ca user
func (s *UserPartner5) GetList(ctx context.Context, in *rpc.GetListRequest) (*rpc.GetListResponse, error) {
	// list := []*rpc.UserPartner{}
	_, err := s.Db.ListUser()
	if err != nil {
		return nil, err
	}
	return &rpc.GetListResponse{}, nil
}
