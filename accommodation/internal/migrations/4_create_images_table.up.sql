CREATE TABLE images
(
    id               int PRIMARY KEY,
    accommodation_id int          NOT NULL,
    path             varchar(255) NOT NULL,

    FOREIGN KEY accommodation_id REFERENCES `accommodations`(id) ON DELETE cascade
);