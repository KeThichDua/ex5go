package db

import (
	"errors"

	"github.com/KeThichDua/ex5go/rpc"
)

type UserPartner struct {
	Id          string
	UserId      string
	PartnerId   string
	AliasUserId string
	Apps        map[string]int64
	Phone       string
	Created     int64
	UpdatedAt   int64
}

// InsertUser de them du lieu user
func (s *Database) InsertUser(user UserPartner) error {
	c, err := s.Engine.Insert(user)
	if c == 0 {
		return errors.New("Loi insert")
	}
	if err != nil {
		return err
	}
	return err
}

// UpdateUser de sua du lieu user
func (s *Database) UpdateUser(user UserPartner, conditions UserPartner) error {
	c, err := s.Engine.Update(user, conditions)
	if err != nil {
		return err
	}
	if c == 0 {
		return errors.New("Khong tim thay user")
	}
	return err
}

// GetUser tim kiem 1 user
func (s *Database) GetUser(id string) (*UserPartner, error) {
	user := &UserPartner{Id: id}
	c, err := s.Engine.Get(user)
	if err != nil {
		return nil, err
	}
	if !c {
		return nil, errors.New("Khong tim thay")
	}
	return user, nil
}

// ListUser de liet ke tat ca user
func (s *Database) ListUser() ([]*UserPartner, error) {
	var users []*UserPartner
	err := s.Engine.Desc("id").Find(&users)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("Database rong")
	}
	return users, nil
}

// DeleteUser xoa user
func (s *Database) DeleteUser(conditions UserPartner) error {
	c, err := s.Engine.Delete(conditions)
	if err != nil {
		return err
	}
	if c == 0 {
		return errors.New("Xoa that bai")
	}
	return nil
}

// Request thuc hien lay user theo  user_id, phone
func (s *Database) Request(in *rpc.UserPartnerRequest) (*rpc.UserPartnerResponse, error) {
	var list []*UserPartner
	condi := UserPartner{
		UserId: in.UserId,
		Phone:  in.Phone,
	}
	err := s.Engine.Limit(int(in.Limit), 0).Find(&list, condi)
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
	return &rpc.UserPartnerResponse{
		UserPartners: users,
	}, nil
}
