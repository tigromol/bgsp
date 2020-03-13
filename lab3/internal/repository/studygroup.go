package repository

import (
	"database/sql"
	"grig/internal/errors"
	"grig/internal/model"
)

func (r *Repository) GetGroupByID(id int) (model.StudyGroup, error) {
	var group model.StudyGroup
	err := r.db.Get(&group, `SELECT * FROM study_group WHERE study_group.id=$1`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return group, errors.ErrNotFound
	}
	return group, err
}

func (r *Repository) GetGroups() ([]model.StudyGroup, error) {
	var groups []model.StudyGroup
	err := r.db.Select(&groups, `SELECT * FROM study_group`)
	return groups, err
}

func (r *Repository) UpsertGroup(group model.StudyGroup) error {
	query := `
		INSERT INTO study_group(id, name) VALUES ($1, $2)
		ON CONFLICT (id) DO 
		UPDATE SET name=$2
	`
	_, err := r.db.Exec(query, group.ID, group.Name)
	return err
}

func (r *Repository) CreateGroup(group model.StudyGroup) error {
	_, err := r.db.Exec(`INSERT INTO study_group(name) VALUES ($1)`, group.Name)
	return err
}

func (r *Repository) DeleteGroup(id int) error {
	_, err := r.db.Exec(`DELETE FROM study_group WHERE study_group.id=$1`, id)
	return err
}
