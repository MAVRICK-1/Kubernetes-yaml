# **Redis Commands and Data Types (Detailed Guide with Examples)**  

## **1. Introduction to Redis**  
Redis (Remote Dictionary Server) is an open-source, in-memory data structure store primarily used as a database, cache, and message broker. It supports various data structures and provides high-performance operations.  

## **2. Redis Data Types**  

Redis supports multiple data types, including:  
1. **Strings**  
2. **Lists**  
3. **Sets**  
4. **Sorted Sets (ZSets)**  
5. **Hashes**  
6. **Bitmaps**  
7. **HyperLogLogs**  
8. **Geospatial Indexes (Geo)**  
9. **Streams**  

---

## **3. Redis Commands by Data Type with Examples**  

### **1. String Data Type**  
A string in Redis is a binary-safe sequence of bytes. It can store text, numbers, or even serialized objects.  

#### **Basic Commands**
| Command | Description |
|---------|------------|
| `SET key value` | Stores a value under a key |
| `GET key` | Retrieves the value of a key |
| `INCR key` | Increments a numeric value |
| `DECR key` | Decrements a numeric value |
| `APPEND key value` | Appends data to an existing string |
| `STRLEN key` | Returns the length of a string |

#### **Example**
```sh
SET name "Rishi"
GET name          # Output: "Rishi"

SET counter 10
INCR counter      # Output: 11
DECR counter      # Output: 10
APPEND name " Mondal"
GET name          # Output: "Rishi Mondal"
```

---

### **2. List Data Type**  
A list in Redis is an ordered collection of strings, similar to a linked list.

#### **Basic Commands**
| Command | Description |
|---------|------------|
| `LPUSH key value` | Pushes a value to the left (beginning) of the list |
| `RPUSH key value` | Pushes a value to the right (end) of the list |
| `LPOP key` | Removes and returns the first element |
| `RPOP key` | Removes and returns the last element |
| `LRANGE key start stop` | Retrieves elements within a range |

#### **Example**
```sh
LPUSH fruits "apple"
LPUSH fruits "banana"
RPUSH fruits "cherry"
LRANGE fruits 0 -1  # Output: ["banana", "apple", "cherry"]

LPOP fruits         # Output: "banana"
RPOP fruits         # Output: "cherry"
```

---

### **3. Set Data Type**  
A set in Redis is an unordered collection of unique strings.

#### **Basic Commands**
| Command | Description |
|---------|------------|
| `SADD key value` | Adds a value to the set |
| `SREM key value` | Removes a value from the set |
| `SMEMBERS key` | Returns all members of the set |
| `SISMEMBER key value` | Checks if a value exists in the set |
| `SCARD key` | Returns the number of elements in the set |

#### **Example**
```sh
SADD colors "red"
SADD colors "blue"
SADD colors "green"
SMEMBERS colors  # Output: ["red", "blue", "green"]

SISMEMBER colors "blue"  # Output: 1 (true)
SCARD colors             # Output: 3
```

---

### **4. Sorted Set (ZSet) Data Type**  
A sorted set is similar to a set but with scores assigned to each element, enabling sorting.

#### **Basic Commands**
| Command | Description |
|---------|------------|
| `ZADD key score value` | Adds a value with a score |
| `ZRANGE key start stop` | Retrieves elements sorted by score |
| `ZREM key value` | Removes a value from the sorted set |
| `ZINCRBY key increment value` | Increments a score |

#### **Example**
```sh
ZADD leaderboard 100 "Alice"
ZADD leaderboard 200 "Bob"
ZRANGE leaderboard 0 -1 WITHSCORES
# Output: ["Alice", 100, "Bob", 200]

ZINCRBY leaderboard 50 "Alice"
ZRANGE leaderboard 0 -1 WITHSCORES
# Output: ["Alice", 150, "Bob", 200]
```

---

### **5. Hash Data Type**  
A hash is a key-value store inside a Redis key, similar to a dictionary.

#### **Basic Commands**
| Command | Description |
|---------|------------|
| `HSET key field value` | Sets a field in a hash |
| `HGET key field` | Gets the value of a field |
| `HGETALL key` | Gets all fields and values |
| `HDEL key field` | Deletes a field |

#### **Example**
```sh
HSET user:1 name "Rishi"
HSET user:1 age 25
HGET user:1 name      # Output: "Rishi"
HGETALL user:1        # Output: {"name": "Rishi", "age": "25"}
```

---

### **6. Bitmap Data Type**  
Bitmaps are used for bitwise operations.

#### **Basic Commands**
| Command | Description |
|---------|------------|
| `SETBIT key offset value` | Sets a bit at a specific position |
| `GETBIT key offset` | Gets the bit at a specific position |

#### **Example**
```sh
SETBIT user_activity 5 1
GETBIT user_activity 5  # Output: 1
```

---

### **7. HyperLogLog Data Type**  
Used for approximating unique counts.

#### **Basic Commands**
| Command | Description |
|---------|------------|
| `PFADD key value` | Adds an element to HyperLogLog |
| `PFCOUNT key` | Returns an approximate count |

#### **Example**
```sh
PFADD unique_visitors "user1"
PFADD unique_visitors "user2"
PFCOUNT unique_visitors  # Output: 2
```

---

### **8. Geospatial Data Type**  
Used for storing and querying locations.

#### **Basic Commands**
| Command | Description |
|---------|------------|
| `GEOADD key longitude latitude member` | Adds a geospatial point |
| `GEODIST key member1 member2` | Gets distance between two points |
| `GEORADIUS key longitude latitude radius unit` | Finds members in a given radius |

#### **Example**
```sh
GEOADD places 13.361389 38.115556 "Palermo"
GEOADD places 15.087269 37.502669 "Catania"
GEODIST places "Palermo" "Catania" km  # Output: ~166.27 km
```

---

### **9. Stream Data Type**  
Used for message queues.

#### **Basic Commands**
| Command | Description |
|---------|------------|
| `XADD key * field value` | Adds an entry to a stream |
| `XRANGE key start end` | Retrieves entries in a range |

#### **Example**
```sh
XADD mystream * user "Alice" message "Hello!"
XRANGE mystream - +
# Output: [{"user": "Alice", "message": "Hello!"}]
```

---

## **4. Redis Transactions and Scripting**  

### **Transactions**
Commands:  
- `MULTI` - Start a transaction  
- `EXEC` - Execute the transaction  
- `DISCARD` - Cancel the transaction  

### **Lua Scripting**
Execute scripts using `EVAL "script" num_keys key1 key2 ... arg1 arg2 ...`

Example:
```sh
EVAL "return redis.call('GET', KEYS[1])" 1 mykey
```

---

## **5. Conclusion**  
Redis offers powerful data structures and commands that make it an excellent choice for caching, real-time analytics, message queues, and more. With its high-speed performance and scalability, Redis is widely used in production environments. 🚀
