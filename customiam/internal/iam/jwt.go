package iam

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

// https://docs.apigee.com/api-platform/reference/policies/oauth-http-status-code-reference
var InvalidAccessToken = errors.New("invalid access token")
var AccessTokenExpired = errors.New("access token expired")

// JwtAuthentication for jwt authorization
type JwtAuthentication struct{}

// GenerateJwtToken generate jwt token with userclaim
func (j *JwtAuthentication) GenerateJwtToken(claim *UserClaim) (*JwtToken, error) {
	// create access
	claimMap, err := claim.ConvertMap()
	if err != nil {
		return nil, err
	}

	// access token
	accessTokenClaim := jwt.MapClaims{}
	for k, v := range claimMap {
		accessTokenClaim[k] = v
	}

	accessTokenClaim["exp"] = time.Now().Add(time.Minute*30).Unix() * 1000
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaim).
		SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	// refresh token
	refreshTokenClaim := jwt.MapClaims{}
	refreshTokenClaim["exp"] = time.Now().Add(time.Hour*24*7).Unix() * 1000
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaim).
		SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &JwtToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// ConvertTokenUserCliam tokenstring to usercliam
func (j *JwtAuthentication) ConvertTokenUserCliam(Accesstoken string) (*UserClaim, error) {
	// https://pkg.go.dev/github.com/golang-jwt/jwt#example-Parse-ErrorChecking
	parsedToken, err := jwt.Parse(Accesstoken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.WithError(err).Error("JWT parsing error")
		if ve, ok := err.(*jwt.ValidationError); ok {
			// Token is either expired or not active yet
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, AccessTokenExpired
			}
		}
		return nil, InvalidAccessToken
	}

	// alg check
	// Don't forget to validate the alg is what you expect:
	// https://pkg.go.dev/github.com/golang-jwt/jwt#example-Parse-Hmac
	if jwt.SigningMethodHS256.Alg() != parsedToken.Header["alg"] {
		log.Error(fmt.Sprintf("Error: jwt token is expected %s signing method but token specified %s",
			jwt.SigningMethodHS256.Alg(), parsedToken.Header["alg"]))
	}

	if !parsedToken.Valid {
		return nil, InvalidAccessToken
	}

	claimInfo, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		log.Error("Can't get jwt.MapClaims")
	}

	userClaim, err := NewUserClaim(claimInfo)
	if err != nil {
		return nil, err
	}

	return userClaim, nil
}

// JwtToken contain token info
type JwtToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// UserClaim jwt use cliam
type UserClaim struct {
	ID   int64  `json:"id"`
	Name string `json:"string"`
}

// ConvertMap convert UserClaim struct to hash map
func (u *UserClaim) ConvertMap() (map[string]interface{}, error) {
	bytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	var resultMap map[string]interface{}
	if err := json.Unmarshal(bytes, &resultMap); err != nil {
		return nil, err
	}

	return resultMap, nil
}

// NewUserClaim convert hash map to UserClaim
func NewUserClaim(mapUserCliam map[string]interface{}) (*UserClaim, error) {
	bytes, err := json.Marshal(mapUserCliam)
	if err != nil {
		return nil, err
	}

	var claim UserClaim
	err = json.Unmarshal(bytes, &claim)
	if err != nil {
		return nil, err
	}

	return &claim, nil
}
