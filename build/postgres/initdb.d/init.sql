CREATE FUNCTION set_updated_at() RETURNS OPAQUE AS '
    begin
        new.updated_at := ''now'';
        return new;
    end;
' LANGUAGE plpgsql;

CREATE TABLE users
(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_tri_users BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

INSERT INTO users (name, email, password) VALUES ('user1', 'user1@example.com', 'user1'), ('user2', 'user2@example.com', 'user2');

CREATE TABLE tasks
(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) on DELETE CASCADE
);

CREATE TRIGGER update_tri_users BEFORE UPDATE ON tasks FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

INSERT INTO tasks (user_id, title) VALUES (1, '洗い物'), (2, '洗濯');
