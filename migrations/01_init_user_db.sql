-- +goose Up
create table users (
		id uuid primary key default gen_random_uuid(),
		name text,
		phone text,
		email text,
		password text,
		registration_date date
		-- user_rating int as sum of column in user_ride table
);

CREATE TYPE taxi AS ENUM ('econom', 'business');

create table user_ride (
		order_id uuid primary key,
		user_id uuid, 
		driver_id uuid,

		driver_name text,

		ride_from text,
		ride_to text,
		taxi_type taxi,

		-- ride_rating 
		foreign key(user_id) references users(id) 
);

-- +goose Down
DROP TABLE user_ride;
DROP TABLE users;
DROP TYPE taxi;