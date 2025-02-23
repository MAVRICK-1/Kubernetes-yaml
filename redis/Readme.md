# **Redis with Go: Comprehensive Guide with Functions & Examples**  

## **1. Introduction**  
Redis is a high-performance in-memory key-value store, and Go provides an efficient way to interact with Redis using the **"github.com/redis/go-redis/v9"** package.

In this guide, we will cover:  
1. Installing Redis and Go client  
2. Connecting Go with Redis  
3. Performing CRUD operations  
4. Working with different Redis data types (Strings, Lists, Sets, Hashes, Sorted Sets, etc.)  
5. Using Redis pipelines and transactions  
6. Redis Pub/Sub  
7. Redis Streams  

---

## **2. Install Redis and Go Client**  

### **Install Redis Server** (Linux/macOS)
```sh
sudo apt install redis-server  # Ubuntu
brew install redis             # macOS
redis-server                   # Start Redis
```

### **Install Go Redis Package**
```sh
go get github.com/redis/go-redis/v9
```

---

## **3. Connecting to Redis in Go**
```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background() // Required for Redis operations

func main() {
	// Create Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password
		DB:       0,                // Default database
	})

	// Test Connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis successfully!")
}
```

---

## **4. Redis Data Types & Commands in Go**
### **1. Strings**
```go
// Set a key
err := client.Set(ctx, "name", "Rishi", 0).Err()
if err != nil {
	log.Fatalf("Error setting key: %v", err)
}

// Get the key
val, err := client.Get(ctx, "name").Result()
if err != nil {
	log.Fatalf("Error getting key: %v", err)
}
fmt.Println("Value:", val)

// Increment and Decrement
client.Set(ctx, "counter", 10, 0)
client.Incr(ctx, "counter") // Increments counter by 1
client.Decr(ctx, "counter") // Decrements counter by 1
```

---

### **2. Lists**
```go
// Push elements to list
client.RPush(ctx, "fruits", "apple", "banana", "cherry")

// Retrieve all list items
fruits, _ := client.LRange(ctx, "fruits", 0, -1).Result()
fmt.Println("Fruits List:", fruits)

// Pop an element
firstFruit, _ := client.LPop(ctx, "fruits").Result()
fmt.Println("Removed First Fruit:", firstFruit)
```

---

### **3. Sets**
```go
// Add elements to a set
client.SAdd(ctx, "colors", "red", "blue", "green")

// Get all elements from the set
colors, _ := client.SMembers(ctx, "colors").Result()
fmt.Println("Colors Set:", colors)

// Check if a value exists
exists, _ := client.SIsMember(ctx, "colors", "blue").Result()
fmt.Println("Is 'blue' in set:", exists)

// Remove a value from the set
client.SRem(ctx, "colors", "red")
```

---

### **4. Hashes**
```go
// Set multiple fields in a hash
client.HSet(ctx, "user:1", "name", "Rishi", "age", "25")

// Get a specific field
name, _ := client.HGet(ctx, "user:1", "name").Result()
fmt.Println("User Name:", name)

// Get all fields and values
userData, _ := client.HGetAll(ctx, "user:1").Result()
fmt.Println("User Data:", userData)
```

---

### **5. Sorted Sets (ZSet)**
```go
// Add elements to sorted set
client.ZAdd(ctx, "leaderboard", 
	&redis.Z{Score: 100, Member: "Alice"},
	&redis.Z{Score: 200, Member: "Bob"},
)

// Get elements with scores
leaderboard, _ := client.ZRangeWithScores(ctx, "leaderboard", 0, -1).Result()
fmt.Println("Leaderboard:")
for _, entry := range leaderboard {
	fmt.Printf("User: %s, Score: %.0f\n", entry.Member, entry.Score)
}

// Increment score
client.ZIncrBy(ctx, "leaderboard", 50, "Alice")
```

---

### **6. HyperLogLog**
```go
// Add elements to HyperLogLog
client.PFAdd(ctx, "unique_visitors", "user1", "user2")

// Get approximate count
count, _ := client.PFCount(ctx, "unique_visitors").Result()
fmt.Println("Unique Visitors:", count)
```

---

### **7. GeoSpatial Indexing**
```go
// Add locations
client.GeoAdd(ctx, "places",
	&redis.GeoLocation{Longitude: 13.361389, Latitude: 38.115556, Name: "Palermo"},
	&redis.GeoLocation{Longitude: 15.087269, Latitude: 37.502669, Name: "Catania"},
)

// Get distance between two places
dist, _ := client.GeoDist(ctx, "places", "Palermo", "Catania", "km").Result()
fmt.Println("Distance:", dist, "km")
```

---

## **5. Redis Transactions**
Transactions allow executing multiple commands atomically.
```go
tx := client.TxPipeline()
tx.Set(ctx, "counter", 0, 0)
tx.Incr(ctx, "counter")
tx.Incr(ctx, "counter")
_, err := tx.Exec(ctx)
if err != nil {
	log.Fatalf("Transaction failed: %v", err)
}
```

---

## **6. Redis Pub/Sub (Message Queue)**
### **Publisher**
```go
client.Publish(ctx, "channel1", "Hello from Go!")
```

### **Subscriber**
```go
sub := client.Subscribe(ctx, "channel1")
ch := sub.Channel()

for msg := range ch {
	fmt.Println("Received:", msg.Payload)
}
```

---

## **7. Redis Streams**
### **Adding Data to Stream**
```go
client.XAdd(ctx, &redis.XAddArgs{
	Stream: "mystream",
	Values: map[string]interface{}{
		"user":    "Alice",
		"message": "Hello!",
	},
})
```

### **Reading Data from Stream**
```go
msgs, _ := client.XRead(ctx, &redis.XReadArgs{
	Streams: []string{"mystream", "0"},
	Count:   2,
	Block:   0,
}).Result()

for _, msg := range msgs {
	fmt.Println("Stream Data:", msg.Messages)
}
```

---

## **8. Closing the Redis Connection**
```go
defer client.Close()
```

---

## **9. Conclusion**
- Redis provides fast in-memory storage for various use cases like caching, session management, leaderboards, and message queues.  
- The **go-redis/v9** library makes it easy to work with Redis in Go.  
- With support for Strings, Lists, Sets, Hashes, Sorted Sets, HyperLogLogs, and more, Redis is a versatile database for high-performance applications.

🚀 **Start integrating Redis into your Go applications today!** 🚀
