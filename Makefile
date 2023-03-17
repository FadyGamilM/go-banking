migrateUp:
	docker run -v "E:/1- freelancing path/Courses/golang stack/go-banking/migrations:/migrations" --network host migrate/migrate -path ./migrations/ -database="postgresql://[YOUR_USERNAME]:[YOUR_PASSWORD]@localhost:[YOUR_PORT]/[YOUR_DB_NAME]?sslmode=disable" -verbose up	

.PHONY: migrateUp migrateDown