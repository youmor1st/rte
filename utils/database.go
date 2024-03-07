// В файле utils/utils.go

package utils

import (
	"context"
	"finalgo/config"
	"fmt"
	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func InitDB() error {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		config.DBUser, config.DBPassword, config.DBName, config.DBHost, config.DBPort)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}

	db = conn
	return nil
}

func GetDB() *pgx.Conn {
	return db
}

func GetOrCreateClass(className string) (int, error) {
	var classID int
	err := db.QueryRow(context.Background(), "SELECT id FROM classes WHERE class_name = $1", className).Scan(&classID)
	if err == nil {
		return classID, nil
	}

	err = db.QueryRow(context.Background(), "INSERT INTO classes (class_name) VALUES ($1) RETURNING id", className).Scan(&classID)
	if err != nil {
		return 0, err
	}

	return classID, nil
}

func CreateUser(user User) error {
	_, err := db.Exec(context.Background(),
		"INSERT INTO users (role, username, password, f_name, s_name, class_id, points) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		user.Role, user.Username, user.Password, user.FName, user.SName, user.ClassID, user.Points)
	if err != nil {
		return err
	}

	return nil
}

func ValidateUser(username, password string) (bool, error) {
	var count int
	err := db.QueryRow(context.Background(), "SELECT COUNT(*) FROM users WHERE username = $1 AND password = $2", username, password).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
func GetStudents() ([]User, error) {
	rows, err := db.Query(context.Background(), "SELECT id, username FROM users WHERE role = 'Student'")
	if err != nil {
		return nil, err

	}
	defer rows.Close()

	var students []User
	for rows.Next() {
		var student User
		if err := rows.Scan(&student.ID, &student.Username); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}
func DeleteUser(userID int) error {
	_, err := db.Exec(context.Background(), "DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		return err
	}

	return nil
}
func GetClasses() ([]Class, error) {
	rows, err := db.Query(context.Background(), "SELECT id, class_name FROM classes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []Class
	for rows.Next() {
		var class Class
		if err := rows.Scan(&class.ID, &class.ClassName); err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	return classes, nil
}
func GetStudentsByClassID(classID int) ([]User, error) {
	rows, err := db.Query(context.Background(), "SELECT id, role, username, f_name, s_name, class_id, points FROM users WHERE class_id = $1", classID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []User
	for rows.Next() {
		var student User
		if err := rows.Scan(&student.ID, &student.Role, &student.Username, &student.FName, &student.SName, &student.ClassID, &student.Points); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}
