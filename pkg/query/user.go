package query

import (
	"context"
	"database/sql"
	"strings"
)

// GetAllUsers returns all usernames in the database and their corresponding userIDs
func GetAllUsers(db *sql.DB) ([]*User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.UserName, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if len(users) == 0 {
		return nil, sql.ErrNoRows
	}

	return users, nil
}

// GetUserByID gets a user that matches the specified user id
func GetUser(id int, db *sql.DB, ctx context.Context) (*User, error) {
	var user User
	userRow := db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", id)
	if err := userRow.Scan(&user.ID, &user.UserName, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser deletes the user
func DeleteUser(id int, db *sql.DB, ctx context.Context) error {
	user, err := GetUser(id, db, ctx)
	if err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, "DELETE FROM users WHERE id = ?", user.ID); err != nil {
		return err
	}

	return nil
}

// CreateUser creates a user entry in the users table
func CreateUser(userName string, db *sql.DB, ctx context.Context) error {
	userName = strings.ToLower(userName)
	email := userName + "@example.com"
	if _, err := db.ExecContext(ctx, "INSERT INTO users (username, email) VALUES (?, ?)", userName, email); err != nil {
		return err
	}

	return nil
}
