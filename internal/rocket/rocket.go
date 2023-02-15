//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/saiful-4321/grpcGO/internal/rocket Store

package rocket

import "context"

// Rocket-should contain the definition of out rocket
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// Store - Defines the interface we expect
// Our database implementation to follow
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service - our rocket service is responsible for
// updating the rocket
type Service struct {
	Store Store
}

// New - returns a new instance of our rocket service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketByID - retrives a rocket based on ID from the store
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)

	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

// InsertRocket - used to insert a new rocket into the store
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)

	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

// DeleteRocket - Delete a rocket from the inventory
func (s Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
