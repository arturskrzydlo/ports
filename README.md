# ports microservices

This repo contains two microservices which process json file with ports and store them in in-memory database

### Web app service

it's simple http server which have two endpoints :

```
POST /ports
GET /ports
```

where `POST` take as param multipart form with json, i.e :

```json

{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  }
}
```

This service just handle rest requests and pass it to `ports` service

### Ports service

This service runs grpc server and process store and fetch requrest from `webapp`
For the sake of simplicity it stores ports in in memory database which is simple map

## Requirements

- [Go](https://golang.org/doc/install) >= Go 1.20
- [GNU Make](https://www.gnu.org/software/make/)
- [Docker](https://docs.docker.com/engine/install)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting started

To run this service just call

```shell
docker compose up -d 
```

It should start two services and be able to process request on localhost:8080

You can verify it by running simple `curl` command :

```shell
curl --location 'localhost:8080/ports' \
--form 'ports=@"/Users/arturskrzydlo/Desktop/ports.json"'
```

Remember to include your location to sample ports.json which you want to test

## Testing

There are few levels of tests. There is e2e test in `webapp` which is tagged as `integration`
This is simple, happy path test which checks if everything works correctly together. To run this test call :

```shell
make all-tests
```

It will run all unit test and this integration test. This will spin up `ports` service and should close it after test

Other test are simple domain test (parametrized) and service test (only ports tests are written). There could be more
tests
like those (i.e webapp is poorly tested) but it gives just example what kind of test I would expect

## Development

* **setup** - only to download linter binary useful for linting during development
* **clean** - delete binary downloaded in setup
* **clean-integration-tests** - cleaning after integration tests (shutting down ports-service)
* **generate-proto** - generating grpc code
* **lint** - running set of linters for static code analysis
* **prepare-integration-test** running `docker-compose-test.yml` with ports service to be able to run integration tests.
* **all-tests** runs all tests unit and integration ones
* **tests** running all tests without integration tests. Doesn't need preparation step
* **tidy** runs go tidy
