package application

type Facades struct {
	UserApplication   *userApplication
	DidApplication    *didApplication
	RegionApplication *regionApplication
}

func NewFacades() (facades *Facades) {
	return &Facades{
		UserApplication:   newUserApplication(),
		DidApplication:    newDidApplication(),
		RegionApplication: newRegionApplication(),
	}
}
