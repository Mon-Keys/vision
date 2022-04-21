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
const WebSiteURL = "localhost:5000"

func CreateAuthSessionUUID(userID string) *domain.Session {
	session := new(domain.Session)

	session.ID = userID
	session.Expiration = CookieDefaultMaxAge
	session.Cookie = uuid.New().String()

	return session
}

func CreateFastHTTPCookie(s domain.Session) fasthttp.Cookie {
	cook := fasthttp.Cookie{}
	cook.SetKey("sessionid")
	cook.SetValue(s.Cookie)
	cook.SetMaxAge(int(s.Expiration))
	cook.SetDomain(WebSiteURL)
	cook.SetPath(("/"))
	cook.SetHTTPOnly(true)
	cook.SetExpire(time.Now().Add(10 * time.Minute))
	cook.SetSecure(false)

	return cook
}
