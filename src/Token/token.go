/*
	The token package is used to handle token related options:
	1. Issue token. Use the resource owner pasword credential flow. Get client credentials and pass to authorization server to get token.
	2. Token validation.
	3. Token refresh.
*/

package Token

import (
	"time"
)

type TokenInfo struct {
	AccessToken, RefreshToken, UserId string
	ExpireIn                          int
	IsValid                           bool
	AddTime                           time.Time
}
