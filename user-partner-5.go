package main

import (
	"context"
	"errors"
	"time"

	"github.com/KeThichDua/ex5go/db"
	"github.com/KeThichDua/ex5go/rpc"
	"github.com/rs/xid"
)

// UserPartner5 la trinh trien khai cua UserPartnerService5
type UserPartner5 struct {
	Db db.Database
}

func (s *UserPartner5) Update(ctx context.Context, in *rpc.UpdateUserRequest) (*rpc.UpdateUserResponse, error) {
	userPartner := rpc.UserPartner5{UpdatedAt: time.Now().Unix()}
	c, err := s.Db.Engine.Update(&userPartner, rpc.UserPartner5{UserId: in.UserId})
	if err != nil {
		return nil, err
	}
	if c == 0 {
		return nil, errors.New("Khong update dc user")
	}
	return &rpc.UpdateUserResponse{}, nil
}

func (s *UserPartner5) Create(ctx context.Context, in *rpc.CreateUserRequest) (*rpc.CreateUserResponse, error) {
	guid := xid.New()
	guid1 := xid.New()
	userPartner := rpc.UserPartner5{Id: guid.String(), UserId: guid.String(), Phone: guid1.String()}
	c, err := s.Db.Engine.Insert(userPartner)
	if err != nil {
		return nil, err
	}
	if c == 0 {
		return nil, errors.New("Khong the them user")
	}
	return &rpc.CreateUserResponse{}, nil
}

func (s *UserPartner5) GetList(ctx context.Context, in *rpc.GetListRequest) (*rpc.GetListResponse, error) {
	list := []*rpc.UserPartner{}
	err := s.Db.Engine.Find(&list)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("Du lieu trong")
	}
	return &rpc.GetListResponse{}, nil
}
