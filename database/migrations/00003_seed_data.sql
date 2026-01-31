
-- +goose Up
-- ROLES
INSERT INTO public.roles (id, name) VALUES
(1, 'admin'),
(2, 'teacher'),
(3, 'student')
ON CONFLICT (name) DO NOTHING;

-- STATUS
INSERT INTO public.status (id, name) VALUES
(1, 'active'),
(2, 'inactive')
ON CONFLICT DO NOTHING;

-- GENDERS
INSERT INTO public.genders (id, name) VALUES
(1, 'male'),
(2, 'female')
ON CONFLICT DO NOTHING;

-- SUBJECTS
INSERT INTO public.subjects (id, name) VALUES
(1, 'Mathematics'),
(2, 'Physics'),
(3, 'Programming'),
(4, 'History'),
(5, 'Physical Education')
ON CONFLICT (name) DO NOTHING;

-- GROUPS
INSERT INTO public."groups" (id, name, direction) VALUES
(1, 'ENG-251', 'Engineering'),
(2, 'CS-251', 'Computer Science'),
(3,'HUM-251','Human Science'),
(4,'ENG-252','Engineering')
ON CONFLICT (name) DO NOTHING;

-- USERS
INSERT INTO public.users (id, email, password_hash, created_at, role_id, updated_at, status_id) VALUES
(1, 'ruslan@student.com', '$2a$10$xmLXu4o6wKDPp9mn98K5ded7g4OrGi7Bl30ijdbY7oWRMMhWkg/Z6', '2026-01-25 15:30:11.077712+00', 3, '2026-01-25 15:30:11.077712+00', 1),
(2, 'dana@student.com', '$2a$10$CGHaXwlTTtmX4tLd2k87H.H2w89p8rhruZKWHRyvdeh5AcNufLsTK', '2026-01-25 15:30:27.60352+00', 3, '2026-01-25 15:30:27.60352+00', 1),
(3, 'arman@student.com', '$2a$10$JYPeIPIFVmFaRzm0h0v9nuz019B3Gu5Nst/boAn52Cxg8uS/zTmsK', '2026-01-25 15:30:36.797964+00', 3, '2026-01-25 15:30:36.797964+00', 1),
(4, 'alina@student.com', '$2a$10$2j1lq9BL75/C.pAjs/cxC.XGP3Cj5c9eARNNabyioWOFMHWFeVbCe', '2026-01-25 15:30:51.717324+00', 3, '2026-01-25 15:30:51.717324+00', 1),
(5, 'teacher1@university.com', '$2a$10$bmcjGs2JggCndvB4UuhNoOHFTXPAWylf9/5gigMf55nYitgOmEQxG', '2026-01-25 15:31:04.215848+00', 2, '2026-01-25 15:31:04.215848+00', 1),
(6, 'admin@university.com', '$2a$10$x1SfuFx7DZzAZxicUQJ6sO2ryJZAuGI3fXyzla/SlKHmMHrVyRL1e', '2026-01-25 15:31:22.953342+00', 1, '2026-01-25 15:31:22.953342+00', 1),
(7, 'aigerim@student.com', '$2a$10$dFS1hXTu3h9Sk2VyaQnheupnDhIP7mQ0Qx6ruXF5U/56Y6zERUeyG', '2026-01-25 15:43:35.713687+00', 3, '2026-01-25 15:43:35.713687+00', 1),
(8, 'teacher2@university.com', '$2a$10$eTMZpiOOWk48ufDJVOKe1OY5sSGNA8PJlDDEB6OGT1HO/YR7frqMO', '2026-01-25 15:49:42.669108+00', 2, '2026-01-25 15:49:42.669108+00', 1)
ON CONFLICT (email) DO NOTHING;

-- STUDENTS
INSERT INTO public.students (id, firstname,surname, birth_date, year_of_study, gender_id, group_id, user_id) VALUES
(1, 'Ruslan','Utepbergen', '2004-10-02', 2, 1, 2, 1),
(2, 'Dana','Kim', '2005-01-25', 1, 2, 3, 2),
(3, 'Arman','Armanbek', '2003-09-14', 3, 1, 2, 3),
(4, 'Alina','Alinova', '2004-12-01', 2, 2, 4, 4),
(5, 'Aigerim','Shakanova', '2005-03-12', 1, 2, 1, 7)
ON CONFLICT DO NOTHING;

-- TEACHERS
INSERT INTO public.teachers (id, full_name, department, user_id) VALUES
(1, 'Ben Tyler', 'Computer Science', 5),
(2, 'Denis Ktototam', 'Engineering', 8)
ON CONFLICT (user_id) DO NOTHING;

SELECT setval('users_id_seq', (SELECT COALESCE(MAX(id), 0) FROM users));
SELECT setval('students_id_seq', (SELECT COALESCE(MAX(id), 0) FROM students));
SELECT setval('teachers_id_seq', (SELECT COALESCE(MAX(id), 0) FROM teachers));
SELECT setval('roles_id_seq', (SELECT COALESCE(MAX(id), 0) FROM roles));
SELECT setval('status_id_seq', (SELECT COALESCE(MAX(id), 0) FROM status));
SELECT setval('genders_id_seq', (SELECT COALESCE(MAX(id), 0) FROM genders));
SELECT setval('subjects_id_seq', (SELECT COALESCE(MAX(id), 0) FROM subjects));
SELECT setval('groups_id_seq', (SELECT COALESCE(MAX(id), 0) FROM "groups"));

-- +goose Down 

-- Delete seed data in reverse order
DELETE FROM public."groups" WHERE id IN (1, 2, 3, 4);
DELETE FROM public.subjects WHERE id IN (1, 2, 3, 4, 5);
DELETE FROM public.genders WHERE id IN (1, 2);
DELETE FROM public.status WHERE id IN (1, 2);
DELETE FROM public.roles WHERE id IN (1, 2, 3);
