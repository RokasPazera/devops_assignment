package models

type Goal struct {
    Name             string `gorm:"primaryKey" json:"name"`
    TargetPercentage int    `gorm:"column:target_percentage" json:"targetPercentage"`
}
