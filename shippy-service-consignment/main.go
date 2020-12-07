// shippy/shippy-service-consignment/main.go
package main

import (
	"log"

	pb "github.com/jipram017/go-shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/jipram017/go-shippy/shippy-service-vessel/proto/vessel"

	"github.com/micro/go-micro/v2"
)

const (
	defaultHost = "mongodb://127.0.0.1:27017"
)

// AuthWrapper is a high-order function which takes a HandlerFunc
// and returns a function, which takes a context, request and response interface.
// The token is extracted from the context set in our consignment-cli, that
// token is then sent over to the user service to be validated.
// If valid, the call is passed along to the handler. If not,
// an error is returned.
// func AuthWrapper(fn servo.HandlerFunc) servo.HandlerFunc {
// 	return func(ctx context.Context, req servo.Request, resp interface{}) error {
// 		meta, ok := metadata.FromContext(ctx)
// 		var token string
// 		if !ok {
// 			// return errors.New("no auth meta-data found in request")

// 			// Instead of return error, we temporarily hardcode it if context is not propagated
// 			token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImVtYWlsIjoicHBwcm1uMkBnbWFpbC5jb20iLCJwYXNzd29yZCI6ImFwanAyMSJ9LCJleHAiOjE1OTM3MDM2NjgsImlzcyI6InNoaXBweS5zZXJ2aWNlLnVzZXIifQ.iYO7KTb-Pz5ZxynM79XHDEQwCykxZTqP-XYAjhjzfZ4"
// 		} else {
// 			// Note this is now uppercase (not entirely sure why this is...)
// 			token = meta["Token"]
// 		}

// 		log.Println("Authenticating with token: ", token)
// 		// Auth here
// 		authClient := userService.NewUserService("go.micro.srv.user", client.DefaultClient)
// 		_, err := authClient.ValidateToken(context.Background(), &userService.Token{
// 			Token: token,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		err = fn(ctx, req, resp)
// 		return err
// 	}
// }

func main() {
	repository := &Repository{}
	service := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		//micro.WrapHandler(AuthWrapper),
	)

	// Initialize service
	service.Init()
	vesselClient := vesselProto.NewVesselService("go.micro.srv.vessel", service.Client())

	h := &handler{repository, vesselClient}

	// Register our implementation with
	if err := pb.RegisterShippingServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
