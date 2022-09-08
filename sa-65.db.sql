BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "users" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"email"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "activities" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"acname"	text,
	"user_id"	integer,
	CONSTRAINT "fk_users_activitys" FOREIGN KEY("user_id") REFERENCES "users"("id"),
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "students" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"sname"	text,
	"surname"	text,
	"saddress"	text,
	"sparent"	text,
	"sidentity_card"	integer,
	"sgrade"	real,
	"sphone_num"	text,
	"seducation_qualification"	text,
	"sgraduated"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "ac_his" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"achour"	integer,
	"activity_id"	integer,
	"user_id"	integer,
	"student_id"	integer,
	CONSTRAINT "fk_activities_ac_hiss" FOREIGN KEY("activity_id") REFERENCES "activities"("id"),
	CONSTRAINT "fk_students_ac_hiss" FOREIGN KEY("student_id") REFERENCES "students"("id"),
	CONSTRAINT "fk_users_ac_hiss" FOREIGN KEY("user_id") REFERENCES "users"("id"),
	PRIMARY KEY("id")
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_users_email" ON "users" (
	"email"
);
CREATE INDEX IF NOT EXISTS "idx_users_deleted_at" ON "users" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_activities_deleted_at" ON "activities" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_students_deleted_at" ON "students" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_ac_his_deleted_at" ON "ac_his" (
	"deleted_at"
);
COMMIT;

