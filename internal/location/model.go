package location

type Country struct {
	ID     int    `gorm:"column:country_id"`
	Name   string `gorm:"column:country_name"`
	Status string `gorm:"column:country_state"`
}
