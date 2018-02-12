-- COMMON STUFF
DROP TABLE IF EXISTS user_profile;


CREATE TABLE user_profile
(
  id INT PRIMARY KEY,
  name VARCHAR(56) NOT NULL
);


INSERT INTO user_profile
VALUES
  (1, 'dominik');
-- wight with unit

