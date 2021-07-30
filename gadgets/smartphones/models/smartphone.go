package models

//Smartphone model stracture for smatphones
type Smartphone struct {
	Id            int64
	Name          string
	Price         string
	CountryOrigin string
	Os            string
}

// CreateSmartphoneCMD
type CreateSmartphoneCMD struct {
	Name          string `json:"name"`
	Price         string `json:"price"`
	CountryOrigin string `json:"country_origin"`
	Os            string `json:"os"`
}
