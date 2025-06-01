CREATE TABLE users (
  chat_id      INTEGER PRIMARY KEY,
  name         VARCHAR(255) NOT NULL,
  created_at   BIGINT NOT NULL DEFAULT 0
);