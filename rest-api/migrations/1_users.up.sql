DO $$ BEGIN
  CREATE EXTENSION pgcrypto;
EXCEPTION
  WHEN duplicate_object THEN null;
END $$;

CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    data JSON NOT NULL,
    role_id TEXT NOT NULL
);

CREATE UNIQUE INDEX role_id ON users (role_id);

CREATE TABLE roles (
  id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
  data JSON NOT NULL
);

INSERT INTO roles (id, data) VALUES ('00000000-0000-0000-0000-000000000001', '{"roles-name": "ADMIN", "created-time":1596516731}');
INSERT INTO users (id, data, role_id) VALUES ('00000000-0000-0000-0000-000000000002', '{"status": "ACTIVE", "username": "bambang", "email":"bambang@getnada.com"}', '00000000-0000-0000-0000-000000000001');