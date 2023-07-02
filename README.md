# Servus API

Servus! Servus API is a lightweight **dummy** REST API with run-time routing management. It was created with UI development and integration tests in mind.

The idea behind the project is to have an API that can add/remove routes in run-time with matching system for different methods, parameters and headers.

There is no user/access management and de-duplication of entries.

## Install

### Binary

1. Download the binary
    - For AMD64 / x86_64

    `curl -Lo ./servus-api https://github.com/arielmorelli/servus-api/releases/download/latest/servus-api-amd64`

    - For ARM64

    `curl -Lo ./servus-api https://github.com/arielmorelli/servus-api/releases/download/latest/servus-api-arm64`

2. Move the valid PATH folder
```
chmod +x ./servus-api
sudo mv ./servus-api /usr/local/bin/servus-api
```

### Docker

Check DockerHub: https://hub.docker.com/r/arielmorelli/servus-api

## Usage

`servus-api [-p port] [-f JSONfile] [-d debugmode]`

* -d, --debug         Debug mode
* -f, --file string   Input file
* -p, --port string   Port to run. (default "8080")

## Load from file

Use the flag `--file` (`-f`) with a JSON using using a list of entries using the [register schema](#Register).

### Example

`file.JSON`
```JSON
[
    {
        "route": "/rouet",
        "methods": ["get"],
        "response": {"hello": "world"},
        "response_code": 200
    }
]
```

`servus-api -f file-JSON`

## API

### Info
Call `/_info/` with a GET to see stored routes and methods.

### Register

Call `/_register/` with a POST request to register a new route, or overwrite an existing one.

* route (string): route to be matched
* methods (list[string]): methods that this route accept
* response code (integer): response status code
* headers (JSON): map with key and values to be matched with headers
* parameters (JSON): map with key and values to be matched with parameters
* response (any): response data
* response_headers (JSON): map with key and values to used as headers

#### Example
```JSON
{
    "route": "/rouet",
    "methods": ["get", "post"],
    "response": {"hello": "world"},
    "response_code": 200
}
```

### Removing existing route

Call `/_remove/`  with a PUT request to remove route or a specific method for a route.

* route (string): route to be removed. Route must match exactly as registered
* methods (list[string]): specific methods to be deleted. If a empty list or not provided, all methods will be removed

#### Example
```JSON
{
    "route": "/rouet",
    "methods": ["get"]
}
```
