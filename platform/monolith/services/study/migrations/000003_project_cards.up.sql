CREATE TABLE IF NOT EXISTS project_cards (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    project_id UUID NOT NULL REFERENCES projects(id),
    question TEXT NOT NULL,
    answer TEXT NOT NULL
);
