package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KeThichDua/ex5go/db"
	"github.com/rs/xid"

	"github.com/KeThichDua/ex5go/rpc"
	"github.com/julienschmidt/httprouter"
)

var us UserPartnerService
var ctx context.Context

// Run4 : kết hợp với kiến thức trên viết 1 server sử dụng `grpc` generate 2 message ở bài 1, kết hợp server bài 2 và
// kiến thức từ bài trước `xorm`. Viết 1 route: POST `/user-partner` tạo mới 1 partner, GET `/user-partner` lấy danh sách partner,
// DELETE `/user-partner/{id}` theo id,  GET `/user-partner/{id}` lấy theo 1 id cụ thể.
func Run4() {
	// ket noi mysql
	var db db.Database
	err := db.Connect()
	ThrowError(err)
	defer db.Engine.Close()
	ctx = context.Background()
	us = UserPartnerService{Db: db}

	if err := RunServer(ctx, us); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// RunServer chay http server
func RunServer(ctx context.Context, us UserPartnerService) error {
	router := httprouter.New()
	router.POST("/user-partner", CreateUser)
	router.GET("/user-partner", ListUser)
	router.DELETE("/user-partner/:id", DeleteUser)
	router.GET("/user-partner/:id", GetUser)

	// start server
	log.Println("starting server...")
	return http.ListenAndServe(":3000", router)
}

// CreateUser tao moi userpartner
func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	guid := xid.New()
	guid1 := xid.New()
	user := rpc.UserPartner{Id: guid.String(), UserId: guid.String(), Phone: guid1.String()}
	req := rpc.CreateRequest{UserPartner: &user}
	res, err := us.Create(ctx, &req)
	ThrowError(err)
	if res != nil {
		fmt.Fprint(w, "Tao user voi id = ", res.UserId)
	}
}

// ListUser liet ke tat ca user
func ListUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := us.ReadAll(ctx, nil)
	ThrowError(err)
	if res != nil {
		fmt.Fprint(w, "Tat ca user: ", res.UserPartners)
	}
}

// DeleteUser xoa user theo id
func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := rpc.DeleteRequest{UserId: ps.ByName("id")}
	res, err := us.Delete(ctx, &req)
	ThrowError(err)
	if res != nil {
		fmt.Fprint(w, "Da xoa user voi id = ", res.UserId)
	}
}

// GetUser lay user theo id
func GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := rpc.ReadRequest{UserId: ps.ByName("id")}
	res, err := us.Read(ctx, &req)
	ThrowError(err)
	if res != nil {
		fmt.Fprint(w, "Tim thay ser: ", res.UserPartner)
	}
}
