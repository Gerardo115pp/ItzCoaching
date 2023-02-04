PRAGMA foreign_keys = ON;
.headers ON

DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
  `id` TEXT,
  `username` TEXT NOT NULL UNIQUE,
  `name` TEXT NOT NULL,
  `email` TEXT NOT NULL UNIQUE,
  `phone` TEXT NOT NULL,
  `password` TEXT NOT NULL,
  `created_at` TEXT DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS `restaurants`;
CREATE TABLE `restaurants` (
  `id` TEXT,
  `email` TEXT NOT NULL UNIQUE,
  `password` TEXT NOT NULL,
  `created_at` TEXT DEFAULT CURRENT_TIMESTAMP
);



