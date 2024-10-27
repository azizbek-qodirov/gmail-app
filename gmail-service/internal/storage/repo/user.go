package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pb "gmail-service/internal/pkg/genproto"
	"time"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) ChangeUserPFP(ctx context.Context, req *pb.UserChangePFPReq) (*pb.Void, error) {
	if req.PhotoUrl == "" {
		return nil, errors.New("photo url is empty")
	}
	query := "UPDATE users SET pfp_url = $1 WHERE email = $2"
	_, err := r.db.Exec(query, req.PhotoUrl, req.Email)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *UserRepo) ChangeUserPassword(ctx context.Context, req *pb.UserRecoverPasswordReq) (*pb.Void, error) {
	query := "UPDATE users SET password = $1 WHERE email = $2"
	_, err := r.db.Exec(query, req.NewPassword, req.Email)
	return nil, err
}

func (m *UserRepo) ConfirmUser(ctx context.Context, req *pb.ByEmail) (*pb.Void, error) {
	query := "UPDATE users SET is_confirmed = true, confirmed_at = $1 WHERE email = $2"
	_, err := m.db.Exec(query, time.Now(), req.Email)
	return nil, err
}

func (r *UserRepo) CreateUser(ctx context.Context, req *pb.UserCreateReq) (*pb.Void, error) {
	query := `
		INSERT INTO users (first_name, last_name, dob, gender, email, password, pfp_url, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(
		query,
		req.FirstName,
		req.LastName,
		req.Dob,
		req.Gender,
		req.Email,
		req.Password,
		req.PfpUrl,
		time.Now(),
	)

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := "UPDATE users SET deleted_at = $1 WHERE id = $2"
	_, err := r.db.Exec(query, time.Now().Unix(), req.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, req *pb.ByEmail) (*pb.UserGetRes, error) {
	var user pb.UserGetRes
	query := `
		SELECT id, first_name, last_name, dob, email, gender, pfp_url
		FROM users
		WHERE email = $1 and deleted_at = 0
	`
	row := r.db.QueryRowContext(ctx, query, req.Email)

	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Dob,
		&user.Email,
		&user.Gender,
		&user.PfpUrl,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByID(ctx context.Context, req *pb.ByID) (*pb.UserGetRes, error) {
	var user pb.UserGetRes
	query := `
		SELECT id, first_name, last_name, dob, email, gender, pfp_url
		FROM users
		WHERE id = $1 and deleted_at = 0
	`
	row := r.db.QueryRowContext(ctx, query, req.Id)
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Dob,
		&user.Email,
		&user.Gender,
		&user.PfpUrl,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserSecurityByEmail(ctx context.Context, req *pb.ByEmail) (*pb.UserGetSecurityRes, error) {
	var user pb.UserGetSecurityRes
	query := `
		SELECT id, email, password, is_confirmed
		FROM users
		WHERE email = $1 and deleted_at = 0
	`
	row := r.db.QueryRowContext(ctx, query, req.Email)
	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.IsConfirmed,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) IsEmailExists(ctx context.Context, req *pb.ByEmail) (*pb.UserEmailCheckRes, error) {
	query := "SELECT COUNT(*) FROM users WHERE email = $1"
	var count int
	err := r.db.QueryRow(query, req.Email).Scan(&count)
	if err != nil {
		return nil, errors.New("server error: " + err.Error())
	}
	if count > 0 {
		return &pb.UserEmailCheckRes{Exists: true}, nil
	}
	return &pb.UserEmailCheckRes{Exists: false}, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, req *pb.UserUpdateReq) (*pb.Void, error) {
	query := `UPDATE users SET`
	var params []interface{}
	paramCounter := 1

	if req.Body.FirstName != "" {
		query += fmt.Sprintf(" first_name = $%d,", paramCounter)
		params = append(params, req.Body.FirstName)
		paramCounter++
	}

	if req.Body.LastName != "" {
		query += fmt.Sprintf(" last_name = $%d,", paramCounter)
		params = append(params, req.Body.LastName)
		paramCounter++
	}

	if req.Body.Dob != "" {
		query += fmt.Sprintf(" dob = $%d,", paramCounter)
		params = append(params, req.Body.Dob)
		paramCounter++
	}

	if req.Body.Gender != "" {
		query += fmt.Sprintf(" gender = $%d,", paramCounter)
		params = append(params, req.Body.Gender)
		paramCounter++
	}

	if paramCounter > 1 {
		query = query[:len(query)-1]
	} else {
		return nil, nil
	}

	query += fmt.Sprintf(" WHERE id = $%d", paramCounter)
	params = append(params, req.Id)

	res, err := r.db.ExecContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, sql.ErrNoRows
	}

	return nil, nil
}
