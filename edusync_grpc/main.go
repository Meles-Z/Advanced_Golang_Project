package main

import (
	"context"
	"log"
	"time"

	"github.com/meles-z/edusysnc_grpc/proto/student"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the gRPC server.
	address := "localhost:50051" // Replace with your server's address
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := student.NewStudentServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Example: Creating a student
	createReq := &student.CreateStudentRequest{
		Student: &student.Student{
			Name:        "John Doe",
			Email:       "john.doe@example.com",
			PhoneNumber: "123-456-7890",
		},
	}

	createResp, err := c.CreateStudent(ctx, createReq)
	if err != nil {
		log.Fatalf("could not create student: %v", err)
	}
	log.Printf("Create Student Response: %s", createResp)

	// Example: Getting all students
	getAllReq := &student.GetAllStudentsRequest{}
	getAllResp, err := c.GetAllStudent(ctx, getAllReq)
	if err != nil {
		log.Fatalf("could not get all students: %v", err)
	}
	log.Printf("All Students: %s", getAllResp)
}