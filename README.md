# Installation

```
go get -u github.com/cclin81922/fakerestful/cmd/fakerestful
export PATH=$PATH:~/go/bin
```

# Commmand Line Usage

```
fakerestful

// then you can use curl to test 
```

Access user resource by curl
* List all users
* Get a user
* Create a user
* Delete a user
* Update a user
* Get a HTML form for user creation
* Get a HTML form for user modification

```
# List all users
curl -X GET http://localhost:8080/users

# Get a user
curl -X GET http://localhost:8080/users/1

# Create a user
curl -X POST http://localhost:8080/users/ -F 'name=cclin'

# Delete a user
curl -X DELETE http://localhost:8080/users/1/delete

# Update a user
curl -X PUT http://localhost:8080/users/1/update -F 'name=CCLIN'

# Get a HTML form for user creation
curl -X GET http://localhost:8080/users/new

# Get a HTML form for user modification
curl -X GET http://localhost:8080/users/1/edit
```
