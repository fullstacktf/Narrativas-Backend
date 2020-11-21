ALTER TABLE `rollify`.`event_relation`
ADD CONSTRAINT `fk_initial_event`
  FOREIGN KEY (`initial_event`)
  REFERENCES `rollify`.`event` (`id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION;

ALTER TABLE `rollify`.`event_relation`
ADD CONSTRAINT `fk_final_event`
  FOREIGN KEY (`final_event`)
  REFERENCES `rollify`.`event` (`id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION;