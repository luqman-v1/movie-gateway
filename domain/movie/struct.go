package movie

import (
	"log"
	pb "movie-gateway/proto/movie"

	"github.com/jinzhu/copier"
)

type Data struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdb_id"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

type DetailMovie struct {
	Title      string     `json:"title"`
	Year       string     `json:"year"`
	Rated      string     `json:"rated"`
	Released   string     `json:"released"`
	Runtime    string     `json:"runtime"`
	Genre      string     `json:"genre"`
	Director   string     `json:"director"`
	Writer     string     `json:"writer"`
	Actors     string     `json:"actors"`
	Plot       string     `json:"plot"`
	Language   string     `json:"language"`
	Country    string     `json:"country"`
	Awards     string     `json:"awards"`
	Poster     string     `json:"poster"`
	Ratings    []*Ratings `json:"ratings"`
	Metascore  string     `json:"metascore"`
	ImdbRating string     `json:"imdb_rating"`
	ImdbVotes  string     `json:"imdb_votes"`
	ImdbID     string     `json:"imdb_id"`
	Type       string     `json:"type"`
	DVD        string     `json:"dvd"`
	BoxOffice  string     `json:"box_office"`
	Production string     `json:"production"`
	Website    string     `json:"website"`
	Response   string     `json:"response"`
}

type Ratings struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}

func (d *Data) Map(res *pb.ListResponse) []Data {
	var result []Data
	err := copier.Copy(&result, &res.Data)
	if err != nil {
		log.Println("error from copier:", err)
	}
	if len(result) <= 0 {
		return make([]Data, 0)
	}
	return result
}
func (d *DetailMovie) Map(res *pb.DetailMovie) *DetailMovie {
	err := copier.Copy(&d, &res)
	if err != nil {
		log.Println("error from copier:", err)
	}
	return d
}
