CREATE TABLE IF NOT EXISTS `rollify`.`character_event` (
  `character_id` INT UNSIGNED NOT NULL,
  `event_id` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP(2) NULL,
  `updated_at` TIMESTAMP(2) NULL,
  PRIMARY KEY (`event_id`, `character_id`),
  INDEX `fk_character_has_event_event1_idx` (`event_id` ASC) VISIBLE,
  INDEX `fk_character_has_event_character1_idx` (`character_id` ASC) VISIBLE,
  CONSTRAINT `fk_character_has_event_character1`
    FOREIGN KEY (`character_id`)
    REFERENCES `rollify`.`actor` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_character_has_event_event1`
    FOREIGN KEY (`event_id`)
    REFERENCES `rollify`.`event` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION);