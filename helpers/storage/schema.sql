CREATE TABLE public.roles (
	id serial4 NOT NULL,
	"name" varchar(20) NOT NULL,
	CONSTRAINT roles_name_key UNIQUE (name),
	CONSTRAINT roles_pkey PRIMARY KEY (id)
);

-- public.genders definition

-- Drop table

-- DROP TABLE public.genders;

CREATE TABLE public.genders (
	id serial4 NOT NULL,
	"name" varchar(20) NOT NULL,
	CONSTRAINT genders_pkey PRIMARY KEY (id)
);

-- public.subjects definition

-- Drop table

-- DROP TABLE public.subjects;

CREATE TABLE public.subjects (
	id serial4 NOT NULL,
	"name" varchar(30) NULL,
	CONSTRAINT subjects_name_key UNIQUE (name),
	CONSTRAINT subjects_pkey PRIMARY KEY (id)
);
-- public.status definition

-- Drop table

-- DROP TABLE public.status;

CREATE TABLE public.status (
	id serial4 NOT NULL,
	"name" varchar(30) NOT NULL,
	CONSTRAINT status_pkey PRIMARY KEY (id)
);

-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id serial4 NOT NULL,
	email varchar(255) NOT NULL,
	password_hash varchar(255) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	role_id int4 NOT NULL,
	updated_at timestamptz DEFAULT now() NULL,
	status_id int4 NULL,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);


-- public.users foreign keys

ALTER TABLE public.users ADD CONSTRAINT fk_users_role FOREIGN KEY (role_id) REFERENCES public.roles(id);
ALTER TABLE public.users ADD CONSTRAINT fk_users_status FOREIGN KEY (status_id) REFERENCES public.status(id);

-- public."groups" definition

-- Drop table

-- DROP TABLE public."groups";

CREATE TABLE public."groups" (
	id serial4 NOT NULL,
	"name" varchar(20) NOT NULL,
	direction varchar(50) NOT NULL,
	CONSTRAINT groups_name_key UNIQUE (name),
	CONSTRAINT groups_pkey PRIMARY KEY (id)
);




-- public.students definition

-- Drop table

-- DROP TABLE public.students;

CREATE TABLE public.students (
	id serial4 NOT NULL,
	"name" varchar(100) NOT NULL,
	birth_date date NOT NULL,
	year_of_study int2 NOT NULL,
	gender_id int4 NOT NULL,
	group_id int4 NOT NULL,
	user_id int4 NULL,
	CONSTRAINT students_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_students_group_id ON public.students USING btree (group_id);
CREATE INDEX idx_students_user_id ON public.students USING btree (user_id);


-- public.students foreign keys

ALTER TABLE public.students ADD CONSTRAINT fk_student_users FOREIGN KEY (user_id) REFERENCES public.users(id);
ALTER TABLE public.students ADD CONSTRAINT students_gender_id_fkey FOREIGN KEY (gender_id) REFERENCES public.genders(id);
ALTER TABLE public.students ADD CONSTRAINT students_group_id_fkey FOREIGN KEY (group_id) REFERENCES public."groups"(id);

-- public.teachers definition

-- Drop table

-- DROP TABLE public.teachers;

CREATE TABLE public.teachers (
	id serial4 NOT NULL,
	full_name varchar(100) NOT NULL,
	department varchar(100) NULL,
	user_id int4 NOT NULL,
	CONSTRAINT teachers_pkey PRIMARY KEY (id),
	CONSTRAINT teachers_user_id_key UNIQUE (user_id)
);


-- public.teachers foreign keys

ALTER TABLE public.teachers ADD CONSTRAINT teachers_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


-- public.class_schedule definition

-- Drop table

-- DROP TABLE public.class_schedule;

CREATE TABLE public.class_schedule (
	id serial4 NOT NULL,
	group_id int4 NULL,
	day_of_week int4 NOT NULL,
	starts_at time NOT NULL,
	ends_at time NOT NULL,
	subject_id int4 NULL,
	teacher_id int4 NULL,
	CONSTRAINT class_schedule_pkey PRIMARY KEY (id)
);


-- public.class_schedule foreign keys

ALTER TABLE public.class_schedule ADD CONSTRAINT class_schedule_group_id_fkey FOREIGN KEY (group_id) REFERENCES public."groups"(id);
ALTER TABLE public.class_schedule ADD CONSTRAINT class_schedule_subject_id_fkey FOREIGN KEY (subject_id) REFERENCES public.subjects(id);
ALTER TABLE public.class_schedule ADD CONSTRAINT fk_teacher_id FOREIGN KEY (teacher_id) REFERENCES public.teachers(id);


-- public.attendance definition

-- Drop table

-- DROP TABLE public.attendance;

CREATE TABLE public.attendance (
	id serial4 NOT NULL,
	student_id int4 NOT NULL,
	subject_id int4 NOT NULL,
	visit_day date NOT NULL,
	visited bool NOT NULL,
	CONSTRAINT attendance_pkey PRIMARY KEY (id),
	CONSTRAINT attendance_student_id_subject_id_visit_day_key UNIQUE (student_id, subject_id, visit_day)
);
CREATE INDEX idx_attendance_subject ON public.attendance USING btree (subject_id);


-- public.attendance foreign keys

ALTER TABLE public.attendance ADD CONSTRAINT attendance_student_id_fkey FOREIGN KEY (student_id) REFERENCES public.students(id);
ALTER TABLE public.attendance ADD CONSTRAINT attendance_subject_id_fkey FOREIGN KEY (subject_id) REFERENCES public.subjects(id);
