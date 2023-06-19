# Servus API


## Motivation

Servus! Servus API is a dummy rest API with run-time routing managment. It was made to make UI development and integration tests easier.


## Usage

### Register

Call `/_register/` with a POST request to register a new route, or overwrite an existing one.
* Route: route to be matched
* Methods: list of methods that this route accept
* Response code: integer
* Parameters: json
* Match_type: bool 
* Response: Any

### Removing existing route
Call `/_remove/{id or router}` with a PUT request to deregister an existing route.


### Load from JSON

TDB
