CREATE TYPE OrderStatus AS ENUM('pending', 'delivered', 'cancelled');
CREATE TYPE PayStatus AS ENUM('card', 'cash');


CREATE TABLE IF NOT EXISTS "users"(
    "id" UUID NOT NULL PRIMARY KEY,
    "email" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "curiers"(
    "id" UUID NOT NULL PRIMARY KEY,
    "email" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

INSERT INTO "users"("id","email","password") VALUES('dbecf401-64b3-4b9b-829a-c8b061431286', 'isomiddinovbahodir7@gmail.com', '1234');
INSERT INTO "curiers"("id","email","password") VALUES('690d15b1-b3bf-416f-83e1-02b183ccb2f2', 'azam1222', '938791222');

CREATE TABLE IF NOT EXISTS "products"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "comment" VARCHAR NOT NULL,
    "price" VARCHAR NOT NULL,
    "product_img" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "orders"(
    "id" UUID NOT NULL PRIMARY KEY,
    "user_id" UUID NOT NULL REFERENCES "users"("id"),
    "curier_id" UUID NOT NULL REFERENCES  "curiers"("id"),
    "product_id" UUID NOT NULL REFERENCES  "products"("id"),
    "total_price" VARCHAR NOT NULL, 
    "status" OrderStatus DEFAULT 'pending', 
    "delivery_address" TEXT NOT NULL,  
    "payment_method" PayStatus NOT NULL, 
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "basket_user"(
    "id" UUID NOT NULL PRIMARY KEY,
    "order_id" UUID NOT NULL REFERENCES "orders"("id"),
    "user_id" UUID NOT NULL REFERENCES "users"("id"),
    "product_id" UUID NOT NULL REFERENCES "products"("id")
);


CREATE TABLE IF NOT EXISTS "addition"(
    "product_id" UUID NOT NULL REFERENCES "products"("id"),
    "user_id" UUID NOT NULL REFERENCES  "users"("id"),
    "order_id" UUID NOT NULL REFERENCES "orders"("id"),
    "thing" VARCHAR NOT NULL,
    "thing_price" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "history_user"(
    "id" UUID NOT NULL PRIMARY KEY,
    "user_id" UUID NOT NULL REFERENCES "users"("id"),
    "order_id" UUID NOT NULL REFERENCES "orders"("id"),
    "product_id" UUID NOT NULL REFERENCES "products"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "history_curier"(
    "id" UUID NOT NULL PRIMARY KEY,
    "curier_id" UUID NOT NULL REFERENCES "curiers"("id"),
    "order_id" UUID NOT NULL REFERENCES "orders"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);