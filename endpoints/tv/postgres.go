package tv

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

// Postgre store tv
func PostgreTv(config PostgreConfig) TvDb {
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

//	GetTv...
func (p postgreStore) GetTv(id int) (*TV, error) {
	tv := &TV{ID: id}
	err := p.db.Select(tv)

	if err != nil {
		return nil, err
	}
	return tv, nil
}
//	CreateTv...
func (p postgreStore) CreateTv(tv *TV) (*TV, error) {
	err := p.db.Insert(tv)
	return tv, err
}
//	DeleteTv...
func (p postgreStore) DeleteTv(id int) error {
	panic("implement me")
}
//	GetMyTv...
func (p postgreStore) GetMyTv(user_id int) ([]*TV, error) {
	var tvs []*TV
	err := p.db.Model(&tvs).Where("user_id=?", user_id).Select()
	if err != nil {
		return nil, err
	}
	return tvs, nil
}
// DeleteMyTv...
func (p postgreStore) DeleteMyTv(id int, user_id int) error {
	panic("implement me")
}
