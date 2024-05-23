package postgres

import (
	"context"
	"durger/models"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) userRepo {
	return userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, user models.AddUser) (int64, error) {
	query := `
	INSERT INTO 
		users (id, first_name, last_name, username, is_premium, language_code)
	VALUES
		($1, $2, $3, $4, $5, $6)
	ON CONFLICT
		(id)
	DO UPDATE SET
		first_name = $2,
		last_name = $3,
		username = $4,
		is_premium = $5,
		language_code = $6,
		updated_at = NOW();`

	_, err := u.db.Exec(ctx, query, user.Id, user.FirstName, user.LastName, user.Username, user.IsPremium, user.LanguageCode)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return user.Id, nil
}

func (u *userRepo) GetAll(ctx context.Context) (models.AllUsers, error) {
	rows, err := u.db.Query(ctx, `SELECT * FROM users;`)
	resp := models.AllUsers{}
	if err != nil {
		log.Fatal("could't get all users' data", err)
		return resp, err
	}

	for rows.Next() {
		var user models.User

		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.IsPremium,
			&user.LanguageCode,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			log.Fatal("could't scan users' data")
		}
		resp.Users = append(resp.Users, user)
	}
	defer rows.Close()

	row := u.db.QueryRow(ctx, `SELECT COUNT(*) FROM users;`)
	err = row.Scan(&resp.Count)

	if err != nil {
		log.Fatal("error scan count: ", err)
		return resp, err
	}

	return resp, nil
}

func (u *userRepo) Delete(ctx context.Context, id int64) error {
	panic("unimplemented")
}

func (u *userRepo) GetUser(ctx context.Context, id int64) (models.User, error) {
	panic("unimplemented")
}