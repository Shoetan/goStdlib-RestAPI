package utils

import  (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func getSecretKey(key string) string {
	err := godotenv.Load(".env")

	if err != nil{
		log.Fatal(err)
	}

	return os.Getenv(key)
}

//get the secret key from the environment

var secretKey = []byte(getSecretKey("SECRET_KEY"))


func CreateToken( email string) (string , error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userEmail": email,
		"expires": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}