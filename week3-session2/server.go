package main

import (
	"context"
	"encoding/json"
	pb "user" // Need to add this package to your GO Path
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

)

var (
	apiVersion					= "v1"
	port       					= flag.Int("port", 10001, "The server port")
	statusSuccess 	int32		= 1
	statusFail 		int32		= 2
)


type userServer struct {
	mutex sync.Mutex // Make sure only 1 can add / delete at a time
	savedUser []*pb.User
}


// checkAPI checks if the API version requested by client is supported by server
func (s *userServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// Check if user is email is already exists
func (s *userServer) checkEmail(email string) error {
	for _, value := range s.savedUser {
		if email == value.GetEmail() {
			return status.Errorf(codes.AlreadyExists, "'%s' is already existed", email)
		}
	}
	return nil
}

// Remove element from a slice by index, used to remove passenger feedback
func removeIndexFromSlice(slice []*pb.User, index int) []*pb.User {
	zeroValue := &pb.User{}
	copy(slice[index:], slice[index+1:]) // Shift a[i+1:] left one index.
	slice[len(slice)-1] = zeroValue      // Erase last element (write zero value).
	slice = slice[:len(slice)-1]         // Truncate slice.
	return slice
}

// Implement create new User
func (s *userServer) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {

	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}
	// Check if the email is already exists
	if err := s.checkEmail(req.User.GetEmail()); err != nil {
		return nil, err
	}

	if req.User.GetEmail() == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"Email field cannot be empty")
	}
	// Add created time to User Req.
	timeNow := time.Now()
	timeAsString := timeNow.Format("2006-01-02 15:04:05")
	req.User.CreatedAt = timeAsString

	// Add new User to savedUser slice
	// Make sure only one can process at a time
	s.mutex.Lock()
	s.savedUser = append(s.savedUser, req.User)
	s.mutex.Unlock()

	return &pb.CreateUserRes{
		Api: apiVersion,
		Result:  "User added successfully!",
	}, nil
	
}
// Get user detail by email
func (s *userServer) ReadUser(ctx context.Context, req *pb.ReadUserReq) (*pb.ReadUserRes, error) {

	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}

	for _, value := range s.savedUser {
		if req.GetEmail() == value.GetEmail() {
			return &pb.ReadUserRes{
				Api: apiVersion,
				User:  value,
			}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "'%s'", req.GetEmail())
}
// Update user age and hobby
func (s *userServer) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {

	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}

	for _, value := range s.savedUser {
		if req.User.GetEmail() == value.GetEmail() {
			if req.User.Hobby != "" {
				value.Hobby = req.User.Hobby
			}
			if req.User.Age > 0 {
				value.Age = req.User.Age
			}
			return &pb.UpdateUserRes{
				Api: apiVersion,
				Status:  statusSuccess,
			}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "'%s'", req.User.GetEmail())
}
// Delete user by Email
func (s *userServer) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {

	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}

	for index, value := range s.savedUser {
		if req.GetEmail() == value.GetEmail() {
			s.mutex.Lock()
			s.savedUser = removeIndexFromSlice(s.savedUser, index)
			s.mutex.Unlock()
				
			return &pb.DeleteUserRes{
				Api: apiVersion,
				Status:  statusSuccess,
			}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "'%s'", req.GetEmail())
}
// List all User
func (s *userServer) ListUsers(ctx context.Context, req *pb.ListUsersReq) (*pb.ListUsersRes, error) {

	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}

	return &pb.ListUsersRes{
		Api: apiVersion,
		Users:  s.savedUser,
	}, nil 
}

// loadFeedbacks loads feedbacks from a JSON file.
func (s *userServer) loadExampleUsers() {
	var data []byte
	data = exampleData
	if err := json.Unmarshal(data, &s.savedUser); err != nil {
		log.Fatalf("Erorr in unmarshal json file: %v", err)
	}
}

func newUserServer() *userServer {
	s := &userServer{}
	s.loadExampleUsers()
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Unable to listen to : %v", err)
	}

	// Init the gRPC server options - empty for now
	opts := []grpc.ServerOption{}
	// Create gRPC server
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterUserServiceServer(grpcServer, newUserServer())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Create a channel to receive OS signals
	c := make(chan os.Signal)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-c

	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	grpcServer.Stop()
	lis.Close()
}

// Create a exampleData
var exampleData = []byte(`[{
	"Email": "ndhduy7798@gmail.com",
	"Hobby": "nope",
	"Age" : 23,
	"CreatedAt" : "2019-08-04 22:56:00"
}, {
	"Email": "example1@gmail.com",
	"Hobby": "singing",
	"Age" : 20,
	"CreatedAt" : "2019-08-04 22:56:00"
}, {
	"Email": "example2@gmail.com",
	"Hobby": "coding",
	"Age" : 25,
	"CreatedAt" : "2019-08-04 22:56:00"
} , {
	"Email": "example3@gmail.com",
	"Hobby": "footbal",
	"Age" : 21,
	"CreatedAt" : "2019-08-04 22:56:00"
}, {
	"Email": "example4@gmail.com",
	"Hobby": "outing",
	"Age" : 20,
	"CreatedAt" : "2019-08-04 22:56:00"
}]`)
