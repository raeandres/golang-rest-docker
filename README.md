# golang-rest-product
Simple webservice that performs CRUD on a product model built using golang and postgresql

## Requirements:
### Model struct:
    - [ ] Name
    - [ ] Type
    - [ ] Picture
    - [ ] Price
    - [ ] Description

### Functionalities:
    - [ ] Get All products
    - [ ] Add product
    - [ ] Edit product



## create a go module dependency file
```go mod init github.com/golang-rest-product```

## Dependencies
```
go get github.com/gofiber/fiber/v2    - Fiber
go get github.com/lib/pq              - go pq
```


## Commands:
### boots up the initial docker container you are building
```docker compose up``` 		


### install dependencies inside container 
```docker compose run —service-ports web bash```

### check go version 
``` go version ```


## Containers
### Postgresql
```
db:
    image: postgres:alpine
    environment:
      - POSTGRES_USER=${DB_USER}
      - POSTGRES_PASSWORD=${DB_PASSWORD}
      - POSTGRES_DB=${DB_NAME}
    ports:
      - 5432:5432
    volumes:
      - postgres-db:/var/lib/postgresql/data
```      

### Web
```
web:
    build: .
    env_file:
      - .env
    ports:
      - 3000:3000
    volumes:
      - .:/usr/src/app
    command: air main/main.go -b 0.0.0.0
```