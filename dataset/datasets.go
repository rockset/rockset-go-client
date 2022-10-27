package dataset

const (
	// RocksetPublicDatasets is the name of the AWS S3 bucket which holds the sample datasets.
	RocksetPublicDatasets = "rockset-public-datasets"
	// Region is the AWS region for the RocksetPublicDatasets bucket.
	Region = "us-west-2"
)

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

// Lookup retrieves the S3 pattern to use for the given dataset. Returns an empty string if the sample dataset
// doesn't exist.
func Lookup(dataset Sample) string {
	pattern := datasets[dataset]
	return pattern
}

// All returns a list of all datasets.
func All() []Sample {
	var list []Sample
	for k := range datasets {
		ds := k
		list = append(list, ds)
	}

	return list
}
