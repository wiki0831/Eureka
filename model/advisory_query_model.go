package model

import (
	"fmt"

	"github.com/twpayne/go-geom/encoding/wkt"
)

type AdvisoryQueryModel struct {
	Geometry string `json:"geometry,omitempty"`
	Layers   []string `json:"layers,omitempty"`
	Note     string `json:"note,omitempty"`
}

func ValidateAdvisoryQueryModel(input *AdvisoryQueryModel) error {
	_, err := wkt.Unmarshal(input.Geometry)
	if err != nil {
		return fmt.Errorf("invalid geometry input")
	}
	return nil
}
