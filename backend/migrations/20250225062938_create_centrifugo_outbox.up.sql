CREATE TABLE IF NOT EXISTS centrifugo_outbox (
     id BIGSERIAL PRIMARY KEY,
     method text NOT NULL,
     payload JSONB NOT NULL,
     partition INTEGER NOT NULL default 0,
     created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);