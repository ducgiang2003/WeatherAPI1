package config

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

// Initialize key variable
var S *securecookie.SecureCookie

func InitSecureCookie() {
	// ... existing code ...
	store := sessions.NewCookieStore(securecookie.GenerateRandomKey(13))
	gothic.Store = store

}
