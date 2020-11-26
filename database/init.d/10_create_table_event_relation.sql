CREATE TABLE IF NOT EXISTS `rollify`.`event_relation` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `initial_event` INT UNSIGNED NOT NULL,
  `final_event` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP(2) NULL,
  `updated_at` TIMESTAMP(2) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  UNIQUE INDEX `relation_UNIQUE` (`initial_event`, `final_event`) VISIBLE,
  INDEX `fk_initial_event_idx` (`initial_event` ASC) VISIBLE,
  INDEX `fk_final_event_idx` (`final_event` ASC) VISIBLE);