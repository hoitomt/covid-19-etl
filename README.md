# Covid-19 ETL
Extract data from the NY Times data set of covid-19 cases ([github link](https://github.com/nytimes/covid-19-data)). Transform the data into a loadable data format. Load the data into a Postgres database.

## Development

- Copy the environment file and set the missing variables. `cp .env.example .env`
- Install dependencies: `go mod vendor`
- Set up the database. This app uses a Postgres database. The app assumes that we are connecting to the database defined by the docker-compose file. The database runs on port 5433 (as opposed to the standard Postgres port of 5432) so as not to interfere with a local instance of Postgres that may be running.
  - `docker-compose up` Start the database
  - `brew install golang-migrate` Install migration tool [link](https://github.com/golang-migrate/migrate)
  - `make migrate` Run migrations.
    - **NOTE** POSTGRESQL_URL is defined in .env

### Helpful commands

- PSQL `docker-compose exec postgres psql -U covid_user -d covid_development`
- Create a migration `migrate create -ext sql -dir db/migrations -seq create_counties_table`
- Rollback a migration `migrate -database ${POSTGRESQL_URL} -path db/migrations down`
- Force version `migrate -database ${POSTGRESQL_URL} -path db/migrations force VERSION`


### Notes
- Attempt to set up the package as recommended [here](https://github.com/golang-standards/project-layout)
- Initialize Dependencies: `go mod init github.com/hoitomt/covid-19-etl`
