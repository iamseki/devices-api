CREATE TABLE IF NOT EXISTS
  devices (
    id SERIAL PRIMARY KEY,
    name TEXT,
    brand TEXT,
    state TEXT DEFAULT 'AVAILABLE' NOT NULL CHECK (state IN ('AVAILABLE', 'IN_USE', 'INACTIVE')),
    creation_time timestamptz NOT NULL DEFAULT NOW()
  );
