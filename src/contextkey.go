package src

type contextKey string

func (c contextKey) String() string {
	return "context key " + string(c)
}

var UserIDContextKey = contextKey("user_id")
