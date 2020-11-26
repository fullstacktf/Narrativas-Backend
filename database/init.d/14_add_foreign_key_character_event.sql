ALTER TABLE `rollify`.`character_event`
ADD CONSTRAINT `fk_character_has_event_character`
  FOREIGN KEY (`character_id`)
  REFERENCES `rollify`.`actor` (`id`)
  ON DELETE CASCADE
  ON UPDATE NO ACTION;

ALTER TABLE `rollify`.`character_event`
ADD CONSTRAINT `fk_character_has_event_event`
  FOREIGN KEY (`event_id`)
  REFERENCES `rollify`.`event` (`id`)
  ON DELETE CASCADE
  ON UPDATE NO ACTION;