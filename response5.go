package main

import (
	"github.com/KeThichDua/ex5go/rpc"
	"google.golang.org/grpc"
)

type FormUpdate struct{
	UserId string `json:"user_id" bingding:"required"`
}

type GRPCClient struct {
	UserRPCClient rpc.UserPartnerService5Client
}

var GRPC_CLI *GRPCClient

func CreateGRPCClient() {
	conn, err := grpc.Dial(":3001", grpc.WithInsecure())
	ThrowError(err)

	GRPC_CLI = &GRPCClient{
		UserRPCClient: rpc.NewUserPartnerService5Client(conn),
	}
}

func Start(){
	createGRPCClient()
	cli := GRPC_CLI.UserRPCClient
	app := gin.Default()
	u := app.Group("/"){
		u.Get("/get-list", func(c *gin.Context){ 
			res, err := cli.GetList()
			if err !=nil{
				c.JSON(500, gin.H{
					"Ok": false,
					"Msg": err.Error(),
				})
				return
			}
			c.JSON(200,res)
		})

		u.Get("/create-user", func(c *gin.Context){
			res, err := cli.Create()
			if err !=nil{
				c.JSON(500, gin.H{
					"Ok": false,
					"Msg": err.Error(),
				})
				return
			}
			c.JSON(200,res)
		 })

		u.Get("/update-user", func(c *gin.Context){
			var info FormUpdate
			err := c.BindJSON(&info)
			if err != nil{
				fmt.Println(err)
				c.JSON(400, gin.H{
					"Ok": false,
					"Msg": err.Error(),
				})
				return
			}
			res, err := cli.Update(context.Background(), &rpc.FormUpdate{
				UserId: info.UserId,
			})
			if err !=nil{
				c.JSON(500, gin.H{
					"Ok": false,
					"Msg": err.Error(),
				})
				return
			}
			c.JSON(200,res)
		 })
	}
	app.Run(":1003")
}