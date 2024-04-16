package location

type Country struct {
	ID     int    `gorm:"column:country_id"`
	Name   string `gorm:"column:country_name"`
	Status string `gorm:"column:country_state"`
}

type CountryRequest struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"state"`
}

type UpdateCountryRequest struct {
	ID int `json:"country_id"`
}
