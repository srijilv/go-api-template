CREATE TYPE common_status as ENUM('active', 'inactive', 'blocked', 'deleted', 'deactivated');

CREATE TABLE books (
    id smallserial primary key,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    genre VARCHAR(100),
    published_year INT,
    isbn VARCHAR(36) UNIQUE,
    price DECIMAL(10,2) NOT NULL,
    status common_status not null default 'active',
    created_by bigint NOT NULL,
    created_at timestamp without time zone not null,
    updated_by bigint,
    updated_at timestamp without time zone not null default current_timestamp
);
