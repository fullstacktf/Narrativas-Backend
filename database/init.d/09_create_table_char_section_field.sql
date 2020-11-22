CREATE TABLE IF NOT EXISTS `rollify`.`character_section_field` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `section_id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(50) NOT NULL,
  `value` VARCHAR(255) NOT NULL,
  `description` LONGTEXT NULL,
  `created_at` TIMESTAMP(2) NULL,
  `updated_at` TIMESTAMP(2) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  INDEX `fk_character_section_idx` (`section_id` ASC) VISIBLE);