
![Logo](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png)


# Simple Rest API Using Golang



Table of Contents
-----------------

  * [Requirements](#requirements)
  * [Tech Stack](#tech-stack)
  * [Environment Variables](#environment-variables)
  * [Run Project](#run-project)
  * [Run Unit Test](#run-unit-test)
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file.

```
DB_HOST=your_db_host                   
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name  
DB_PORT=your_db_port
BASE_URL=localhost:8008
TOKEN=3cdcnTiBsl
```


## Run Project

Clone the project.

```bash
  git clone https://github.com/juanvaleriand/go-test.git
```

Go to the project directory.

```bash
  cd go-test
```

Install GO packages.

```bash
  go get github.com/juanvaleriand/go-test
```

Start the server.

```bash
  go run main.go
```


## Run Unit Test

To run unit test, run the following command

```bash
  go test -v ./unit-tests
```

Please read official website to read more about unit testing.

https://go.dev/doc/tutorial/add-a-test


## Endpoints

### Get All Users
Method: `GET` \
Endpoint: `{{base_url}}/api/v1/users?page=1&limit=10`\
Response:
```json
{
    "page": 1,
    "limit": 10,
    "total": 12,
    "data": [
        {
            "id": 12,
            "email": "rachel.howell@reqres.in",
            "first_name": "Rachel",
            "last_name": "Howell",
            "avatar": "https://reqres.in/img/faces/12-image.jpg",
            "created_at": "2023-04-16T14:36:16.364754+07:00",
            "updated_at": "2023-04-16T14:36:16.364754+07:00",
            "deleted_at": null
        },
        {
            "id": 11,
            "email": "george.edwards@reqres.in",
            "first_name": "George",
            "last_name": "Edwards",
            "avatar": "https://reqres.in/img/faces/11-image.jpg",
            "created_at": "2023-04-16T14:36:16.363725+07:00",
            "updated_at": "2023-04-16T14:36:16.363725+07:00",
            "deleted_at": null
        },
        {
            "id": 10,
            "email": "byron.fields@reqres.in",
            "first_name": "Byron",
            "last_name": "Fields",
            "avatar": "https://reqres.in/img/faces/10-image.jpg",
            "created_at": "2023-04-16T14:36:16.36265+07:00",
            "updated_at": "2023-04-16T14:36:16.36265+07:00",
            "deleted_at": null
        },
        {
            "id": 9,
            "email": "tobias.funke@reqres.in",
            "first_name": "Tobias",
            "last_name": "Funke",
            "avatar": "https://reqres.in/img/faces/9-image.jpg",
            "created_at": "2023-04-16T14:36:16.36139+07:00",
            "updated_at": "2023-04-16T14:36:16.36139+07:00",
            "deleted_at": null
        },
        {
            "id": 8,
            "email": "lindsay.ferguson@reqres.in",
            "first_name": "Lindsay",
            "last_name": "Ferguson",
            "avatar": "https://reqres.in/img/faces/8-image.jpg",
            "created_at": "2023-04-16T14:36:16.360263+07:00",
            "updated_at": "2023-04-16T14:36:16.360263+07:00",
            "deleted_at": null
        },
        {
            "id": 7,
            "email": "michael.lawson@reqres.in",
            "first_name": "Michael",
            "last_name": "Lawson",
            "avatar": "https://reqres.in/img/faces/7-image.jpg",
            "created_at": "2023-04-16T14:36:16.359172+07:00",
            "updated_at": "2023-04-16T14:36:16.359172+07:00",
            "deleted_at": null
        },
        {
            "id": 6,
            "email": "tracey.ramos@reqres.in",
            "first_name": "Tracey",
            "last_name": "Ramos",
            "avatar": "https://reqres.in/img/faces/6-image.jpg",
            "created_at": "2023-04-16T14:36:16.357622+07:00",
            "updated_at": "2023-04-16T14:36:16.357622+07:00",
            "deleted_at": null
        },
        {
            "id": 5,
            "email": "charles.morris@reqres.in",
            "first_name": "Charles",
            "last_name": "Morris",
            "avatar": "https://reqres.in/img/faces/5-image.jpg",
            "created_at": "2023-04-16T14:36:16.356563+07:00",
            "updated_at": "2023-04-16T14:36:16.356563+07:00",
            "deleted_at": null
        },
        {
            "id": 4,
            "email": "eve.holt@reqres.in",
            "first_name": "Eve",
            "last_name": "Holt",
            "avatar": "https://reqres.in/img/faces/4-image.jpg",
            "created_at": "2023-04-16T14:36:16.355424+07:00",
            "updated_at": "2023-04-16T14:36:16.355424+07:00",
            "deleted_at": null
        },
        {
            "id": 3,
            "email": "emma.wong@reqres.in",
            "first_name": "Emma",
            "last_name": "Wong",
            "avatar": "https://reqres.in/img/faces/3-image.jpg",
            "created_at": "2023-04-16T14:36:16.354222+07:00",
            "updated_at": "2023-04-16T14:36:16.354222+07:00",
            "deleted_at": null
        }
    ]
}
```
### Get User By Id
Method: `GET` \
Endpoint: `{{base_url}}/api/v1/user/:id`\
Response:
```json
{
    "data": {
        "id": 1,
        "email": "george.bluth@reqres.in",
        "first_name": "George",
        "last_name": "Bluth Update From Unit Testing",
        "avatar": "https://reqres.in/img/faces/1-image.jpg",
        "created_at": "2023-04-16T14:20:28.711782+07:00",
        "updated_at": "2023-04-16T14:24:52.637686+07:00",
        "deleted_at": null
    }
}
```

### Create User
Method: `POST` \
Endpoint: `{{base_url}}/api/v1/user`\
Payload:
```json
{
    "email": "juanvaleriand@gmail.com",
    "first_name": "Juan Valerian",
    "last_name": "Delima",
    "avatar": "https://www.anakui.com/wp-content/uploads/2023/01/karakter-anime-ter-cool-1.webp"
}
```
Response:
```json
{
    "message": "Data created successfully!"
}
```

### Update User
Method: `PUT` \
Endpoint: `{{base_url}}/api/v1/user/:id`\
Payload:
```json
{
    "email": "juanvaleriand@gmail.com",
    "first_name": "Juan",
    "last_name": "Delima",
    "avatar": "https://www.anakui.com/wp-content/uploads/2023/01/karakter-anime-ter-cool-1.webp"
}
```
Response:
```json
{
    "message": "Data updated successfully!"
}
```

### Delete User
Method: `DELETE` \
Endpoint: `{{base_url}}/api/v1/user/:id`\
Response:
```json
{
    "message": "Data deleted successfully!"
}
```
