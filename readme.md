# Uniqlo 

Uniqlo is an markeplace inventory system that allow customer to buy products, and staff to add product inside the system

![Unqilo](UniqloImage.jpg)

## 📜Table of Contents

- [Uniqlo Inventory System](#Uniqlo)
  - [📜Table of Contents](#table-of-contents)
  - [🔍Requirements](#requirements)
  - [🛠️Installation](#️installation)
  - [🌟Features](#features)
  - [🚀Usage](#usage)
  - [💾Database Migration](#database-migration)

## 🛠️Installation

To install the boilerplate, follow these steps:

1. Make sure you download Golang, PostgreSQL, and migration package (go-migrate)

2. Clone this repository, make sure ssh is enable:

   ```bash
   git clone git@github.com:Steve246/EniqloStore.git
   ```

3. Navigate to the project directory:

   ```bash
   cd EniqloStore
   ```

4. Run the vendor command to install dependencies:
   ```bash
   go mod vendor
   ```

5. Run the program, after migration process is finished:
   ```bash
   go mod run .
   ```

## 🌟Features

Uniqlo offers the following features:

- **Authentication**:
  - Staff registration
  - Staff login
- **Product Management (CRUD)**:
  - Create new products
  - Get all products
  - Update exisiting products
  - Delete exisiting products
- **Search SKU**:
  - Get All Product from Customer

- **Checkout**:
  - Customer registration
  - Get All customer data
  - Customer Checkout
  - Customer Checkout History (Coming Soon)

## 💾Database Migration

Database migration must use golang-migrate as a tool to manage database migration

1. Direct your terminal to your project folder first
2. Initiate folder
   ```bash
   mkdir db/migrations
   ```
3. Create migration

   ```bash
   migrate create -ext sql -dir db/migrations add_user_table
   ```

   This command will create two new files named `add_user_table.up.sql` and `add_user_table.down.sql` inside the `db/migrations` folder

   - `.up.sql` can be filled with database queries to create / delete / change the table
   - `.down.sql` can be filled with database queries to perform a `rollback` or return to the state before the table from `.up.sql` was created

4. Insert this querry, inside add_user_table.up.sql

``` bash

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

```


5. Insert this querry, inside add_user_table.down.sql

```bash 

    DROP TABLE IF EXISTS "staffdata", "authentication";

```

6. Execute migration

   ```bash
   migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations up
   ```

7. Rollback migration

   ```bash
   migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations dow
   ```

8. View the current migration state
   ```bash
    migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" version
   ```
