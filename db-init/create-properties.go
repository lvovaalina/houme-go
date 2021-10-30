package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Project struct {
	gorm.Model
	ProjectId                int    `gorm:"primary_key;autoIncrement"`
	Name                     string `gorm:"unique"`
	BucketName               string
	Filename                 string `gorm:"unique"`
	LivingArea               string
	RoomsNumber              int
	ConstructionDuraton      int
	ConstructionCost         int
	ConstructonWorkersNumber string
	FoundationMaterial       string
	WallMaterial             string
	FinishMaterial           string
	RoofingMaterial          string
}

type Property struct {
	gorm.Model
	PropertyId   int    `gorm:"primary_key;autoIncrement"`
	PropertyName string `gorm:"unique"`
	PropertyUnit string
	PropertyCode string
}

type Job struct {
	gorm.Model
	JobId              int    `gorm:"primary_key;autoIncrement"`
	JobName            string `gorm:"unique"`
	StageName          string
	PropertyCode       string
	WallMaterial       string
	FinishMaterial     string
	FoundationMaterial string
	RoofingMaterial    string
	InteriorMaterial   string
	Property           Property `gorm:"foreignKey:PropertyCode"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(
		"postgres",
		"host=localhost port=5432 user=postgres dbname=houmly sslmode=disable password=l8397040")

	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}

	db.DropTable(&Property{})
	db.DropTable(&Job{})

	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Property{})
	db.AutoMigrate(&Job{})

	var properties = []Property{
		{PropertyName: "Foundation volume", PropertyUnit: "sq.m.", PropertyCode: "FV"},
		{PropertyName: "Floor area at the base", PropertyUnit: "sq.m.", PropertyCode: "FA"},
		{PropertyName: "Total floor area", PropertyUnit: "sq.m.", PropertyCode: "TFA"},
		{PropertyName: "Walls volume", PropertyUnit: "sq.m.", PropertyCode: "WV"},
		{PropertyName: "Number of walls", PropertyUnit: "ps", PropertyCode: "WN"},
		{PropertyName: "Roof area", PropertyUnit: "sq.m.", PropertyCode: "RA"},
		{PropertyName: "Number of windows", PropertyUnit: "ps", PropertyCode: "WWN"},
		{PropertyName: "Number of kitchens", PropertyUnit: "ps", PropertyCode: "KN"},
		{PropertyName: "Number of doors", PropertyUnit: "ps", PropertyCode: "DN"},
		{PropertyName: "Number of stairs", PropertyUnit: "ps", PropertyCode: "SN"},
		{PropertyName: "House perimeter", PropertyUnit: "sq.m.", PropertyCode: "HM"},
		{PropertyName: "Exterior finishing area", PropertyUnit: "sq.m.", PropertyCode: "EFA"},
		{PropertyName: "Interior finishing area", PropertyUnit: "sq.m.", PropertyCode: "IFA"},
		{PropertyName: "Tile area", PropertyUnit: "sq.m.", PropertyCode: "TA"},
	}

	var jobs = []Job{
		{JobName: "Removal of the fertile layer", StageName: "Excavation", PropertyCode: "FA"},
		{JobName: "Axis markings", StageName: "Excavation", PropertyCode: "FA"},
		{JobName: "Pile pouring", StageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Pile"},
		{JobName: "Pile Shrinkage 1", StageName: "Foundation", FoundationMaterial: "Pile"},
		{JobName: "Pile Grillage", StageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Pile"},
		{JobName: "Pile Shrinkage 2", StageName: "Foundation", FoundationMaterial: "Pile"},
		{JobName: "Ribbon Digging", StageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Ribbon"},
		{JobName: "Ribbon Tying reinforcement + formwork", StageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Ribbon"},
		{JobName: "Ribbon Pouring tape", StageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Ribbon"},
		{JobName: "Ribbon Shrinkage 3", StageName: "Foundation", FoundationMaterial: "Ribbon"},
		{JobName: "Plate Excavation", StageName: "Foundation", PropertyCode: "FV", FoundationMaterial: "Plate"},
		{JobName: "Plate Backfilling ASG + gravel + rammer", StageName: "Foundation", PropertyCode: "FV", FoundationMaterial: "Plate"},
		{JobName: "Plate Styrofoam, foil, formwork + reinforcement", StageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Plate"},
		{JobName: "Plate Fill", StageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Plate"},
		{JobName: "Plate Shrinkage 4", StageName: "Foundation", FoundationMaterial: "Plate"},
		{JobName: "Backfilling of the earth", StageName: "Foundation", PropertyCode: "FA"},
		{JobName: "Communications (piping)", StageName: "Foundation"},
		{JobName: "Foam block", StageName: "Walls", PropertyCode: "WV", WallMaterial: "Foam block"},
		{JobName: "Brick", StageName: "Walls", PropertyCode: "WV", WallMaterial: "Brick"},
		{JobName: "CLT", StageName: "Walls", PropertyCode: "FA", WallMaterial: "CLT"},
		{JobName: "Frame", StageName: "Walls", PropertyCode: "WN", WallMaterial: "Frame"},
		{JobName: "Roof frame", StageName: "Roof", PropertyCode: "RA"},
		{JobName: "Fold", StageName: "Roof", PropertyCode: "RA", RoofingMaterial: "Fold"},
		{JobName: "Soft roof", StageName: "Roof", PropertyCode: "RA", RoofingMaterial: "Soft roof"},
		{JobName: "Roof tiles (metal/soft)", StageName: "Roof", PropertyCode: "RA", RoofingMaterial: "Roof tiles"},
		{JobName: "Windows and windowsills", StageName: "Windows and windowsills", PropertyCode: "WWN"},
		{JobName: "Warming", StageName: "Warming", PropertyCode: "EFA"},
		{JobName: "Subfloor/Floor System", StageName: "Subfloor/Floor System", PropertyCode: "FA"},
		{JobName: "Stairs", StageName: "Stairs", PropertyCode: "SN"},
		{JobName: "Plaster", StageName: "Exterior decoration of the house", PropertyCode: "WV", FinishMaterial: "Plaster"},
		{JobName: "Ventfacade", StageName: "Exterior decoration of the house", PropertyCode: "WV", FinishMaterial: "Ventfacade"},
		{JobName: "Floor", StageName: "Floor", PropertyCode: "TFA"},
		{JobName: "Electrical wiring", StageName: "Electrical wiring", PropertyCode: "TFA"},
		{JobName: "Plaster + Painting", StageName: "Interior decoration of the house", PropertyCode: "WV", InteriorMaterial: "Plaster + Painting"},
		{JobName: "Tile", StageName: "Interior decoration of the house", PropertyCode: "TA", InteriorMaterial: "Tile"},
		{JobName: "Doors", StageName: "Doors", PropertyCode: "DN"},
		{JobName: "Kitchen assembly, equipment installation", StageName: "Kitchen assembly, equipment installation", PropertyCode: "KN"},
		{JobName: "Plumbing", StageName: "Plumbing", PropertyCode: "TFA"},
		{JobName: "Lighting, switches", StageName: "Lighting, switches"},
		{JobName: "Furnishing", StageName: "Furnishing"},
		{JobName: "Commissioning works", StageName: "Commissioning works"},
	}

	//log.Println(properties)
	for _, prop := range properties {
		db.Create(&prop)
	}

	for _, job := range jobs {
		db.Create(&job)
	}
}
