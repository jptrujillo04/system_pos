package location

type Country struct {
	ID     int    `gorm:"column:country_id"`
	Name   string `gorm:"column:country_name"`
	Status string `gorm:"column:country_state"`
}

type CountryRequest struct {
	Name   string `json:"country_name"`
	Status string `json:"country_state"`
}
