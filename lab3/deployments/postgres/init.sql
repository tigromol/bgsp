CREATE TABLE IF NOT EXISTS study_group(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS student(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    surname VARCHAR(50),
    second_name VARCHAR(50),
    study_group_id integer REFERENCES study_group(id)
);

CREATE TABLE IF NOT EXISTS subject(
  id SERIAL PRIMARY KEY,
  name VARCHAR(50),
  short_name VARCHAR(50)
);


CREATE TABLE IF NOT EXISTS exam_type(
    id SERIAL PRIMARY KEY,
    type VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS study_plan(
     id SERIAL PRIMARY KEY,
     subject_id integer REFERENCES subject(id),
     exam_type_id integer REFERENCES exam_type(id)
);

CREATE TABLE IF NOT EXISTS mark(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    value VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS journal(
    id SERIAL PRIMARY KEY,
    student_id integer REFERENCES student(id),
    study_plan_id integer REFERENCES study_plan(id),
    in_time boolean,
    count integer,
    mark_id integer REFERENCES mark(id)
);

INSERT INTO subject(name, short_name) VALUES
('Проектирование информационных систем', 'ПрИС'),
('Системы искусственного интеллекта', 'СИИ'),
('Программная инженерия', 'ПИ'),
('Национальная система информационной безопасности', 'НСИБ'),
('Системный анализ', 'СисАнал'),
('Распределенные базы данных', 'РБД'),
('Системное программное обеспечение', 'СПО');

INSERT INTO exam_type(type) VALUES
('Экзамен'),
('Зачет'),
('Зачет с оценкой'),
('Курсовая');

INSERT INTO study_plan(subject_id, exam_type_id) VALUES
(1, 1),
(1, 4),
(2, 1),
(3, 1),
(4, 2),
(5, 1),
(6, 2),
(7, 1);

INSERT INTO mark(name, value) VALUES
('Отлично', 5),
('Хорошо', 4),
('Удовлетворительно', 3),
('Неудовлетворительно', 2),
('Зачет', 'з'),
('Незачет', 'н'),
('Неявка', '');

INSERT INTO study_group(name) VALUES
('икбо-01-16'),
('икбо-02-16'),
('икбо-03-16');

INSERT INTO student(name, surname, second_name, study_group_id) VALUES
('Иван', 'Иванов', 'Иванович', 1),
('Андрей', 'Иванов', 'Иванович', 2),
('Дмитрий', 'Иванов', 'Иванович', 3),
('Анатолий', 'Иванов', 'Иванович', 1),
('Борис', 'Иванов', 'Иванович', 2);

INSERT INTO journal(student_id, study_plan_id, in_time, count, mark_id) VALUES
(1, 1, true, 1, 1),
(2, 2, true, 2, 2),
(3, 3, false, 3, 3),
(4, 4, true, 4, 4);
