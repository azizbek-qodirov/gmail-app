DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_type') THEN
        CREATE TYPE gender_type AS ENUM ('male', 'female');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'message_type') THEN
        CREATE TYPE message_type AS ENUM ('to', 'cc', 'bcc');
    END IF;

END $$;


CREATE TABLE IF NOT EXISTS users 
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    dob DATE NOT NULL,
    gender gender_type NOT NULL,
    email VARCHAR(320) UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    pfp_url VARCHAR(255),

    created_at TIMESTAMP,
    confirmed_at TIMESTAMP,
    is_confirmed BOOLEAN DEFAULT FALSE,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS outbox
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    subject VARCHAR(1000),
    body TEXT,
    attachment_ids UUID[],
    sender_id UUID REFERENCES users(id),
    receiver_to_emails VARCHAR(320)[],
    receiver_cc_emails VARCHAR(320)[],
    receiver_bcc_emails VARCHAR(320)[],
    
    is_draft BOOLEAN DEFAULT FALSE,
    is_archived BOOLEAN DEFAULT FALSE,
    is_starred BOOLEAN DEFAULT FALSE,

    sent_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS inbox 
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    outbox_id UUID REFERENCES outbox(id),
    receiver_id UUID REFERENCES users(id),
    type message_type NOT NULL,
    is_spam BOOLEAN DEFAULT FALSE,
    is_archived BOOLEAN DEFAULT FALSE,
    is_starred BOOLEAN DEFAULT FALSE,

    read_at INT DEFAULT 0,
    deleted_at INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS attachments
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    file_url VARCHAR(255) NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_size FLOAT NOT NULL,
    mime_type VARCHAR(255) NOT NULL,
    user_id UUID REFERENCES users(id),
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);