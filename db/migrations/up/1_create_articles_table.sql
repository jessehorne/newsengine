CREATE TABLE articles (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255),
  body TEXT,
  created_by INT,
  created_at TIMESTAMP,
  updated_by INT,
  updated_at TIMESTAMP
);