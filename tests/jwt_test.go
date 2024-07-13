package tests

import (
	`testing`
	`time`
	
	`github.com/golang-jwt/jwt/v5`
)

func TestJwt(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS512)
	token.Claims = jwt.MapClaims{
		"foo": "bar",
		"exp": time.Date(2024, 7, 6, 0, 0, 0, 0, time.UTC).Unix(),
		"iat": time.Date(2024, 7, 5, 0, 0, 0, 0, time.UTC).Unix(),
	}
	value, err := token.SignedString([]byte("admin"))
	if err != nil {
		t.Error(err)
	}
	v, err := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
		return []byte("admin"), nil
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(v.Claims.GetIssuedAt())
	t.Log(v.Claims)
}
