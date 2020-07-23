package reqres5

import (
	"context"
	"time"

	"github.com/KeThichDua/ex5go/db"
	"github.com/rs/xid"
)

// UserPartnerService la trinh trien khai cua UserPartnerService
type UserPartnerService struct {
	Db db.Database
}

// Create tao user
func (s *UserPartnerService) Create(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, error) {
	guid := xid.New()
	guid1 := xid.New()
	temp := db.UserPartner{Id: guid.String(), UserId: guid.String(), Phone: guid1.String()}
	err := s.Db.InsertUser(temp)
	if err != nil {
		return nil, err
	}
	return &CreateUserResponse{}, nil
}

// GetList lay ve tat ca user
func (s *UserPartnerService) GetList(ctx context.Context, in *GetListRequest) (*GetListResponse, error) {
	_, err := s.Db.ListUser()
	if err != nil {
		return nil, err
	}
	return &GetListResponse{}, nil
}

// Update user theo id
func (s *UserPartnerService) Update(ctx context.Context, in *UpdateUserRequest) (*UpdateUserResponse, error) {
	temp := db.UserPartner{
		Id: in.UserId,
	}
	user := db.UserPartner{UpdatedAt: time.Now().Unix()}
	err := s.Db.UpdateUser(user, temp)
	if err != nil {
		return nil, err
	}
	return &UpdateUserResponse{}, nil
}
