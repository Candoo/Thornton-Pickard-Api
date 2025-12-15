package models

import (
    "time"
    "fmt" 
    "encoding/json"
    "gorm.io/gorm"
)

type Camera struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	Name                string    `gorm:"not null" json:"name"`
	Manufacturer        string    `gorm:"not null" json:"manufacturer"`
	YearIntroduced      int       `json:"year_introduced"`
	YearDiscontinued    *int      `json:"year_discontinued,omitempty"`
	Format              string    `json:"format"`
	PlateSizes          string    `json:"plate_sizes"` // Store as JSON string or comma-separated
	Lens                string    `json:"lens"`
	Shutter             string    `json:"shutter"`
	Features            string    `json:"features"` // Store as JSON string
	Description         string    `gorm:"type:text" json:"description"`
	ImageURLs           string    `json:"image_urls"` // Store as JSON string
	Rarity              string    `json:"rarity"`
	EstimatedValueMin   *float64  `json:"estimated_value_min,omitempty"`
	EstimatedValueMax   *float64  `json:"estimated_value_max,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"-"`
}

type CameraResponse struct {
	ID                  uint      `json:"id"`
	Name                string    `json:"name"`
	Manufacturer        string    `json:"manufacturer"`
	YearIntroduced      int       `json:"year_introduced"`
	YearDiscontinued    *int      `json:"year_discontinued,omitempty"`
	Format              string    `json:"format"`
	PlateSizes          []string  `json:"plate_sizes"`
	Lens                string    `json:"lens"`
	Shutter             string    `json:"shutter"`
	Features            []string  `json:"features"`
	Description         string    `json:"description"`
	ImageURLs           []string  `json:"image_urls"`
	Rarity              string    `json:"rarity"`
	EstimatedValueRange string    `json:"estimated_value_range,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

// ToCameraResponse converts a Camera model to the client-friendly CameraResponse model.
func (c *Camera) ToCameraResponse() CameraResponse {
    resp := CameraResponse{
        ID:             c.ID,
        Name:           c.Name,
        Manufacturer:   c.Manufacturer,
        YearIntroduced: c.YearIntroduced,
        YearDiscontinued: c.YearDiscontinued,
        Format:         c.Format,
        Lens:           c.Lens,
        Shutter:        c.Shutter,
        Description:    c.Description,
        Rarity:         c.Rarity,
        CreatedAt:      c.CreatedAt,
        UpdatedAt:      c.UpdatedAt,
    }

    json.Unmarshal([]byte(c.PlateSizes), &resp.PlateSizes)
    json.Unmarshal([]byte(c.Features), &resp.Features)
    json.Unmarshal([]byte(c.ImageURLs), &resp.ImageURLs)

    if c.EstimatedValueMin != nil && c.EstimatedValueMax != nil {
        resp.EstimatedValueRange = fmt.Sprintf("$%.0f - $%.0f", *c.EstimatedValueMin, *c.EstimatedValueMax)
    }
    
    return resp
}