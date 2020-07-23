package db

import (
	"xorm.io/xorm"
)

// Database bao gom Engine, bang User, bang Point
type Database struct {
	Engine *xorm.Engine
}

// Connect tao ket noi database
func (s *Database) Connect() error {
	engine, err := xorm.NewEngine("mysql", "root:1@tcp(0.0.0.0:3306)/ex5go")
	if err != nil {
		return err
	}
	s.Engine = engine
	return nil
}

// CreateTable la phuong thuc tao bang User
func (s *Database) CreateTable() error {
	err := s.Engine.CreateTables(UserPartner{})
	if err != nil {
		return err
	}
	return nil
}

// Sync2 de anh xa bang
func (s *Database) Sync2() error {
	err := s.Engine.Sync2(new(UserPartner))
	if err != nil {
		return err
	}
	return nil
}
