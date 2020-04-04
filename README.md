# Covid-19 ETL
Extract data from the NY Times data set of covid-19 cases ([github link](https://github.com/nytimes/covid-19-data)). Transform the data into a loadable data format. Load the data into a Postgres database.

## Development

- Install dependencies: `go mod vendor`

### Notes
- Attempt to set up the package as recommended [here](https://github.com/golang-standards/project-layout)
- Initialize Dependencies: `go mod init github.com/hoitomt/covid-19-etl`
