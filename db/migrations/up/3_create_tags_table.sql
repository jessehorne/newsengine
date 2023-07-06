CREATE TABLE tags (
  id SERIAL PRIMARY KEY,
  tag VARCHAR(100),
  article_id INT
);