-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_update_column()
  RETURNS TRIGGER AS $$
BEGIN
  IF row (NEW.*) IS DISTINCT FROM row (OLD.*) THEN
    NEW.updated = now();
RETURN NEW;
ELSE
    RETURN OLD;
END IF;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_airport_modtime
  BEFORE UPDATE
  ON airports
  FOR EACH ROW
  EXECUTE PROCEDURE update_update_column();
CREATE TRIGGER update_airport_events_modtime
  BEFORE UPDATE
  ON airport_events
  FOR EACH ROW
  EXECUTE PROCEDURE update_update_column();
-- +goose StatementEnd

-- +goose Down
DROP TRIGGER update_airport_modtime on airports;
DROP TRIGGER update_airport_events_modtime on airport_events;
DROP FUNCTION update_update_column;

