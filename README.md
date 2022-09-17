# go-gincookieauth
This Go package providing cookie authentication middleware for Gin. CookieAuth provides simple (and basic) access to store user login in the session cookies.

## Example

```go
func GetLogout(c *gin.Context) {
    // ...
    gincookieauth.LogoutUser(c)
    // ...
}

func PostLogin(c *gin.Context) {
    // Do some auth checks ...
    gincookieauth.LoginUser(c, userID)
    // ...
}

func main() {
	r := gin.Default()

	store := cookie.NewStore(secretKey)
	session := sessions.Sessions("AUTH", store)
	cookieAuth := gincookieauth.CookieAuth(false)
    r.Use(session)
	r.Use(cookieAuth)

    // ...
```
