package main

import (
	"context"
	"time"
	pb "user" // Need to add this package to your GO Path
	"net/http"
	"strconv"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/gin-gonic/gin"
)

const (
	address     = "127.0.0.1:10001"
	apiVersion = "v1"
	httpPort	= ":8080"
)
var (
	client pb.UserServiceClient
)

// Handle error return from server
func handleErrorFromServer(err error, g *gin.Context) {
	grpcErr, ok := status.FromError(err)
	if ok {
		switch grpcErr.Code() {
		case codes.NotFound:
			g.JSON(http.StatusNotFound, gin.H{"Message": "Not Found: "+grpcErr.Message(), "Status": http.StatusNotFound})
		case codes.Unimplemented:
			g.JSON(http.StatusMethodNotAllowed, gin.H{"Message": "Method not allowed: "+grpcErr.Message(), "Status": http.StatusMethodNotAllowed})
		case codes.AlreadyExists, codes.InvalidArgument:
			g.JSON(http.StatusBadRequest, gin.H{"Message": "Bad Request: "+grpcErr.Message(), "Status": http.StatusBadRequest})
		default:
			g.JSON(http.StatusInternalServerError, gin.H{"Message": "Unexpected error: "+grpcErr.Message(), "Status": http.StatusInternalServerError})
		}
	} else {
		g.String(500, "failed to call server")
	}
}
// Add User
func runCreateUser(g *gin.Context) {
	userAge, err := strconv.ParseInt(g.PostForm("age"), 10, 32)
	convertUserAge := int32(userAge)
	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	}
	request := &pb.CreateUserReq {
		Api: apiVersion,
		User: &pb.User {
			Email: g.PostForm("email"),
			Age: convertUserAge,
			Hobby: g.PostForm("hobby"),
		},
	}

	serverReturn, err := client.CreateUser(context.Background(), request)
	if err != nil {
		handleErrorFromServer(err,g)
		return
	}

	g.JSON( http.StatusOK, serverReturn)

}

// Get User by Email
func runReadUser(g *gin.Context) {
	request := &pb.ReadUserReq {
		Api: apiVersion,
		Email: g.Param("email"),
	}

	serverReturn, err := client.ReadUser(context.Background(), request)
	if err != nil {
		handleErrorFromServer(err, g)
		return
	}
	g.JSON(http.StatusOK, serverReturn)
}

// Update hobby and age of existed user
func runUpdateUser(g *gin.Context) {

	userAge, err := strconv.ParseInt(g.PostForm("age"), 10, 32)
	convertUserAge := int32(userAge)
	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	}
	request := &pb.UpdateUserReq {
		Api: apiVersion,
		User: &pb.User {
			Email: g.PostForm("email"),
			Age: convertUserAge,
			Hobby: g.PostForm("hobby"),
		},
	}

	serverReturn, err := client.UpdateUser(context.Background(), request)
	if err != nil {
		handleErrorFromServer(err, g)
		return
	}
	
	g.JSON(http.StatusOK, serverReturn)
}

//Remove User by email
func runDeleteUser(g *gin.Context) {
	request := &pb.DeleteUserReq {
		Api: apiVersion,
		Email: g.PostForm("email"),
	}

	serverReturn, err := client.DeleteUser(context.Background(), request)
	if err != nil {
		handleErrorFromServer(err, g)
		return
	}

	g.JSON(http.StatusOK, serverReturn)

}

// Get all saved User
func runListUsers(g *gin.Context) {
	request := &pb.ListUsersReq {
		Api: apiVersion,
	}

	serverReturn, err := client.ListUsers(context.Background(), request)
	if err != nil {
		handleErrorFromServer(err, g)
		return
	}
	g.JSON(http.StatusOK, serverReturn)
}

func main() {
	// init grpc connection options
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	// Connect to 127.0.0.1:10001
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	// create a client
	client = pb.NewUserServiceClient(conn)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		// Create User Route
		v1.POST("/user", runCreateUser)
		// Get User Route
		v1.GET("/user/:email", runReadUser)
		// Update User Route
		v1.PUT("/user", runUpdateUser)
		// Delete User Route
		v1.POST("/user/delete", runDeleteUser)
		// List Users Route
		v1.GET("/users", runListUsers)
	}

	router.Run(httpPort)
}
