DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  address_id VARCHAR(255) DEFAULT NULL NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DO $$ 
DECLARE 
  tbl TEXT;
BEGIN
  FOR tbl IN 
      SELECT tablename FROM pg_tables 
      WHERE schemaname = 'public' AND tablename IN ('users') 
  LOOP
      EXECUTE format(
          'CREATE TRIGGER trigger_%I_updated_at
          BEFORE UPDATE ON %I
          FOR EACH ROW
          EXECUTE FUNCTION update_timestamp()', 
          tbl, tbl
      );
  END LOOP;
END $$;

INSERT INTO users (email, name, password, address_id) VALUES
('luqman@mail.com', 'Luqman', '$2a$10$ySQb0xIYeyOuGR2.rADgcO8NGjSrdfVidcutjsIvFFnvLFaVaNLZO', '67e122fe715004e3bbdce21e'),
('yosua@mail.com', 'Yosua', '$2a$10$ySQb0xIYeyOuGR2.rADgcO8NGjSrdfVidcutjsIvFFnvLFaVaNLZO', '67e122fe715004e3bbdce21e'),
('nafatul@mail.com', 'Nafatul', '$2a$10$ySQb0xIYeyOuGR2.rADgcO8NGjSrdfVidcutjsIvFFnvLFaVaNLZO', '67e122fe715004e3bbdce21e');
