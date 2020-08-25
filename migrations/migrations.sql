-- add timestamps
ALTER TABLE threads ALTER COLUMN created_on SET default NOW();
ALTER TABLE threads ALTER COLUMN bumped_on SET default NOW();
ALTER TABLE replies ALTER COLUMN created_on SET default NOW();

-- add not null constraints
ALTER TABLE replies ALTER COLUMN delete_password SET NOT NULL;
ALTER TABLE threads ALTER COLUMN delete_password SET NOT NULL;

-- add foreign keys
ALTER TABLE threads ADD FOREIGN KEY (board_id) REFERENCES boards(id);
ALTER TABLE replies ADD FOREIGN KEY (thread_id) REFERENCES threads(id);