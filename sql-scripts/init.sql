CREATE TABLE notes (
  id SERIAL PRIMARY KEY,
  title VARCHAR(50) NOT NULL,
  content VARCHAR(255) NOT NULL
  );
  
INSERT INTO notes (title, content) VALUES ('First Note', 'This is my first note');
INSERT INTO notes (title, content) VALUES ('Second Note', 'This is my second note');
INSERT INTO notes (title, content) VALUES ('Third Note', 'This is my third note');