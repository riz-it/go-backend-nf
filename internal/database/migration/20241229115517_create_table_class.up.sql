CREATE TABLE "classes" (
    "id" SERIAL NOT NULL,
    "name" VARCHAR(125) NOT NULL,
    "is_active" BOOLEAN NOT NULL DEFAULT true,
    "leader" INT,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "classes_pkey" PRIMARY KEY ("id")
);