package db

import (
	"errors"
	"log"

	"github.com/KeThichDua/ex5go/rpc"
	"xorm.io/xorm"
)

// Database dung xorm de ket noi mysql
type Database struct {
	Engine *xorm.Engine
}

// Connect tao ket noi database
func (s *Database) Connect(driverName string, dataSourceName string) error {
	var err error
	s.Engine, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		log.Println("fail")
		return err
	}
	return nil
}

// CreateTable la phuong thuc tao bang UserPartner
func (s *Database) CreateTable() error {
	err := s.Engine.CreateTables(rpc.UserPartner{})
	if err != nil {
		return err
	}
	return nil
}

// Sync2 de anh xa bang
func (s *Database) Sync2() error {
	err := s.Engine.Sync2(new(rpc.UserPartner))
	if err != nil {
		return err
	}
	return nil
}

// InsertUser de them du lieu user
func (s *Database) InsertUser(user *rpc.UserPartner) error {
	c, err := s.Engine.Insert(user)
	if c == 0 {
		return errors.New("Loi khong the insert")
	}
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser de sua du lieu user
func (s *Database) UpdateUser(user *rpc.UserPartner, conditions *rpc.UserPartner) error {
	c, err := s.Engine.Update(user, conditions)
	if c == 0 {
		return errors.New("Khong tim thay user")
	}
	if err != nil {
		return err
	}
	return err
}

// FindUser tim kiem 1 user
func (s *Database) FindUser(id string) (*rpc.UserPartner, error) {
	user := &rpc.UserPartner{UserId: id}
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
func (s *Database) ListUser() ([]*rpc.UserPartner, error) {
	var users []*rpc.UserPartner
	err := s.Engine.Desc("id").Find(&users)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("Database rong")
	}
	return users, nil
}

// DeleteUser tim kiem 1 user
func (s *Database) DeleteUser(id string) error {
	user := rpc.UserPartner{UserId: id}
	c, err := s.Engine.Delete(&user)
	if err != nil {
		return err
	}
	if c == 0 {
		return errors.New("Khong tim thay")
	}
	return nil
}

// ListUserConditions dung liet ke user theo dieu kien
func (s *Database) ListUserConditions(userId string, phone string, limit int64) ([]*rpc.UserPartner, error) {
	list := []*rpc.UserPartner{}
	err := s.Engine.Limit(int(limit), 0).Find(&list, rpc.UserPartner{UserId: userId, Phone: phone})
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("Khong tim thay")
	}
	return list, nil
}

// Request thuc hien yeu cau UserPartnerRequest ma ko co server
func (s *Database) Request(in *rpc.UserPartnerRequest) (*rpc.UserPartnerResponse, error) {
	list, err := s.ListUserConditions(in.UserId, in.Phone, in.Limit)
	if err != nil {
		return nil, err
	}
	return &rpc.UserPartnerResponse{
		UserPartners: list,
	}, nil
}
