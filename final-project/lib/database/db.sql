-- CREATE user TABLE
CREATE TABLE IF NOT EXISTS user_ (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    age INTEGER CHECK(age >= 8) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_user_username ON user_(username);
CREATE INDEX IF NOT EXISTS idx_user_email ON user_(email);

-- CREATE photo TABLE
CREATE TABLE IF NOT EXISTS photo (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    caption VARCHAR(100),
    url VARCHAR(50) NOT NULL,
    user_id INTEGER REFERENCES user_(id) ON DELETE CASCADE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- CREATE comment TABLE
CREATE TABLE IF NOT EXISTS comment (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INTEGER REFERENCES user_(id) ON DELETE CASCADE,
    photo_id INTEGER REFERENCES photo(id) ON DELETE CASCADE,
    message TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- CREATE social_media TABLE
CREATE TABLE IF NOT EXISTS social_media (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    url VARCHAR(50) NOT NULL,
    user_id INTEGER REFERENCES user_(id) ON DELETE CASCADE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);