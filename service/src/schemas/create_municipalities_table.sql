CREATE TABLE `MUNICIPALITIES` (
  `ID`   INT         NOT NULL AUTO_INCREMENT,
  `NAME` VARCHAR(20) NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE (`NAME`)
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB;

CREATE TABLE `TAXES` (
  `ID`              INT            NOT NULL AUTO_INCREMENT,
  `MUNICIPALITY_ID` INT            NOT NULL,
  `FROM`            DATE           NOT NULL,
  `TO`              DATE           NOT NULL,
  `TYPE`            VARCHAR(20)    NOT NULL,
  `VALUE`           DECIMAL(10, 2) NOT NULL,
  PRIMARY KEY (`ID`),
  FOREIGN KEY (`MUNICIPALITY_ID`) REFERENCES `MUNICIPALITIES` (`ID`)
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB;