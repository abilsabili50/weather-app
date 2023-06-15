package dto

import (
	"net/http"
	"time"
)

type IResponse interface {
	Error() string
	GetCode() int
	GetData() WeatherResponse
}

type WeatherRequest struct {
	Water int `json:"water" binding:"required"`
	Wind  int `json:"wind" binding:"required"`
}

type WeatherResponse struct {
	ID          uint       `json:"id"`
	Water       int        `json:"water"`
	Wind        int        `json:"wind"`
	WaterStatus string     `json:"water_status"`
	WindStatus  string     `json:"wind_status"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Time        string     `json:"time,omitempty"`
	Degree      int        `json:"degree,omitempty"`
}

type Response struct {
	Err        bool            `json:"error"`
	Message    string          `json:"message"`
	Status     string          `json:"status"`
	StatusCode int             `json:"code"`
	Data       WeatherResponse `json:"data,omitempty"`
}

func (r *Response) Error() string {
	return r.Status
}

func (r *Response) GetCode() int {
	return r.StatusCode
}

func (r *Response) GetData() WeatherResponse {
	return r.Data
}

func NewCreatedResponse(err bool, message string, payload WeatherResponse) IResponse {
	return &Response{
		Err:        err,
		Message:    message,
		Status:     "CREATED",
		StatusCode: http.StatusCreated,
		Data:       payload,
	}
}

func NewOKResponse(err bool, message string, payload WeatherResponse) IResponse {
	return &Response{
		Err:        err,
		Message:    message,
		Status:     "OK",
		StatusCode: http.StatusOK,
		Data:       payload,
	}
}
