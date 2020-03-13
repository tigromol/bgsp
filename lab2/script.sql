CREATE TABLE IF NOT EXISTS study_group(
    id integer PRIMARY KEY,
    name VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS student(
    id integer PRIMARY KEY,
    name VARCHAR(50),
    surname VARCHAR(50),
    second_name VARCHAR(50),
    study_group_id integer REFERENCES study_group(id)
);

CREATE TABLE IF NOT EXISTS subject(
  id integer PRIMARY KEY,
  name VARCHAR(50),
  short_name VARCHAR(50)
);


CREATE TABLE IF NOT EXISTS exam_type(
    id integer PRIMARY KEY,
    type VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS study_plan(
     id integer PRIMARY KEY,
     subject_id integer REFERENCES subject(id),
     exam_type_id integer REFERENCES exam_type(id)
);

CREATE TABLE IF NOT EXISTS mark(
    id integer PRIMARY KEY,
    name VARCHAR(50),
    value VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS journal(
    id integer PRIMARY KEY,
    student_id integer REFERENCES student(id),
    study_plan_id integer REFERENCES study_plan(id),
    in_time boolean,
    count integer,
    mark_id integer REFERENCES mark(id)
);

INSERT INTO subject VALUES
(1, 'Проектирование информационных систем', 'ПрИС'),
(2, 'Системы искусственного интеллекта', 'СИИ'),
(3, 'Программная инженерия', 'ПИ'),
(4, 'Национальная система информационной безопасности', 'НСИБ'),
(5, 'Системный анализ', 'СисАнал'),
(6, 'Распределенные базы данных', 'РБД'),
(7, 'Системное программное обеспечение', 'СПО');

INSERT INTO exam_type VALUES
(1, 'Экзамен'),
(2, 'Зачет'),
(3, 'Зачет с оценкой'),
(4, 'Курсовая');

INSERT INTO study_plan VALUES
(1, 1, 1),
(2, 1, 4),
(3, 2, 1),
(4, 3, 1),
(5, 4, 2),
(6, 5, 1),
(7, 6, 2),
(8, 7, 1);

INSERT INTO mark VALUES
(1, 'Отлично', 5),
(2, 'Хорошо', 4),
(3, 'Удовлетворительно', 3),
(4, 'Неудовлетворительно', 2),
(5, 'Зачет', 'з'),
(6, 'Незачет', 'н'),
(7, 'Неявка', '');

INSERT INTO study_group VALUES
(1, 'икбо-01-16'),
(2, 'икбо-02-16'),
(3, 'икбо-03-16'),
(4, 'икбо-04-16'),
(5, 'икбо-05-16'),
(6, 'икбо-06-16'),
(7, 'икбо-07-16');