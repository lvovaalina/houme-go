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
	JobCode            string `gorm:"unique,index"`
	StageName          string
	PropertyCode       string
	WallMaterial       string
	FinishMaterial     string
	FoundationMaterial string
	RoofingMaterial    string
	InteriorMaterial   string
	Property           Property `gorm:"foreignKey:PropertyCode"`
}

type ConstructionJobProperty struct {
	gorm.Model
	ConstructionJobPropertyId      int `gorm:"primary_key;autoIncrement"`
	ConstructionSpeed              float32
	ConstructionCost               float32
	ConstructionFixDurationInHours float32
	MaxWorkers                     int
	OptWorkers                     int
	MinWorkers                     int
	JobCode                        string
	Job                            Job `gorm:"foreignKey:JobCode"`
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

	db.DropTableIfExists(&Property{})
	db.DropTableIfExists(&Job{})
	db.DropTableIfExists(&ConstructionJobProperty{})

	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Property{})
	db.AutoMigrate(&Job{})
	db.AutoMigrate(&ConstructionJobProperty{})

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
		{JobCode: "rem-fert-lay", JobName: "Removal of the fertile layer", StageName: "Excavation", PropertyCode: "FA"},
		{JobCode: "ax-mark", JobName: "Axis markings", StageName: "Excavation", PropertyCode: "FA"},
		{JobCode: "pile-pour", JobName: "Pile pouring", StageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Pile"},
		{JobCode: "pile-shr-1", JobName: "Pile Shrinkage 1", StageName: "Foundation", FoundationMaterial: "Pile"},
		{JobCode: "pile-grill", JobName: "Pile Grillage", StageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Pile"},
		{JobCode: "pile-shr-2", JobName: "Pile Shrinkage 2", StageName: "Foundation", FoundationMaterial: "Pile"},
		{JobCode: "ribbon-dig", JobName: "Ribbon Digging", StageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Ribbon"},
		{JobCode: "ribbon-tying-formwork", JobName: "Ribbon Tying reinforcement + formwork", StageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Ribbon"},
		{JobCode: "ribbon-pour", JobName: "Ribbon Pouring tape", StageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Ribbon"},
		{JobCode: "ribbon-shr-3", JobName: "Ribbon Shrinkage 3", StageName: "Foundation", FoundationMaterial: "Ribbon"},
		{JobCode: "plate-exv", JobName: "Plate Excavation", StageName: "Foundation", PropertyCode: "FV", FoundationMaterial: "Plate"},
		{JobCode: "plate-backfl-grvl-ramr", JobName: "Plate Backfilling ASG + gravel + rammer", StageName: "Foundation", PropertyCode: "FV", FoundationMaterial: "Plate"},
		{JobCode: "plate-styrofoam-foil-form-reinf", JobName: "Plate Styrofoam, foil, formwork + reinforcement", StageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Plate"},
		{JobCode: "plate-fill", JobName: "Plate Fill", StageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Plate"},
		{JobCode: "plate-shr-4", JobName: "Plate Shrinkage 4", StageName: "Foundation", FoundationMaterial: "Plate"},
		{JobCode: "backfl-earth", JobName: "Backfilling of the earth", StageName: "Foundation", PropertyCode: "FA"},
		{JobCode: "commun", JobName: "Communications (piping)", StageName: "Foundation"},
		{JobCode: "foam-blck", JobName: "Foam block", StageName: "Walls", PropertyCode: "WV", WallMaterial: "Foam block"},
		{JobCode: "brick", JobName: "Brick", StageName: "Walls", PropertyCode: "WV", WallMaterial: "Brick"},
		{JobCode: "clt", JobName: "CLT", StageName: "Walls", PropertyCode: "FA", WallMaterial: "CLT"},
		{JobCode: "framt", JobName: "Frame", StageName: "Walls", PropertyCode: "WN", WallMaterial: "Frame"},
		{JobCode: "roof-frame", JobName: "Roof frame", StageName: "Roof", PropertyCode: "RA"},
		{JobCode: "fold", JobName: "Fold", StageName: "Roof", PropertyCode: "RA", RoofingMaterial: "Fold"},
		{JobCode: "soft-roof", JobName: "Soft roof", StageName: "Roof", PropertyCode: "RA", RoofingMaterial: "Soft roof"},
		{JobCode: "roof-tiles", JobName: "Roof tiles (metal/soft)", StageName: "Roof", PropertyCode: "RA", RoofingMaterial: "Roof tiles"},
		{JobCode: "wind-windsills", JobName: "Windows and windowsills", StageName: "Windows and windowsills", PropertyCode: "WWN"},
		{JobCode: "warming", JobName: "Warming", StageName: "Warming", PropertyCode: "EFA"},
		{JobCode: "floor-sys", JobName: "Subfloor/Floor System", StageName: "Subfloor/Floor System", PropertyCode: "FA"},
		{JobCode: "stairs", JobName: "Stairs", StageName: "Stairs", PropertyCode: "SN"},
		{JobCode: "plaster", JobName: "Plaster", StageName: "Exterior decoration of the house", PropertyCode: "WV", FinishMaterial: "Plaster"},
		{JobCode: "ventfacade", JobName: "Ventfacade", StageName: "Exterior decoration of the house", PropertyCode: "WV", FinishMaterial: "Ventfacade"},
		{JobCode: "floor", JobName: "Floor", StageName: "Floor", PropertyCode: "TFA"},
		{JobCode: "elecrt-wiring", JobName: "Electrical wiring", StageName: "Electrical wiring", PropertyCode: "TFA"},
		{JobCode: "plast-paint", JobName: "Plaster + Painting", StageName: "Interior decoration of the house", PropertyCode: "WV", InteriorMaterial: "Plaster + Painting"},
		{JobCode: "tile", JobName: "Tile", StageName: "Interior decoration of the house", PropertyCode: "TA", InteriorMaterial: "Tile"},
		{JobCode: "doors", JobName: "Doors", StageName: "Doors", PropertyCode: "DN"},
		{JobCode: "kitchen-assbly-eq-inst", JobName: "Kitchen assembly, equipment installation", StageName: "Kitchen assembly, equipment installation", PropertyCode: "KN"},
		{JobCode: "plumbing", JobName: "Plumbing", StageName: "Plumbing", PropertyCode: "TFA"},
		{JobCode: "light-switches", JobName: "Lighting, switches", StageName: "Lighting, switches"},
		{JobCode: "furnish", JobName: "Furnishing", StageName: "Furnishing"},
		{JobCode: "comiss-works", JobName: "Commissioning works", StageName: "Commissioning works"},
	}

	var constructionProperties = []ConstructionJobProperty{
		{JobCode: "rem-fert-lay", ConstructionSpeed: 25.0, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},
		{JobCode: "ax-mark", ConstructionSpeed: 6.25, ConstructionCost: 0, MinWorkers: 4, MaxWorkers: 4, OptWorkers: 4},

		{JobCode: "pile-pour", ConstructionSpeed: 0.56, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 6, OptWorkers: 3},
		{JobCode: "pile-shr-1", ConstructionFixDurationInHours: 14, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},
		{JobCode: "pile-grill", ConstructionSpeed: 0.83, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{JobCode: "pile-shr-2", ConstructionFixDurationInHours: 26, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},

		{JobCode: "ribbon-dig", ConstructionSpeed: 1.67, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{JobCode: "ribbon-tying-formwork", ConstructionSpeed: 0.67, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{JobCode: "ribbon-pour", ConstructionSpeed: 0.42, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{JobCode: "ribbon-shr-3", ConstructionFixDurationInHours: 26, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},

		{JobCode: "plate-exv", ConstructionSpeed: 1.88, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},
		{JobCode: "plate-backfl-grvl-ramr", ConstructionSpeed: 0.94, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},
		{JobCode: "plate-styrofoam-foil-form-reinf", ConstructionSpeed: 0.42, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},
		{JobCode: "plate-fill", ConstructionSpeed: 2.5, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},
		{JobCode: "plate-shr-4", ConstructionFixDurationInHours: 26, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},

		{JobCode: "backfl-earth", ConstructionSpeed: 2.22, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{JobCode: "commun", ConstructionFixDurationInHours: 24, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},

		{JobCode: "foam-blck", ConstructionSpeed: 2.67, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 12, OptWorkers: 6},
		{JobCode: "brick", ConstructionSpeed: 0.13, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 12, OptWorkers: 6},
		{JobCode: "clt", ConstructionSpeed: 0.78, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 16, OptWorkers: 4},
		{JobCode: "framt", ConstructionSpeed: 0.13, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},

		{JobCode: "roof-frame", ConstructionSpeed: 0.31, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},
		{JobCode: "fold", ConstructionSpeed: 0.83, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},
		{JobCode: "soft-roof", ConstructionSpeed: 1.25, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},
		{JobCode: "roof-tiles", ConstructionSpeed: 1.67, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},

		{JobCode: "wind-windsills", ConstructionSpeed: 0.2, ConstructionCost: 0, MinWorkers: 2, MaxWorkers: 6, OptWorkers: 4},
		{JobCode: "warming", ConstructionSpeed: 2.0, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},
		{JobCode: "floor-sys", ConstructionSpeed: 0.63, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},
		{JobCode: "stairs", ConstructionSpeed: 0.01, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},

		{JobCode: "plaster", ConstructionSpeed: 1.67, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 8},
		{JobCode: "ventfacade", ConstructionSpeed: 0.42, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 16, OptWorkers: 8},

		{JobCode: "floor", ConstructionSpeed: 1.25, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},
		{JobCode: "elecrt-wiring", ConstructionSpeed: 1.25, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 2},

		{JobCode: "plast-paint", ConstructionSpeed: 1.25, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},
		{JobCode: "tile", ConstructionSpeed: 1.00, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},

		{JobCode: "doors", ConstructionSpeed: 0.25, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 2},
		{JobCode: "kitchen-assbly-eq-inst", ConstructionSpeed: 0.03, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 2, OptWorkers: 2},
		{JobCode: "plumbing", ConstructionSpeed: 3.33, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 2},
		{JobCode: "light-switches", ConstructionFixDurationInHours: 32, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 2, OptWorkers: 2},
		{JobCode: "furnish", ConstructionFixDurationInHours: 64, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 6, OptWorkers: 4},
		{JobCode: "comiss-works", ConstructionFixDurationInHours: 48, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 2},
	}

	//log.Println(properties)
	for _, prop := range properties {
		db.Create(&prop)
	}

	for _, job := range jobs {
		db.Create(&job)
	}

	for _, constr := range constructionProperties {
		db.Create(&constr)
	}
}
