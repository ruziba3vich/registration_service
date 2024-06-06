CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS admins (
    admin_id UUID DEFAULT uuid_generate_v4(),
    admin_name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS users (
    user_id UUID DEFAULT uuid_generate_v4(),
    username VARCHAR(255),
    data TEXT
);
