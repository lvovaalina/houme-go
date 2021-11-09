package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	PropertyId   int    `gorm:"primary_key;autoIncrement"`
	PropertyCode string `gorm:"index"`
	PropertyName string `gorm:"unique"`
	PropertyUnit string
}

type Job struct {
	gorm.Model
	JobId              int    `gorm:"unique;primary_key;autoIncrement"`
	JobName            string `gorm:"unique"`
	StageName          string
	SubStageName       string
	WallMaterial       string
	FinishMaterial     string
	FoundationMaterial string
	RoofingMaterial    string
	InteriorMaterial   string
	Required           bool
	PropertyCode       string
	InParallel         bool
	ParallelGroupCode  string
	JobCode            string `gorm:"unique"`
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
	JobID                          string
	Job                            Job `gorm:"references:JobCode"`
	CompanyName                    string
}

var db *gorm.DB

func main() {
	var err error
	// db, err = gorm.Open(
	// 	"postgres",
	// 	"host=ec2-52-22-81-147.compute-1.amazonaws.com port=5432 user=soxoxijvmbhqiv dbname=ddnmu64tjqh9ju password=0ab277b623defd4ca7a72cba84bc60f06d7cabb6a8b311bc7580250bcef78b69")

	// db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=houmly password=l8397040 sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// 	panic("failed to connect database")
	// }

	dbHost, dbUser, dbPassword, dbName := "localhost", "postgres", "l8397040", "houmly"

	var connectionString = fmt.Sprintf(
		"host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbUser, dbName, dbPassword,
	)

	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{QueryFields: true})
	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}

	db.Migrator().DropTable(&Property{})
	db.Migrator().DropTable(&Job{})
	db.Migrator().DropTable(&ConstructionJobProperty{})

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
		{JobCode: "rem-fert-lay", StageName: "Excavation", JobName: "Removal of the fertile layer", SubStageName: "Excavation", PropertyCode: "FA", Required: true},
		{JobCode: "ax-mark", StageName: "Excavation", JobName: "Axis markings", SubStageName: "Excavation", PropertyCode: "FA", Required: true},
		{JobCode: "pile-pour", StageName: "Foundation", JobName: "Pile pouring", SubStageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Pile", Required: false},
		{JobCode: "pile-shr-1", StageName: "Foundation", JobName: "Pile Shrinkage 1", SubStageName: "Foundation", FoundationMaterial: "Pile", Required: false},
		{JobCode: "pile-grill", StageName: "Foundation", JobName: "Pile Grillage", SubStageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Pile", Required: false},
		{JobCode: "pile-shr-2", StageName: "Foundation", InParallel: true, ParallelGroupCode: "par-group-1", JobName: "Pile Shrinkage 2", SubStageName: "Foundation", FoundationMaterial: "Pile", Required: false},
		{JobCode: "ribbon-dig", StageName: "Foundation", JobName: "Ribbon Digging", SubStageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Ribbon", Required: false},
		{JobCode: "ribbon-tying-formwork", StageName: "Foundation", JobName: "Ribbon Tying reinforcement + formwork", SubStageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Ribbon", Required: false},
		{JobCode: "ribbon-pour", StageName: "Foundation", JobName: "Ribbon Pouring tape", SubStageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Ribbon", Required: false},
		{JobCode: "ribbon-shr-3", StageName: "Foundation", InParallel: true, ParallelGroupCode: "par-group-1", JobName: "Ribbon Shrinkage 3", SubStageName: "Foundation", FoundationMaterial: "Ribbon", Required: false},
		{JobCode: "plate-exv", StageName: "Foundation", JobName: "Plate Excavation", SubStageName: "Foundation", PropertyCode: "FV", FoundationMaterial: "Plate", Required: false},
		{JobCode: "plate-backfl-grvl-ramr", StageName: "Foundation", JobName: "Plate Backfilling ASG + gravel + rammer", SubStageName: "Foundation", PropertyCode: "FV", FoundationMaterial: "Plate", Required: false},
		{JobCode: "plate-styrofoam-foil-form-reinf", StageName: "Foundation", JobName: "Plate Styrofoam, foil, formwork + reinforcement", SubStageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Plate", Required: false},
		{JobCode: "plate-fill", StageName: "Foundation", JobName: "Plate Fill", SubStageName: "Foundation", PropertyCode: "FA", FoundationMaterial: "Plate", Required: false},
		{JobCode: "plate-shr-4", StageName: "Foundation", InParallel: true, ParallelGroupCode: "par-group-1", JobName: "Plate Shrinkage 4", SubStageName: "Foundation", FoundationMaterial: "Plate", Required: false},
		{JobCode: "backfl-earth", StageName: "Foundation", InParallel: true, ParallelGroupCode: "par-group-1", JobName: "Backfilling of the earth", SubStageName: "Foundation", PropertyCode: "FA", Required: true},
		{JobCode: "commun", StageName: "Foundation", InParallel: true, ParallelGroupCode: "par-group-1", JobName: "Communications (piping)", SubStageName: "Foundation", Required: true},
		{JobCode: "foam-blck", StageName: "Box", JobName: "Foam block", SubStageName: "Walls", PropertyCode: "WV", WallMaterial: "Foam block", Required: false},
		{JobCode: "brick", StageName: "Box", JobName: "Brick", SubStageName: "Walls", PropertyCode: "WV", WallMaterial: "Brick", Required: false},
		{JobCode: "clt", StageName: "Box", JobName: "CLT", SubStageName: "Walls", PropertyCode: "FA", WallMaterial: "CLT", Required: false},
		{JobCode: "framt", StageName: "Box", JobName: "Frame", SubStageName: "Walls", PropertyCode: "WN", WallMaterial: "Frame", Required: false},
		{JobCode: "roof-frame", StageName: "Box", JobName: "Roof frame", SubStageName: "Roof", PropertyCode: "RA", Required: true},
		{JobCode: "fold", StageName: "Box", InParallel: true, ParallelGroupCode: "par-group-2", JobName: "Fold", SubStageName: "Roof", PropertyCode: "RA", RoofingMaterial: "Fold", Required: false},
		{JobCode: "soft-roof", StageName: "Box", InParallel: true, ParallelGroupCode: "par-group-2", JobName: "Soft roof", SubStageName: "Roof", PropertyCode: "RA", RoofingMaterial: "Soft roof", Required: false},
		{JobCode: "roof-tiles", StageName: "Box", InParallel: true, ParallelGroupCode: "par-group-2", JobName: "Roof tiles (metal/soft)", SubStageName: "Roof", PropertyCode: "RA", RoofingMaterial: "Roof tiles", Required: false},
		{JobCode: "wind-windsills", StageName: "Box", InParallel: true, ParallelGroupCode: "par-group-2", JobName: "Windows and windowsills", SubStageName: "Windows and windowsills", PropertyCode: "WWN", Required: true},
		{JobCode: "warming", StageName: "Box", InParallel: true, ParallelGroupCode: "par-group-2", JobName: "Warming", SubStageName: "Warming", PropertyCode: "EFA", Required: true},
		{JobCode: "floor-sys", StageName: "Box", InParallel: true, ParallelGroupCode: "par-group-3", JobName: "Subfloor/Floor System", SubStageName: "Subfloor/Floor System", PropertyCode: "FA", Required: true},
		{JobCode: "stairs", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-3", JobName: "Stairs", SubStageName: "Stairs", PropertyCode: "SN", Required: true},
		{JobCode: "plaster", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-3", JobName: "Plaster", SubStageName: "Exterior decoration of the house", PropertyCode: "WV", FinishMaterial: "Plaster", Required: false},
		{JobCode: "ventfacade", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-3", JobName: "Ventfacade", SubStageName: "Exterior decoration of the house", PropertyCode: "WV", FinishMaterial: "Ventfacade", Required: false},
		{JobCode: "floor", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-4", JobName: "Floor", SubStageName: "Floor", PropertyCode: "TFA", Required: true},
		{JobCode: "elecrt-wiring", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-4", JobName: "Electrical wiring", SubStageName: "Electrical wiring", PropertyCode: "TFA", Required: true},
		{JobCode: "plumbing", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-4", JobName: "Plumbing", SubStageName: "Plumbing", PropertyCode: "TFA", Required: true},
		{JobCode: "plast-paint", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-5", JobName: "Plaster + Painting", SubStageName: "Interior decoration of the house", PropertyCode: "WV", InteriorMaterial: "Plaster + Painting", Required: false},
		{JobCode: "tile", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-5", JobName: "Tile", SubStageName: "Interior decoration of the house", PropertyCode: "TA", InteriorMaterial: "Tile", Required: false},
		{JobCode: "doors", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-5", JobName: "Doors", SubStageName: "Doors", PropertyCode: "DN", Required: true},
		{JobCode: "kitchen-assbly-eq-inst", StageName: "Furnishing", InParallel: true, ParallelGroupCode: "par-group-6", JobName: "Kitchen assembly, equipment installation", SubStageName: "Kitchen assembly, equipment installation", PropertyCode: "KN", Required: true},
		{JobCode: "light-switches", StageName: "Furnishing", InParallel: true, ParallelGroupCode: "par-group-6", JobName: "Lighting, switches", SubStageName: "Lighting, switches", Required: true},
		{JobCode: "furnish", StageName: "Furnishing", InParallel: true, ParallelGroupCode: "par-group-6", JobName: "Furnishing", SubStageName: "Furnishing", Required: true},
		{JobCode: "comiss-works", StageName: "Commissioning works", JobName: "Commissioning works", SubStageName: "Commissioning works", Required: true},
	}

	var constructionProperties = []ConstructionJobProperty{
		{CompanyName: "Construction", Job: Job{JobCode: "rem-fert-lay"}, ConstructionSpeed: 25.0, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},
		{CompanyName: "Construction", Job: Job{JobCode: "ax-mark"}, ConstructionSpeed: 6.25, ConstructionCost: 0, MinWorkers: 4, MaxWorkers: 4, OptWorkers: 4},

		{CompanyName: "Construction", Job: Job{JobCode: "pile-pour"}, ConstructionSpeed: 0.56, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 6, OptWorkers: 3},
		{CompanyName: "Construction", Job: Job{JobCode: "pile-shr-1"}, ConstructionFixDurationInHours: 14, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},
		{CompanyName: "Construction", Job: Job{JobCode: "pile-grill"}, ConstructionSpeed: 0.83, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{CompanyName: "Construction", Job: Job{JobCode: "pile-shr-2"}, ConstructionFixDurationInHours: 26, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},

		{CompanyName: "Construction", Job: Job{JobCode: "ribbon-dig"}, ConstructionSpeed: 1.67, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{CompanyName: "Construction", Job: Job{JobCode: "ribbon-tying-formwork"}, ConstructionSpeed: 0.67, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{CompanyName: "Construction", Job: Job{JobCode: "ribbon-pour"}, ConstructionSpeed: 0.42, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{CompanyName: "Construction", Job: Job{JobCode: "ribbon-shr-3"}, ConstructionFixDurationInHours: 26, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},

		{CompanyName: "Construction", Job: Job{JobCode: "plate-exv"}, ConstructionSpeed: 1.88, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "plate-backfl-grvl-ramr"}, ConstructionSpeed: 0.94, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "plate-styrofoam-foil-form-reinf"}, ConstructionSpeed: 0.42, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "plate-fill"}, ConstructionSpeed: 2.5, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "plate-shr-4"}, ConstructionFixDurationInHours: 26, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},

		{CompanyName: "Construction", Job: Job{JobCode: "backfl-earth"}, ConstructionSpeed: 2.22, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{CompanyName: "Construction", Job: Job{JobCode: "commun"}, ConstructionFixDurationInHours: 24, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},

		{CompanyName: "Construction", Job: Job{JobCode: "foam-blck"}, ConstructionSpeed: 2.67, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 12, OptWorkers: 6},
		{CompanyName: "Construction", Job: Job{JobCode: "brick"}, ConstructionSpeed: 0.13, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 12, OptWorkers: 6},
		{CompanyName: "Construction", Job: Job{JobCode: "clt"}, ConstructionSpeed: 0.78, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 16, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "framt"}, ConstructionSpeed: 0.13, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},

		{CompanyName: "Construction", Job: Job{JobCode: "roof-frame"}, ConstructionSpeed: 0.31, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "fold"}, ConstructionSpeed: 0.83, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "soft-roof"}, ConstructionSpeed: 1.25, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "roof-tiles"}, ConstructionSpeed: 1.67, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},

		{CompanyName: "Construction", Job: Job{JobCode: "wind-windsills"}, ConstructionSpeed: 0.2, ConstructionCost: 0, MinWorkers: 2, MaxWorkers: 6, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "warming"}, ConstructionSpeed: 2.0, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "floor-sys"}, ConstructionSpeed: 0.63, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "stairs"}, ConstructionSpeed: 0.01, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},

		{CompanyName: "Construction", Job: Job{JobCode: "plaster"}, ConstructionSpeed: 1.67, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 8},
		{CompanyName: "Construction", Job: Job{JobCode: "ventfacade"}, ConstructionSpeed: 0.42, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 16, OptWorkers: 8},

		{CompanyName: "Construction", Job: Job{JobCode: "floor"}, ConstructionSpeed: 1.25, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "elecrt-wiring"}, ConstructionSpeed: 1.25, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 2},

		{CompanyName: "Construction", Job: Job{JobCode: "plast-paint"}, ConstructionSpeed: 1.25, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "tile"}, ConstructionSpeed: 1.00, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},

		{CompanyName: "Construction", Job: Job{JobCode: "doors"}, ConstructionSpeed: 0.25, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 2},
		{CompanyName: "Construction", Job: Job{JobCode: "kitchen-assbly-eq-inst"}, ConstructionSpeed: 0.03, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 2, OptWorkers: 2},
		{CompanyName: "Construction", Job: Job{JobCode: "plumbing"}, ConstructionSpeed: 3.33, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 2},
		{CompanyName: "Construction", Job: Job{JobCode: "light-switches"}, ConstructionFixDurationInHours: 32, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 2, OptWorkers: 2},
		{CompanyName: "Construction", Job: Job{JobCode: "furnish"}, ConstructionFixDurationInHours: 64, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 6, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "comiss-works"}, ConstructionFixDurationInHours: 48, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 2},
	}

	log.Println(constructionProperties[0].CompanyName)

	db.Create(&properties)

	db.Create(&jobs)

	db.Create(&constructionProperties)

}
