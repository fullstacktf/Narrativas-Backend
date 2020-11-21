CREATE TABLE IF NOT EXISTS `rollify`.`character_event` (
  `character_id` INT UNSIGNED NOT NULL,
  `event_id` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP(2) NULL,
  `updated_at` TIMESTAMP(2) NULL,
  PRIMARY KEY (`event_id`, `character_id`),
  INDEX `fk_character_has_event_event_idx` (`event_id` ASC) VISIBLE,
  INDEX `fk_character_has_event_character_idx` (`character_id` ASC) VISIBLE);