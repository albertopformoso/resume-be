#!/bin/bash
set -e
source .env

echo -e "\n[+] populating profile table..."
psql "${DATABASE_URL}" -c "\COPY profile FROM './data/profile.csv' DELIMITER ',' CSV HEADER"

echo -e "\n[+] populating experience table..."
psql "${DATABASE_URL}" -c "\COPY experience FROM './data/experience.csv' DELIMITER ',' CSV HEADER"

echo -e "\n[+] populating projects table..."
psql "${DATABASE_URL}" -c "\COPY projects FROM './data/projects.csv' DELIMITER ',' CSV HEADER"

echo -e "\n[+] populating education table..."
psql "${DATABASE_URL}" -c "\COPY education FROM './data/education.csv' DELIMITER ',' CSV HEADER"

echo -e "\n[+] populating certifications table..."
psql "${DATABASE_URL}" -c "\COPY certifications FROM './data/certifications.csv' DELIMITER ',' CSV HEADER"
