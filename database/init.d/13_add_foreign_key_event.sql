ALTER TABLE `rollify`.`event`
ADD CONSTRAINT `fk_event_story_id`
  FOREIGN KEY (`story_id`)
  REFERENCES `rollify`.`story` (`id`)
  ON DELETE CASCADE
  ON UPDATE NO ACTION;