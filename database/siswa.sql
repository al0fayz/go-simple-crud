CREATE TABLE `siswa` ( 
	`id` Int( 8 ) AUTO_INCREMENT NOT NULL,
	`nim` BigInt( 14 ) NOT NULL,
	`name` VarChar( 255 ) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
	`semester` VarChar( 255 ) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
	`created_at` DateTime NOT NULL,
	`updated_at` DateTime NOT NULL,
	CONSTRAINT `unique_id` UNIQUE( `id` ) )
CHARACTER SET = utf8
COLLATE = utf8_general_ci
ENGINE = InnoDB
AUTO_INCREMENT = 9;
-- -------------------------------------------------------------


-- Dump data of "siswa" --------------------------------
INSERT INTO `siswa`(`id`,`nim`,`name`,`semester`,`created_at`,`updated_at`) VALUES 
( '1', '0001', 'Alfa Code', '4', '2020-10-01 00:00:00', '2020-10-01 00:00:00' ),
( '2', '0002', 'Beta Hack', '4', '2020-10-01 00:00:00', '2020-10-01 00:00:00' ),
( '3', '0003', 'Cika Cika', '9', '2020-10-01 00:00:00', '2020-10-01 00:00:00' ),
( '4', '0004', 'Choco Late', '9', '2020-10-01 00:00:00', '2020-10-01 00:00:00' );
-- ---------------------------------------------------------