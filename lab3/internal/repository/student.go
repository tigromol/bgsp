package repository

import (
	"database/sql"
	"grig/internal/errors"
	"grig/internal/model"
	"log"
)

func (r *Repository) CreateStudent(student model.Student) error {
	query := `INSERT INTO student(name, surname, second_name, study_group_id) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, student.Name, student.Surname, student.SecondName, student.StudyGroup.ID)
	return err
}

func (r *Repository) GetStudentByID(id int) (model.Student, error) {
	var student model.Student
	row := r.db.QueryRowx(`
		SELECT 
			student.id,
			student.name,
			student.surname,
			student.second_name,
			student.study_group_id,
			study_group.name
		FROM student 
		JOIN study_group ON student.study_group_id=study_group.id
		WHERE student.id=$1
	`, id)
	err := row.Scan(&student.ID, &student.Name, &student.Surname, &student.SecondName, &student.StudyGroup.ID, &student.StudyGroup.Name)
	if errors.Is(err, sql.ErrNoRows) {
		return student, errors.ErrNotFound
	}
	return student, err
}

func (r *Repository) GetStudentsByGroup(id int) ([]model.Student, error) {
	var students []model.Student
	rows, err := r.db.Queryx(`
		select
			student.id,
			student.name, 
			student.surname, 
			student.second_name,
			student.study_group_id,
			study_group.name 
		from student 
		join study_group on student.study_group_id=study_group.id
		where student.study_group_id=$1;
	`, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var student model.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.SecondName, &student.StudyGroup.ID, &student.StudyGroup.Name)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, err
}

func (r *Repository) GetStudents() ([]model.Student, error) {
	var students []model.Student
	rows, err := r.db.Queryx(`
		select
			student.id,
			student.name, 
			student.surname, 
			student.second_name,
			student.study_group_id,
			study_group.name 
		from student 
		join study_group on student.study_group_id=study_group.id;
	`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var student model.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.SecondName, &student.StudyGroup.ID, &student.StudyGroup.Name)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, err
}

func (r *Repository) DeleteStudent(id int) error {
	log.Println(id)
	_, err := r.db.Exec(`DELETE FROM student WHERE student.id=$1`, id)
	return err
}

func (r *Repository) UpsertStudent(student model.Student) error {
	query := `
		INSERT INTO student(id, name, surname, second_name, study_group_id) VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO 
		UPDATE SET name=$2, surname=$3, second_name=$4, study_group_id=$5
	`
	_, err := r.db.Exec(query, student.ID, student.Name, student.Surname, student.SecondName, student.StudyGroup.ID)
	return err
}
