CREATE DATABASE bookstore_n9;

CREATE TABLE author (
    author_id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    bith_date DATE,
    biography TEXT,
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) NOT NULL
  );

INSERT INTO author (first_name, birth_date, biography, created_at, updated_at) VALUES ($1, $2, $3, $4, $5);

SELECT * FROM author;

SELECT * FROM author WHERE author_id = $1;

DELETE FROM author WHERE author_id = $1;

UPDATE author
	    SET name = $1, birth_date = $2, biography = $3, updated_at = $4
    	WHERE author_id = $5
    	RETURNING author_id, name, birth_date, biography, created_at, updated_at;

CREATE TABLE book (
    book_id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    author_id INT,
    publication_date DATE,
    isbn VARCHAR(255),
    b_description TEXT, 
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) NOT NULL,
    FOREIGN KEY (author_id) REFERENCES author (author_id)
  );

DELETE FROM book WHERE book_id = $1;

SELECT * FROM book;

SELECT * FROM book WHERE book_id = $1

UPDATE book
	    SET title = $1, publication_date = $2, isbn = $3, b_description = $4, updated_at = $5
    	WHERE book_id = $6
    	RETURNING book_id, title, author_id, publication_date, isbn, b_description, created_at, updated_at