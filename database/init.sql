CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(80) UNIQUE NOT NULL,
    password_hash VARCHAR(128) NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE
);

CREATE TABLE chores (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    user_id INTEGER REFERENCES users(id) NOT NULL
);

-- Insert some sample data
INSERT INTO users (username, password_hash, is_admin) VALUES
('admin', 'pbkdf2:sha256:150000$...', TRUE),
('kid1', 'pbkdf2:sha256:150000$...', FALSE),
('kid2', 'pbkdf2:sha256:150000$...', FALSE);

INSERT INTO chores (title, description, user_id) VALUES
('Clean room', 'Make your bed and tidy up', 2),
('Do dishes', 'Wash and dry the dishes after dinner', 2),
('Take out trash', 'Empty all trash bins and take to the curb', 3);