CREATE TABLE customers (
       id BIGSERIAL PRIMARY KEY,
       name varchar(255) NOT NULL,
       email varchar (255) NOT NULL UNIQUE ,
       password varchar (255) NOT NULL,
       address varchar (255) NOT NULL,
       created_at date NOT NULL,
       updated_at date
);

-- name: GetCustomer :one
SELECT * FROM customers WHERE id = $1 LIMIT 1;

-- name: FindCustomers :many
SELECT * FROM customers;

-- name: CreateCustomer :one
INSERT INTO customers (name, email, password, address, created_at)
    VALUES (
        $1, $2, $3, $4, $5
    )
    RETURNING *;

-- name: UpdateCustomer :exec
UPDATE customers SET name=$1, email=$2, password=$3, address=$4, updated_at=$5 WHERE id=$6;

-- name: DeleteCustomer :exec
DELETE FROM customers WHERE id=$1;