package change_password

type ChangePasswordDTO struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

type ChangePasswordCommandType int8

const (
	ChangePassword ChangePasswordCommandType = iota
	RollbackPassword
	PrintSuccessful
	UnknownCommand
)

type ChangePasswordCommand struct {
	ChagePasswordDTO ChangePasswordDTO
	Type             ChangePasswordCommandType
}

type ChangePasswordReplyType int8

const (
	PasswordChanged ChangePasswordReplyType = iota
	PasswordNotChanged
	OldPasswordReturned
	SuccessfulyFinished
	UnknownReply
)

type ChangePasswordReply struct {
	ChangePasswordDTO ChangePasswordDTO
	Type              ChangePasswordReplyType
}
