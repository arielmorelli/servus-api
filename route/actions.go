package route

import (
	"fmt"
	"log"
	"strings"
)

// Routes is a global variable to store all routes
var Routes Route = make(Route)

// AsRouteName get a string and transform into a internal route name representation
func AsRouteName(raw string) string {
	return strings.Trim(raw, "/")
}

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

func mapIsSubset(baseMap, mapToCheck map[string]string) bool {
	// Check headers
	for key, value := range mapToCheck {
		key = strings.ToLower(key)
		headerValue, exists := baseMap[key]
		if !exists {
			return false
		}
		if headerValue != value {
			return false
		}
	}
	return true

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
		// Check headers
		if mapIsSubset(headers, methodValue.Headers) && mapIsSubset(params, methodValue.Parameters) {
			return methodValue, true
		}
	}

	return MethodValue{}, false
}
