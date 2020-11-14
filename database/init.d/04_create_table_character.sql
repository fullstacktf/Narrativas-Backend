CREATE TABLE IF NOT EXISTS `rollify`.`actor` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(50) NOT NULL,
  `biography` LONGTEXT NULL,
  `created_at` TIMESTAMP(2) NULL,
  `updated_at` TIMESTAMP(2) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  INDEX `user_id_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_user_actor`
    FOREIGN KEY (`user_id`)
    REFERENCES `rollify`.`user` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION);