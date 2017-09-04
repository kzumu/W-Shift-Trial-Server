-- +migrate Up
CREATE TABLE `users` (
  `id` INTEGER AUTO_INCREMENT NOT NULL,
  `name` VARCHAR(255),
  `uuid` VARCHAR(255) NOT NULL,
  `partner_id` INTEGER,
  `created` timestamp NOT NULL DEFAULT NOW(),
  `updated` timestamp NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`partner_id`) REFERENCES `users`(`id`)
);

-- +migrate Down
DROP TABLE `users`;