--  WEALTH STUFF
DROP VIEW IF EXISTS user_ws;
DROP VIEW IF EXISTS ws_with_type;
DROP VIEW IF EXISTS ws_currency;
DROP VIEW IF EXISTS wsr_currency;

DROP TABLE IF EXISTS ws_user_connector;
DROP TABLE IF EXISTS ws_currency_connector;
DROP TABLE IF EXISTS ws_record;
DROP TABLE IF EXISTS wsource;
DROP TABLE IF EXISTS ws_type;
DROP TABLE IF EXISTS currency_rate;
DROP TABLE IF EXISTS currency;

-- TABLES

CREATE TABLE currency
(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL
);

-- CREATE TABLE currency_rate (
--   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--   from_currency UUID REFERENCES currency(id) NOT NULL,
--   to_currency UUID REFERENCES currency(id) NOT NULL,
--   value REAL NOT NULL,
--   timestamp TIMESTAMP NOT NULL
-- );

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


CREATE VIEW wsr_currency AS
SELECT
  wsr.id,
  wsr.name,
  wsr.value,
  wsr.timestamp,
  wsr.ws_id,
  wsr.user_id,
  c.name as currency_name
from ws_record wsr
INNER JOIN currency c ON c.id = wsr.currency_id;
-- INSERTS

INSERT INTO ws_type
VALUES
  ('1cbd81c0-9f94-44cc-8eac-6798720bbf61', 'p2p'),
  ('1cbd81c0-9f94-44cc-8eac-6798720bbf62', 'cryptocurrency'),
  ('1cbd81c0-9f94-44cc-8eac-6798720bbf63', 'bank_account'),
  ('1cbd81c0-9f94-44cc-8eac-6798720bbf64', 'saving_account'),
  ('1cbd81c0-9f94-44cc-8eac-6798720bbf65', 'cash'),
  ('1cbd81c0-9f94-44cc-8eac-6798720bbf66', 'investment_account');

INSERT INTO currency
VALUES
  ('2cbd81c0-9f94-44cc-8eac-6798720bbf61', 'CZK'),
  ('2cbd81c0-9f94-44cc-8eac-6798720bbf62', 'EUR');

INSERT INTO wsource
VALUES
  ('3cbd81c0-9f94-44cc-8eac-6798720bbf61', 'mintos', '1cbd81c0-9f94-44cc-8eac-6798720bbf61'),
  ('3cbd81c0-9f94-44cc-8eac-6798720bbf62', 'coinbase', '1cbd81c0-9f94-44cc-8eac-6798720bbf62');

INSERT INTO ws_currency_connector
VALUES
  ('5cbd81c0-9f94-44cc-8eac-6798720bbf61', '2cbd81c0-9f94-44cc-8eac-6798720bbf61', '3cbd81c0-9f94-44cc-8eac-6798720bbf61'),
  ('5cbd81c0-9f94-44cc-8eac-6798720bbf62', '2cbd81c0-9f94-44cc-8eac-6798720bbf61', '3cbd81c0-9f94-44cc-8eac-6798720bbf62');

INSERT INTO ws_user_connector
VALUES
  ('4cbd81c0-9f94-44cc-8eac-6798720bbf61', '8cbd81c0-9f94-44cc-8eac-6798720bbf67', '3cbd81c0-9f94-44cc-8eac-6798720bbf61'),
  ('4cbd81c0-9f94-44cc-8eac-6798720bbf62', '8cbd81c0-9f94-44cc-8eac-6798720bbf67', '3cbd81c0-9f94-44cc-8eac-6798720bbf62');

INSERT INTO ws_record
VALUES
  ('6cbd81c0-9f94-44cc-8eac-6798720bbf61'
  ,'first mintos'
  , 300
  ,'3cbd81c0-9f94-44cc-8eac-6798720bbf61'
  ,'8cbd81c0-9f94-44cc-8eac-6798720bbf67'
  ,'2cbd81c0-9f94-44cc-8eac-6798720bbf61'
  , current_timestamp),
  ('6cbd81c0-9f94-44cc-8eac-6798720bbf62'
  ,'first coinbase'
  , 301
  ,'3cbd81c0-9f94-44cc-8eac-6798720bbf62'
  ,'8cbd81c0-9f94-44cc-8eac-6798720bbf67'
  ,'2cbd81c0-9f94-44cc-8eac-6798720bbf61'
  , '2018-03-16 13:55:50.171906');