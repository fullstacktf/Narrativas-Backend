ALTER TABLE `rollify`.`actor`
ADD CONSTRAINT `fk_user_actor`
  FOREIGN KEY (`user_id`)
  REFERENCES `rollify`.`user` (`id`)
  ON DELETE CASCADE
  ON UPDATE NO ACTION;