CREATE TABLE announcements
(
    id            int PRIMARY KEY,
    accommodation int NOT NULL,
    option        int NOT NULL,
    location      int NOT NULL,
    created_at    timestamp DEFAULT current_timestamp,
    updated_at    timestamp DEFAULT current_timestamp,

    FOREIGN KEY (accommodation) REFERENCES accommodations (id) ON DELETE cascade,
    FOREIGN KEY (option) REFERENCES options (id) ON DELETE cascade,
    FOREIGN KEY (location) REFERENCES locations (id) ON DELETE cascade
);