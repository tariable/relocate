CREATE TABLE options
(
    id                int PRIMARY KEY,
    rentTerm          varchar(255) DEFAULT 'long',
    wcType            varchar(255) DEFAULT 'combined',
    bathType          varchar(255) DEFAULT 'shower',
    balcony           boolean      DEFAULT false,
    isAllowedPets     boolean      DEFAULT false,
    isAllowedSmoking  boolean      DEFAULT false,
    isAllowedChildren boolean      DEFAULT false,
    elevator          boolean      DEFAULT false,
    conditioner       boolean      DEFAULT false,
    fridge            boolean      DEFAULT false,
    dishWasher        boolean      DEFAULT false,
    laundry           boolean      DEFAULT false,
    internet          boolean      DEFAULT false,
    tv                boolean      DEFAULT false
);