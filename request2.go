package main

import (
	"fmt"
	"log"

	"github.com/KeThichDua/ex5go/db"
	"github.com/KeThichDua/ex5go/rpc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/xid"
)

// Run2 Viết một message UserPartnerRequest nhằm tạo 1 query xorm. Bao gồm lấy userpartner theo
// user_id, phone, với limit là số lượng row lớn nhất được quét ra. Với id được genere ngẫu nhiên với xid
func Run2() {
	var d db.Database
	err := d.Connect()
	ThrowError(err)

	// anh xa bang
	err = d.CreateTable()
	ThrowError(err)
	err = d.Sync2()
	ThrowError(err)
	// var err error
	// insert 1 vai ban ghi
	for i := 0; i < 10; i++ {
		guid := xid.New()
		guid1 := xid.New()
		user := db.UserPartner{Id: guid.String(), UserId: guid.String(), Phone: guid1.String()}
		err = d.InsertUser(user)
		ThrowError(err)
	}

	// thuc hien UserPartnerRequest
	userId := "1"
	phone := ""
	limit := int64(5)
	in := rpc.UserPartnerRequest{UserId: userId, Phone: phone, Limit: limit}
	out, err := d.Request(&in)
	ThrowError(err)
	// Kiem tra out
	if err == nil {
		for i := range out.UserPartners {
			fmt.Println(out.UserPartners[i])
		}
	}

}

// ThrowError nem ra loi neu khac nil
func ThrowError(err error) {
	if err != nil {
		log.Println(err)
	}
}
