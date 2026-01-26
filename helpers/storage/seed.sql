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


INSERT INTO public."groups" (id, name, direction) VALUES
(1, 'ENG-251', 'Engineering'),
(2, 'CS-251', 'Computer Science'),
(3,'HUM-251','Human Science'),
(4,'ENG-252','Engineering')
ON CONFLICT (name) DO NOTHING;





