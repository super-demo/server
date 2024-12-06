# Server

This repository is api of Super Office (Senior Project)
Tech stacks are `Golang` / `Postgresql`

Refer to Golang Template Documentation:

## Prerequisite

1. Install Go - https://go.dev/doc/install
2. Install **Goose**

```
brew install goose
```

3. Install Docker engine - https://docs.docker.com/engine/install/ or Docker desktop - https://docs.docker.com/desktop/install/mac-install/ or OrbStack - https://orbstack.dev/download

## How to set up your project locally

1. create .env

```
cp ./.env.example ./.env
```

2. run (this will run api and postgres container)

```
make docker-dev-up
```

## Migration

- Run the migration

```
  make migrate-up
```

- Reset the migrations

```
  make migrate-down
```

- Check migrations status

```
  make migrate-status
```

> Ensure that your API and DB containers are running before executing the migration script.

## Mocks

- Install mockery
  ```sh
  brew install mockery
  ```
- Create a mock file
  ```sh
  mockery --name [interfaceName] --filename [fileName].go
  ```

## Running Tests

### Basic Test

```sh
go test ./...
```

### With Coverage

1. Run tests and generate coverage:

```sh
go test ./... -coverprofile=coverage.txt
```
