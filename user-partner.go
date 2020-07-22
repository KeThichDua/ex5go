package main

import (
	"context"

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
	c, err := s.Db.FindUser(in.UserId)
	if err != nil {
		return nil, err
	}
	return &rpc.ReadResponse{
		UserPartner: c,
	}, nil
}

// Create tao user
func (s *UserPartnerService) Create(ctx context.Context, in *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	err := s.Db.InsertUser(in.UserPartner)
	if err != nil {
		return nil, err
	}
	return &rpc.CreateResponse{
		UserId: "success",
	}, nil
}

// ReadAll lay ve tat ca user
func (s *UserPartnerService) ReadAll(ctx context.Context, in *rpc.ReadAllRequest) (*rpc.ReadAllResponse, error) {
	list, err := s.Db.ListUser()
	if err != nil {
		return nil, err
	}
	return &rpc.ReadAllResponse{
		UserPartners: list,
	}, nil
}

// Delete xoa user theo id
func (s *UserPartnerService) Delete(ctx context.Context, in *rpc.DeleteRequest) (*rpc.DeleteResponse, error) {
	err := s.Db.DeleteUser(in.UserId)
	if err != nil {
		return nil, err
	}
	return &rpc.DeleteResponse{
		UserId: in.UserId,
	}, nil
}
