-- +goose Up

CREATE TABLE airports
(
  id       BIGSERIAL PRIMARY KEY,
  name     CHAR(3)      NOT NULL,
  location VARCHAR(255) NOT NULL,
  removed  BOOLEAN      NOT NULL DEFAULT FALSE,
  created  timestamp             DEFAULT current_timestamp,
  updated  timestamp
);



CREATE TABLE airport_events
(
  id         BIGSERIAL PRIMARY KEY,
  airport_id BIGINT  NOT NULL,
  type       SMALLINT NOT NULL,
  status     SMALLINT NOT NULL DEFAULT 1,
  is_locked  BOOLEAN NOT NULL DEFAULT FALSE,
  payload    jsonb,
  updated    timestamp,
  CONSTRAINT fk_airport FOREIGN KEY (airport_id) REFERENCES airports (id) ON DELETE CASCADE

);

-- +goose Down
DROP TABLE airports;
DROP TABLE airport_events;
