BEGIN;

ALTER TABLE areas ADD COLUMN latitude decimal, ADD COLUMN longitude decimal;

COMMIT;