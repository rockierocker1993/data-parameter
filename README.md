# DATA PARAMETER
This repository is for handling parameterized data, configuration, system values, templates, and the data in these parameters can be used for all service needs.

# SPESIFICATION
- Golang Version go1.24.0 - min (1.20)
- Gin
- GORM
- REST API
- Database PostgreSQL

 ## Project Structure Description
 ```
    ├── config          # Package containing for all configartion
    ├── controllers     # Package containing for controller
    ├── dto             # Package containing for dto(data transfer object), containing request & response object
    ├── migration       # Package containing for sql script migration
    ├── models          # Package containing for orm database
    ├── repositories    # Package containing for query to database
    ├── routes          # Package containing for api list
    ├── service         # Package containing for bisnis logic
    
 ```

# DATABASE MIGRATION
cmd : 
- installation
    - go get -u github.com/golang-migrate/migrate/v4
    - go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
- create migration file
    - migrate create -ext sql -dir migration -seq create_data_parameter
- run migration
    - run migrate
        - migrate -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" -path database/ddl up
    - rollback
        - migrate -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" -path database/ddl down


# HOW TO RUN
 - duplicate .env to .env-local
 - fill the required environment value
 - go run main.go