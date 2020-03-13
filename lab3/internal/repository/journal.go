package repository

import "grig/internal/model"

func (r *Repository) GetJournals() ([]model.Journal, error) {
	var journals []model.Journal
	query := `
		SELECT 
			journal.id,
		    student.id,
		    student.name,
		    student.surname,
		    student.second_name,
		    study_group.id,
		    study_group.name,
		    study_plan.id,
		    subject.id,
		    subject.name,
		    subject.short_name,
		    exam_type.id,
		    exam_type.type,
		    journal.in_time,
		    journal.count,
		    mark.id,
		    mark.name,
		    mark.value
		FROM journal
		JOIN student ON journal.student_id=student.id
		JOIN study_plan ON journal.study_plan_id=study_plan.id
		JOIN mark ON journal.mark_id=mark.id
		JOIN study_group ON student.study_group_id=study_group.id
		JOIN subject ON study_plan.subject_id=subject.id
		JOIN exam_type ON study_plan.exam_type_id=exam_type.id
	`
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var journal model.Journal
		err := rows.Scan(
			&journal.ID,
			&journal.Student.ID,
			&journal.Student.Name,
			&journal.Student.Surname,
			&journal.Student.SecondName,
			&journal.StudyGroup.ID,
			&journal.StudyGroup.Name,
			&journal.StudyPlan.ID,
			&journal.Subject.ID,
			&journal.Subject.Name,
			&journal.Subject.ShortName,
			&journal.ExamType.ID,
			&journal.ExamType.Type,
			&journal.InTime,
			&journal.Count,
			&journal.Mark.ID,
			&journal.Mark.Name,
			&journal.Mark.Value,
		)
		if err != nil {
			return nil, err
		}
		journals = append(journals, journal)
	}
	return journals, nil
}

func (r *Repository) GetJournalByStudent(id int) ([]model.Journal, error) {
	var journals []model.Journal
	query := `
		SELECT 
			journal.id,
		    student.id,
		    student.name,
		    student.surname,
		    student.second_name,
		    study_group.id,
		    study_group.name,
		    study_plan.id,
		    subject.id,
		    subject.name,
		    subject.short_name,
		    exam_type.id,
		    exam_type.type,
		    journal.in_time,
		    journal.count,
		    mark.id,
		    mark.name,
		    mark.value
		FROM journal
		JOIN student ON journal.student_id=student.id AND student.id=$1
		JOIN study_plan ON journal.study_plan_id=study_plan.id
		JOIN mark ON journal.mark_id=mark.id
		JOIN study_group ON student.study_group_id=study_group.id
		JOIN subject ON study_plan.subject_id=subject.id
		JOIN exam_type ON study_plan.exam_type_id=exam_type.id
	`
	rows, err := r.db.Queryx(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var journal model.Journal
		err := rows.Scan(
			&journal.ID,
			&journal.Student.ID,
			&journal.Student.Name,
			&journal.Student.Surname,
			&journal.Student.SecondName,
			&journal.StudyGroup.ID,
			&journal.StudyGroup.Name,
			&journal.StudyPlan.ID,
			&journal.Subject.ID,
			&journal.Subject.Name,
			&journal.Subject.ShortName,
			&journal.ExamType.ID,
			&journal.ExamType.Type,
			&journal.InTime,
			&journal.Count,
			&journal.Mark.ID,
			&journal.Mark.Name,
			&journal.Mark.Value,
		)
		if err != nil {
			return nil, err
		}
		journals = append(journals, journal)
	}
	return journals, nil
}

func (r *Repository) GetJournalByStudyGroup(id int) ([]model.Journal, error) {
	var journals []model.Journal
	query := `
		SELECT 
			journal.id,
		    student.id,
		    student.name,
		    student.surname,
		    student.second_name,
		    study_group.id,
		    study_group.name,
		    study_plan.id,
		    subject.id,
		    subject.name,
		    subject.short_name,
		    exam_type.id,
		    exam_type.type,
		    journal.in_time,
		    journal.count,
		    mark.id,
		    mark.name,
		    mark.value
		FROM journal
		JOIN student ON journal.student_id=student.id
		JOIN study_plan ON journal.study_plan_id=study_plan.id
		JOIN mark ON journal.mark_id=mark.id
		JOIN study_group ON student.study_group_id=study_group.id AND study_group.id=$1
		JOIN subject ON study_plan.subject_id=subject.id
		JOIN exam_type ON study_plan.exam_type_id=exam_type.id
	`
	rows, err := r.db.Queryx(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var journal model.Journal
		err := rows.Scan(
			&journal.ID,
			&journal.Student.ID,
			&journal.Student.Name,
			&journal.Student.Surname,
			&journal.Student.SecondName,
			&journal.StudyGroup.ID,
			&journal.StudyGroup.Name,
			&journal.StudyPlan.ID,
			&journal.Subject.ID,
			&journal.Subject.Name,
			&journal.Subject.ShortName,
			&journal.ExamType.ID,
			&journal.ExamType.Type,
			&journal.InTime,
			&journal.Count,
			&journal.Mark.ID,
			&journal.Mark.Name,
			&journal.Mark.Value,
		)
		if err != nil {
			return nil, err
		}
		journals = append(journals, journal)
	}
	return journals, nil
}

func (r *Repository) UpsertJournal(journal model.Journal) error {
	query := `
		INSERT INTO journal VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO 
		UPDATE SET student_id=$2, study_plan_id=$3, in_time=$4, count=$5, mark_id=$6
	`
	_, err := r.db.Exec(query, journal.ID, journal.Student.ID, journal.StudyPlan.ID, journal.InTime, journal.Count, journal.Mark.ID)
	return err
}
