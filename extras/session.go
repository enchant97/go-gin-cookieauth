// Copyright (c) 2022 Leo Spratt. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package extras

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Make a session quickly
func MakeSession(name string, secretKey []byte) gin.HandlerFunc {
	store := cookie.NewStore(secretKey)
	return sessions.Sessions(name, store)
}
