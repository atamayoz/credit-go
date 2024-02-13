# credit-go

Golang Pet project credit simulator created with Gin-gonic, Ent and PosgreSQL.

## TODO
- Put some tests

## How to run the project

- Go to the foler `tests/docker` the folder has the `docker-compose.yml` with PostgreSql and PgAdmin.
- I'm assuming you have installed docker and docker-compose.
- run the command `docker compose up -d` this command will pull the images and initiate the containers.
- Run this Atlas command to create the tables in the postgres DB
```shell
atlas migrate apply --dir file://ent/migrate/migrations \
  --url "postgres://user:admin@localhost:54320/credit_golang?search_path=public&sslmode=disable"
```
- After the containers are initiated go to the root folder of the project where the `main.go` is located and run the command `go run main.go`.
This command will start the server in the port 8081

## Atlas commands

The project uses Atlas for generating migrations

This command generates the first migration the Simulation table
```
atlas migrate diff create_simulation_table \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "docker://postgres/15/test?search_path=public"
```
```
atlas migrate apply --dir file://ent/migrate/migrations \
  --url "postgres://user:admin@localhost:54320/credit_golang?search_path=public&sslmode=disable"
```
