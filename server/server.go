package main

import (
	"context"
	"encoding/json"
	pb "feedback" // Need to add this package to your GO Path
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"

	"google.golang.org/grpc"
)

// gRPC server port
var port	= 10000


type feedbackServer struct {
	mutex sync.Mutex
	feedbackRecord []*pb.PassengerFeedback
}

// GetFeedback lists all feedbacks match the BookingCode or PassengerID
func (s *feedbackServer) GetFeedback(req *pb.GetFeedbackReq, stream pb.Feedback_GetFeedbackServer) error {
	for _, value := range s.feedbackRecord {
		if (req.GetPassengerID() != 0 && req.GetPassengerID() == value.PassengerID) || (len(req.GetBookingCode()) > 0 && req.GetBookingCode() == value.BookingCode) {
			returnValue := &pb.GetFeedbackRes{
				PsgFeedback: value,
			}
			
			if err := stream.Send(returnValue); err != nil {
				return err
			}
		}
	}
	return nil
}

// Check if the BookingCode is used in file before creating the feedback
func checkForCreatedFeedback(savedFeedback []*pb.PassengerFeedback, bookingCode string) bool {
	for _, value := range savedFeedback {
		if bookingCode == value.GetBookingCode() {
			return true
		}
	}
	return false
}

// Create passenger feedback
func (s *feedbackServer) AddPassengerFeedback(ctx context.Context, req *pb.AddFeedbackReq) (*pb.AddFeedbackRes, error) {
	feedBack := req.GetPsgFeedback()
	isBookingCodeUsed := checkForCreatedFeedback(s.feedbackRecord, feedBack.GetBookingCode())

	returnValue := &pb.AddFeedbackRes{
		Status:  2,
		Message: "The booking code is incorrect, try again later!",
	}
	if isBookingCodeUsed == true {
		return returnValue, nil
	}
	// Make sure only one can process at a time
	s.mutex.Lock()
	s.feedbackRecord = append(s.feedbackRecord, feedBack)
	s.mutex.Unlock()

	returnValue.Status, returnValue.Message = 1, "Feedback created successfully!"

	return returnValue, nil
}

// Remove element from a slice by index, used to remove passenger feedback
func removeIndexFromSlice(slice []*pb.PassengerFeedback, index int) []*pb.PassengerFeedback {
	zeroValue := &pb.PassengerFeedback{}
	copy(slice[index:], slice[index+1:]) // Shift a[i+1:] left one index.
	slice[len(slice)-1] = zeroValue      // Erase last element (write zero value).
	slice = slice[:len(slice)-1]         // Truncate slice.
	return slice
}

// Remove passenger feedback by passengerID
func (s *feedbackServer) RemoveFeedback(ctx context.Context, req *pb.RemoveFeedbackReq) (*pb.RemoveFeedbackRes, error) {
	deleteFeedbackCont := &pb.AddFeedbackRes{
		Status:  2,
		Message: "Feedback Deleted Fail!",
	}
	returnValue := &pb.RemoveFeedbackRes{
		DeleteFeedback: deleteFeedbackCont,
	}
	for index, value := range s.feedbackRecord {

		if req.GetPassengerID() > 0 && req.GetPassengerID() == value.PassengerID {
			
			s.mutex.Lock()
			s.feedbackRecord = removeIndexFromSlice(s.feedbackRecord, index)
			s.mutex.Unlock()
			deleteFeedbackCont.Status, deleteFeedbackCont.Message = 1, "Feedback Deleted!"
			return returnValue, nil
		}
	}
	return returnValue, nil
}

// loadFeedbacks loads feedbacks from a JSON file.
func (s *feedbackServer) loadFeedbacks() {
	var data []byte
	data = exampleData
	if err := json.Unmarshal(data, &s.feedbackRecord); err != nil {
		log.Fatalf("Erorr in unmarshal json file: %v", err)
	}
}

func newServer() *feedbackServer {
	s := &feedbackServer{}
	s.loadFeedbacks()
	return s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("Unable to listen to : %v", err)
	}

	// Init the gRPC server options - empty for now
	opts := []grpc.ServerOption{}
	// Create gRPC server
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterFeedbackServer(grpcServer, newServer())

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
	"bookingCode": "BC1",
	"passengerID": 17,
	"feedback" : "Good"
}, {
	"bookingCode": "BC2",
	"passengerID": 17,
	"feedback" : "Average"
}, {
	"bookingCode": "BC3",
	"passengerID": 12,
	"feedback" : "Love it"
} , {
	"bookingCode": "BC4",
	"passengerID": 17,
	"feedback" : "Poor"
}, {
	"bookingCode": "BC5",
	"passengerID": 96,
	"feedback" : "Good work!"
}]`)
