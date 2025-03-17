package utils

var AuthenticationCodes = map[string]string{
	"success_login":      "S-000",
	"success_logout":     "S-001",
	"error_invalid_pwd":  "E-001",
	"error_expired_tkn":  "E-002",
	"error_invalid_user": "E-003",
	"error_blocked_acc":  "E-004",
	"warning_attempt":    "W-001",
}
