package auth

import (
	"github.com/vektah/gqlparser/gqlerror"
)

var AUTH_REQUIRED_GQL_ERROR = &gqlerror.Error{
	Message: "Authentication required to perform this action",
	Extensions: map[string]interface{}{
		"code": 401,
	},
}

var PERMISSION_DENIED_GQL_ERROR = &gqlerror.Error{
	Message: "User is not authorized to perform this action",
	Extensions: map[string]interface{}{
		"code": 403,
	},
}
