CREATE TABLE IF NOT EXISTS `rollify`.`event` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `story_id` INT UNSIGNED NOT NULL,
  `title` VARCHAR(50) NOT NULL,
  `description` LONGTEXT NULL,
  `created_at` TIMESTAMP(2) NULL,
  `updated_at` TIMESTAMP(2) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  INDEX `fk_story_idx` (`story_id` ASC) VISIBLE,
  CONSTRAINT `fk_story`
    FOREIGN KEY (`story_id`)
    REFERENCES `rollify`.`story` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION);