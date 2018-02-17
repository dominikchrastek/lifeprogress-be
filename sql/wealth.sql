--  WEALTH STUFF
DROP VIEW IF EXISTS user_ws;
DROP VIEW IF EXISTS ws_with_type;
DROP VIEW IF EXISTS ws_currency;

DROP TABLE IF EXISTS ws_user_connector;
DROP TABLE IF EXISTS ws_currency_connector;
DROP TABLE IF EXISTS ws_record;
DROP TABLE IF EXISTS wsource;
DROP TABLE IF EXISTS ws_type;
DROP TABLE IF EXISTS currency;

-- TABLES
CREATE TABLE currency
(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL
);

CREATE TABLE ws_type
(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL
);

CREATE TABLE wsource (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL,
  ws_type UUID REFERENCES ws_type (id) NOT NULL
);

CREATE TABLE ws_user_connector (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID REFERENCES user_profile (id) NOT NULL,
  ws_id UUID REFERENCES wsource (id) NOT NULL
);

CREATE TABLE ws_currency_connector (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  currency_ID UUID REFERENCES currency (id) NOT NULL,
  ws_id UUID REFERENCES wsource (id) NOT NULL
);

CREATE TABLE ws_record (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL,
  value real NOT NULL,
  ws_id UUID REFERENCES wsource (id) NOT NULL,
  user_id UUID REFERENCES user_profile (id) NOT NULL,
  currency_id UUID REFERENCES currency (id) NOT NULL,
  timestamp TIMESTAMP NOT NULL
);

-- VIEWS

CREATE VIEW ws_with_type AS
SELECT
  ws.id,
  ws.name,
  ws_type.name as ws_type
from wsource ws
INNER JOIN ws_type ON ws_type.id = ws.ws_type;

CREATE VIEW user_ws AS
SELECT
  ws.id,
  ws.name,
  ws.ws_type,
  wuc.user_id
from ws_with_type ws
INNER JOIN ws_user_connector wuc ON ws.id = wuc.ws_id
INNER JOIN user_profile u ON u.id = wuc.user_id;



CREATE VIEW ws_currency AS
SELECT
  c.id,
  wcc.ws_id,
  c.name as name
from wsource ws
INNER JOIN ws_currency_connector wcc ON wcc.ws_id = ws.id
INNER JOIN currency c ON c.id = wcc.currency_id;
-- INSERTS

INSERT INTO ws_type
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

