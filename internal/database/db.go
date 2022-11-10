package database

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/mitchellh/hashstructure"
)

var ctx = context.Background()

type Dbredis struct {
	db    *redis.Client
	mutex sync.Mutex
}

/*InitRedis - Init a new Redis to use*/
func InitRedis() (newdbase *Dbredis) {
	host := os.Getenv("hostredis")
	port := os.Getenv("portredis")
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Dbredis{
		db: rdb,
	}
}
func (c *Dbredis) Set(input []string, valid bool) {
	opts := hashstructure.HashOptions{IgnoreZeroValue: false, UseStringer: true}
	hashid, err := hashstructure.Hash(input, &opts)
	value := c.Get(strconv.Itoa(int(hashid)))
	if len(value) > 0 {
		return
	}
	var valueToSave []interface{}
	valueToSave = append(valueToSave, input)
	valueToSave = append(valueToSave, valid)
	toredis, err := json.Marshal(valueToSave)
	err = c.db.Set(ctx, strconv.Itoa(int(hashid)), toredis, 0).Err()
	if err != nil {
		panic(err)
	}

}

func (c *Dbredis) Get(hash string) (returned string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, err := c.db.Get(ctx, hash).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return ""
		} else {
			panic("redis error: " + err.Error())
		}
	}
	return val
}

var cursor uint64

func (c *Dbredis) ReturnStats() (valid int, notvalid int, ratio float64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	var keys []string
	var err error
	keys, cursor, err = c.db.Scan(ctx, cursor, "", 0).Result()
	if err != nil {

		panic(err)
	}
	qtditens := len(keys)
	_ = qtditens
	itensvalid := 0
	itensNotValid := 0
	for _, key := range keys {
		val, err := c.db.Get(ctx, key).Result()
		if err != nil {

			panic(err)
		}
		var raw []interface{}
		if err := json.Unmarshal([]byte(val), &raw); err != nil {
			panic(err)
		}
		if raw[1] == true {
			itensvalid++
		} else {
			itensNotValid++
		}

	}
	if itensvalid == 0 {
		return itensvalid, itensNotValid, 1.0
	}
	total := float64(itensvalid + itensNotValid)
	ratio = float64(itensvalid) / total

	return itensvalid, itensNotValid, ratio

}

func (c *Dbredis) Reset() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	_, err := c.db.FlushAll(ctx).Result()
	if err != nil {
		panic(err.Error())
	}
}
