package spatial_test

import (
	"testing"

	"github.com/golang/geo/r2"
	"github.com/nvnamsss/goinf/spatial"
)

type EmergencyMaintenance struct {
	Id  string  `gorm:"column:id"`
	Lat float64 `gorm:"column:lat"`
	Lon float64 `gorm:"column:lon"`
}

const EmergencyMaintenanceTableName = "emergency_maintenance"

var eme_maintenance spatial.Quadtree
var map_eme map[string]*spatial.Quadtree

func (EmergencyMaintenance) TableName() string {
	return EmergencyMaintenanceTableName
}

func (e EmergencyMaintenance) GetId() string {
	return e.Id
}

func (e EmergencyMaintenance) Location() r2.Point {
	var location r2.Point
	location.X = e.Lat
	location.Y = e.Lon
	return location
}

func (e EmergencyMaintenance) CreateBound() (b spatial.Bounds) {
	b.X = e.Lat
	b.Y = e.Lon
	b.Height = 0.01
	b.Width = 0.01
	b.Item = e

	return
}

func TestRemoveItem(t *testing.T) {
	var s spatial.Quadtree
	b := spatial.Bounds{X: 10, Y: 10, Height: 1, Width: 1}
	b2 := spatial.Bounds{X: 10, Y: 10, Height: 1, Width: 1}
	s.Insert(b)
	s.Insert(b2)

	s.RemoveItem(b.Item)

	if len(s.Objects) != 1 {
		t.Error("Expected 1")
	} else {
		t.Log("Completed")
	}
}
