package auth

import (
	"encoding/json"
	"fmt"
	"github.com/clerk/clerk-sdk-go/v2"
)

type User struct {
	ID string `json:"user_id"`
}

func GetUserFromClerkUser(user *clerk.User) (User, error) {
	userObj := User{}
	err := json.Unmarshal(user.UnsafeMetadata, &userObj)
	if err != nil {
		return User{}, fmt.Errorf("failed to unmarshal user metadata: %w", err)
	}
	return userObj, nil
}
