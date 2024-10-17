// executed 2024-10-17 20:46
CREATE TABLE articles (
  id TEXT PRIMARY KEY,
  title TEXT NOT NULL,
  body TEXT NOT NULL
);

// executed 2024-10-17 20:46
CREATE TABLE comments (
  id TEXT PRIMARY KEY,
  article_id TEXT NOT NULL,
  body TEXT NOT NULL,
  FOREIGN KEY (article_id) REFERENCES article(id)
);
