--  WEALTH STUFF
DROP VIEW IF EXISTS ws_with_asset;
DROP TABLE IF EXISTS w_user_connector;
DROP TABLE IF EXISTS w_currency_connector;
DROP TABLE IF EXISTS w_record;
DROP TABLE IF EXISTS w_source;
DROP TABLE IF EXISTS asset_type;
DROP TABLE IF EXISTS currency;

-- TABLES
CREATE TABLE currency
(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL
);

CREATE TABLE asset_type
(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL
);

CREATE TABLE w_source (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL,
  asset_type_id UUID REFERENCES asset_type (id) NOT NULL
);

CREATE TABLE w_user_connector (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id INT REFERENCES user_profile (id) NOT NULL,
  ws_id UUID REFERENCES w_source (id) NOT NULL
);

CREATE TABLE ws_currency_connector (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  currency_ID UUID REFERENCES currency (id) NOT NULL,
  ws_id UUID REFERENCES w_source (id) NOT NULL
);

CREATE TABLE ws_record (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL,
  value real NOT NULL,
  w_source_id UUID REFERENCES w_source (id) NOT NULL,
  user_id INT REFERENCES user_profile (id) NOT NULL
);

-- VIEWS

CREATE VIEW ws_with_asset AS
SELECT
  ws.id,
  ws.name,
  asset_type.name as asset_type
from w_source ws
INNER JOIN asset_type ON asset_type.id = ws.asset_type_id;

-- VIEWS

INSERT INTO asset_type
VALUES
  (gen_random_uuid(), 'p2p'),
  (gen_random_uuid(), 'cryptocurrency'),
  (gen_random_uuid(), 'bank_account'),
  (gen_random_uuid(), 'saving_account'),
  (gen_random_uuid(), 'cash'),
  (gen_random_uuid(), 'investment_account');

INSERT INTO currency
VALUES
  (gen_random_uuid(), 'czk'),
  (gen_random_uuid(), 'eur');
