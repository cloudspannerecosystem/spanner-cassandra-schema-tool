# Spanner Cassandra Schema Tool

This CLI tool streamlines Cassandra to Spanner schema migration by automating the following key steps:
1. Read the Cassandra DDL from a specified CQL file.
2. Translate those statements into Spanner-compatible DDL.
3. Apply the resulting schema to your target Spanner database.

## Note on the DDL Translation

* Only `CREATE TABLE` statements are supported.
* Keyspace name is optional. If it is specified, it should match the database name.
* Table options in the `CREATE TABLE` statements are not supported.
* Static columns are not supported.
* The following data types are not supported: Duration, Tuple, and Frozen
* Mapping of [Cassandra native type](https://cassandra.apache.org/doc/stable/cassandra/cql/types.html#native-types) to [Spanner Type](https://cloud.google.com/spanner/docs/reference/standard-sql/data-types#data_type_list):
    | Cassandra | Spanner    |
    | --------- | ---------- |
    | BIGINT    | INT64      |
    | INT       | INT64      |
    | SMALLINT  | INT64      |
    | TINYINT   | INT64      |
    | TIME      | INT64      |
    | COUNTER   | INT64      |
    | ASCII     | STRING(MAX)|
    | TEXT      | STRING(MAX)|
    | VARCHAR   | STRING(MAX)|
    | INET      | STRING(MAX)|
    | UUID      | STRING(MAX)|
    | TIMEUUID  | STRING(MAX)|
    | BOOLEAN   | BOOL       |
    | TIMESTAMP | TIMESTAMP  |
    | DATE      | DATE       |
    | FLOAT     | FLOAT32    |
    | DOUBLE    | FLOAT64    |
    | DECIMAL   | NUMERIC    |
    | VARINT    | NUMERIC    |
    | BLOB      | BYTES(MAX) |

* Mapping of [Cassandra collection type](https://cassandra.apache.org/doc/stable/cassandra/cql/types.html#collections) to [Spanner Type](https://cloud.google.com/spanner/docs/reference/standard-sql/data-types#data_type_list):
    | Cassandra  | Spanner      |
    | ---------- | ------------ |
    | LIST\<cassandra_type\> | ARRAY\<spanner_type\>  |
    | SET\<cassandra_type\>  | ARRAY\<spanner_type\>  |
    | MAP<type, type> | JSON    |

* The original Cassandra type will be added to the options in each column definition in the translated `CREATE TABLE` statement. See the example section below for details.
* Cassandra (composite) partition key and clustering columns are combined to form the composite primary key in Spanner. See the example section below for details.
* The translation process primarily focuses on **data type** and **primary key clause** conversion. It **does not** perform comprehensive semantic validation of the DDL. For instance, translation of the CREATE TABLE stmt will succeed even without a primary key, but Spanner will return an error upon application.

## Requirements

- **Go**: Ensure that Go is installed on your system. The required go version is 1.25+.
- **Google Cloud SDK**: Ensure `gcloud` is installed and authenticated with proper permissions.
- **Spanner Database**: Ensure the target Spanner database is created.

## Setup

### Clone the Repository

```bash
git clone https://github.com/cloudspannerecosystem/spanner-cassandra-schema-tool.git
cd spanner-cassandra-schema-tool
```

### Install Dependencies

Ensure that all necessary Go modules are installed:

```bash
go mod download
```

### Set Up Google Cloud Credentials

This tool uses [Application Default Credentials](https://cloud.google.com/docs/authentication/production?hl=en#providing_credentials_to_your_application) as the credential source for connecting to Spanner databases. Set the `GOOGLE_APPLICATION_CREDENTIALS` environment variable to the path of your service account key file.

```bash
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/your/service-account-file.json"
```

## Usage

### SYNOPSIS

```bash
go run schema_converter.go \
    --project <PROJECT_ID> \
    --instance <INSTANCE_ID> \
    --database <DATABASE_ID> \
    --cql <PATH_TO_CQL_FILE> \
    [--dry-run]
```

- `--project`: Google Cloud project ID.
- `--instance`: Spanner instance ID.
- `--database`: Spanner database ID.
- `--cql`: Path to the CQL file containing DDL statements.
- `[--dry-run]`: Print the converted DDL statements without applying to the Spanner database. This flag is optional.

### EXAMPLE

#### Command
```bash
go run schema_converter.go \
    --project cassandra-to-spanner \
    --instance spanner-instance-dev \
    --database testdb \
    --cql /path/to/cql-file.cql \
    --dry-run
```

#### Cassandra DDL in the input CQL file
```sql
CREATE TABLE IF NOT EXISTS testdb.example_table (
    id UUID PRIMARY KEY,
    ascii_value ASCII,
    bigint_value BIGINT,
    blob_value BLOB,
    boolean_value BOOLEAN,
    counter_value COUNTER,
    decimal_value DECIMAL,
    double_value DOUBLE,
    float_value FLOAT,
    inet_value INET,
    int_value INT,
    smallint_value SMALLINT,
    text_value TEXT,
    timestamp_value TIMESTAMP,
    timeuuid_value TIMEUUID,
    tinyint_value TINYINT,
    varchar_value VARCHAR,
    varint_value VARINT,
    list_value LIST<TEXT>,
    map_value MAP<TEXT, INT>,
    set_value SET<UUID>,
    date_value DATE,
    time_value TIME
);

CREATE TABLE IF NOT EXISTS sensor_data (
    sensor_id UUID,
    reading_time TIMESTAMP,
    location TEXT,
    temperature DECIMAL,
    humidity DECIMAL,
    pressure DECIMAL,
    PRIMARY KEY ((sensor_id, location), reading_time)
);
```

#### Translated Spanner DDL
```sql
CREATE TABLE IF NOT EXISTS example_table (
 id STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'uuid'),
 ascii_value STRING(MAX) OPTIONS (cassandra_type = 'ascii'),
 bigint_value INT64 OPTIONS (cassandra_type = 'bigint'),
 blob_value BYTES(MAX) OPTIONS (cassandra_type = 'blob'),
 boolean_value BOOL OPTIONS (cassandra_type = 'boolean'),
 counter_value INT64 OPTIONS (cassandra_type = 'counter'),
 decimal_value NUMERIC OPTIONS (cassandra_type = 'decimal'),
 double_value FLOAT64 OPTIONS (cassandra_type = 'double'),
 float_value FLOAT32 OPTIONS (cassandra_type = 'float'),
 inet_value STRING(MAX) OPTIONS (cassandra_type = 'inet'),
 int_value INT64 OPTIONS (cassandra_type = 'int'),
 smallint_value INT64 OPTIONS (cassandra_type = 'smallint'),
 text_value STRING(MAX) OPTIONS (cassandra_type = 'text'),
 timestamp_value TIMESTAMP OPTIONS (cassandra_type = 'timestamp'),
 timeuuid_value STRING(MAX) OPTIONS (cassandra_type = 'timeuuid'),
 tinyint_value INT64 OPTIONS (cassandra_type = 'tinyint'),
 varchar_value STRING(MAX) OPTIONS (cassandra_type = 'varchar'),
 varint_value NUMERIC OPTIONS (cassandra_type = 'varint'),
 list_value ARRAY<STRING(MAX)> OPTIONS (cassandra_type = 'list<text>'),
 map_value JSON OPTIONS (cassandra_type = 'map<text,int>'),
 set_value ARRAY<STRING(MAX)> OPTIONS (cassandra_type = 'set<uuid>'),
 date_value DATE OPTIONS (cassandra_type = 'date'),
 time_value INT64 OPTIONS (cassandra_type = 'time'),
) PRIMARY KEY (id)
CREATE TABLE IF NOT EXISTS sensor_data (
 sensor_id STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'uuid'),
 reading_time TIMESTAMP NOT NULL OPTIONS (cassandra_type = 'timestamp'),
 location STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),
 temperature NUMERIC OPTIONS (cassandra_type = 'decimal'),
 humidity NUMERIC OPTIONS (cassandra_type = 'decimal'),
 pressure NUMERIC OPTIONS (cassandra_type = 'decimal'),
) PRIMARY KEY (sensor_id, location, reading_time)
```
