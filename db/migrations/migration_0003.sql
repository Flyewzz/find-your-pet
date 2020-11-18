DROP TRIGGER tsvectorupdatelost on lost;
DROP TRIGGER tsvectorupdatefound on found;

DROP FUNCTION lost_found_trigger;

update lost
set textsearchable_index_col =
        setweight(to_tsvector('russian', coalesce(breed, '')), 'A') ||
	setweight(to_tsvector('russian', coalesce(address, '')), 'A') ||
        setweight(to_tsvector('russian', coalesce(description, '')), 'C');

update found
set textsearchable_index_col =
        setweight(to_tsvector('russian', coalesce(breed, '')), 'A') ||
	setweight(to_tsvector('russian', coalesce(address, '')), 'A') ||
        setweight(to_tsvector('russian', coalesce(description, '')), 'C');

CREATE FUNCTION lost_found_trigger() RETURNS trigger AS $$
begin
  new.textsearchable_index_col :=
        setweight(to_tsvector('russian', coalesce(new.breed,'')), 'A') ||
	setweight(to_tsvector('russian', coalesce(new.address, '')), 'A') ||
        setweight(to_tsvector('russian', coalesce(new.description,'')), 'C');
  return new;
end
$$ LANGUAGE plpgsql;

CREATE TRIGGER tsvectorupdatelost BEFORE INSERT OR UPDATE
  ON lost FOR EACH ROW EXECUTE PROCEDURE lost_found_trigger();
CREATE TRIGGER tsvectorupdatefound BEFORE INSERT OR UPDATE
  ON found FOR EACH ROW EXECUTE PROCEDURE lost_found_trigger();
