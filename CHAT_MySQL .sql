CREATE database mystoredbfoodpanda;
CREATE USER 'user1' @'localhost' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user1' @'localhost';
USE mystoredbfoodpanda;
CREATE TABLE Users (
  UserName VARCHAR(30) NOT NULL PRIMARY KEY,
  Password VARCHAR(256),
  FirstName VARCHAR(30),
  LastName VARCHAR(30),
  Language VARCHAR(30)
);

