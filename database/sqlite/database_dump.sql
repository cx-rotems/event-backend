PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "names" (
	"id"	INTEGER,
	"firstName"	TEXT NOT NULL,
	"lastName"	TEXT NOT NULL,
	"arrived"	BOOLEAN DEFAULT 0,
	PRIMARY KEY("id" AUTOINCREMENT)
);
INSERT INTO names VALUES(1,'rotem','sheps',1);
INSERT INTO sqlite_sequence VALUES('names',8);
COMMIT;
