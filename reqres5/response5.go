package reqres5

import (
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// GRPCClient la grpc client
type GRPCClient struct {
	UserRPCClient UserPartnerServiceClient
}

// GRPC_CLI tao con tro grpc client
var GRPC_CLI *GRPCClient

// CreateGRPCClient tao grpc client
func CreateGRPCClient() {
	conn, err := grpc.Dial(":3001", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	GRPC_CLI = &GRPCClient{
		UserRPCClient: NewUserPartnerServiceClient(conn),
	}
}

// Start se tao va chay grpc client cung voi router client
func Start() {
	CreateGRPCClient()
	cli := GRPC_CLI.UserRPCClient
	app := gin.Default()
	u := app.Group("/")
	{
		u.GET("/get-list", func(c *gin.Context) {
			res, err := cli.GetList(ctx, &GetListRequest{})
			if err != nil {
				c.JSON(500, gin.H{
					"Ok":  false,
					"Msg": err.Error(),
				})
				return
			}
			c.JSON(200, res)
		})

		u.POST("/create-user", func(c *gin.Context) {
			res, err := cli.Create(ctx, &CreateUserRequest{})
			if err != nil {
				c.JSON(500, gin.H{
					"Ok":  false,
					"Msg": err.Error(),
				})
				return
			}
			c.JSON(200, res)
		})

		u.PUT("/update-user/:id", func(c *gin.Context) {
			id := c.Params.ByName("id")
			res, err := cli.Update(ctx, &UpdateUserRequest{
				UserId: id,
			})
			if err != nil {
				c.JSON(500, gin.H{
					"Ok":  false,
					"Msg": err.Error(),
				})
				return
			}
			c.JSON(200, res)
		})
	}
	app.Run(":3002")
}
