package route

// MethodValue stores the matching rules and payload for an specific method
type MethodValue struct {
	// Matching variables
	Headers    map[string]string
	Parameters map[string]string
	// Return value
	ResponseCode int
	Response     any
}

// Method contains a method and a list of MethodValue for multiple matching for the same route
type Method map[string][]MethodValue

// Route contains the route to be matched and a list of MethodValue for multiple matching for the same route
type Route map[string]Method
