ALTER TABLE story
ADD COLUMN `initial_event_id` INT UNSIGNED NULL
AFTER `user_id`;