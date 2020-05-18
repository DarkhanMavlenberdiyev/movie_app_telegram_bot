package movie

import (
	"github.com/go-pg/pg"
)
//struct for user
type PostgreConfig struct {
	User     string
	Password string
	Port     string
	Host     string
}

// Postgre store movies
func PostgreMovies(config PostgreConfig) MoviesDb {
	db := pg.Connect(&pg.Options{
		Addr:     config.Host + ":" + config.Port,
		User:     config.User,
		Password: config.Password,
		Database: "movies",
	})
	return &postgreStore{db: db}
}


// postgreStore...
type postgreStore struct {
	db     *pg.DB
}


//	GetMovie...
func (p postgreStore) GetMovie(id int) (*Movie, error) {
	movie := &Movie{ID: id}
	err := p.db.Select(movie)

	if err != nil {
		return nil, err
	}
	return movie, nil
}
//	CreateMovie...
func (p postgreStore) CreateMovie(movie *Movie) (*Movie, error) {
	err := p.db.Insert(movie)
	return movie, err
}
//	DeleteMovie...
func (p postgreStore) DeleteMovie(id int) error {
	panic("implement me")
}
// GetMyMovie ...
func (p postgreStore) GetMyMovie(userId int) ([]*Movie, error) {
	var movies []*Movie
	err := p.db.Model(&movies).Where("user_id=?", userId).Select()
	if err != nil {
		return nil, err
	}
	return movies, nil
}

// DeleteMyMovie ...
func (p postgreStore) DeleteMyMovie(id int, user_id int) error {
	movie := &Movie{ID: id, UserID: user_id}
	err := p.db.Delete(movie)
	if err != nil {
		return err
	}

	return nil
}
