# OAuth Service

## Project description

*    Create a Golang http server that issues JWT Access Tokens (rfc7519) using Client Credentials Grant with Basic Authentication (rfc6749) in endpoint /token
*    Sign the tokens with a self-made RS256 key
*    Provide an endpoint to list the signing keys (rfc7517)
*    Provide deployment manifests to deploy the server in Kubernetes cluster
*    Create an Introspection endpoint (rfc7662) to introspect the issued JWT Access Tokens

## Setup

### Local test

## Steps

* Generate the [Public and Private keys](#generate-public-and-private-keys)
* Download all dependencies and run `go mod tidy`
* Run the command `make test` or `go clean -testcache && go test -v -cover ./...`
* Run the command `make run` or `go run cmd/main.go`

### Docker and Kubernetes

#### Prerequisites

Before proceeding, ensure that you have the following prerequisites:

- You need to have installed [docker](https://docs.docker.com/engine/install/) and [kubectl](https://kubernetes.io/docs/tasks/tools/)
- Access to a Kubernetes cluster
- Docker image of the server built and pushed to a container registry accessible from the cluster

## Steps

* Generate the [Public and Private keys](#generate-public-and-private-keys)
* Run the command `make docker-build` or `docker build -t oauth-service .`
* Push the image to a container registry, for this example the `deployment.yaml` is using [this one](docker.io/jsovalles/oauth-service:latest)
* Run the command `make kube-deploy` or `kubectl apply -f deployment.yaml`
* Run the command `make kube-service` or `kubectl apply -f service.yaml`
* Verify the deployment with `kubectl get pods`
* Verify the services with `kubectl get services`
* Port-forward the service with `kubectl port-forward {pod_name} 8080:8080`


#### Cleanup
To clean up the server deployment and associated resources, run the following command:

```bash
kubectl delete -f deployment.yaml
```

#### Generate Public and Private Keys

```bash
ssh-keygen -t rsa -b 4096 -m PEM -f cert/jwtRS256.key
# Don't add passphrase
openssl rsa -in cert/jwtRS256.key -pubout -outform PEM -out cert/jwtRS256.key.pub
```

# Golang Clean architecture

The solution architecture we can divide the code in 5 main layers:

- Models: Is a set of data structures.
- Services: Contains application specific business rules. It encapsulates and implements all the use cases of the system.
- Controllers: Is a set of adapters that convert data from the format most convenient for the use cases and models.
- Utils and Config: Is generally composed of frameworks and tools.
- Token: Provides an implementation for generating and verifying JWTs using RSA-based signing methods

In Clean Architecture, each layer of the application (use case, data service and domain model) only depends on interface of other layers instead of concrete types. 
Dependency Injection is one of the SOLID principles, a rule about the constraint between modules that abstraction should not depend on details. 
Clean Architecture uses this rule to keep the dependency direction from outside to inside.

## Routes

### List Signing Keys

- Method: GET
- URL: `/api/signing-keys`
- Description: Retrieves the list of signing keys.
- Response:
    - Status Code: 200 OK
    - Body: JSON representation of the signing keys.

#### Example

```bash
curl --location 'localhost:8080/api/signing-keys'
```

### Create Token

- Method: POST
- URL: `/api/token`
- Description: Creates a new token.
- Request:
    - Headers:
        - Authorization: Basic authentication header with valid credentials.
          - Credentials:
            - Username: `endava`
            - Password: `secretpass`
            - OR
            - Username: `user`
            - Password: `root`
    - Body: None
- Response:
    - Status Code: 200 OK
    - Body: JSON representation of the created token.

#### Example

```bash
curl --location --request POST 'localhost:8080/api/token' \
--header 'Authorization: Basic dXNlcjpyb290'
```

### Verify Token

- Method: GET
- URL: `/api/introspect`
- Description: Verifies the validity of a token.
- Request:
    - Headers:
        - Authorization: Bearer authentication header with a valid token. 
          - This token can be generated from the [create token endpoint](#create-token)
    - Body: None
- Response:
    - Status Code: 200 OK

#### Example

```bash
curl --location 'localhost:8080/api/introspect' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRJZCI6InVzZXIiLCJpc3N1ZWRBdCI6IjIwMjMtMDYtMTRUMTA6MjU6NTcuNDg1NzQ2Mi0wNTowMCIsImV4cGlyZWRBdCI6IjIwMjMtMDYtMTVUMTA6MjU6NTcuNDg1NzQ2Ni0wNTowMCJ9.1PiL_J6Y5DbsFOtJZYwvLzdCTU0WQuPzplHbPbugDrHGqAUYcbCLr40w-Bu3JLbqUXtWSB7uaJCKvHnZHreyoVCF8XcszKyAYRdU-tlHoPpKBiE_bXk-zD9DKIpJGFtqLXA_qIbwj0G8jwjdObfoz1dBcqVwLdF4Dqzi-JovJ5FCtaNGYv8mikgZmrcE7Nz8h-_l3bJCkMP2X-uGg1abT8wH3gfEkkimuhycPug9lHWuRmKvBOQH8agEpa6Vp2wrL-t3jeUY3Fx_6F0FnIeZ17zChn-6ZCTpwlyW6JGWL4PJWZNoXqtJzJwVGesllQXx9VqaAFgbnb6TOjINgk0Kx7qD57flj-k0XIgnl3S0SUwaIf6l8AdkLftWQjlWrfEO7e_4eCy5qW0f01DUF5UxMyr8n6VDt-3v4kIDtnL4rhVhn5CBbl70k4nHPFng2jxnz-sZtJv5i7xoEQry4mZspo2n7yB7mWprUxhVXxZjdBpvq--0whQPxyBiQGHRty9BV9QXFooMn8f8wd06cekuM8er6IRLd9DJO-cYSKtJLAt9yHPAWKJgpZL0DKFzdadER8UbpXa_QOYnhnZJ16pqKP_Ta_B9ddelTolC0hPUQFSixPN0jlDzHJGMulVy98kfZ0rK0BwyBTKiueukf-E42fqG7MOCEI09WThcmHu_YBQ'
```

## Authentication

Some routes require authentication. The following authentication methods are used:

- List Signing Keys: No authentication required.
- Create Token: Basic Authentication with valid credentials.
- Verify Token: Bearer Authentication with a valid token.

## Built With

- [Go](https://go.dev/) - version 1.20.2
- [Testify](https://github.com/stretchr/testify)
- [Viper](https://github.com/spf13/viper)
- [Makefile](https://www.gnu.org/software/make/manual/make.html#Introduction)
- [UberFx](https://github.com/uber-go/fx)
- [Gin](https://github.com/gin-gonic-gin)
- [JWT](https://github.com/golang-jwt/jwt/v4)