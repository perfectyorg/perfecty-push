package application

import "github.com/rwngallego/perfecty-push/internal/domain/user"

type (
	PreferenceService struct {
		userRepository UserRepository
	}
)

//NewPreferenceService Creates a Preference service
func NewPreferenceService(userRepository UserRepository) (service *PreferenceService) {
	service = &PreferenceService{
		userRepository: userRepository,
	}
	return
}

func (p *PreferenceService) Update(u *user.User) (err error) {
	return p.userRepository.Update(u)
}
