-- SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
-- SPDX-License-Identifier: MIT
--
-- name: create-patches
CREATE TABLE `patches` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);
-- name: create-prefixes
CREATE TABLE `prefixes` (
    `prefix` VARCHAR(20) NOT NULL,
    `count` INT NOT NULL DEFAULT 1,
    PRIMARY KEY (`prefix`)
);