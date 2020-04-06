USE DATABASE covid_development;

CREATE TABLE IF NOT EXISTS counties (
  id INT PRIMARY KEY,
  name VARCHAR,
  zip_code VARCHAR,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);

CREATE INDEX IF NOT EXISTS zip_code_idx ON counties (zip_code);

CREATE TABLE IF NOT EXISTS states (
  id INT PRIMARY KEY,
  name VARCHAR,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS county_data (
  id INT PRIMARY KEY,
  date timestamp without time zone NOT NULL,
  county_id INT,
  cases INT,
  deaths INT,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS state_data (
  id INT PRIMARY KEY,
  date timestamp without time zone NOT NULL,
  state_id INT,
  cases INT,
  deaths INT,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
)
