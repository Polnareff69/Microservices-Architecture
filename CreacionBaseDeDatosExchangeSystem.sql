CREATE TABLE accounts (
    idaccount UUID PRIMARY KEY,
    iduser UUID NOT NULL,
    credits INTEGER,
    creationDate TIMESTAMP WITHOUT time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
    idtransaction UUID PRIMARY KEY,
    from_acc UUID NOT NULL,
    to_acc UUID NOT NULL,
    amount INTEGER,
    currency INTEGER,
	creationDate TIMESTAMP WITHOUT time zone DEFAULT CURRENT_TIMESTAMP
);

