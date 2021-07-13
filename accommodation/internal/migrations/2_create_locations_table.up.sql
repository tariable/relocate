CREATE TABLE locations
(
    id       int PRIMARY KEY,
    long     float        NOT NULL,
    lat      float        NOT NULL,
    street   varchar(255) NOT NULL,
    house    int          NOT NULL,
    building varchar(255) NOT NULL,
    flat     int          NOT NULL
);