package model

type Day struct {
	Datetime      string  `json:"datetime"`
	DatetimeEpoch int64   `json:"datetimeEpoch"`
	TempMax       float64 `json:"tempmax"`
	TempMin       float64 `json:"tempmin"`
	Temp          float64 `json:"temp"`
	FeelsLikeMax  float64 `json:"feelslikemax"`
	FeelsLikeMin  float64 `json:"feelslikemin"`
	FeelsLike     float64 `json:"feelslike"`
	Dew           float64 `json:"dew"`
	// Thêm các trường khác nếu cần thiết
}
