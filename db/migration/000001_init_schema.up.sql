CREATE TABLE staffdata (
    id SERIAL PRIMARY KEY,
    user_unique_id varchar(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL, 
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE authentication(
	id SERIAL PRIMARY KEY,
	user_unique_id varchar(255),
	token_auth varchar(255),
	expire TIMESTAMP WITHOUT TIME ZONE
);

