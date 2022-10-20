
## App
`go run app.go`

## Docker

Build: `docker build . -t go-rat`
Run docker image: `docker run -d -p 3001:3000 go-rat`


# Load testing
Creates 1000 concurrent sessions (virtual users), and runs this single request 10 times for each virtual user, hence 10,000 iterations.
`cat test.js | docker run --rm -i grafana/k6 run - --vus 1000 --iterations 10000`

# DB

Start Postgres
`docker run --name postgresql -e POSTGRES_USER=dreo -e POSTGRES_PASSWORD=mypassword -p 5432:5432 -v /data:/var/lib/postgresql/data -d postgres`

Start PGAdmin
`docker run --name my-pgadmin -p 82:80 -e 'PGADMIN_DEFAULT_EMAIL=user@local.dev' -e 'PGADMIN_DEFAULT_PASSWORD=postgresmaster'-d dpage/pgadmin4`

Get postgres local IP
`docker inspect postgresql -f “{{json .NetworkSettings.Networks }}”` 