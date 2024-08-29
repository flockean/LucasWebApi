CREATE TABLE IF NOT EXISTS project(
    project_id      SERIAL PRIMARY KEY,
    name            VARCHAR(64),
    description     TEXT,
    logo            bytea
);


CREATE TABLE IF NOT EXISTS service(
    service_id      SERIAL PRIMARY KEY,
    name            VARCHAR(64),
    lang            VARCHAR(64),
    focus           VARCHAR(64),
    project         INT,
    FOREIGN KEY (project) REFERENCES project(project_id)
    ON DELETE CASCADE
);
