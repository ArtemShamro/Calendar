CREATE TABLE events (
    id              SERIAL,
    owner           TEXT,
    created_at      DATE,
    name            TEXT,
    description     TEXT,
    when_occurs     DATE,
    when_remind     DATE
);