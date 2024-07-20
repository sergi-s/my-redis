# My Redis Implementation in Go

I haven't learned a new programming language in a while, and what better way to do so than with a small project? I built a simple Redis implementation from scratch in Go! 🚀

## Project Overview

This project is a basic implementation of Redis in Go. It includes key features such as command parsing, data structure handling for in-memory storage, client connections, and Append-Only Files (AOF) for data persistence.

## Features

- **Command Parsing and Execution:** Efficiently parse and execute commands sent by clients.
- **Data Structures:** Implementation of in-memory storage using appropriate data structures.
- **Append-Only Files (AOF):** Data persistence using AOF.

## Learning Resources

This project was inspired and guided by the article [Build Redis from Scratch](https://www.build-redis-from-scratch.dev/en/aof). The article provides an excellent overview and step-by-step instructions for building a Redis clone from scratch.

## Topics Covered

- Append-Only Files (AOF)
- Command parsing and execution
- Data structures for in-memory storage
- Handling client connections

## Installation

To get started, clone the repository:

```bash
git clone https://github.com/sergi-s/my-redis.git
```

Navigate to the project directory:

```bash
cd my-redis
```

Run the project:

```bash
go run .
```

## Usage

Once the server is running, you can interact with it using a Redis client or a simple telnet connection:

```bash
telnet localhost 6379
```
or 
```bash
redis-cli
```
Note: If you are going to use redis-cli, make sure to stop any running Redis service to avoid conflicts.

### Example Commands

Here are some examples of how to use the implemented Redis commands:

```redis-cli
127.0.0.1:6379> ping
PONG

127.0.0.1:6379> hset users u1 sergi
OK

127.0.0.1:6379> hgetall users
1) "u1"
2) "sergi"

127.0.0.1:6379> hset users u2 sergi2
OK

127.0.0.1:6379> hgetall users
1) "u1"
2) "sergi"
3) "u2"
4) "sergi2"

127.0.0.1:6379> HGET users u1
"sergi"

127.0.0.1:6379> HGET users u2
"sergi2"
```
