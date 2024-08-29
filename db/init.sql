
CREATE TABLE IF NOT EXISTS service(
    service_id      SERIAL PRIMARY KEY,
    name            VARCHAR(64),
    lang            VARCHAR(64),
    focus           VARCHAR(64),
    project         INT,
    FOREIGN KEY (service) REFERENCES service(service_id)
    ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS project(

    project_id      SERIAL PRIMARY KEY,
    name            VARCHAR(64),
    description     TEXT,
    logo            bytea
);
