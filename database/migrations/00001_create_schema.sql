-- +goose Up

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

-- +goose Down 

DROP TABLE IF EXISTS public.students CASCADE;
DROP TABLE IF EXISTS public."groups" CASCADE;
DROP TABLE IF EXISTS public.users CASCADE;
DROP TABLE IF EXISTS public.status CASCADE;
DROP TABLE IF EXISTS public.subjects CASCADE;
DROP TABLE IF EXISTS public.genders CASCADE;
DROP TABLE IF EXISTS public.roles CASCADE;


