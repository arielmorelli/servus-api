# Servus API


## Motivation

Servus! Servus API is a dummy REST API with run-time routing managment. It was made with UI development and integration tests in mind.

The idea behind the project is to have an API that can add/remvoe routes and also have an matching system for different methods, parameters and headers.


## Usage

TBD

## API

### Info
Call `/_info/` with a GET to see stored routes and methods.


### Register

Call `/_register/` with a POST request to register a new route, or overwrite an existing one.

* route (string): route to be matched
* methods (list[string]): methods that this route accept
* response code (integer): response status code
* headers (json): map with key and values to be matched with headers
* parameters (json): map with key and values to be matched with parameters
* response (any): response data
* response_headers (json): map with key and values to used as headers

### Removing existing route

Call `/_remove/`  with a PUT request to remove route or a specific method for a route.

* route (string): route to be removed. Route must match exactly as registered
* methods (list[string]): specific methods to be deleted. If a empty list or not provided, all methods will be removed


### Load from JSON



