CREATE TABLE IF NOT EXISTS
  devices (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    brand TEXT NOT NULL,
    state TEXT NOT NULL DEFAULT 'AVAILABLE' CHECK (state IN ('AVAILABLE', 'IN_USE', 'INACTIVE')),
    creation_time timestamptz NOT NULL DEFAULT NOW()
  );
