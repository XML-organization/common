package create_user

type Role int

const (
	Host Role = iota
	Guest
	NK
)

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
	Name     string `json:"nameuuid`
	Surname  string `json:"surname"`
	Country  string `json:"country" gorm:"not null;type:string"`
	City     string `json:"city" gorm:"not null;type:string"`
	Street   string `json:"street" gorm:"not null;type:string"`
	Number   string `json:"number" gorm:"not null;type:string"`
}

type CreateUserCommandType int8

const (
	SaveUser CreateUserCommandType = iota
	DeleteUserCredentials
	PrintSuccessful
	UnknownCommand
)

type CreateUserCommand struct {
	User User
	Type CreateUserCommandType
}

type CreateUserReplyType int8

const (
	UserSaved CreateUserReplyType = iota
	UserNotSaved
	UserCredentialsDeleted
	SuccessfulyFinished
	UnknownReply
)

type CreateUserReply struct {
	User User
	Type CreateUserReplyType
}
