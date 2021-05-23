


## Setup

Start up docker compose

`docker-compose up -d`

Keycloak is running on [http://localhost:8080](http://localhost:8080)

A database preloaded is provided in **keycloak.mv.db**.

Realm MONITORING
client alerts-ui

user: admin
password: password123

## Code Examples (Get token with client secret)



Enter main container with the command,

`docker-compose exec demoapp /bin/bash`

### C++

Enter cpp code directory

`cd cpp`

Configure and build main.cpp

`cmake -B build/ . && cmake --build build/`

Run app

`./build/main`

### Python

Enter python code directory

`cd python`

Run app

`python3 main.py`

### Golang

Enter golang code directory

`cd go`

Run app

`go run main.go`

# Keycloak preset up

Keycloaks default H2 database file is stored by default in 

`/opt/jboss/keycloak/standalone/data/keycloak.mv.db`

