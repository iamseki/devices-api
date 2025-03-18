CREATE TABLE IF NOT EXISTS
  devices (
    id SERIAL PRIMARY KEY,
    name TEXT,
    brand TEXT,
    state TEXT DEFAULT 'AVAILABLE' NOT NULL CHECK (state IN ('AVAILABLE', 'IN_USE', 'INACTIVE')),
    created_at timestamptz NOT NULL DEFAULT NOW()
  );
