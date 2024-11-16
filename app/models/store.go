package models

type Store struct {
	StoreID   string `gorm:"primaryKey" json:"store_id"`
	StoreName string `json:"store_name"`
	AreaCode  string `json:"area_code"`
}
