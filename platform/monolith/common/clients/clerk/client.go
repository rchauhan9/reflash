package clerk

import (
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwks"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/rchauhan9/reflash/monolith/config"
	"sync"
)

type ClerkClient struct {
	User *user.Client
	Jwks *jwks.Client
}

var (
	clerkClient ClerkClient
	once        sync.Once
)

func GetClient() ClerkClient {
	once.Do(func() {
		conf := config.GetConfig()
		clerkConfig := &clerk.ClientConfig{}
		clerkConfig.Key = clerk.String(conf.Clerk.ApiKey)
		clerkClient = ClerkClient{
			User: user.NewClient(clerkConfig),
			Jwks: jwks.NewClient(clerkConfig),
		}
	})
	return clerkClient
}
