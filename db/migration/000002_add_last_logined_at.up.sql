ALTER TABLE users ADD last_logined_at timestamptz NOT NULL DEFAULT (now())