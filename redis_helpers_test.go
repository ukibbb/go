package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func TestRedisHelpersCreateValues(t *testing.T) {
	user := User{
		Id:        uuid.New(),
		Username:  "TestUser",
		Email:     "test@email.com",
		Password:  "UserPassword",
		CreatedAt: time.Now().Local().Format(time.RFC822),
		IsActive:  true,
		Role:      "role",
	}
	h := NewRedisHelpers[User]()
	v, _ := h.CreateValues(user)

	s, _ := h.RetriveStruct(v)

	if s != user {
		t.Error("stuct retrieved from created values is not the same")
	}

	if !reflect.DeepEqual(s, user) {
		t.Error("stuct retrieved from created values is not the same")
	}
}

type RedisUser struct {
	Username string `redis:"username"` // only for how keys are stored
	Email    string `redis:"email"`
	Role     int    `redis:"role"`
	IsActive bool   `redis:"isActive"`
}

func TestRedisSettingGetting(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":5000",
	})

	user := RedisUser{
		Username: "user",
		Email:    "uki@bb.pl",
		Role:     0,
		IsActive: true,
	}
	id := uuid.NewString()
	if err := rdb.HSet(context.Background(), id, user).Err(); err != nil {
		t.Error("failed to set user")
	}

	var nUser RedisUser

	// u, err := rdb.HGetAll(context.Background(), id).Result() // field-value
	if err := rdb.HGetAll(context.Background(), id).Scan(&nUser); err != nil {
		t.Error("failed to get user")
	}

	if user != nUser {
		t.Error("failed to retrive proper redis user")
	}

	if !reflect.DeepEqual(user, nUser) {
		t.Error("Shit happens")
	}
}

func TestRedisScan(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":5000",
	})

	ctx := context.Background()

	var cursor uint64
	var err error
	var keys []string

	// scan command uses a cursor to iterate over the keyspace
	// The cursor is not a simple numeric index but rather a bitmask
	// that Redis uses internally to traverse its hash table.
	// The cursor value (e.g., 8) is an implementation detail
	// and doesn't directly correspond to the number of keys scanned or returned.
	// Redis uses a hash table to store keys, and the cursor represents the current position in this hash table.
	// The value 8 (or any other non-zero value) indicates that Redis has more keys to scan in the current iteration.
	// When the cursor returns 0, it means the iteration is complete.
	//
	// The COUNT parameter (in your case, 2) is a hint to Redis about how many keys to return per iteration.
	//
	// Redis may return fewer or more keys than the COUNT value, depending on the internal structure of the hash table.

	for {
		keys, cursor, err = rdb.Scan(ctx, cursor, "user:*", 10).Result()
		if cursor == 0 {
			break
		}

		if err != nil {
			break
		}
		_ = keys
	}
}

func TestRedisZRange(t *testing.T) {

	rdb := redis.NewClient(&redis.Options{
		Addr: ":5000",
	})

	ctx := context.Background()
	// ZRANGE: Retrieves elements based on their rank (position in the sorted set).
	// ZRANGEBYSCORE: Retrieves elements based on their score (within a specified score range).

	items := []redis.Z{
		{Score: 2, Member: "item2"},
		{Score: 3, Member: "item3"},
	}
	rdb.ZAdd(ctx, "items", items...)

	// Pagination parameters
	offset := 1 // Start from the second item (0-based index)
	limit := 2  // Number of items per page

	keys, err := rdb.ZRange(ctx, "items", int64(offset), int64(offset+limit-1)).Result()
	if err != nil {
		t.Error("could not get zragne")
	}

	_ = keys

	newItems := []redis.Z{
		{Score: 1, Member: "item1"},
		{Score: 4, Member: "item4"},
		{Score: 5, Member: "item5"},
	}

	offset = 3
	limit = 7

	rdb.ZAdd(ctx, "items", newItems...)

	keys, err = rdb.ZRange(ctx, "items", int64(offset), int64(offset+limit-1)).Result()

	if err != nil {
		t.Error("could not get zragne")
	}

}

func TestRedisList(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":5000",
	})

	ctx := context.Background()

	items := []RedisUser{
		{Username: "username", Email: "email@email.com", Role: 0, IsActive: true},
		{Username: "another", Email: "xd@email.com", Role: 0, IsActive: true},
		//"hello",
		//"world",
	}
	rdb.RPush(ctx, "users", items)

	offset := 1 // Start from the second item (0-based index)
	limit := 2  // Number of items per page

	keys, err := rdb.LRange(ctx, "users", int64(offset), int64(offset+limit-1)).Result()
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Paginated items:", keys)
}
