package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/sichangSun/course-enrollment-manager/domain/model"
	"github.com/sichangSun/course-enrollment-manager/domain/repository"
)

// studentRepository ...
type StudentRepository struct {
	db *sqlx.DB
}

// NewCourseRepository ...
func NewStudentRepository(db *sqlx.DB) *StudentRepository {
	return &StudentRepository{db}
}

// GetStudentByUserEmail
func (s *StudentRepository) GetStudentByUserEmail(con context.Context, email string) (*model.Student, error) {
	query := `SELECT id, name, password,email
	          FROM students WHERE email = ?;`

	var student model.Student
	err := s.db.Get(&student, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrStudentNotFound
		}
		return nil, err
	}
	return &student, nil
}
