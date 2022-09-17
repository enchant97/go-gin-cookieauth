// Copyright (c) 2022 Leo Spratt. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package extras

import (
	"github.com/enchant97/go-gincookieauth"
	"github.com/gin-gonic/gin"
)

// Replaces gin.Context.HTML() to add the auth data into templates
func TemplateWithAuth(c *gin.Context, code int, name string, obj gin.H) {
	authData := c.MustGet(gincookieauth.GlobalDataKey)
	obj[gincookieauth.GlobalDataKey] = authData
	c.HTML(code, name, obj)
}
