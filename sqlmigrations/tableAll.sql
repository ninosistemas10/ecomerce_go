-- Active: 1697216304905@@192.168.1.130@5432@BDEcommerce@public



CREATE TABLE users(
	id 		    	UUID 			NOT NULL,
	nombre	    	VARCHAR(100)	NOT NULL,
	email			VARCHAR(254)	NOT NULL,
	password		VARCHAR(100)	NOT NULL,
	is_admin		BOOLEAN 		NOT NULL DEFAULT FALSE,
	details	    	JSONB 	    	NOT NULL,
	create_at		INTEGER		NOT NULL  DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
	update_at 		INTEGER,
	CONSTRAINT users_id_pk PRIMARY KEY (id),
	CONSTRAINT users_email_uk UNIQUE (email)
);
CREATE TABLE products(
	id 			UUID			NOT NULL,
	product_name	VARCHAR(128)	NOT NULL,
	price			NUMERIC(10,2)	NOT NULL,
	images		JSONB			NOT NULL,
	description		TEXT			NOT NULL,
	features		JSONB			NOT NULL,
	created_at		INTEGER 		NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
	updated_at		INTEGER,
	CONSTRAINT products_id_pk PRIMARY KEY(id)
);

CREATE TABLE purchase_orders (
	id 			UUID NOT NULL,
	user_id 		UUID NOT NULL,
	products 		JSONB NOT NULL,
	created_at 		INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
	updated_at 		INTEGER,
	CONSTRAINT purchase_orders_id_pk PRIMARY KEY (id),
	CONSTRAINT purchase_orders_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);



CREATE TABLE invoices (
	id 			UUID NOT NULL,
	user_id 		UUID NOT NULL,
	purchase_order_id UUID NOT NULL,
	created_at 		INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
	updated_at 		INTEGER,
	CONSTRAINT invoices_id_pk PRIMARY KEY (id),
	CONSTRAINT invoices_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT invoices_purchase_order_id_fk FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);



CREATE TABLE invoice_details (
	id 			UUID NOT NULL,
	invoice_id 		UUID NOT NULL,
	product_id 		UUID NOT NULL,
	amount 		INTEGER NOT NULL,
	unit_price 		NUMERIC(10,2) NOT NULL,
	created_at 		INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
	updated_at 		INTEGER,
	CONSTRAINT invoice_details_id_pk PRIMARY KEY (id),
	CONSTRAINT invoice_details_invoice_id_fk FOREIGN KEY (invoice_id) REFERENCES invoices (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT invoice_details_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);

