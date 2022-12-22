package postgres

import (
	"errors"
	"fmt"
	"time"

	pb "github.com/baxromumarov/my-services/user-service/genproto"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	// "database/sql"
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
	_, err := r.db.Exec(`CREATE TABLE IF NOT EXISTS users (
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

	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error while generating uuid")
		return nil, err
	}
	crtime := time.Now()
	err = r.db.QueryRow(`INSERT INTO users (id, first_name, last_name, email, bio, phone_numbers, typeid, status, createdat,user_name,password, email_code)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11,$12) Returning id,first_name,last_name,email,
		bio`, id, user.FirstName, user.LastName, user.Email,
		user.Bio, pq.Array(user.PhoneNumbers),
		user.TypeId, user.Status, crtime, user.UserName,user.Password,user.EmailCode).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.Bio,
	)
	if err != nil {
		return &pb.User{}, err
	}
	return &res, nil
}

//insert into addresses
func (r *userRepo) InsertAd(ad *pb.Address) (*pb.Address, error) {
	var add pb.Address
	idd, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error while generating uuid")
		return nil, err
	}
	err = r.db.QueryRow(`INSERT INTO addresses (id, city, country,
		district,postal_code) VALUES ($1, $2, $3, $4, $5) Returning id,city,country, district,postal_code`, idd, ad.City,
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

func (r *userRepo) GetUserList(limit, page int64) ([]*pb.User, int64, error) {
	var (
		users []*pb.User
		count int64
	)
	offset := (page - 1) * limit

	query := `SELECT id, first_name, last_name,email,bio,phonenumbers,
	status,createdat FROM users ORDER BY first_name OFFSET $1 LIMIT $2`

	rows, err := r.db.Query(query, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		var user pb.User
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Bio,
			pq.Array(&user.PhoneNumbers),
			&user.Status,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, &user)
	}
	countQuery := `SELECT count(*) FROM users`
	err = r.db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	return users, count, nil
}

func (r *userRepo) UserList(limit, page int64) ([]*pb.User, int64, error) {

	var users []*pb.User

	offset := (page - 1) * limit

	listQuery := `SELECT id, first_name, last_name, bio, email, status, created_at, phone_number FROM users LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(listQuery, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var user pb.User
		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Bio,
			&user.Email,
			&user.Status,
			&user.CreatedAt,
			pq.Array(&user.PhoneNumbers),
		)

		if err != nil {
			return nil, 0, err
		}

		var address pb.Address

		addressQuery := `SELECT city, country, district, postal_code FROM adress WHERE user_id = $1`

		rows1, err := r.db.Query(addressQuery, user.Id)

		if err != nil {
			return nil, 0, err
		}

		for rows1.Next() {
			err := rows1.Scan(
				&address.City,
				&address.Country,
				&address.District,
				&address.PostalCode,
			)

			if err != nil {
				return nil, 0, err
			}
		}
		user.Addresses = append(user.Addresses, &address)
		users = append(users, &user)
	}

	var count int64
	err = r.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)

	if err != nil {
		return nil, 0, err
	}
	return users, count, nil
}

func (r *userRepo) CheckFeild(field, value string) (bool, error) {
	var cle int
	if field == "username" {
		err := r.db.QueryRow("SELECT COUNT(1) FROM users WHERE user_name = $1 AND deletedat = NULL", value).Scan(&cle)
		if err != nil {
			return false, err
		}
	} else if field == "email" {
		err := r.db.QueryRow("SELECT COUNT(1) FROM users WHERE user_name = $1 AND deletedat = NULL", value).Scan(&cle)
		if err != nil {
			return false, err
		}
	} else {
		err := errors.New("ERROR IN CheckField")
		return false, err
	}

	if cle == 0 {
		return false, nil
	}

	return true, nil
}
