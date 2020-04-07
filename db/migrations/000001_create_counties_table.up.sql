CREATE TABLE IF NOT EXISTS counties (
  id INT PRIMARY KEY,
  name VARCHAR,
  fips VARCHAR,
  state_id INT,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);

CREATE INDEX IF NOT EXISTS county_fips_idx ON counties (fips);

CREATE TABLE IF NOT EXISTS states (
  id INT PRIMARY KEY,
  name VARCHAR,
  fips INT,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);

CREATE INDEX IF NOT EXISTS state_fips_idx ON states (fips);

CREATE TABLE IF NOT EXISTS county_data (
  id INT PRIMARY KEY,
  date timestamp without time zone NOT NULL,
  county_id INT,
  cases INT,
  deaths INT,
  sha256_hash VARCHAR,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);

CREATE INDEX IF NOT EXISTS county_data_hash_idx ON county_data (sha256_hash);

CREATE TABLE IF NOT EXISTS state_data (
  id INT PRIMARY KEY,
  date timestamp without time zone NOT NULL,
  state_id INT,
  cases INT,
  deaths INT,
  sha256_hash VARCHAR,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);
CREATE INDEX IF NOT EXISTS state_data_hash_idx ON state_data (sha256_hash);
