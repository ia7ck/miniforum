drop TABLE IF EXISTS users CASCADE;
drop TABLE IF EXISTS sessions;
drop TABLE IF EXISTS threads CASCADE;
drop TABLE IF EXISTS posts;

CREATE TABLE users (
  id          SERIAL PRIMARY KEY,
  screen_name VARCHAR(256) NOT NULL UNIQUE,
  password    VARCHAR(256) NOT NULL
);

CREATE TABLE sessions (
  id         SERIAL PRIMARY KEY,
  session_id VARCHAR(64) NOT NULL UNIQUE,
  user_id    INT REFERENCES users(id),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE threads (
  id         SERIAL PRIMARY KEY,
  content    TEXT,
  user_id    INT REFERENCES users(id), -- threadを作ったuser
  created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE posts (
  id         SERIAL PRIMARY KEY,
  body       TEXT,
  user_id    INT REFERENCES users(id),
  thread_id  INT REFERENCES threads(id),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL
);
