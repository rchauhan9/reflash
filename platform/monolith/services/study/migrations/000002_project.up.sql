CREATE TABLE IF NOT EXISTS project (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    icon TEXT,
    PRIMARY KEY (id)
);