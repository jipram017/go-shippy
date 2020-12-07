// shippy-service-vessel/repository.go
package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/jipram017/go-shippy/shippy-service-vessel/proto/vessel"
)

type repository interface {
	FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error)
	Create(ctx context.Context, vessel *Vessel) error
}

type Repository struct {
	vessels []*Vessel
}

type Specification struct {
	Capacity  int32
	MaxWeight int32
}

func MarshalSpecification(spec *pb.Specification) *Specification {
	return &Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

func UnmarshalSpecification(spec *Specification) *pb.Specification {
	return &pb.Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

func MarshalVessel(vessel *pb.Vessel) *Vessel {
	return &Vessel{
		ID:        vessel.Id,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerID:   vessel.OwnerId,
	}
}

func UnmarshalVessel(vessel *Vessel) *pb.Vessel {
	return &pb.Vessel{
		Id:        vessel.ID,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerId:   vessel.OwnerID,
	}
}

type Vessel struct {
	ID        string
	Capacity  int32
	MaxWeight int32
	Name      string
	Available bool
	OwnerID   string
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repository *Repository) FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error) {
	log.Println(repository.vessels)
	for _, vessel := range repository.vessels {
		if vessel.Capacity >= spec.Capacity && vessel.MaxWeight >= spec.MaxWeight {
			return vessel, nil
		}
	}

	return nil, errors.New("Could not find matching vessel")
}

// Create a new vessel
func (repository *Repository) Create(ctx context.Context, vessel *Vessel) error {
	repository.vessels = append(repository.vessels, vessel)
	return nil
}
