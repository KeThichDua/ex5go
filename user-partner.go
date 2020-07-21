package main

import (
	"context"
	"errors"

	"github.com/KeThichDua/ex5go/db"
	"github.com/KeThichDua/ex5go/rpc"
)

// UserPartnerService la trinh trien khai cua UserPartnerService
type UserPartnerService struct {
	Db db.Database
}

func (s *UserPartnerService) Request(ctx context.Context, in *rpc.UserPartnerRequest) (*rpc.UserPartnerResponse, error) {
	return nil, nil
}

func (s *UserPartnerService) Read(ctx context.Context, in *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	userPartner := rpc.UserPartner{UserId: in.UserId}
	c, err := s.Db.Engine.Get(userPartner)
	if err != nil {
		return nil, err
	}
	if !c {
		return nil, errors.New("Khong tim thay user")
	}
	return &rpc.ReadResponse{
		UserPartner: &userPartner,
	}, nil
}

func (s *UserPartnerService) Create(ctx context.Context, in *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	userPartner := in.UserPartner
	c, err := s.Db.Engine.Insert(userPartner)
	if err != nil {
		return nil, err
	}
	if c == 0 {
		return nil, errors.New("Khong the them user")
	}
	return &rpc.CreateResponse{
		UserId: userPartner.UserId,
	}, nil
}

func (s *UserPartnerService) ReadAll(ctx context.Context, in *rpc.ReadAllRequest) (*rpc.ReadAllResponse, error) {
	list := []*rpc.UserPartner{}
	err := s.Db.Engine.Find(&list)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("Du lieu trong")
	}
	return &rpc.ReadAllResponse{
		UserPartners: list,
	}, nil
}

func (s *UserPartnerService) Delete(ctx context.Context, in *rpc.DeleteRequest) (*rpc.DeleteResponse, error) {
	userPartner := rpc.UserPartner{UserId: in.UserId}
	c, err := s.Db.Engine.Delete(userPartner)
	if err != nil {
		return nil, err
	}
	if c == 0 {
		return nil, errors.New("Khong the them user")
	}
	return &rpc.DeleteResponse{
		UserId: userPartner.UserId,
	}, nil
}
