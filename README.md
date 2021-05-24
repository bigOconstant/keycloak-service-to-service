



# Keycloak Service to Service examples and example code

## Setup

Start up docker compose

`docker-compose up -d`

Keycloak is running on [http://localhost:8080](http://localhost:8080)

A database preloaded is provided in **keycloak.mv.db**.

**Realm** MONITORING
**client** alerts-ui

**user**: admin
**password**: password123

## Code Examples (Get token with client secret)

Enter main container with the command,

`docker-compose exec demoapp /bin/bash`

### Curl

```bash
curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d 'grant_type=client_credentials&client_id=alerts-ui&client_secret=fbe69472-563d-4604-9336-1ac39cf1efa3' \
"http://keycloak:8080/auth/realms/MONITORING/protocol/openid-connect/token" 

```

If you have jq installed you can make it pretty and pipe it to jq. It comes installed in the container
bash
```
curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d 'grant_type=client_credentials&client_id=alerts-ui&client_secret=fbe69472-563d-4604-9336-1ac39cf1efa3' \
"http://keycloak:8080/auth/realms/MONITORING/protocol/openid-connect/token" | jq
```

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

## Test if a token would authenticate

After you get a token, you can test whether or not it is a valid token. 
A small http server is running on port 8083 and listening to the endpoint `v1/alerts`

### Curl command to validate token. 

```bash

curl -X POST localhost:8093/v1/alerts \
-H 'Authorization: Bearer Foo'
```

Where **Foo** is the token from above

If the token fails to authenticate the post will return the reason for failure. For example,

```bash 
Auth Failed
oidc: token is expired (Token Expiry: 2021-05-24 01:28:36 +0000 UTC)
```

Success will look like 

```
Verrification Successfull
```



