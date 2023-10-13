CREATE purchase_order (
	id 		UUID 		NOT NULL,
	user_id	UUID	 	NOT NULL,
	products	JSONB		NOT NULL,
	created_at	INTEGER	NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
	update_at	INTEGER,

	CONSTRAINT purchase_order_id_pk PRIMARY KEY (id),
	CONSTRAINT purchase_orders_user_id_fk FOREIGN KEY (user_id) REFERENCE users(id) ON UPDATE RESTRICT ON DELETE RESTRICT
);

