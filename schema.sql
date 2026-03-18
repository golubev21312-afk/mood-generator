CREATE TABLE mood_requests (
  id         SERIAL PRIMARY KEY,
  user_input TEXT NOT NULL,
  mood_label VARCHAR(50),
  energy     INTEGER,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE mood_results (
  id           SERIAL PRIMARY KEY,
  request_id   INTEGER REFERENCES mood_requests(id),
  palette      JSONB,
  quote        TEXT,
  quote_author VARCHAR(100),
  tracks       JSONB,
  created_at   TIMESTAMP DEFAULT NOW()
);
