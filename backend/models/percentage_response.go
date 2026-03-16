package models

type PercentageResponse struct {
    Percentage int `json:"percentage"`
}

func NewPercentageResponse(value int) PercentageResponse {
    return PercentageResponse{Percentage: value}
}
