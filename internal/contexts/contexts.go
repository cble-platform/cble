package contexts

type contextKey struct {
	field string
}

var USER_CTX_KEY = &contextKey{"user"}
var GIN_CONTEXT_KEY = &contextKey{"gincontext"}
