CREATE TABLE people (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        surname VARCHAR(255) NOT NULL,
                        patronymic VARCHAR(255),
                        age INT,
                        gender VARCHAR(255),
                        nationality VARCHAR(255)
);