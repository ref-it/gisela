CREATE TABLE IF NOT EXISTS items (
    item_id             SERIAL PRIMARY KEY,
	item_title          TEXT,
	item_description    TEXT,
	item_condition      TEXT,
	item_category       INT,
	item_place          INT,
    item_group          INT,
	item_number         INT,
	item_image			TEXT,
	item_image_preview  TEXT,
	date_of_receipt     DATE,
	date_of_retirement  DATE,
	date_of_accounting  DATE,
	acquisition_cost    FLOAT,
	project             INT[] NOT NULL,
	receipt             INT[] NOT NULL,
	supplier            TEXT,
	guarantee_until     DATE,
	borrowable          BOOLEAN NOT NULL,
	parent_item         INT
);

CREATE TABLE IF NOT EXISTS categories (
    category_id             SERIAL PRIMARY KEY,
    category_title          VARCHAR NOT NULL,
    category_description    VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS places (
    place_id            SERIAL PRIMARY KEY,
    place_title         VARCHAR NOT NULL,
    place_description   VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS groups (
    group_id            SERIAL PRIMARY KEY,
    group_title         TEXT NOT NULL,
    group_description   TEXT NOT NULL,
    group_email         TEXT NOT NULL,
	group_roles			TEXT[] NOT NULL
);