\c olx_scrapper_db

CREATE TABLE appartments(
  id serial,
  offered_by text,
  surface integer,
  partitioning text,
  year_of_construction text,
  floor text,
  no_rooms int,
  description int,
  images text[],
  price int
)
