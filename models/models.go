package route

// RegisterSchema struct to represent a new route
type RegisterSchema struct {
	Route           string            `json:"route"`
	Methods         []string          `json:"methods"`
	Headers         map[string]string `json:"headers"`
	Parameters      map[string]string `json:"parameters"`
	Response        any               `json:"response"`
	ResponseCode    int               `json:"response_code,default=200"`
	ResponseHeaders map[string]string `json:"response_headers"`
}

// DeleteSchema struct to represent a new route
type DeleteSchema struct {
	Route   string   `json:"route"`
	Methods []string `json:"methods"`
}

// MethodValue stores the matching rules and payload for an specific method
type MethodValue struct {
	// Matching variables
	Headers    map[string]string
	Parameters map[string]string
	// Response value
	ResponseCode    int
	Response        any
	ResponseHeaders map[string]string
}

// Method contains a method and a list of MethodValue for multiple matching for the same route
type Method map[string][]MethodValue

// Route contains the route to be matched and a list of MethodValue for multiple matching for the same route
type Route map[string]Method
