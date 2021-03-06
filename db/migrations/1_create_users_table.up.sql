CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	email VARCHAR(256) NOT NULL,
	first_name VARCHAR(256) NOT NULL,
	last_name VARCHAR(256) NOT NULL,
	password VARCHAR(256) NOT NULL,
	created_At TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT email_unique UNIQUE (email)
);
