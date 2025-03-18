package utils

var AuthenticationCodes = map[string]string{
	"success":               "S-000",
	"invalid_credentials":   "E-001",
	"invalid_data":          "E-002",
	"invalid_token":         "E-003",
	"email_already_exists":  "E-004",
	"user_not_found":        "E-005",
	"internal_server_error": "E-999",
}

func IsSuccess(code string) bool {
	return code == AuthenticationCodes["success"]
}
