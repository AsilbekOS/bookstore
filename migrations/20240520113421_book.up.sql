CREATE TABLE book (
    book_id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    author_id INT,
    publication_date DATE,
    isbn VARCHAR(255),
    b_description TEXT, 
    created_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES author (author_id)
  );

INSERT INTO book (title, author_id, publication_date, isbn, b_description) VALUES
('The Great Novel', 1, '2020-05-20', '123-4567890123', 'A fascinating tale of adventure and discovery.'),
('Mystery at the Mansion', 2, '2019-08-15', '123-4567890124', 'A thrilling mystery set in an old mansion.'),
('Science Fiction Story', 3, '2018-02-11', '123-4567890125', 'An exciting science fiction journey.'),
('Historical Novel', 4, '2021-09-30', '123-4567890126', 'A deep dive into historical events.'),
('Romantic Novel', 5, '2022-04-22', '123-4567890127', 'A touching romantic story.'),
('Thriller Novel', 6, '2021-12-05', '123-4567890128', 'An edge-of-your-seat thriller.'),
('Award-Winning Book', 7, '2020-03-14', '123-4567890129', 'A book that has won multiple awards.'),
('Famous Poetry', 8, '2019-08-08', '123-4567890130', 'A collection of famous poems.'),
('Mystery Series', 9, '2020-11-25', '123-4567890131', 'The latest in a popular mystery series.'),
('Young Adult Fiction', 10, '2021-01-18', '123-4567890132', 'A captivating young adult fiction book.');
