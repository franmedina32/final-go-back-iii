-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema database_back_iii
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema database_back_iii
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `database_back_iii` DEFAULT CHARACTER SET utf8mb3 ;
USE `database_back_iii` ;

-- -----------------------------------------------------
-- Table `database_back_iii`.`odontologos`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `database_back_iii`.`odontologos` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(100) NULL DEFAULT NULL,
  `apellido` VARCHAR(100) NULL DEFAULT NULL,
  `matricula` VARCHAR(100) NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `database_back_iii`.`pacientes`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `database_back_iii`.`pacientes` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(100) NULL DEFAULT NULL,
  `apellido` VARCHAR(100) NULL DEFAULT NULL,
  `domicilio` VARCHAR(200) NULL DEFAULT NULL,
  `dni` VARCHAR(100) NULL DEFAULT NULL,
  `fecha_alta` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `database_back_iii`.`turnos`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `database_back_iii`.`turnos` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `id_paciente` INT NULL,
  `id_odontologo` INT NULL DEFAULT NULL,
  `fecha` DATETIME NULL DEFAULT NULL,
  `descripcion` VARCHAR(1000) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `OT_idx` (`id_odontologo` ASC) VISIBLE,
  INDEX `PT_idx` (`id_paciente` ASC) VISIBLE,
  CONSTRAINT `OT`
    FOREIGN KEY (`id_odontologo`)
    REFERENCES `database_back_iii`.`odontologos` (`id`),
  CONSTRAINT `PT`
    FOREIGN KEY (`id_paciente`)
    REFERENCES `database_back_iii`.`pacientes` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
