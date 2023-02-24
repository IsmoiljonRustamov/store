CREATE TABLE IF NOT EXISTS "addresses" (
    "id" SERIAL PRIMARY KEY,
    "city" varchar(30) NOT NULL,
    "steet_name" varchar(30),
    "branch_id" INTEGER REFERENCES branches(id)
);