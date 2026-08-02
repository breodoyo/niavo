CREATE TABLE organizations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    slug TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
ALTER TABLE organizations
	ADD COLUMN updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	ADD COLUMN deleted_at TIMESTAMPTZ;

