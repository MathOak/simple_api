package dto

type Films struct {
	MyRowID             uint64  `gorm:"primaryKey;autoIncrement" json:"my_row_id"`
	ID                  string  `json:"id"`
	Title               string  `json:"title"`
	Link                string  `json:"link"`
	Year                int     `json:"year"`
	Duration            string  `json:"duration"`
	RatingMPA           string  `json:"rating_mpa"`
	RatingIMDB          float64 `json:"rating_imdb"`
	Vote                int     `json:"vote"`
	Budget              string  `json:"budget"`
	GrossWorldWide      string  `json:"gross_world_wide"`
	GrossUSCanada       string  `json:"gross_us_canada"`
	GrossOpeningWeekend string  `json:"gross_opening_weekend"`
	Director            string  `json:"director"`
	Writer              string  `json:"writer"`
	Star                string  `json:"star"`
	Genre               string  `json:"genre"`
	CountryOrigin       string  `json:"country_origin"`
	FilmingLocation     string  `json:"filming_location"`
	ProductionCompany   string  `json:"production_company"`
	Language            string  `json:"language"`
	Win                 int     `json:"win"`
	Nomination          int     `json:"nomination"`
	Oscar               int     `json:"oscar"`
}
