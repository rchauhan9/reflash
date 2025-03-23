package auth

import (
	"context"
	"fmt"

	"github.com/clerkinc/clerk-sdk-go/clerk"
)

type Service interface {
	HasPermission(ctx context.Context, userID, permission string) (bool, error)
}

func NewService(client clerk.Client) Service {
	return service{client: &client}
}

type service struct {
	client *clerk.Client
}

func (s service) HasPermission(ctx context.Context, userID, permission string) (bool, error) {
	fmt.Printf("Checking permission %s for user %s", permission, userID)
	return true, nil
}
