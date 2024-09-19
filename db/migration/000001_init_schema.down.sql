ALTER TABLE "products" DROP CONSTRAINT products_category_fkey;
ALTER TABLE "products" DROP CONSTRAINT products_farmer_id_fkey;
ALTER TABLE "farms" DROP CONSTRAINT farms_farmer_id_fkey;
ALTER TABLE "buyers" DROP CONSTRAINT buyers_user_id_fkey;
ALTER TABLE "farmers" DROP CONSTRAINT farmers_user_id_fkey;

DROP TABLE IF EXISTS "categories";
DROP TABLE IF EXISTS "products";
DROP TABLE IF EXISTS "farms";
DROP TABLE IF EXISTS "buyers";
DROP TABLE IF EXISTS "farmers";
DROP TABLE IF EXISTS "users";
