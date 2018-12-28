package basic_date

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "118.190.65.33:6372",
		Password: "Jingbanyun426!426", // no password set
		DB:       0,                   // use default DB
	})
}

func TestRedis(t *testing.T) {
	fmt.Println(client.Get("oks"))
}
