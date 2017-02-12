package main

import (
	"log"
)

const (
	dbVersion = "0.1"
)

func prepareDB() {
	checkMeta()
	createTables()
}

func checkMeta() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Meta (version TEXT)`)
	if err != nil { log.Fatal(err) }
	rows, err := db.Query(`SELECT version FROM Meta`)
	if err != nil { log.Fatal(err) }
	var version string
	cnt := 0
	for rows.Next() {
		cnt++
		if err := rows.Scan(&version); err != nil { log.Fatal(err) }
	}
	rows.Close()
	if cnt > 1 {
		log.Fatalf("Meta table has more than one row (%d)\n", cnt)
	} else if cnt == 0 {
		_, err := db.Exec(
			`INSERT INTO Meta (version) VALUES (` + dbVersion + `)`)
		if err != nil { log.Fatal(err) }
	} else if version != dbVersion {
		log.Fatalf("Database version (%s) is not supported\n", version)
	}
}

func createTables() {
	queries := []struct{name, query string}{
		{"Objects",
			`oid INT NOT NULL PRIMARY KEY, ` +
			`comment TEXT NOT NULL`},
		{"ContactOwner",
			`oid INT NOT NULL PRIMARY KEY, ` +
			`FOREIGN KEY (oid) REFERENCES Objects(oid) ON DELETE CASCADE`},
		{"Persons",
			`oid INT NOT NULL PRIMARY KEY, ` +
			`FOREIGN KEY (oid) REFERENCES ContactOwner(oid) ` +
				`ON DELETE CASCADE, ` +
			`birthdate DATE NULL, ` +
			`sex: ENUM('M', 'F') NULL`},
		{"NameTypes",
			`id INT NOT NULL PRIMARY KEY, ` +
			`name TEXT NOT NULL, ` +
			`canPat      TEXT NOT NULL, ` +  // canonical
			`canShortPat TEXT NOT NULL, ` +  //   short
			`offPat      TEXT NOT NULL, ` +  // official
			`offShortPat TEXT NOT NULL, ` +  //   short
			`dictPat     TEXT NOT NULL, ` +  // dictionary
			}
		{"Names",
			`oid INT NOT NULL, ` +
			`FOREIGN KEY (oid) REFERENCES Objects(oid) ` +
				`ON DELETE CASCADE, ` +
			`since DATE NOT NULL, ` +
			`PRIMARY KEY(oid, since), ` +
			`components TEXT NOT NULL, ` +
			`typeId INT NOT NULL, ` +
			`FOREIGN KEY (typeId) REFERENCES NameTypes(id) ` +
				`ON DELETE RESTRICT`}
	}
	_, err := db.Exec(
}
