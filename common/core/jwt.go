package core

var JWT_SECRET []byte

func NewJWT(jwtSecret string) {
	JWT_SECRET = []byte(jwtSecret)
}
