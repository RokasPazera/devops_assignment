package models

type Percentage struct {
    ID        uint `gorm:"primaryKey" json:"id"`
    Value     int  `json:"percentage"`
}
