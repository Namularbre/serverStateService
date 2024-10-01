# Server state service

This a console app to test the connection to the server with a simple http request

## Install
Just clone the repository and add a .env file like this:
````
# Set the environment variable for the MariaDB connection string
CONN_STR="namu:2261@tcp(REPLACE_WITH_DB_HOST:3306)/server_monitoring"
# Set the environment variable for the http port
PORT=8080
# Set the environment variable for the http address
ADDRESS=0.0.0.0
````
Then you can run this command to install all dependencies:
````
go mod tidy
````
## Run
This app run like a basic go app, with this command
````
go run serverStateService.go
````
If you prefer to build it first:
````
go build -o serverStateService
````
And then execute it:
````
./serverStateService
````
## Docker
To run this app in docker, use this command to build the image
````
docker build . -t server_state_service
````
And run it as usual with other containers:
````
docker run --rm -d server_state_service
````
## Authors
[Namulabre](https://github.com/Namularbre)
