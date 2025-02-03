BEGIN;

CREATE TABLE IF NOT EXISTS `tab_address` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `zipcode` VARCHAR(8) NULL,
  `street` VARCHAR(255) NULL,
  `complement` VARCHAR(255) NULL,
  `neighborhood` VARCHAR(255) NULL,
  `city` VARCHAR(120) NULL,
  `pa` VARCHAR(2) NULL,
  `state` VARCHAR(120) NULL,
  `region` VARCHAR(120) NULL,
  `created_at` TIMESTAMP DEFAULT NOW(),
  `updated_at` TIMESTAMP DEFAULT NOW(),
  PRIMARY KEY (`id`));

COMMIT;