ALTER TABLE `rollify`.`story`
ADD CONSTRAINT `fk_user_story`
  FOREIGN KEY (`user_id`)
  REFERENCES `rollify`.`user` (`id`)
  ON DELETE CASCADE
  ON UPDATE NO ACTION;

ALTER TABLE `rollify`.`story`
ADD CONSTRAINT `fk_initial_event_story`
  FOREIGN KEY (`initial_event_id`)
  REFERENCES `rollify`.`event` (`id`)
  ON DELETE CASCADE
  ON UPDATE NO ACTION;