CREATE TABLE IF NOT EXISTS `rollify`.`character_section` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `character_id` INT UNSIGNED NOT NULL,
  `title` VARCHAR(50) NOT NULL,
  `created_at` TIMESTAMP(2) NULL,
  `updated_at` TIMESTAMP(2) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  UNIQUE INDEX `title_UNIQUE` (`title` ASC) VISIBLE,
  UNIQUE INDEX `character_id_UNIQUE` (`character_id` ASC) VISIBLE,
  CONSTRAINT `fk_character`
    FOREIGN KEY (`character_id`)
    REFERENCES `rollify`.`actor` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION);