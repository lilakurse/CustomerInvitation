package model

type Customer struct {
	Latitude  *float64 `json:"latitude,string,omitempty"`
	UserID    int      `json:"user_id"`
	Name      string   `json:"name"`
	Longitude *float64 `json:"longitude,string,omitempty"`
}

const (
	R               float64 = 6371		// Radius of the earth
	DistanceLimit   float64 = 100		// Distance limit for deciding whether to invite the customer or not
	OfficeLatitude  float64 = 53.339428	// Latitude of the office in decimal degrees
	OfficeLongitude float64 = -6.257664 // Longitude of the office in decimal degrees
)
