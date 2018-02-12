CREATE EXTENSION
IF NOT EXISTS pgcrypto;
DROP VIEW IF EXISTS user_weight;
DROP VIEW IF EXISTS weight_with_unit;
DROP TABLE IF EXISTS user_weight_connector;
DROP TABLE IF EXISTS weight;
DROP TABLE IF EXISTS user_profile;
DROP TABLE IF EXISTS weight_unit;
DROP TABLE IF EXISTS asset_type;
-- aads
CREATE TABLE weight_unit
(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(8) NOT NULL
);
CREATE TABLE asset_type
(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL
);
CREATE TABLE user_profile
(
  id INT PRIMARY KEY,
  name VARCHAR(56) NOT NULL
);
CREATE TABLE weight
(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  value REAL NOT NULL,
  unit UUID REFERENCES weight_unit (id) NOT NULL,
  timestamp TIMESTAMP NOT NULL
);
CREATE TABLE user_weight_connector
(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id INT REFERENCES user_profile (id) NOT NULL,
  weight_id UUID REFERENCES weight (id) NOT NULL
);
-- inserts
INSERT INTO weight_unit
VALUES
  (gen_random_uuid(), 'kg'),
  (gen_random_uuid(), 'lb');
INSERT INTO asset_type
VALUES
  (gen_random_uuid(), 'p2p'),
  (gen_random_uuid(), 'cryptocurrency'),
  (gen_random_uuid(), 'bank_account'),
  (gen_random_uuid(), 'saving_account'),
  (gen_random_uuid(), 'cash'),
  (gen_random_uuid(), 'investment_account');


INSERT INTO user_profile
VALUES
  (1, 'dominik');
-- wight with unit


CREATE VIEW weight_with_unit AS
SELECT
  w.id,
  w.value,
  wu.name as unit,
  timestamp
from weight w
INNER JOIN weight_unit wu ON wu.id = w.unit;

CREATE VIEW user_weight AS
SELECT
  w.id,
  u.id as user_id,
  w.value,
  w.unit,
  timestamp
from weight_with_unit w
INNER JOIN user_weight_connector uwc ON w.id = uwc.weight_id
INNER JOIN user_profile u ON u.id = uwc.user_id