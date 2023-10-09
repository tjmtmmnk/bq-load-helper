DROP SCHEMA IF EXISTS sample;
CREATE SCHEMA sample;
USE sample;

DROP TABLE IF EXISTS test;

CREATE TABLE test
(
    id BIGINT UNSIGNED NOT NULL,
    bn BINARY,
    b  BIT,
    bb BLOB,
    bl BOOL,
    c  CHAR,
    d  DATE,
    dt DATETIME,
    dc DECIMAL,
    db DOUBLE,
    e  ENUM ('a'),
    f  FLOAT,
    i  INTEGER,
    j  JSON,
    lb LONGBLOB,
    lt LONGTEXT,
    mb MEDIUMBLOB,
    mi MEDIUMINT,
    mt MEDIUMTEXT,
    s  SET ('a', 'b'),
    si SMALLINT,
    t  TEXT,
    tm TIME,
    ts TIMESTAMP,
    tb TINYBLOB,
    ti TINYINT,
    tt TINYTEXT,
    vb VARBINARY(20),
    vc VARCHAR(20),
    y  YEAR(4),
    PRIMARY KEY (`id`)
);

INSERT INTO test (id)
VALUES (uuid_short());
INSERT INTO test (id)
VALUES (uuid_short());
