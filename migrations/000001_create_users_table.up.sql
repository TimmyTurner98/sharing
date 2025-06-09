CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   username VARCHAR (50),
   number VARCHAR(50) NOT NULL,
   password VARCHAR (255) NOT NULL,
   email VARCHAR (300) UNIQUE,
   updated_at TIMESTAMP DEFAULT now(),
   deleted_at TIMESTAMP NULL
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();