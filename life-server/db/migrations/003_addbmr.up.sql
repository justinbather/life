CREATE TABLE bmr (
	id SERIAL PRIMARY KEY,
	user_id VARCHAR NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	total_calories INT NOT NULL,
	num_workouts INT NOT NULL
);
