CREATE TABLE IF NOT EXISTS courses(
    id UUID PRIMARY KEY NOT NULL,
    is_active BOOLEAN NOT NULL,
    course_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS modules(
    id UUID PRIMARY KEY NOT NULL,
    course_id UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    experience_level SMALLINT NOT NULL
);

CREATE TABLE IF NOT EXISTS material(
    id UUID PRIMARY KEY NOT NULL,
    module_id UUID NOT NULL REFERENCES courses(id) ON
    DELETE CASCADE,
    description TEXT NOT NULL,
    explanation TEXT,
    object_url TEXT
);
