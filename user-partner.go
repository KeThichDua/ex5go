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

// Request UserPartnerService
func (s *UserPartnerService) Request(ctx context.Context, in *rpc.UserPartnerRequest) (*rpc.UserPartnerResponse, error) {
	return nil, nil
}

// Read UserPartnerService
func (s *UserPartnerService) Read(ctx context.Context, in *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	c, err := s.Db.GetUser(in.UserId)
	if err != nil {
		return nil, err
	}
	temp := rpc.UserPartner{
		Id:          c.Id,
		UserId:      c.UserId,
		PartnerId:   c.PartnerId,
		AliasUserId: c.AliasUserId,
		Apps:        c.Apps,
		Phone:       c.Phone,
		Created:     c.Created,
		UpdatedAt:   c.UpdatedAt,
	}
	return &rpc.ReadResponse{
		UserPartner: &temp,
	}, nil
}

// Create tao user
func (s *UserPartnerService) Create(ctx context.Context, in *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	temp := db.UserPartner{
		Id:          in.UserPartner.Id,
		UserId:      in.UserPartner.UserId,
		PartnerId:   in.UserPartner.PartnerId,
		AliasUserId: in.UserPartner.AliasUserId,
		Apps:        in.UserPartner.Apps,
		Phone:       in.UserPartner.Phone,
		Created:     in.UserPartner.Created,
		UpdatedAt:   in.UserPartner.UpdatedAt,
	}
	err := s.Db.InsertUser(temp)
	if err != nil {
		return nil, err
	}
	return &rpc.CreateResponse{
		UserId: temp.UserId,
	}, nil
}

// ReadAll lay ve tat ca user
func (s *UserPartnerService) ReadAll(ctx context.Context, in *rpc.ReadAllRequest) (*rpc.ReadAllResponse, error) {
	list, err := s.Db.ListUser()
	if err != nil {
		return nil, err
	}
	var users []*rpc.UserPartner
	for i := range list {
		temp := rpc.UserPartner{
			Id:          list[i].Id,
			UserId:      list[i].UserId,
			PartnerId:   list[i].PartnerId,
			AliasUserId: list[i].AliasUserId,
			Apps:        list[i].Apps,
			Phone:       list[i].Phone,
			Created:     list[i].Created,
			UpdatedAt:   list[i].UpdatedAt,
		}
		users = append(users, &temp)
	}
	return &rpc.ReadAllResponse{
		UserPartners: users,
	}, nil
}

// Delete xoa user theo id
func (s *UserPartnerService) Delete(ctx context.Context, in *rpc.DeleteRequest) (*rpc.DeleteResponse, error) {
	temp := db.UserPartner{
		Id: in.UserId,
	}
	err := s.Db.DeleteUser(temp)
	if err != nil {
		return nil, err
	}
	return &rpc.DeleteResponse{
		UserId: in.UserId,
	}, nil
}
