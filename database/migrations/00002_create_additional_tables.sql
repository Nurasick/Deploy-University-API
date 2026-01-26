-- +goose Up

-- Create attendance table
CREATE TABLE public.attendance (
    id integer NOT NULL,
    student_id integer NOT NULL,
    subject_id integer NOT NULL,
    visit_day date NOT NULL,
    visited boolean NOT NULL
);

CREATE SEQUENCE public.attendance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.attendance_id_seq OWNED BY public.attendance.id;

-- Create class_schedule table
CREATE TABLE public.class_schedule (
    id integer NOT NULL,
    group_id integer,
    day_of_week integer NOT NULL,
    starts_at time without time zone NOT NULL,
    ends_at time without time zone NOT NULL,
    subject_id integer,
    teacher_id integer
);

CREATE SEQUENCE public.class_schedule_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.class_schedule_id_seq OWNED BY public.class_schedule.id;

-- Create teachers table
CREATE TABLE public.teachers (
    id integer NOT NULL,
    full_name varchar(100) NOT NULL,
    user_id integer not null unique,
    department varchar(100)
);

CREATE SEQUENCE public.teachers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.teachers_id_seq OWNED BY public.teachers.id;

-- +goose Down 

DROP TABLE IF EXISTS public.teachers CASCADE;
DROP SEQUENCE IF EXISTS public.teachers_id_seq;

DROP TABLE IF EXISTS public.class_schedule CASCADE;
DROP SEQUENCE IF EXISTS public.class_schedule_id_seq;

DROP TABLE IF EXISTS public.attendance CASCADE;
DROP SEQUENCE IF EXISTS public.attendance_id_seq;
