package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rchauhan9/reflash/monolith/common/auth"
	"net/http"
	"strings"
	"sync"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	clerkclient "github.com/rchauhan9/reflash/monolith/common/clients/clerk"
)

func RequiresAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		clerkClient := clerkclient.GetClient()

		// Get the session JWT from the Authorization header
		authHeader := c.GetHeader("Authorization")
		sessionToken := strings.TrimPrefix(authHeader, "Bearer ")

		if sessionToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"access": "unauthorized"})
			c.Abort()
			return
		}

		store := GetJWKStore()

		// Attempt to get the JSON Web Key from your store.
		jwk := store.GetJWK()
		if jwk == nil {
			// Decode the session JWT to get the Key ID
			unsafeClaims, err := jwt.Decode(c.Request.Context(), &jwt.DecodeParams{
				Token: sessionToken,
			})
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"access": "unauthorized"})
				c.Abort()
				return
			}

			// Fetch the JSON Web Key using the Key ID
			jwk, err = jwt.GetJSONWebKey(c.Request.Context(), &jwt.GetJSONWebKeyParams{
				KeyID:      unsafeClaims.KeyID,
				JWKSClient: clerkClient.Jwks,
			})
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"access": "unauthorized"})
				c.Abort()
				return
			}
			// Cache the JWK
			store.SetJWK(jwk)
		}

		// Verify the session
		claims, err := jwt.Verify(c.Request.Context(), &jwt.VerifyParams{
			Token: sessionToken,
			JWK:   jwk,
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"access": "unauthorized"})
			c.Abort()
			return
		}

		clerkUser, err := clerkClient.User.Get(c.Request.Context(), claims.Subject)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
			c.Abort()
			return
		}
		user, err := auth.GetUserFromClerkUser(clerkUser)
		fmt.Printf("User: %v\n", user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"access": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

type JWKStore interface {
	GetJWK() *clerk.JSONWebKey
	SetJWK(*clerk.JSONWebKey)
}

// inMemoryJWKStore is a thread-safe in-memory implementation of JWKStore
type inMemoryJWKStore struct {
	mu  sync.RWMutex
	jwk *clerk.JSONWebKey
}

var inMemoryJWKStoreInstance inMemoryJWKStore

// NewJWKStore creates a new instance of an in-memory JWK store
func GetJWKStore() JWKStore {
	return &inMemoryJWKStoreInstance
}

// GetJWK returns the currently stored JSON Web Key
func (s *inMemoryJWKStore) GetJWK() *clerk.JSONWebKey {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.jwk
}

// SetJWK sets a new JSON Web Key into the store
func (s *inMemoryJWKStore) SetJWK(jwk *clerk.JSONWebKey) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jwk = jwk
}
