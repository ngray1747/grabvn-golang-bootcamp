package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "user" // Need to add this package to your GO Path
)

const (
	address     = "127.0.0.1:10001"
	apiVersion = "v1"
)

// Handle error return from server
func handleErrorFromServer(err error) {
	grpcErr, ok := status.FromError(err)
	if ok {
		switch grpcErr.Code() {
		case codes.NotFound:
			log.Println("Not Found: ", grpcErr.Message())
			// return
		case codes.Unimplemented:
			log.Println("Api not found: ", grpcErr.Message())
		case codes.AlreadyExists:
			log.Println("Email already exists: ", grpcErr.Message())
		case codes.InvalidArgument:
			log.Println("The input value is incorrect: ", grpcErr.Message())
		default:
			log.Println("Unexpected error: ", grpcErr.Code())
		}
	} else {
		log.Println("Failed to call server", err)
	}
}
// Add User
func runCreateUser(client pb.UserServiceClient, request *pb.CreateUserReq) {
	log.Printf("Your Request: %v", request)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverReturn, err := client.CreateUser(ctx, request)
	if err != nil {
		handleErrorFromServer(err)
		return
	}

	log.Printf("Status : %s", serverReturn.Result)
	
}

// Get User by Email
func runReadUser(client pb.UserServiceClient, request *pb.ReadUserReq) {
	log.Printf("Your Request: %v", request)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverReturn, err := client.ReadUser(ctx, request)
	if err != nil {
		handleErrorFromServer(err)
		return
	}
	
	log.Printf("Read result: <%+v>\n\n", serverReturn)
	
}

// Update hobby and age of existed user
func runUpdateUser(client pb.UserServiceClient, request *pb.UpdateUserReq) {
	log.Printf("Your Request: %v", request)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverReturn, err := client.UpdateUser(ctx, request)
	if err != nil {
		handleErrorFromServer(err)
		return
	}
	
	log.Printf("Update result: <%+v>\n\n", serverReturn)
}

// Remove User by email
func runDeleteUser(client pb.UserServiceClient, request *pb.DeleteUserReq) {
	log.Printf("Your Request: %v", request)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverReturn, err := client.DeleteUser(ctx, request)
	if err != nil {
		handleErrorFromServer(err)
		return
	}

	log.Printf("Delete User result : %d", serverReturn.Status)

}

// Get all saved User
func runListUsers(client pb.UserServiceClient, request *pb.ListUsersReq) {
	log.Printf("Your Request: %v", request)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverReturn, err := client.ListUsers(ctx, request)
	if err != nil {
		handleErrorFromServer(err)
		return
	}
	for userIndex, userDetail := range serverReturn.Users {
		log.Printf("Users %d : Email -> '%s', Hobby: -> '%s', Age -> '%d', 'CreatedAt' -> '%s' ", userIndex, userDetail.Email, userDetail.Hobby, userDetail.Age, userDetail.CreatedAt)
	}
	
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
	client := pb.NewUserServiceClient(conn)

	//Add User to server 
	addUserReq := &pb.CreateUserReq {
		Api: apiVersion,
		User: &pb.User{
			Email: "danhnguyen@gmail.com",
			Hobby: "Reading",
			Age: 23,
		},
	}
	runCreateUser(client, addUserReq)

	// Read User has email "example4@gmail.com"
	readUserReq := &pb.ReadUserReq {
		Api: apiVersion,
		Email: "example4@gmail.com",
	}
	runReadUser(client, readUserReq)
	
	// Update User has email "example1@gmail.com"
	updateUserReq := &pb.UpdateUserReq {
		Api: apiVersion,
		User: &pb.User {
			Email: "example1@gmail.com",
			Hobby: "not singing",
			Age : 21,
		},
	}
	runUpdateUser(client, updateUserReq)

	// Delete User has email "example2@gmail.com"
	deleteUserReq := &pb.DeleteUserReq {
		Api: apiVersion,
		Email: "example2@gmail.com",
	}
	runDeleteUser(client, deleteUserReq)

	// Get all saved User from server
	listUserReq := &pb.ListUsersReq {
		Api: apiVersion,
	}
	runListUsers(client, listUserReq)
}