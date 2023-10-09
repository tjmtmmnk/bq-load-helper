DROP SCHEMA IF EXISTS sample;
CREATE SCHEMA sample;
USE sample;

DROP TABLE IF EXISTS test;

CREATE TABLE test
(
    id        BIGINT UNSIGNED NOT NULL,
    name      VARCHAR(40),
    is_bool   BOOL DEFAULT TRUE,
    small_int SMALLINT,
    i         integer,
    PRIMARY KEY (`id`)
);

INSERT INTO test (id, name)
VALUES (uuid_short(), 'Tanaka');
INSERT INTO test (id, name)
VALUES (uuid_short(), 'Tarou');
