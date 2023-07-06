CREATE TABLE authors (
  id SERIAL PRIMARY KEY,
  article_id INT,
  user_id INT,
  created_at TIMESTAMP,
  created_by TIMESTAMP
);