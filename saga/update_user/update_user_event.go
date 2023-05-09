package update_user

type UpdateUserDTO struct {
	OldEmail string `json:"old_email"`
	NewEmail string `json:"new_email"`
}

type UpdateUserCommandType int8

const (
	UpdateUser UpdateUserCommandType = iota
	RollbackEmail
	PrintSuccessful
	UnknownCommand
)

type UpdateUserCommand struct {
	UpdateUserDTO UpdateUserDTO
	Type          UpdateUserCommandType
}

type UpdateUserReplyType int8

const (
	UserUpdated UpdateUserReplyType = iota
	UserNotUpdated
	OldEmailReturned
	SuccessfulyFinished
	UnknownReply
)

type UpdateUserReply struct {
	UpdateUserDTO UpdateUserDTO
	Type          UpdateUserReplyType
}
