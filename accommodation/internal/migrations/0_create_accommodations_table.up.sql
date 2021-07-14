CREATE TABLE accommodations
(
    id          int PRIMARY KEY,
    status      varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    space       int NOT NULL,
    price       int NOT NULL,
    deposit     int NOT NULL,
    floor       int NOT NULL,
    max_floor   int NOT NULL,
    room_count  int NOT NULL
);

