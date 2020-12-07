package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	pb "github.com/jipram017/go-shippy/shippy-service-consignment/proto/consignment"
	"github.com/micro/go-micro/metadata"
	micro "github.com/micro/go-micro/v2"
)

const (
	defaultFilename = "consignment.json"
	defaultToken    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImVtYWlsIjoiamlwcmFtMDE5QGdtYWlsLmNvbSIsInBhc3N3b3JkIjoidG9ob2t1MjAxNCJ9LCJleHAiOjE1MDAwLCJpc3MiOiJzaGlwcHkuc2VydmljZS51c2VyIn0.3f2YEW8hAoqhwvKyVt6p7uINVSqwdDjvA5J-uY0krKo"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {

	//cmd.Init()

	token := defaultToken
	file := defaultFilename

	// Create a new context which contains our given token.
	// This same context will be passed into both the calls we make
	// to our consignment-service.
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"Token": token,
	})

	service := micro.NewService()

	// Initialize service
	service.Init()

	client := pb.NewShippingService("go.micro.srv.consignment", service.Client())

	// if len(os.Args) < 3 {
	// 	log.Fatal(errors.New(
	// 		"Not enough arguments, expecing file and token",
	// 	))
	// }

	// file = os.Args[1]
	// token = os.Args[2]

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse error: %v", err)
	}

	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	// Second call
	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
