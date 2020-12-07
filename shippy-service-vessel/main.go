// shippy-service-vessel/main.go
package main

import (
	"log"

	pb "github.com/jipram017/go-shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

const (
	defaultHost = "mongodb://127.0.0.1:27017"
)

func main() {

	service := micro.NewService(micro.Name("go.micro.srv.vessel"))
	service.Init()

	vesselCollection := []*Vessel{
		&Vessel{
			ID:        "vessel001",
			Capacity:  10,
			MaxWeight: 500000,
		},
	}
	repository := &Repository{vesselCollection}

	h := &handler{repository}

	// Register our implementation with
	if err := pb.RegisterVesselServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
