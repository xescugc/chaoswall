package migrations

// V0Initial is the first migration
var V0Initial = Migration{
	Name: "Initial",
	SQL: `
		SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

		CREATE TABLE gyms (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255),
			canonical VARCHAR(30),

			UNIQUE(canonical)
		);

		CREATE TABLE walls (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255),
			canonical VARCHAR(30),
			gym_id INT UNSIGNED NOT NULL,

			CONSTRAINT fk__walls__gyms
				FOREIGN KEY (gym_id) REFERENCES gyms (id)
				ON DELETE CASCADE
				ON UPDATE RESTRICT,

			UNIQUE(gym_id, canonical)
		);

		CREATE TABLE holds (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
			x SMALLINT UNSIGNED,
			y SMALLINT UNSIGNED,
			wall_id INT UNSIGNED NOT NULL,

			CONSTRAINT fk__holds__walls
				FOREIGN KEY (wall_id) REFERENCES walls (id)
				ON DELETE CASCADE
				ON UPDATE RESTRICT,

			UNIQUE(wall_id, x, y)
		);

		CREATE TABLE routes (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255),
			canonical VARCHAR(30),
			description TEXT,
			type ENUM('boulder', 'lead'),
			wall_id INT UNSIGNED NOT NULL,

			CONSTRAINT fk__routes__walls
				FOREIGN KEY (wall_id) REFERENCES walls (id)
				ON DELETE CASCADE
				ON UPDATE RESTRICT,

			UNIQUE(wall_id, canonical)
		);

		CREATE TABLE routes_holds (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
			route_id INT UNSIGNED NOT NULL,
			hold_id INT UNSIGNED NOT NULL,

			CONSTRAINT fk__routes_holds__routes
				FOREIGN KEY (route_id) REFERENCES routes (id)
				ON DELETE CASCADE
				ON UPDATE RESTRICT,

			CONSTRAINT fk__routes_holds__holds
				FOREIGN KEY (hold_id) REFERENCES holds (id)
				ON DELETE CASCADE
				ON UPDATE RESTRICT
		);
	`,
}
