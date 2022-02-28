package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	PropertyId     int    `gorm:"primary_key;autoIncrement"`
	PropertyCode   string `gorm:"unique"`
	PropertyName   string `gorm:"unique"`
	PropertyNamePL string
	PropertyUnit   string
}

type Job struct {
	gorm.Model
	JobId              int    `gorm:"primary_key;autoIncrement"`
	JobName            string `gorm:"unique"`
	StageName          string
	SubStageName       string
	JobNamePL          string
	StageNamePL        string
	SubStageNamePL     string
	WallMaterial       string
	FinishMaterial     string
	FoundationMaterial string
	RoofingMaterial    string
	InteriorMaterial   string
	Required           bool
	InParallel         bool
	ParallelGroupCode  string
	JobCode            string `gorm:"unique"`

	PropertyID *string
	Property   Property `gorm:"references:PropertyCode"`
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
	//sslmode=disable
	// dbHost, dbUser, dbPassword, dbName :=
	// 	"ec2-3-209-38-221.compute-1.amazonaws.com",
	// 	"tqykcafyttxirk",
	// 	"610c4d0e7a853449d21c4e1344b432b7ebff5a71466166728ee9dca963958fb8",
	// 	"d440nml0a86pe8"

	var connectionString = fmt.Sprintf(
		"host=%s port=5432 user=%s dbname=%s password=%s",
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
		{PropertyName: "Foundation volume", PropertyNamePL: "Objętość podkładu", PropertyUnit: "sq.m.", PropertyCode: "FV"},
		{PropertyName: "Floor area at the base", PropertyNamePL: "Powierzchnia fundamentów", PropertyUnit: "sq.m.", PropertyCode: "FA"},
		{PropertyName: "Total floor area", PropertyNamePL: "Całkowita powierzchnia podłogi", PropertyUnit: "sq.m.", PropertyCode: "TFA"},
		{PropertyName: "Walls volume", PropertyNamePL: "Objętość ścian", PropertyUnit: "sq.m.", PropertyCode: "WV"},
		{PropertyName: "Number of walls", PropertyNamePL: "Liczba ścian", PropertyUnit: "ps", PropertyCode: "WN"},
		{PropertyName: "Roof area", PropertyNamePL: "Powierzchnia dachu", PropertyUnit: "sq.m.", PropertyCode: "RA"},
		{PropertyName: "Number of windows", PropertyNamePL: "Liczba okien", PropertyUnit: "ps", PropertyCode: "WWN"},
		{PropertyName: "Number of kitchens", PropertyNamePL: "Number of kitchens", PropertyUnit: "ps", PropertyCode: "KN"},
		{PropertyName: "Number of doors", PropertyNamePL: "Number of doors", PropertyUnit: "ps", PropertyCode: "DN"},
		{PropertyName: "Number of stairs", PropertyNamePL: "Number of stairs", PropertyUnit: "ps", PropertyCode: "SN"},
		{PropertyName: "House perimeter", PropertyNamePL: "Obwód domu", PropertyUnit: "sq.m.", PropertyCode: "HM"},
		{PropertyName: "Exterior finishing area", PropertyNamePL: "Zewnętrzna powierzchnia wykończeniowa", PropertyUnit: "sq.m.", PropertyCode: "EFA"},
		{PropertyName: "Interior finishing area", PropertyNamePL: "Powierzchnia wykończenia wnętrz", PropertyUnit: "sq.m.", PropertyCode: "IFA"},
		{PropertyName: "Tile area", PropertyNamePL: "Obszar płytek", PropertyUnit: "sq.m.", PropertyCode: "TA"},
		{PropertyName: "Rooms Number", PropertyNamePL: "Liczba pokoi", PropertyUnit: "ps", PropertyCode: "RN"},
		{PropertyName: "Points Number", PropertyNamePL: "Liczba punktów", PropertyUnit: "ps", PropertyCode: "PN"},
		{PropertyName: "Outlet Number", PropertyNamePL: "Numer wylotu", PropertyUnit: "ps", PropertyCode: "ON"},
		{PropertyName: "Pipes Length", PropertyNamePL: "Długość rur", PropertyUnit: "rm", PropertyCode: "PL"},
	}

	var jobs = []Job{
		{JobCode: "rem-fert-lay", StageName: "Excavation", JobName: "Removal of the fertile layer", SubStageName: "Excavation", StageNamePL: "Wykop", JobNamePL: "Usunięcie płodnej warstwy", SubStageNamePL: "Wykop", Property: Property{PropertyCode: "FA"}, Required: true},
		{JobCode: "ax-mark", StageName: "Excavation", JobName: "Axis markings", SubStageName: "Excavation", StageNamePL: "Wykop", JobNamePL: "Oznaczenia osi", SubStageNamePL: "Wykop", Property: Property{PropertyCode: "FA"}, Required: true},
		// {JobCode: "pile-drill", StageName: "Foundation", JobName: "Pile Drilling", SubStageName: "Foundation", Property: Property{PropertyCode: "FA"}, FoundationMaterial: "Pile", Required: false},

		{JobCode: "pile-pour", StageName: "Foundation", JobName: "Pile Pouring", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Wylewanie stosu", SubStageNamePL: "Fundamenta", Property: Property{PropertyCode: "FA"}, FoundationMaterial: "Pile", Required: false},
		{JobCode: "pile-shr-1", StageName: "Foundation", JobName: "Pile Shrinkage 1", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Skurcz pala 1", SubStageNamePL: "Fundamenta", PropertyID: nil, FoundationMaterial: "Pile", Required: false},
		{JobCode: "pile-grill", StageName: "Foundation", JobName: "Pile Grillage", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Grillowanie palowe", SubStageNamePL: "Fundamenta", Property: Property{PropertyCode: "FA"}, FoundationMaterial: "Pile", Required: false},
		{JobCode: "pile-shr-2", StageName: "Foundation", InParallel: true, ParallelGroupCode: "par-group-1", JobName: "Pile Shrinkage 2", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Skurcz pala 2", SubStageNamePL: "Fundamenta", FoundationMaterial: "Pile", Required: false},
		{JobCode: "ribbon-dig", StageName: "Foundation", JobName: "Ribbon Digging", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Kopanie wstążki", SubStageNamePL: "Fundamenta", Property: Property{PropertyCode: "FA"}, FoundationMaterial: "Ribbon", Required: false},
		{JobCode: "ribbon-tying-formwork", StageName: "Foundation", JobName: "Ribbon Tying reinforcement + formwork", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Zbrojenie wiązania taśmą + szalunek", SubStageNamePL: "Fundamenta", Property: Property{PropertyCode: "FA"}, FoundationMaterial: "Ribbon", Required: false},
		{JobCode: "ribbon-pour", StageName: "Foundation", JobName: "Ribbon Pouring tape", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Taśma do wylewania wstążki", SubStageNamePL: "Fundamenta", Property: Property{PropertyCode: "FA"}, FoundationMaterial: "Ribbon", Required: false},
		{JobCode: "ribbon-shr-3", StageName: "Foundation", InParallel: true, ParallelGroupCode: "par-group-1", JobName: "Ribbon Shrinkage 3", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Skurcz wstążki 3", SubStageNamePL: "Fundamenta", FoundationMaterial: "Ribbon", Required: false},
		{JobCode: "plate-exv", StageName: "Foundation", JobName: "Plate Excavation", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Wykop płyty", SubStageNamePL: "Fundamenta", Property: Property{PropertyCode: "FV"}, FoundationMaterial: "Plate", Required: false},
		{JobCode: "plate-backfl-grvl-ramr", StageName: "Foundation", JobName: "Zasyp płyt ASG + żwir + ubijak", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Pile Pouring", SubStageNamePL: "Fundamenta", Property: Property{PropertyCode: "FV"}, FoundationMaterial: "Plate", Required: false},
		{JobCode: "plate-styrofoam-foil-form-reinf", StageName: "Foundation", JobName: "Plate Styrofoam, foil, formwork + reinforcement", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Płyta styropian, folia, szalunek + wzmocnienie", SubStageNamePL: "Fundamenta", Property: Property{PropertyCode: "FA"}, FoundationMaterial: "Plate", Required: false},
		{JobCode: "plate-fill", StageName: "Foundation", JobName: "Plate Fill", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Wypełnienie płyty", SubStageNamePL: "Fundamenta", Property: Property{PropertyCode: "FA"}, FoundationMaterial: "Plate", Required: false},
		{JobCode: "plate-shr-4", StageName: "Foundation", InParallel: true, ParallelGroupCode: "par-group-1", JobName: "Plate Shrinkage 4", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Skurcz płyty 4", SubStageNamePL: "Fundamenta", FoundationMaterial: "Plate", Required: false},
		{JobCode: "backfl-earth", StageName: "Foundation", InParallel: true, ParallelGroupCode: "par-group-1", JobName: "Backfilling of the earth", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Zasypywanie ziemi", SubStageNamePL: "Fundamenta", Property: Property{PropertyCode: "FA"}, Required: true},
		{JobCode: "commun", StageName: "Foundation", InParallel: true, ParallelGroupCode: "par-group-1", JobName: "Communications (piping)", SubStageName: "Foundation", StageNamePL: "Fundamenta", JobNamePL: "Komunikacja (rurociągi)", SubStageNamePL: "Fundamenta", Property: Property{PropertyCode: "PL"}, Required: true},

		{JobCode: "foam-blck", StageName: "Box", JobName: "Foam block", SubStageName: "Walls", StageNamePL: "Pudełko", JobNamePL: "Blok piankowy", SubStageNamePL: "Ściany",
			Property: Property{PropertyCode: "WV"}, WallMaterial: "Foam block", Required: false},
		{JobCode: "brick", StageName: "Box", JobName: "Brick", SubStageName: "Walls", StageNamePL: "Pudełko", JobNamePL: "Cegła", SubStageNamePL: "Ściany", Property: Property{PropertyCode: "WV"}, WallMaterial: "Brick", Required: false},

		{JobCode: "clt", StageName: "Box", JobName: "CLT", SubStageName: "Walls", StageNamePL: "Pudełko", JobNamePL: "CLT", SubStageNamePL: "Ściany", Property: Property{PropertyCode: "FA"}, WallMaterial: "CLT", Required: false},

		{JobCode: "framt", StageName: "Box", JobName: "Frame", SubStageName: "Walls", StageNamePL: "Pudełko", JobNamePL: "Rama", SubStageNamePL: "Ściany", Property: Property{PropertyCode: "WN"}, WallMaterial: "Frame", Required: false},

		{JobCode: "roof-frame", StageName: "Box", JobName: "Roof frame", SubStageName: "Roof", StageNamePL: "Pudełko", JobNamePL: "Rama dachu", SubStageNamePL: "Dach", Property: Property{PropertyCode: "RA"}, Required: true},

		{JobCode: "fold", StageName: "Box", InParallel: true, ParallelGroupCode: "par-group-2", StageNamePL: "Pudełko", JobNamePL: "Na rąbek", SubStageNamePL: "Dach", JobName: "Fold", SubStageName: "Roof", Property: Property{PropertyCode: "RA"}, RoofingMaterial: "Fold", Required: false},
		{JobCode: "soft-roof", StageName: "Box", InParallel: true, ParallelGroupCode: "par-group-2", StageNamePL: "Pudełko", JobNamePL: "Miękki dach", SubStageNamePL: "Dach", JobName: "Soft roof", SubStageName: "Roof", Property: Property{PropertyCode: "RA"}, RoofingMaterial: "Soft roof", Required: false},
		{JobCode: "roof-tiles", StageName: "Box", InParallel: true, ParallelGroupCode: "par-group-2", StageNamePL: "Pudełko", JobNamePL: "Dachówki (metalowe/miękkie)", SubStageNamePL: "Dach", JobName: "Roof tiles (metal/soft)", SubStageName: "Roof", Property: Property{PropertyCode: "RA"}, RoofingMaterial: "Roof tiles", Required: false},

		{JobCode: "wind-windsills", StageName: "Box", InParallel: true, ParallelGroupCode: "par-group-2", StageNamePL: "Pudełko", JobNamePL: "Okna i parapety", SubStageNamePL: "Okna i parapety", JobName: "Windows and windowsills", SubStageName: "Windows and windowsills", Property: Property{PropertyCode: "WWN"}, Required: true},

		{JobCode: "ins", StageName: "Box", InParallel: true, ParallelGroupCode: "par-group-2", StageNamePL: "Pudełko", JobNamePL: "Izolacja", SubStageNamePL: "Izolacja", JobName: "Insulation", SubStageName: "Insulation", Property: Property{PropertyCode: "EFA"}, Required: true},

		{JobCode: "floor-sys", StageName: "Box", InParallel: true, ParallelGroupCode: "par-group-3", StageNamePL: "Pudełko", JobNamePL: "Podkład podłogowy", SubStageNamePL: "Podkład podłogowy", JobName: "Subfloor/Floor System", SubStageName: "Subfloor/Floor System", Property: Property{PropertyCode: "FA"}, Required: true},
		// {JobCode: "stairs", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-3", JobName: "Stairs", SubStageName: "Stairs", Property: Property{PropertyCode: "SN"}, Required: true},
		// {JobCode: "plaster", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-3", JobName: "Plaster", SubStageName: "Exterior decoration of the house", Property: Property{PropertyCode: "WV"}, FinishMaterial: "Plaster", Required: false},
		// {JobCode: "ventfacade", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-3", JobName: "Ventfacade", SubStageName: "Exterior decoration of the house", Property: Property{PropertyCode: "WV"}, FinishMaterial: "Ventfacade", Required: false},
		// {JobCode: "floor", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-4", JobName: "Floor", SubStageName: "Floor", Property: Property{PropertyCode: "TFA"}, Required: true},
		// {JobCode: "elecrt-wiring", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-4", JobName: "Electrical wiring", SubStageName: "Electrical wiring", Property: Property{PropertyCode: "TFA"}, Required: true},
		// {JobCode: "plumbing", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-4", JobName: "Plumbing", SubStageName: "Plumbing", Property: Property{PropertyCode: "TFA"}, Required: true},
		// {JobCode: "plast-paint", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-5", JobName: "Plaster + Painting", SubStageName: "Interior decoration of the house", Property: Property{PropertyCode: "WV"}, InteriorMaterial: "Plaster + Painting", Required: false},
		// {JobCode: "tile", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-5", JobName: "Tile", SubStageName: "Interior decoration of the house", Property: Property{PropertyCode: "TA"}, InteriorMaterial: "Tile", Required: false},
		// {JobCode: "doors", StageName: "Interior", InParallel: true, ParallelGroupCode: "par-group-5", JobName: "Doors", SubStageName: "Doors", Property: Property{PropertyCode: "DN"}, Required: true},
		// {JobCode: "kitchen-assbly-eq-inst", StageName: "Furnishing", InParallel: true, ParallelGroupCode: "par-group-6", JobName: "Kitchen assembly, equipment installation", SubStageName: "Kitchen assembly, equipment installation", Property: Property{PropertyCode: "KN"}, Required: true},
		// {JobCode: "light", StageName: "Furnishing", InParallel: true, ParallelGroupCode: "par-group-6", JobName: "Lighting", Property: Property{PropertyCode: "PN"}, SubStageName: "Lighting, switches", Required: true},
		// {JobCode: "switches", StageName: "Furnishing", InParallel: true, ParallelGroupCode: "par-group-6", JobName: "Switches", Property: Property{PropertyCode: "ON"}, SubStageName: "Lighting, switches", Required: true},
		// {JobCode: "furnish", StageName: "Furnishing", InParallel: true, ParallelGroupCode: "par-group-6", JobName: "Furnishing", Property: Property{PropertyCode: "RN"}, SubStageName: "Furnishing", Required: true},
		{JobCode: "comiss-works", StageName: "Commissioning works", JobName: "Commissioning works", SubStageName: "Commissioning works", StageNamePL: "Prace uruchomieniowe", JobNamePL: "Prace uruchomieniowe", SubStageNamePL: "Prace uruchomieniowe", Required: true},
	}

	var constructionProperties = []ConstructionJobProperty{
		{CompanyName: "Construction", Job: Job{JobCode: "rem-fert-lay"}, ConstructionSpeed: 25.0, ConstructionCost: 6.0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},
		{CompanyName: "Construction", Job: Job{JobCode: "ax-mark"}, ConstructionSpeed: 6.25, ConstructionCost: 1.0, MinWorkers: 4, MaxWorkers: 4, OptWorkers: 4},

		// {CompanyName: "Construction", Job: Job{JobCode: "pile-drill"}, ConstructionSpeed: 0.56, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 6, OptWorkers: 3},
		{CompanyName: "Construction", Job: Job{JobCode: "pile-pour"}, ConstructionSpeed: 0.56, ConstructionCost: 12.0, MinWorkers: 1, MaxWorkers: 6, OptWorkers: 3},
		{CompanyName: "Construction", Job: Job{JobCode: "pile-shr-1"}, ConstructionFixDurationInHours: 14, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},
		{CompanyName: "Construction", Job: Job{JobCode: "pile-grill"}, ConstructionSpeed: 0.83, ConstructionCost: 5.0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{CompanyName: "Construction", Job: Job{JobCode: "pile-shr-2"}, ConstructionFixDurationInHours: 26, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},

		{CompanyName: "Construction", Job: Job{JobCode: "ribbon-dig"}, ConstructionSpeed: 1.67, ConstructionCost: 3.0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{CompanyName: "Construction", Job: Job{JobCode: "ribbon-tying-formwork"}, ConstructionSpeed: 0.67, ConstructionCost: 6.0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{CompanyName: "Construction", Job: Job{JobCode: "ribbon-pour"}, ConstructionSpeed: 0.42, ConstructionCost: 9.0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{CompanyName: "Construction", Job: Job{JobCode: "ribbon-shr-3"}, ConstructionFixDurationInHours: 26, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},

		{CompanyName: "Construction", Job: Job{JobCode: "plate-exv"}, ConstructionSpeed: 1.88, ConstructionCost: 3.0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "plate-backfl-grvl-ramr"}, ConstructionSpeed: 0.94, ConstructionCost: 6.0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "plate-styrofoam-foil-form-reinf"}, ConstructionSpeed: 0.42, ConstructionCost: 7.0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "plate-fill"}, ConstructionSpeed: 2.5, ConstructionCost: 3.0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "plate-shr-4"}, ConstructionFixDurationInHours: 26, ConstructionCost: 0, MinWorkers: 1, MaxWorkers: 1, OptWorkers: 1},

		{CompanyName: "Construction", Job: Job{JobCode: "backfl-earth"}, ConstructionSpeed: 2.22, ConstructionCost: 3.0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},
		{CompanyName: "Construction", Job: Job{JobCode: "commun"}, ConstructionFixDurationInHours: 24, ConstructionCost: 120.0, MinWorkers: 1, MaxWorkers: 3, OptWorkers: 3},

		{CompanyName: "Construction", Job: Job{JobCode: "foam-blck"}, ConstructionSpeed: 2.67, ConstructionCost: 120.0, MinWorkers: 1, MaxWorkers: 12, OptWorkers: 6},
		{CompanyName: "Construction", Job: Job{JobCode: "brick"}, ConstructionSpeed: 0.13, ConstructionCost: 140.0, MinWorkers: 1, MaxWorkers: 12, OptWorkers: 6},
		{CompanyName: "Construction", Job: Job{JobCode: "clt"}, ConstructionSpeed: 0.78, ConstructionCost: 50.0, MinWorkers: 1, MaxWorkers: 16, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "framt"}, ConstructionSpeed: 0.13, ConstructionCost: 38.0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},

		{CompanyName: "Construction", Job: Job{JobCode: "roof-frame"}, ConstructionSpeed: 0.31, ConstructionCost: 17.0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "fold"}, ConstructionSpeed: 0.83, ConstructionCost: 10.0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "soft-roof"}, ConstructionSpeed: 1.25, ConstructionCost: 8.0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "roof-tiles"}, ConstructionSpeed: 1.67, ConstructionCost: 8.0, MinWorkers: 1, MaxWorkers: 18, OptWorkers: 4},

		{CompanyName: "Construction", Job: Job{JobCode: "wind-windsills"}, ConstructionSpeed: 0.2, ConstructionCost: 100.0, MinWorkers: 2, MaxWorkers: 6, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "ins"}, ConstructionSpeed: 2.0, ConstructionCost: 12.0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "floor-sys"}, ConstructionSpeed: 0.63, ConstructionCost: 12.0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},
		// {CompanyName: "Construction", Job: Job{JobCode: "stairs"}, ConstructionSpeed: 0.01, ConstructionCost: 15.0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 4},

		// {CompanyName: "Construction", Job: Job{JobCode: "plaster"}, ConstructionSpeed: 1.67, ConstructionCost: 13.0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 8},
		// {CompanyName: "Construction", Job: Job{JobCode: "ventfacade"}, ConstructionSpeed: 0.42, ConstructionCost: 15.0, MinWorkers: 1, MaxWorkers: 16, OptWorkers: 8},

		// {CompanyName: "Construction", Job: Job{JobCode: "floor"}, ConstructionSpeed: 1.25, ConstructionCost: 22.0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},
		// {CompanyName: "Construction", Job: Job{JobCode: "elecrt-wiring"}, ConstructionSpeed: 1.25, ConstructionCost: 11.0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 2},

		// {CompanyName: "Construction", Job: Job{JobCode: "plast-paint"}, ConstructionSpeed: 1.25, ConstructionCost: 36.0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},
		// {CompanyName: "Construction", Job: Job{JobCode: "tile"}, ConstructionSpeed: 1.00, ConstructionCost: 36.0, MinWorkers: 1, MaxWorkers: 8, OptWorkers: 4},

		// {CompanyName: "Construction", Job: Job{JobCode: "doors"}, ConstructionSpeed: 0.25, ConstructionCost: 122.0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 2},
		// {CompanyName: "Construction", Job: Job{JobCode: "kitchen-assbly-eq-inst"}, ConstructionSpeed: 0.03, ConstructionCost: 15.0, MinWorkers: 1, MaxWorkers: 2, OptWorkers: 2},
		// {CompanyName: "Construction", Job: Job{JobCode: "plumbing"}, ConstructionSpeed: 3.33, ConstructionCost: 3.0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 2},
		// {CompanyName: "Construction", Job: Job{JobCode: "light"}, ConstructionFixDurationInHours: 16, ConstructionCost: 8.0, MinWorkers: 1, MaxWorkers: 2, OptWorkers: 2},
		// {CompanyName: "Construction", Job: Job{JobCode: "switches"}, ConstructionFixDurationInHours: 16, ConstructionCost: 10.0, MinWorkers: 1, MaxWorkers: 2, OptWorkers: 2},
		// {CompanyName: "Construction", Job: Job{JobCode: "furnish"}, ConstructionFixDurationInHours: 64, ConstructionCost: 30.0, MinWorkers: 1, MaxWorkers: 6, OptWorkers: 4},
		{CompanyName: "Construction", Job: Job{JobCode: "comiss-works"}, ConstructionFixDurationInHours: 48, ConstructionCost: 20.0, MinWorkers: 1, MaxWorkers: 4, OptWorkers: 2},
	}

	log.Println(constructionProperties[0].CompanyName)
	log.Println(jobs[0].JobCode)
	log.Println(properties[0].PropertyCode)

	db.Create(&properties)

	db.Create(&jobs)

	db.Create(&constructionProperties)

}
