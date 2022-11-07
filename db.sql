CREATE TABLE "products"(
    "id" serial primary key NOT NULL,
    "name" VARCHAR,
    "price" DECIMAL(18, 2),
    "image_url" VARCHAR,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE "product_images"(
    "id" serial primary key NOT NULL,
    "image_url" varchar,
    "sequence_number" INTEGER,
    "product_id" INTEGER NOT NULL references products(id) on delete restrict
);