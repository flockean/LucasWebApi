
CREATE TABLE IF NOT EXISTS service(
    service_id      INT PRIMARY KEY,
    name            VARCHAR(64),
    lang            VARCHAR(64),
    focus           VARCHAR(64)
);


CREATE TABLE IF NOT EXISTS project(

    project_id      INT PRIMARY KEY,
    name            VARCHAR(64),
    description     TEXT,
    service         INT,
    FOREIGN KEY (service) REFERENCES service(service_id)
    ON DELETE CASCADE
);
