CREATE TABLE "users" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "email" varchar(100) NOT NULL,
  "hashed_password" text NOT NULL,
  "phone_number" varchar(20) NOT NULL,
  "is_admin" bool NOT NULL
);

CREATE TABLE "farmers" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "user_id" int
);

CREATE TABLE "buyers" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "address" varchar(255),
  "payment_method" varchar(20),
  "user_id" int
);

CREATE TABLE "farms" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "farmer_id" int,
  "address" varchar(255),
  "size" float,
  "government_id" int
);

CREATE TABLE "products" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "farmer_id" int,
  "price" decimal(10,2) NOT NULL,
  "quantity" int NOT NULL,
  "description" text,
  "category" int
);

CREATE TABLE "categories" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "name" varchar(100) NOT NULL
);

ALTER TABLE "farmers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "buyers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "farms" ADD FOREIGN KEY ("farmer_id") REFERENCES "farmers" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("farmer_id") REFERENCES "farmers" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("category") REFERENCES "categories" ("id");