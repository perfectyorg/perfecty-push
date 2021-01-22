package application

import "github.com/rwngallego/perfecty-push/internal/domain/user"

type (
	PreferenceService struct {
		repository UserRepository
	}
)

//NewPreferenceService Creates a Preference service
func NewPreferenceService(userRepository UserRepository) (service *PreferenceService) {
	service = &PreferenceService{
		repository: userRepository,
	}
	return
}

func (p *PreferenceService) Update(u *user.User) (err error) {
	return p.repository.Update(u)
}
