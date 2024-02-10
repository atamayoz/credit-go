# credit-go
Golang Pet project credit simulator


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
