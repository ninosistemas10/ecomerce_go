CREATE TABLE invoices (
	id				UUID			NOT NULL,
	user_id			UUID			NOT NULL,
	purchase_order_id		UUID			NOT NULL,
	create_at			INTEGER		NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
	update_at			INTEGER,

	CONSTRAINT invoice_id_pk PRIMARY KEY (id),
	CONSTRAINT invoice_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT invoices_purchase_order_id_fk FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders(id) ON UPDATE RESTRICT ON DELETE RESTRICT
);


COMMENT ON TABLE invoices IS 'Storage the head of the invoices'
