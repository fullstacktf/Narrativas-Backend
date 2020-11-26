ALTER TABLE `rollify`.`character_section_field`
ADD CONSTRAINT `fk_character_section`
  FOREIGN KEY (`section_id`)
  REFERENCES `rollify`.`character_section` (`id`)
  ON DELETE CASCADE
  ON UPDATE NO ACTION;