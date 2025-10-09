CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_shop_owner BOOLEAN DEFAULT FALSE
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

/*
DB Data Types
-------------
SERIAL - 1, 2, 3, 4, ..... (32 bits)
BIGSERIAL - (64 bits)

SMALLINT - 16 bits
INT - 32 bits
BIGINT - 64 bits

REAL - 32 bits - 10.32
DOUBLE PRECISION - 64 bits

CHAR(n) - Fixed upto n - CHAR(6) => 'a      ' (5xspace)
VARCHAR(n) - Not Fixed upto n
TEXT  -> unlimited

BOOLEAN -> true/false

TIME - without time zone -> Time of the day (9:32)
DATE - 2025-09-09
TIMESTAMP - Without time zone - 2025-09-09 9:32.43
TIMESTAMP WITH TIME ZONE - 2025-09-09 9:32.43+6
*/