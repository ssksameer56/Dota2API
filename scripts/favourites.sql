CREATE DATABASE `DotaDatabase`;

USE DotaDatabase;

CREATE TABLE Favourites (
    UserID INT NOT NULL,
    HeroID VARCHAR(200),
    PRIMARY KEY (UserID)
);

INSERT INTO Favourites VALUES(-1,'122,13,134')