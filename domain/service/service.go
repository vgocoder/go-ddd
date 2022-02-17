package service

type Facades struct {
	UserService userService
}

func NewFacades() *Facades {
	var userService userService
	return &Facades{
		UserService: userService,
	}
}
