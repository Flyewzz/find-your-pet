alter table lost
  add column textsearchable_index_col tsvector;

CREATE INDEX textsearch_idx ON lost USING GIN (textsearchable_index_col);

update lost
set textsearchable_index_col =
        setweight(to_tsvector('russian', coalesce(breed, '')), 'A') ||
        setweight(to_tsvector('russian', coalesce(description)), 'C');

alter table found
  add column textsearchable_index_col tsvector;

CREATE INDEX textsearch_idx_found ON found USING GIN (textsearchable_index_col);

update found
set textsearchable_index_col =
        setweight(to_tsvector('russian', coalesce(breed, '')), 'A') ||
        setweight(to_tsvector('russian', coalesce(description)), 'C');

CREATE FUNCTION lost_found_trigger() RETURNS trigger AS $$
begin
  new.textsearchable_index_col :=
        setweight(to_tsvector('russian', coalesce(new.breed,'')), 'A') ||
        setweight(to_tsvector('russian', coalesce(new.description,'')), 'C');
  return new;
end
$$ LANGUAGE plpgsql;

CREATE TRIGGER tsvectorupdatelost BEFORE INSERT OR UPDATE
  ON lost FOR EACH ROW EXECUTE PROCEDURE lost_found_trigger();
CREATE TRIGGER tsvectorupdatefound BEFORE INSERT OR UPDATE
  ON found FOR EACH ROW EXECUTE PROCEDURE lost_found_trigger();
