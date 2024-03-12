
CREATE TABLE IF NOT EXISTS classes (
                                       id SERIAL PRIMARY KEY,
                                       class_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     role VARCHAR(50) NOT NULL,
                                     username VARCHAR(255) UNIQUE NOT NULL,
                                     password VARCHAR(255) NOT NULL,
                                     f_name VARCHAR(255) NOT NULL,
                                     s_name VARCHAR(255) NOT NULL,
                                     class_id INT REFERENCES classes(id),
                                     points INT DEFAULT 100 CHECK (points >= 0)
);

CREATE TABLE IF NOT EXISTS points_rules (
                                            id SERIAL PRIMARY KEY,
                                            rule_name VARCHAR(255) NOT NULL,
                                            rule_description VARCHAR(255),
                                            rule_point INT NOT NULL,
                                            rule_type VARCHAR(10) CHECK (rule_type IN ('positive', 'negative'))
);

CREATE TABLE IF NOT EXISTS points ( 
                                      id SERIAL PRIMARY KEY,
                                      user_id INT REFERENCES users(id),
                                      awarded_by INT REFERENCES users(id),
                                      rule_id INT REFERENCES points_rules(id),
                                      points INT NOT NULL,
                                      reason VARCHAR(255),
                                      timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
