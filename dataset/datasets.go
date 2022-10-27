package dataset

const RocksetPublicDatasets = "rockset-public-datasets"

type Sample string

const (
	CausesOfDeath Sample = "causes-of-death"
	Cities        Sample = "cities"
	MovieRatings  Sample = "movie-ratings"
	Movies        Sample = "movies"
	PartialCities Sample = "partial-cities"
	Reservations  Sample = "reservations"
	Stocks        Sample = "stocks"
)

var datasets = map[Sample]string{
	Cities:        "cities/*",
	MovieRatings:  "movie-ratings/*",
	Movies:        "movies/*",
	PartialCities: "partial-cities/*",
	CausesOfDeath: "causes-of-death-xml/*.xml",
	Reservations:  "reservations/reservations/*.json.gz",
	Stocks:        "stocks/nyse-listed_csv.csv",
}

func Lookup(dataset Sample) string {
	pattern := datasets[dataset]
	return pattern
}
