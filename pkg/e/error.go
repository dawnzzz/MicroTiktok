package e

const (
	Success = int32(iota)
)

const (
	ErrUnknown = int32(iota) + 1000
	ErrBadRequest
	ErrUserNameOrPasswordWrong
	ErrAuthenticationFailed
	ErrGenerateTokenFailed
	ErrDBError
	ErrUserRpcFailed
	ErrAuthenticationRpcFailed
	ErrDuplicatedUserName

	ErrUserIDGeneratorInitFailed
	ErrVideoIDGeneratorInitFailed
	ErrCommentIDGeneratorInitFailed
)
