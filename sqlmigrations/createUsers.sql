CREATE TABLE users(
	id 		UUID 			NOT NULL,
	nombre	VARCHAR(100)	NOT NULL,
	email		VARCHAR(254)	NOT NULL,
	password	VARCHAR(100)	NOT NULL,
	is_admin	BOOLEAN		NOT NULL DEFAULT FALSE,
	details	JSONB 		NOT NULL,
	create_at	INTEGER		NOT NULL  DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
	update_at 	INTEGER,

	CONSTRAINT users_id_pk PRIMARY KEY (id),
	CONSTRAINT users_email.uk UNIQUE (email)
);


COMMENT ON TABLE users IS 'Storage the admins and customers for the E-COMMERCE';
