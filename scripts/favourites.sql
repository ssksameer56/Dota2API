CREATE DATABASE `DotaDatabase`;

USE DotaDatabase;

CREATE TABLE Favourites (
    UserID INT NOT NULL,
    HeroID VARCHAR(200),
    PRIMARY KEY (UserID)
);
