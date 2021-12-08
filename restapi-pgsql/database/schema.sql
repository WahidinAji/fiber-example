CREATE TABLE books(  
    id BIGSERIAL PRIMARY KEY,
    author VARCHAR(100) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description text NOT NULL,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE books IS 'books table';
COMMENT ON COLUMN books.author IS 'author column';
DROP TABLE books;