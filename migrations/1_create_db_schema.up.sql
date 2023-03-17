CREATE TABLE IF NOT EXISTS  Account(
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  	OwnerName VARCHAR(50) NOT NULL,
  	Balance DECIMAL Not NULL,
  	Currency VARCHAR(20) Not NULL,
  	Deleted BOOLEAN Not NULL DEFAULT FALSE,
  	CreatedAt TIMESTAMPTZ NOt NULL
);


CREATE TABLE IF NOT EXISTS Entry(
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  	AccountId INT,
  	Amount DECIMAL Not NULL,
  	CreatedAt TIMESTAMPTZ NOt NULL,
    CONSTRAINT fk_account_id_constraint
      FOREIGN KEY(AccountId) 
        REFERENCES Account(ID)
);


CREATE TABLE IF NOT EXISTS Transfer(
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  	ToAccount INT,
  	FromAccount INT,
  	Amount DECIMAL Not NULL CHECK(Amount > 0),
  	Deleted BOOLEAN Not NULL DEFAULT FALSE,
  	CreatedAt TIMESTAMPTZ NOt NULL,
    CONSTRAINT fk_to_account_id_constraint
      FOREIGN KEY(ToAccount) 
        REFERENCES Account(ID),
    CONSTRAINT fk_from_account_id_constraint
        FOREIGN KEY (FromAccount)
        REFERENCES Account(ID)
);

create INDEX on Account(id);
create INDEX on Entry(AccountId);
create INDEX on Transfer(ToAccount);
create INDEX on Transfer(FromAccount);
create INDEX on Transfer(ToAccount, FromAccount);
