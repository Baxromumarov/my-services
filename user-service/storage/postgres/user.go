package postgres

import (
	pb "github.com/baxromumarov/my-services/user-service/genproto"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

//create table users
func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
	_, err := r.db.Query(`CREATE TABLE IF NOT EXISTS users (
		id varchar(255),
		first_name varchar(255),
		last_name varchar(255),
		email varchar(255),
		bio varchar(255),
		phoneNumbers text[],
		address_id varchar(255),
		typeId varchar(255),
		Status varchar(255),
		createdAt timestamp,
		updatedAt varchar(255),
		deletedAt varchar(255),
		FOREIGN KEY (address_id) REFERENCES addresses(id) )`)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *userRepo) CreateAd(ad *pb.Address) (*pb.Address, error) {
	var res = pb.Address{}

	_, err := r.db.Exec(`CREATE TABLE IF NOT EXISTS addresses (
		id varchar(255) Primary Key,
		city varchar(255),
		country varchar(255),
		district varchar(255),
		user_id varchar (255) NOT NULL,
    	postal_code varchar(255))`)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

//insert into users
func (r *userRepo) Insert(user *pb.User) (*pb.User, error) {
	var res = pb.User{}

	err := r.db.QueryRow(`INSERT INTO users (id, first_name, last_name, email, bio, phoneNumbers, typeId, Status, createdAt, updatedAt, deletedAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) Returning id,first_name,last_name`, user.Id, user.FirstName, user.LastName, user.Email,
		user.Bio, pq.Array(user.PhoneNumbers),
		user.TypeId, user.Status, user.CreatedAt, user.UpdatedAt, user.DeletedAt).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
	)
	if err != nil {
		return &pb.User{}, err
	}
	return &res, nil
}

//insert into addresses
func (r *userRepo) InsertAd(ad *pb.Address) (*pb.Address, error) {
	var add pb.Address

	err := r.db.QueryRow(`INSERT INTO addresses (id, city, country,
		district,postal_code) VALUES ($1, $2, $3, $4, $5) Returning id,city,country, district,postal_code`, ad.Id, ad.City,
		ad.Country, ad.District, ad.PostalCode).Scan(
		&add.Id,
		&add.City,
		&add.Country,
		&add.District,
		&add.PostalCode,
	)
	if err != nil {
		return nil, err
	}
	return &add, nil

}

func (r *userRepo) Delete(id *pb.ById) (*pb.UserInfo, error) {
	var res pb.UserInfo

	_, err := r.db.Query(`DELETE FROM users where id = $1`, id.Id)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *userRepo) GetById(id *pb.ById) (*pb.User, error) {
	var res pb.User

	err := r.db.QueryRow(`SELECT first_name, last_name
	FROM users where id = $1  `, id.Id).Scan(
		&res.FirstName,
		&res.LastName,
	)

	if err != nil {
		return &pb.User{}, err
	}
	return &res, nil
}

func (r *userRepo) GetAll() ([]*pb.User, error) {
	var ruser1 []*pb.User

	getByIdQuery := `SELECT id, first_name, last_name, email, bio, status, createdat, phonenumbers FROM users`
	rowss, err := r.db.Query(getByIdQuery)

	if err != nil {
		return nil, err
	}

	for rowss.Next() {
		var ruser pb.User
		err = rowss.Scan(
			&ruser.Id,
			&ruser.FirstName,
			&ruser.LastName,
			&ruser.Email,
			&ruser.Bio,
			&ruser.Status,
			&ruser.CreatedAt,
			pq.Array(&ruser.PhoneNumbers),
		)
		if err != nil {
			return nil, err
		}

		getByIdAdressQuery := `SELECT city, country, district, postal_code FROM addresses`
		rows, err := r.db.Query(getByIdAdressQuery)

		if err != nil {
			return nil, err
		}

		var tempUser pb.User
		for rows.Next() {
			var adressById pb.Address
			err = rows.Scan(
				&adressById.City,
				&adressById.Country,
				&adressById.District,
				&adressById.PostalCode,
			)

			if err != nil {
				return nil, err
			}

			tempUser.Addresses = append(tempUser.Addresses, &adressById)
		}
		ruser.Addresses = tempUser.Addresses
		ruser1 = append(ruser1, &ruser)
	}

	return ruser1, nil
}


