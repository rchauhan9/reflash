CREATE TABLE IF NOT EXISTS project_card (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    project_id UUID NOT NULL REFERENCES project(id),
    question TEXT NOT NULL,
    answer TEXT NOT NULL
);
