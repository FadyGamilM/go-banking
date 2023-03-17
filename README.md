# go-banking
Backend service for bank application.

# 1- Database Design

![tables](./account_db_design.png)

#### `relations between tables`
![tables](./Db_relations.png)

#### `To create an instance of postgresql database running on docker container:`
```cmd
➜ docker run -d --name go_banking_container -p 5432:5432 -e POSTGRES_USER=go_banking_username -e POSTGRES_PASSWORD=go_banking_password -e POSTGRES_DB=go_banking_db 6a35e2c987a6
```








