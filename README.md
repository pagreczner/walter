Walter is a utility to quickly generate PostgreSQL type change migrations.

## Usage
```
walter <type_name> <table_name> <col_name> [...enum_values]
```


### Usage without Enum's passed in.
```
$ walter my_type some_table potato

ALTER TYPE my_type RENAME TO __my_type;
CREATE TYPE my_type AS ENUM();

ALTER TABLE some_table RENAME COLUMN potato to _potato;
ALTER TABLE some_table ADD potato my_type NOT NULL DEFAULT '';
UPDATE some_table SET potato = _potato::text::my_type;
ALTER TABLE some_table DROP COLUMN _potato;
DROP TYPE __my_type;
```

### Usage With Enum's
```
$ walter my_type some_table potato ONE TWO THREE

ALTER TYPE my_type RENAME TO __my_type;
CREATE TYPE my_type AS ENUM('ONE','TWO','THREE');

ALTER TABLE some_table RENAME COLUMN potato to _potato;
ALTER TABLE some_table ADD potato my_type NOT NULL DEFAULT 'ONE';
UPDATE some_table SET potato = _potato::text::my_type;
ALTER TABLE some_table DROP COLUMN _potato;
DROP TYPE __my_type;
```