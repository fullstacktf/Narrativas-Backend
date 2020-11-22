ALTER TABLE `roliffy`.`character_section`
ADD CONSTRAINT `fk_section_character_id`
  FOREIGN KEY (`character_id`)
  REFERENCES `rollify`.`actor` (`id`)
  ON DELETE CASCADE
  ON UPDATE NO ACTION;