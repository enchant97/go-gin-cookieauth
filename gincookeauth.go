// Copyright (c) 2022 Leo Spratt. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gincookieauth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	// The key used for storing AuthData
	GlobalDataKey = "AuthData"
	// Name of key to use inside session
	SessionUserIDKey = "AuthUserID"
)

// Used to configure the CookieAuth middleware
type CookieAuthConfig struct {
	// Whether auth is required on all routes
	AuthRequired bool
}

// The global authData that is set for every context
type AuthData struct {
	// The current users id, or nil
	UserID *interface{}
	// Whether current user is authenticated
	IsAuthenticated bool
}

// Login a user (stores in session)
func LoginUser(c *gin.Context, userID interface{}) error {
	session := sessions.Default(c)
	session.Set(SessionUserIDKey, userID)
	return session.Save()
}

// Get the stored user (from session)
func GetUserID(c *gin.Context) *interface{} {
	session := sessions.Default(c)
	if userID := session.Get(SessionUserIDKey); userID != nil {
		return &userID
	}
	return nil
}

// Logout a user, removing stored auth or clearing all session data
func LogoutUser(c *gin.Context, clearAll bool) error {
	session := sessions.Default(c)
	if clearAll {
		session.Clear()
	} else {
		session.Delete(SessionUserIDKey)
	}
	return session.Save()
}

// CookieAuth middleware
// allowing for storing user login in a users session
func CookieAuth(config CookieAuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := GetUserID(c)
		if userID == nil && config.AuthRequired {
			// deny access if authentication is marked as required
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// set global authData data
		c.Set(GlobalDataKey, AuthData{
			UserID:          userID,
			IsAuthenticated: userID != nil,
		})
		c.Next()
	}
}
