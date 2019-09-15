package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/go-redis/redis"
)

func OpenRedisClient() *redis.Client {
	var host string = os.Getenv("REDIS_HOST")
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Failed to start redis server")
	} else {
		log.Printf("Redis started on host %s\nPong response: %s\n", host, pong)
	}

	return client
}

func IssueToken(uuid string) error {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := Claims{
		UUID: uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	client := OpenRedisClient()
	keyString := uuid + "_token"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := os.Getenv("JWT_TOKEN_SIGNING_KEY")
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err == nil {
		err := client.Set(keyString, tokenString, time.Since(expirationTime)).Err()
		if err != nil {
			log.Fatal("Cannot set token to the redis storage: ", err)
		}
		fmt.Println("SETTING: ", tokenString, " WITH DURATION OF ", time.Since(expirationTime))
		readFromDB, err := client.Get(keyString).Result()
		fmt.Println("USER UUID AND CLIENT GET:  ", uuid, readFromDB)
	}

	return err
}
