CREATE TABLE IF NOT EXISTS `rollify`.`story` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `initial_event` INT UNSIGNED NOT NULL,
  `final_event` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP(2) NULL,
  `updated_at` TIMESTAMP(2) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  UNIQUE INDEX `relation_UNIQUE` (`initial_event`, `final_event`) VISIBLE,
  INDEX `fk_initial_event_idx` (`initial_event` ASC) VISIBLE,
  CONSTRAINT `fk_initial_event`
    FOREIGN KEY (`initial_event`)
    REFERENCES `rollify`.`event` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  INDEX `fk_final_event_idx` (`final_event` ASC) VISIBLE,
  CONSTRAINT `fk_final_event`
    FOREIGN KEY (`final_event`)
    REFERENCES `rollify`.`event` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);