CREATE TABLE author (
    author_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    birth_date DATE,
    biography TEXT,
    created_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP
  );

INSERT INTO author (name, birth_date, biography) VALUES
('John Doe', '1980-05-15', 'John Doe is a prolific writer known for his novels.'),
('Jane Smith', '1975-07-20', 'Jane Smith has written numerous bestsellers.'),
('Michael Johnson', '1990-02-11', 'Michael Johnson is known for his thrillers.'),
('Emily Davis', '1985-09-30', 'Emily Davis writes historical fiction.'),
('David Brown', '1978-04-22', 'David Brown specializes in science fiction.'),
('Sarah Wilson', '1982-12-05', 'Sarah Wilson writes romantic novels.'),
('James Taylor', '1983-03-14', 'James Taylor is an award-winning author.'),
('Patricia Miller', '1979-08-08', 'Patricia Miller is a famous poet.'),
('Robert Martinez', '1987-11-25', 'Robert Martinez is known for his mystery novels.'),
('Linda Anderson', '1992-01-18', 'Linda Anderson writes young adult fiction.');
