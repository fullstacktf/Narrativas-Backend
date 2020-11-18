CREATE TABLE IF NOT EXISTS `rollify`.`story` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` INT UNSIGNED NOT NULL,
  `title` VARCHAR(50) NOT NULL,
  `image` VARCHAR(150) NOT NULL,
  `created_at` TIMESTAMP(2) NULL,
  `updated_at` TIMESTAMP(2) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  UNIQUE INDEX `image_UNIQUE` (`image` ASC) VISIBLE,
  INDEX `fk_user_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_user_story`
    FOREIGN KEY (`user_id`)
    REFERENCES `rollify`.`user` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION);