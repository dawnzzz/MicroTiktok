package e

const (
	Success = int32(iota)
	ErrUnknown
	ErrBadRequest
	ErrAuthenticationFailed
	ErrGenerateTokenFailed
	ErrDBError
	ErrUserRpcFailed
	ErrAuthenticationRpcFailed
	ErrDuplicatedUserName
)
