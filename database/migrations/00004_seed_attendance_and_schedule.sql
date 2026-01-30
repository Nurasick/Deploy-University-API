
-- +goose Up
-- ATTENDANCE
INSERT INTO public.attendance (id, student_id, subject_id, visit_day, visited) VALUES
(1, 1, 1, '2026-01-20', true),
(2, 1, 2, '2026-01-20', false),
(3, 2, 1, '2026-01-20', true),
(4, 2, 4, '2026-01-21', true),
(5, 3, 3, '2026-01-21', true),
(6, 3, 1, '2026-01-22', false),
(7, 4, 5, '2026-01-21', true),
(8, 4, 4, '2026-01-20', true),
(9, 5, 1, '2026-01-20', true),
(10, 5, 2, '2026-01-21', true)
ON CONFLICT DO NOTHING;

-- CLASS_SCHEDULE
INSERT INTO public.class_schedule (id, group_id, day_of_week, starts_at, ends_at, subject_id, teacher_id) VALUES
(1, 1, 1, '09:00:00', '10:30:00', 3, 1),
(2, 1, 1, '09:00:00', '10:30:00', 3, 1),
(3, 1, 1, '09:00:00', '10:30:00', 3, 1),
(4, 2, 3, '11:00:00', '12:30:00', 2, 2),
(5, 1, 0, '09:00:00', '10:30:00', 1, 2),
(6, 2, 0, '09:00:00', '10:30:00', 3, 1),
(7, 2, 0, '10:45:00', '12:15:00', 1, 1),
(8, 3, 1, '09:00:00', '10:30:00', 5, 1),
(9, 3, 1, '10:45:00', '12:15:00', 3, 1),
(10, 4, 0, '10:45:00', '12:15:00', 1, 2),
(11, 4, 1, '09:00:00', '10:30:00', 3, 1)
ON CONFLICT DO NOTHING;


SELECT setval('attendance_id_seq', (SELECT COALESCE(MAX(id), 0) FROM attendance));
SELECT setval('class_schedule_id_seq', (SELECT COALESCE(MAX(id), 0) FROM class_schedule));
-- +goose Down 

-- Delete attendance and schedule data
DELETE FROM public.class_schedule WHERE id IN (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11);
DELETE FROM public.attendance WHERE id IN (1, 2, 3, 4, 5, 6, 7, 8, 9, 10);
