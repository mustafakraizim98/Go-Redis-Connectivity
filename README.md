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
