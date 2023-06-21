package route

import (
	"fmt"
	"log"
	"strings"
)

// Routes is a global variable to store all routes
var Routes Route = make(Route)

// RegisterRoute register a new route in the global context
func RegisterRoute(route, method string, responseCode int, headers, params map[string]string, response any) {
	route = AsRouteName(route)
	method = strings.ToLower(method)

	methodValue := MethodValue{
		Headers:      headers,
		Parameters:   params,
		ResponseCode: responseCode,
		Response:     response,
	}

	_, exists := Routes[route]
	if !exists {
		Routes[route] = make(Method)
	}

	_, exists = Routes[route][method]
	if !exists {
		Routes[route][method] = make([]MethodValue, 0)
	}

	if len(headers) == 0 && len(params) == 0 {
		Routes[route][method] = append(Routes[route][method], methodValue)
	} else {
		Routes[route][method] = append([]MethodValue{methodValue}, Routes[route][method]...)
	}
	log.Println("Added", route, strings.ToUpper(method))
}

// FindRoute find route in method in registered Routes
func FindRoute(route, method string, headers, params map[string]string) (MethodValue, bool) {
	route = AsRouteName(route)
	method = strings.ToLower(method)

	_, exists := Routes[route]
	if !exists {
		return MethodValue{}, false
	}

	_, exists = Routes[route][method]
	if !exists {
		return MethodValue{}, false
	}

	for _, methodValue := range Routes[route][method] {
		fmt.Println(methodValue)
		if isMapSubset(headers, methodValue.Headers) && isMapSubset(params, methodValue.Parameters) {
			return methodValue, true
		}
	}

	return MethodValue{}, false
}

// DeleteRoute delete a given route and a method, if provided
func DeleteRoute(route, method string) bool {
	route = AsRouteName(route)
	method = strings.ToLower(method)

	mapOfMethods, exists := Routes[route]
	if !exists {
		return false
	}

	if method == "" {
		delete(Routes, route)
	} else {
		_, exists = mapOfMethods[method]
		if !exists {
			return false
		}
		delete(mapOfMethods, method)
	}

	log.Println("Deleted", route, strings.ToUpper(method))
	return true
}
