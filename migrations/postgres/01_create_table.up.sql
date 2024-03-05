
CREATE TABLE IF NOT EXISTS "branches" (
  "id" uuid PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "address" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "teacher" (
  "id" uuid PRIMARY KEY,
  "full_name" varchar(255) NOT NULL,
  "email" varchar(255) NOT NULL,
  "age" int NOT NULL,
  "status" varchar(60) NOT NULL CHECK("status" IN ('active', 'inactive')) DEFAULT 'active',
  "login" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "admin" (
  "id" uuid PRIMARY KEY,
  "full_name" varchar(255) NOT NULL,
  "email" varchar(255) NOT NULL,
  "age" int NOT NULL,
  "status" varchar(60) NOT NULL CHECK("status" IN ('active', 'inactive')) DEFAULT 'active',
  "login" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "group" (
  "id" uuid PRIMARY KEY,
  "group_id" varchar(255) NOT NULL UNIQUE, -- GR-0000001, 
  "branch_id" uuid NOT NULL REFERENCES "branches"("id"),
  "teacher" uuid NOT NULL REFERENCES "teacher"("id"),
  "type" varchar(255) NOT NULL CHECK ("type" IN ('backend', 'frontend', 'mobile', 'devops', 'qa', 'pm', 'designer')),
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "student" (
    "id" uuid PRIMARY KEY,
    "full_name" varchar(255) NOT NULL,
    "email" varchar(255) NOT NULL,
    "age" int NOT NULL,
    "paid_sum" decimal(10, 2) NOT NULL DEFAULT 0,
    "status" varchar(60) NOT NULL CHECK("status" IN ('active', 'inactive')) DEFAULT 'active',
    "login" varchar(255) NOT NULL,
    "password" varchar(255) NOT NULL,
    "group_id" uuid REFERENCES "group"("id"),
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
    
); 



CREATE TABLE IF NOT EXISTS "schedule" (
  "id" UUID NOT NULL PRIMARY KEY,
  "group_id" UUID REFERENCES "group"("id"),
  "group_type" VARCHAR,
  "start_time" TIME,
  "end_time" TIME,
  "date" varchar(255),
  "branch_id" UUID REFERENCES "branches"("id"),
  "teacher_id" UUID REFERENCES "teacher"("id"),
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "lesson" (
  "id" UUID NOT NULL PRIMARY KEY,
  "schedule_id" UUID REFERENCES "schedule"("id"),
  "group_id" UUID REFERENCES "group"("id"),
  "from" DATE,
  "to" DATE,
  "theme" varchar(255) NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "tasks" (
    "id" UUID NOT NULL PRIMARY KEY,
    "lesson_id" UUID REFERENCES "lesson"("id"),
    "group_id" UUID REFERENCES "group"("id"),
    "task" varchar(255) NOT NULL,
    "score" integer not NULL DEFAULT 0,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
 

CREATE TABLE IF NOT EXISTS "payment" (
  "id" uuid PRIMARY KEY,
  "price" decimal(10, 2) NOT NULL,
  "student_id" uuid NOT NULL REFERENCES "student"("id"),
  "branch_id" uuid NOT NULL REFERENCES "branches"("id"),
  "admin_id" uuid NOT NULL REFERENCES "admin"("id"),
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);