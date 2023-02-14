DROP DATABASE IF EXISTS `itz_coaching`;
CREATE DATABASE IF NOT EXISTS `itz_coaching`;
USE `itz_coaching`;

DROP TABLE IF EXISTS `itz_coaching`.`admins`;
CREATE TABLE IF NOT EXISTS `itz_coaching`.`admins` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(50) NOT NULL UNIQUE,
    `email` varchar(50) NOT NULL,
    `password` varchar(60) NOT NULL,
    `is_active` TINYINT(1) NOT NULL DEFAULT '1',
    `is_superadmin` TINYINT(1) NOT NULL DEFAULT '0',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `last_login` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` INT(11) NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;

ALTER TABLE `admins` ADD CONSTRAINT `admin_ibfk_1` FOREIGN KEY (`created_by`) REFERENCES `admins` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

DROP TABLE IF EXISTS `itz_coaching`.`experts`;
CREATE TABLE IF NOT EXISTS `itz_coaching`.`experts` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(50) NOT NULL UNIQUE,
    `name` varchar(50) NOT NULL,
    `email` varchar(50) NOT NULL,
    `password` varchar(60) NOT NULL,
    `is_active` TINYINT(1) NOT NULL DEFAULT '1',
    `is_available` TINYINT(1) NOT NULL DEFAULT '0',
    `is_email_verified` TINYINT(1) NOT NULL DEFAULT '0',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `last_login` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `type` ENUM('mentor', 'consultant') NOT NULL,
    `created_by` INT(11) NOT NULL,
    CONSTRAINT `expert_ibfk_1` FOREIGN KEY (`created_by`) REFERENCES `admins` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `itz_coaching`.`public_profiles`;
CREATE TABLE IF NOT EXISTS `itz_coaching`.`public_profiles` (
    `public_name` varchar(50) NOT NULL,
    `professional_title` varchar(100) NOT NULL DEFAULT '',
    `description` TEXT NOT NULL DEFAULT '',
    `brief` VARCHAR(255) NOT NULL DEFAULT '',
    `expert` INT(11) NOT NULL,
    `image_href` varchar(255) NOT NULL,
    CONSTRAINT `public_profile_ibfk_1` FOREIGN KEY (`expert`) REFERENCES `experts` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
    PRIMARY KEY (`expert`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `itz_coaching`.`appointments`;
CREATE TABLE `appointments` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `expert` INT(11) NOT NULL,
    `customer_email` VARCHAR(120) NOT NULL,
    `start` timestamp NOT NULL,
    `duration` INT NOT NULL,
    `status` ENUM('unpaid', 'paid', 'finalized','cancelled') NOT NULL DEFAULT 'unpaid',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `appoINTment_ibfk_1` FOREIGN KEY (`expert`) REFERENCES `experts` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `itz_coaching`.`appintment_confirmations`;
CREATE TABLE `appintment_confirmations` (
    `appointment` INT(11) NOT NULL,
    `expert_confirmed` TIMESTAMP NULL DEFAULT NULL,
    `consumer_confirmed` TIMESTAMP NULL DEFAULT NULL,
    CONSTRAINT `appintment_confirmation_ibfk_1` FOREIGN KEY (`appointment`) REFERENCES `appointments` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;

DROP TRIGGER IF EXISTS `itz_coaching`.`public_profile_CREATE`;
DELIMITER //
CREATE TRIGGER `public_profile_CREATE` AFTER INSERT ON `experts` FOR EACH ROW
BEGIN
    INSERT INTO `public_profiles` (`expert`, `public_name`, `image_href`) VALUES (NEW.id, NEW.name, '');
END //
DELIMITER ;

DROP TABLE IF EXISTS `itz_coaching`.`payments`;
CREATE TABLE `payments` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `appointment` INT(11) NOT NULL,
    `amount` DECIMAL(10,2) NOT NULL,
    `currency` VARCHAR(3) NOT NULL,
    `stripe_charge_id` VARCHAR(255) NOT NULL,
    `stripe_customer_id` VARCHAR(255) NOT NULL,
    `stripe_payment_intent_id` VARCHAR(255) NOT NULL,
    `stripe_payment_id` VARCHAR(255) NOT NULL,
    `description` TEXT NOT NULL,
    `status` ENUM('succeded', 'failed', 'pending') NOT NULL DEFAULT 'pending',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `payment_ibfk_1` FOREIGN KEY (`appointment`) REFERENCES `appointments` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;

INSERT INTO `admins`(`username`, `email`, `password`, `is_active`, `is_superadmin`) VALUES ( 'el_maligno', 'theronin115@gmail.com', '$2a$10$yDp75OTteLqYx5Y6bX53OuStUxiv9fu997kT5EHNKYbLOqUV3pXNG', 1, 1); 