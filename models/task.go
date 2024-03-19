package models

type Task struct {
	Id string
}

// CREATE TABLE IF NOT EXISTS "task" (
//   "id" UUID NOT NULL PRIMARY KEY,
//   "lesson_id" UUID REFERENCES "lesson"("id"),
//   "group_id" UUID REFERENCES "group"("id"),
//   "task" varchar(255) NOT NULL,
//   "score" integer not NULL DEFAULT 0,
//   "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//   "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );