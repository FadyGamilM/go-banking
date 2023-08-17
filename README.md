# go-banking
Backend service for bank application.

# Supported Features 
- Create and manage a bank account
    <br>➜ Each bank has an owner, balance, and currency.

- Keep records of all the balance changes of the bank account
    <br>➜ Each bank operation done to the account, a new account entry will be created and recoreded.

- Money transfer transacitons between bank accounts
    <br>➜ Keep the money transfer consistent between 2 accounts with a transaction.

# Database Design

![tables](./account_db_design.png)

#### `relations between tables`
![tables](./Db_relations.png)

#### `To create an instance of postgresql database running on docker container:`
```cmd
➜ docker run -d --name go_banking_container -p 5432:5432 -e POSTGRES_USER=go_banking_username -e POSTGRES_PASSWORD=go_banking_password -e POSTGRES_DB=go_banking_db 6a35e2c987a6
```









├── Dockerfile
├── cmd
│   └── runner.go
├── conf
│   └── app.yaml
├── internal
│   ├── controller
│   │   └── controller.go
│   ├── core
│   │   ├── common
│   │   │   ├── router
│   │   │   │   └── router.go
│   │   │   └── utils
│   │   │       └── logger.go
│   │   ├── dto
│   │   │   └── user.go
│   │   ├── entity
│   │   ├── model
│   │   │   ├── request
│   │   │   │   └── request.go
│   │   │   └── response
│   │   │       └── response.go
│   │   ├── port
│   │   │   ├── repository
│   │   │   │   ├── db.go
│   │   │   │   └── user.go
│   │   │   └── service
│   │   │       └── user.go
│   │   ├── server
│   │   │   └── http_server.go
│   │   └── service
│   │       └── user.go
│   └── infra
│       ├── config
│       │   └── config.go
│       └── repository
│           ├── db.go
│           └── user.go
├── schema
│   └── schema.sql
└── script
    └── run.sh