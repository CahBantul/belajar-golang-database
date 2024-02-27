CREATE TABLE customer(
    id VARCHAR(100) NOT NULL,
    NAME VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
) ENGINE = Innodb;

DELETE FROM customer;

ALTER TABLE customer
    ADD COLUMN email VARCHAR(100),
    ADD COLUMN balance INT DEFAULT 0,
    ADD COLUMN rating DOUBLE DEFAULT 0.0,
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN birth_date DATE,
    ADD COLUMN married BOOLEAN DEFAULT false;

DESC customer;

INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES 
('noza', 'nozami', 'nozami@gmail.com', 1000, 90.0, '1992-08-10', true),
('aji', 'ajitama', 'ajitama@gmail.com', 2000, 80.0, '1994-08-10', false),
('tama', 'tama', 'tama@gmail.com', 10000, 95.0, '1996-10-18', true);

INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES 
('fardan', 'fardan', NULL, 1000, 90.0, NULL, true);

SELECT * FROM customer