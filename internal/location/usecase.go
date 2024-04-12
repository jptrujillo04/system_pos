package location

import "context"

type UseCase interface {
	GetAllCountries(ctx context.Context) ([]Country, error)
	SaveCountry(ctx context.Context, country CountryRequest) error
}

type UseCaseLocation struct {
	RepositoryLocation RepositoryLocation
}

func NewUseCaseLocation(repositoryLocation RepositoryLocation) *UseCaseLocation {
	return &UseCaseLocation{
		RepositoryLocation: repositoryLocation,
	}
}

func (u *UseCaseLocation) GetAllCountries(ctx context.Context) ([]Country, error) {
	var countries []Country
	err := u.RepositoryLocation.GetAllCountries(&countries)
	if err != nil {
		return nil, err
	}
	return countries, nil
}

func (u *UseCaseLocation) SaveCountry(ctx context.Context, countryReq CountryRequest) error {
	country := MapCountryRequestToModelCountry(countryReq)
	return u.RepositoryLocation.SaveCountry(&country)
}
