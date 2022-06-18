package postgres

import (
	// "os/user"

	pb "github.com/baxromumarov/my-services/post-service/genproto"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type postRepo struct {
	db *sqlx.DB
}

//NewPostRepo ...
func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}

func (r *postRepo) CreatePost(post *pb.Post) (*pb.Post, error) {
	var res = pb.Post{}
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	crtime := time.Now()

	err = r.db.QueryRow(`INSERT INTO posts(id,user_id,name,createdat)
	VALUES($1,$2,$3,$4) RETURNING id,name`,
		id, post.UserId, post.Name, crtime).Scan(
		&res.Id,
		&res.Name,
		//&res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	for _, val := range post.Medias {
		Id, _ := uuid.NewV4()
		_, err := r.db.Exec(`INSERT INTO medias(id,post_id,type,link) VALUES($1,$2,$3,$4)`, Id,
			res.Id, val.Type, val.Link)
		if err != nil {
			return nil, err
		}
		res.Medias = append(res.Medias, val)

	}
	if err != nil {
		return nil, err

	}
	return &res, nil
}

func (r *postRepo) GetByIdPost(ID string) (*pb.Post, error) {
	var (
		rPost = pb.Post{}
	)

	err := r.db.QueryRow("SELECT user_id, name, createdat from posts WHERE id = $1", ID).Scan(
		&rPost.UserId,
		&rPost.Name,
		&rPost.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	var medias []*pb.Media
	rows, err := r.db.Query("SELECT id, type, link from medias where post_id = $1", ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var media pb.Media
		err := rows.Scan(
			&media.Id,
			&media.Type,
			&media.Link,
		)

		if err != nil {
			return nil, err
		}
		medias = append(medias, &media)
	}
	rPost.Medias = medias

	return &rPost, nil
}

func (r *postRepo) GetAllUserPosts(ID string) ([]*pb.Post, error) {
	var (
		posts []*pb.Post
	)

	rows, err := r.db.Query("SELECT id, user_id, name, createdat from posts WHERE user_id = $1", ID)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post pb.Post
		err := rows.Scan(
			&post.Id,
			&post.UserId,
			&post.Name,
			&post.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		var medias []*pb.Media
		rows, err := r.db.Query("SELECT id, type, link from medias ")

		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var media pb.Media
			err := rows.Scan(
				&media.Id,
				&media.Type,
				&media.Link,
			)
			if err != nil {
				return nil, err
			}

			post.Medias = append(medias, &media)
		}
		posts = append(posts, &post)
	}

	return posts, nil
}
