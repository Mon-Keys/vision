package cookie

import (
	"time"

	"github.com/google/uuid"
	"github.com/perlinleo/vision/internal/domain"
	"github.com/valyala/fasthttp"
)

// TO CONFIG
const CookieDefaultMaxAge = 3600000

// TO CONFIG
const WebSiteURL = "http://vision.leonidperl.in:5000"

func CreateAuthSessionUUID(userID int32) *domain.UserSession {
	session := new(domain.UserSession)

	session.UserID = userID
	session.Expiration = CookieDefaultMaxAge
	session.Cookie = uuid.New().String()

	return session
}

func CreateAccountSessionUUID(userID int32) *domain.AccountSession {
	session := new(domain.AccountSession)

	session.AccountID = userID
	session.Expiration = CookieDefaultMaxAge
	session.Cookie = uuid.New().String()

	return session
}

func CreateFastHTTPCookie(cookie string, name string) fasthttp.Cookie {
	cook := fasthttp.Cookie{}
	cook.SetExpire(time.Now().Add(360 * time.Hour))
	cook.SetKey(name)
	cook.SetValue(cookie)
	// cook.SetMaxAge(int(s.Expiration))
	// cook.SetDomain(WebSiteURL)
	cook.SetPath(("/api/v1"))
	cook.SetHTTPOnly(true)
	cook.SetSecure(true)
	// cook.SetSecure(false)
	cook.SetSameSite(fasthttp.CookieSameSiteNoneMode)

	return cook
}
