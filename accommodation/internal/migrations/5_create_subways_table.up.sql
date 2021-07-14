CREATE TABLE subways
(
    id             int PRIMARY KEY,
    location_id    int          NOT NULL,
    station_name   varchar(255) NOT NULL,
    line_color     varchar(255) NOT NULL,
    transport_type varchar(255) NOT NULL,
    transport_time varchar(255) NOT NULL,

    FOREIGN KEY (location_id) REFERENCES locations (id) ON DELETE cascade
);