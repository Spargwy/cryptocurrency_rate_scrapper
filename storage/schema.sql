CREATE TABLE IF NOT EXISTS cryptocurrency_rate (
    raw json NOT NULL,
    display json NOT NULL,
    ID int NOT NULL AUTO_INCREMENT PRIMARY KEY
);
