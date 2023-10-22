BEGIN TRANSACTION;

    CREATE TABLE IF NOT EXISTS user_records(
        id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
        name VARCHAR,
        data BYTEA,
        data_type VARCHAR,
        version SERIAL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        user_id INT NOT NULL,
        CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
    );

    CREATE UNIQUE INDEX IF NOT EXISTS user_id_name_idx on user_records(user_id, name);

COMMIT;