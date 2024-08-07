CREATE TABLE users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE statistics(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    activity VARCHAR(255) NOT NULL,
    comment TEXT
);

CREATE TABLE links(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users ON DELETE CASCADE,
    ref TEXT NOT NULL,
    description TEXT,
    UNIQUE(user_id, ref)
);

CREATE TABLE categories(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    UNIQUE(user_id, name)
);

CREATE TABLE links_categories(
    link_id UUID REFERENCES links ON DELETE CASCADE,
    category_id UUID REFERENCES categories ON DELETE CASCADE,
    UNIQUE(link_id, category_id)
);
