package repository

type Facades struct {
	UserRepository   userRepository
	DidRepository    didRepository
	RegionRepository regionRepository
}

func NewFacades() (facades *Facades) {
	var u userRepository
	var d didRepository
	var r regionRepository
	return &Facades{
		UserRepository:   u,
		DidRepository:    d,
		RegionRepository: r,
	}
}
