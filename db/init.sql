CREATE TABLE IF NOT EXISTS project(
    id      SERIAL PRIMARY KEY,
    name            VARCHAR(64),
    description     TEXT
);


CREATE TABLE IF NOT EXISTS service(
    id      SERIAL PRIMARY KEY,
    name            VARCHAR(64),
    lang            VARCHAR(64),
    focus           VARCHAR(64),
    project         INT,
    FOREIGN KEY (project) REFERENCES project(project_id)
    ON DELETE CASCADE
);
