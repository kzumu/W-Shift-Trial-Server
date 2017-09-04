-- +migrate Up
CREATE TABLE `temp_keys` (
  `id` INTEGER AUTO_INCREMENT NOT NULL,
  `user_id` INTEGER NOT NULL,
  `temp_key` INTEGER NOT NULL,
  `created` timestamp NOT NULL DEFAULT NOW(),
  `updated` timestamp NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

-- +migrate Down
DROP TABLE `temp_keys`;