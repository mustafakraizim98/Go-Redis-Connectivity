# Go-Redis-Connectivity

![Go-Docker-Redis](https://user-images.githubusercontent.com/113289516/193471975-65924bbe-a3f6-45ef-ab9c-b8c81afbfaeb.png)

Redis is an in-memory data store used as a database, cache, or message broker. Go-redis/redis is a type-safe, Redis client library for Go with support for features like Pub/Sub, sentinel, and pipelining.

## Prerequisites
- Go
- Docker
- Redis installed (In this example i used Docker image)
- RedisInsight (Optional)

## Running Redis With Docker Locally
> docker pull redis

> docker run -d -i -p 6379:6379 --name redis-instance redis

The first pull command retrieves the Redis image from DockerHub so that we can then run it as a container using the second command. In the second command, we specify the name of our Redis container and we also map our local port 6379 to the port that Redis is running against within the container using the -p flag.

## Test Cases
### Connecting to our Redis Instance
> RedisAddr = "localhost:6379"

> database, err := db.NewClient(RedisAddr)

### Setting Values to Redis
> db.Set(database, "creator", "Mustafa", 0)

### Getting Values from Redis
> db.Get(database, "creator")
#### Output
> "Mustafa"

![Screenshot 2022-10-02 223524](https://user-images.githubusercontent.com/113289516/193472908-039225c1-2d56-43ce-bc4d-85ce896c1f63.png)

### Deleting Values from Redis
> // db.Del(database, "creator")

### Storing Composite Values
```
type Message struct {
	Sender      string    `json:"sender,omitempty" binding:"required"`
	Receiver    string    `json:"receiver,omitempty" binding:"required"`
	MessageBody string    `json:"message,omitempty" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func JsonMarshal(msg Message) []byte {
	json, err := json.Marshal(msg)
	db.Must(err)
	return json
}
```

```
var message = jsonhandling.Message{
    Sender:      "Bob",
		Receiver:    "Alice",
		MessageBody: "Welcome to Redis in-memory caching.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
}
```
> db.Set(database, "message", jsonhandling.JsonMarshal(message), 0)

## Proven Results
![Screenshot 2022-10-02 224133](https://user-images.githubusercontent.com/113289516/193473116-13382647-5006-44dd-9c7a-c40dc0d1132d.png)

