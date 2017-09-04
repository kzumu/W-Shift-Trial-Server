-- +migrate Up
CREATE TABLE `shifts` (
  `id` INTEGER AUTO_INCREMENT NOT NULL,
  `user_id` INTEGER NOT NULL,
  `name` VARCHAR(256),
  `color` VARCHAR(6),
  `date` DATE NOT NULL,
  `created` timestamp NOT NULL DEFAULT NOW(),
  `updated` timestamp NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

-- +migrate Down
DROP TABLE `shifts`;