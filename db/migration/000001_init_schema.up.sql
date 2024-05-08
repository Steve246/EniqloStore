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

CREATE TABLE ProductList (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    sku VARCHAR(30) NOT NULL,
    category VARCHAR(60) NOT NULL,
    imageUrl VARCHAR(100) NOT NULL,
    notes VARCHAR(200) NOT NULL,
    price INT NOT NULL CHECK (price >= 0),
    stock INT NOT NULL CHECK (stock >= 0 AND stock <= 100000),
    location VARCHAR(200) NOT NULL,
    isAvailable BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);