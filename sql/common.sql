CREATE EXTENSION IF NOT EXISTS pgcrypto;
-- COMMON STUFF
DROP TABLE IF EXISTS user_profile;


CREATE TABLE user_profile
(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(56) NOT NULL
);


INSERT INTO user_profile
VALUES
  ('8cbd81c0-9f94-44cc-8eac-6798720bbf67', 'dominik');

