CREATE TABLE "profile" (
    "id" serial NOT NULL PRIMARY KEY,
    "title" varchar(25) NOT NULL,
    "name" varchar(50) NOT NULL,
    "description" text NOT NULL,
    "linkedin" varchar,
    "github" varchar,
    "code_wars" varchar,
    "certs" text[][],
    "technologies" text[],
    "img" varchar,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp
);

CREATE TABLE "experience" (
    "id" serial NOT NULL PRIMARY KEY,
    "profile_id" int REFERENCES profile(id),
    "position" varchar(50) NOT NULL,
    "company" varchar(50) NOT NULL,
    "time" varchar(25) NOT NULL,
    "description" text,
    "img" varchar,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp
);

CREATE TABLE "projects" (
    "id" serial NOT NULL PRIMARY KEY,
    "profile_id" int REFERENCES profile(id),
    "title" varchar(50) NOT NULL,
    "date" date NOT NULL,
    "description" text,
    "language" varchar(25) NOT NULL,
    "repo_link" varchar NOT NULL,
    "site" varchar,
    "img" varchar,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp
);

CREATE TABLE "education" (
    "id" serial NOT NULL PRIMARY KEY,
    "profile_id" int REFERENCES profile(id),
    "institution" varchar(50) NOT NULL,
    "degree" varchar NOT NULL,
    "time" varchar(50) NOT NULL,
    "description" text,
    "img" varchar,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp
);

CREATE TABLE "certifications" (
    "id" serial NOT NULL PRIMARY KEY,
    "profile_id" int REFERENCES profile(id),
    "title" varchar NOT NULL,
    "institution" varchar(25) NOT NULL,
    "time" varchar(50) NOT NULL,
    "credential_id" varchar NOT NULL,
    "link" varchar,
    "img" varchar,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp
);
