# the container for the database
DB_DOCKER_CONTAINER=gobanking_db_container


# to create a running container instance for postgres using the name of the container we specified above
postgres:
	docker run -d --name ${DB_DOCKER_CONTAINER} -p 1472:5432 -e POSTGRES_USER=gobanking -e POSTGRES_PASSWORD=gobanking postgres:14

# now lets create an actual datbase inside this container 
createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=gobanking --owner=gobanking gobankingdb

# for migration commands
migrate_up:
	docker run -i -v "H:\1- freelancing path\Courses\golang stack\projects\go-banking\internal\adapters\repository\postgres\migrations:/migrations" --network host migrate/migrate -path=/migrations/ -database "postgresql://gobanking:gobanking@127.0.0.1:1472/gobankingdb?sslmode=disable" up 1

migrate_down:
	docker run -i -v "H:\1- freelancing path\Courses\golang stack\projects\go-banking\internal\adapters\repository\postgres\migrations:/migrations" --network host migrate/migrate -path=/migrations/ -database "postgresql://gobanking:gobanking@127.0.0.1:1472/gobankingdb?sslmode=disable" down 1

# to start the docker container for the databsae 
start_db_container:
	docker start ${DB_DOCKER_CONTAINER}