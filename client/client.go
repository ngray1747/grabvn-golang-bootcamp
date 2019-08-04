package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "feedback" // Need to add this package to your GO Path
)

const (
	address     = "127.0.0.1:10000"
)

// Get Passenger Feedback by BookingCode or PassengerID
func runGetFeedback(client pb.FeedbackClient, request *pb.GetFeedbackReq) {
	log.Printf("Your Request : %v", request)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.GetFeedback(ctx, request)
	if err != nil {
		log.Fatalf("%v.List Feedback(_) = _, %v", client, err)
	}
	for {
		feedback, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.List Feedback(_) = _, %v", client, err)
		}
		log.Println(feedback)
	}
}

// Add passenger Feedback
func runAddFeedback(client pb.FeedbackClient, request *pb.AddFeedbackReq) {
	log.Printf("Your Request: %v", request)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverReturn, err := client.AddPassengerFeedback(ctx, request)
	if err != nil {
		log.Fatalf("%v.Error occured  _, %v", client, err)
	}

	log.Printf("Status : %d - Message: %v", serverReturn.Status, serverReturn.Message)
	
}
// Remove Passenger Feedback by PassengerID
func runRemoveFeedback(client pb.FeedbackClient, request *pb.RemoveFeedbackReq) {
	log.Printf("Your Request: %v", request)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverReturn, err := client.RemoveFeedback(ctx, request)
	if err != nil {
		log.Fatalf("%v.Error occured  _, %v", client, err)
	}

	log.Printf("Status : %d - Message: %v", serverReturn.DeleteFeedback.Status, serverReturn.DeleteFeedback.Message)

}
func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewFeedbackClient(conn)

	// Get feedback of passenger got ID = 12
	getByPassengerID := &pb.GetFeedbackReq {
		PassengerID: 12,
	}
	runGetFeedback(client, getByPassengerID)

	// Get feedback of booking code = BC2
	getByBookingCode := &pb.GetFeedbackReq {
		BookingCode: "BC2",
	}
	runGetFeedback(client, getByBookingCode)

	// Add feedback to server 
	feedBackContent := &pb.PassengerFeedback {
		BookingCode: "DN",
		PassengerID: 99,
		Feedback: "Same same",
	}
	// Get the new added feedback
	addFeedback := &pb.AddFeedbackReq {
		PsgFeedback: feedBackContent,
	}
	runAddFeedback(client, addFeedback)

	// Looking for the created feedback
	newGetByPassengerID := &pb.GetFeedbackReq {
		PassengerID: 99,
	}
	runGetFeedback(client, newGetByPassengerID)

	// Delete created feedback by customerID
	deleteFeedback := &pb.RemoveFeedbackReq {
		PassengerID: 96,
	}
	runRemoveFeedback(client, deleteFeedback)
}