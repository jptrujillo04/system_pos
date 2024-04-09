package location

import "context"

type UseCase interface {
	GetAllContries(ctx context.Context) ([]Country, error)
}

type UseCaseLocation struct {
	RepositoryLocation RepositoryLocation
}

func NewUseCaseLocation(repositoryLocation RepositoryLocation) *UseCaseLocation {
	return &UseCaseLocation{
		RepositoryLocation: repositoryLocation,
	}
}

func (u *UseCaseLocation) GetAllContries(ctx context.Context) ([]Country, error) {
	var countries []Country
	err := u.RepositoryLocation.GetAllContries(&countries)
	if err != nil {
		return nil, err
	}
	return countries, nil
}
