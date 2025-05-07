// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package translator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSpannerStmt_CreateTable(t *testing.T) {
	tests := []struct {
		name                string
		cqlStmt             string
		expectedSpannerStmt string
		expectError         bool
		expectedErrorMsg    string
	}{
		{
			name: "Only Text type",
			cqlStmt: `CREATE TABLE ks.t_test (
							pk_text  TEXT PRIMARY KEY,
							col_text text
						)`,
			expectedSpannerStmt: "CREATE TABLE t_test (\n" +
				" pk_text STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_text STRING(MAX) OPTIONS (cassandra_type = 'text'),\n" +
				") PRIMARY KEY (pk_text)",
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Has IF NOT EXISTS",
			cqlStmt: `CREATE TABLE IF NOT EXISTS ks.t_test (
									pk_text  TEXT PRIMARY KEY,
									col_text text
								)`,
			expectedSpannerStmt: "CREATE TABLE IF NOT EXISTS t_test (\n" +
				" pk_text STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_text STRING(MAX) OPTIONS (cassandra_type = 'text'),\n" +
				") PRIMARY KEY (pk_text)",
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Has primary key clause with partition key",
			cqlStmt: `CREATE TABLE ks.t_test (
									col_1 TEXT ,
									col_2 TEXT,
									col_3 TEXT ,
									col_4 TEXT ,
									primary key (col_1)
								)`,
			expectedSpannerStmt: "CREATE TABLE t_test (\n" +
				" col_1 STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_2 STRING(MAX) OPTIONS (cassandra_type = 'text'),\n" +
				" col_3 STRING(MAX) OPTIONS (cassandra_type = 'text'),\n" +
				" col_4 STRING(MAX) OPTIONS (cassandra_type = 'text'),\n" +
				") PRIMARY KEY (col_1)",
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Has primary key clause with partition key and clustering key",
			cqlStmt: `CREATE TABLE ks.t_test (
									col_1 TEXT ,
									col_2 TEXT,
									col_3 TEXT ,
									col_4 TEXT ,
									primary key (col_1, col_2, col_3)
								)`,
			expectedSpannerStmt: "CREATE TABLE t_test (\n" +
				" col_1 STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_2 STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_3 STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_4 STRING(MAX) OPTIONS (cassandra_type = 'text'),\n" +
				") PRIMARY KEY (col_1, col_2, col_3)",
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Has primary key clause with composite partition key and clustering key",
			cqlStmt: `CREATE TABLE ks.t_test (
									col_1 TEXT ,
									col_2 TEXT,
									col_3 TEXT ,
									col_4 TEXT ,
									primary key ((col_1, col_2), col_3, col_4)
								)`,
			expectedSpannerStmt: "CREATE TABLE t_test (\n" +
				" col_1 STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_2 STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_3 STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_4 STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				") PRIMARY KEY (col_1, col_2, col_3, col_4)",
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Table Options are ignored",
			cqlStmt: `CREATE TABLE ks.t_test (
							pk_text  TEXT PRIMARY KEY,
							col_text text
						) WITH 	COMPACT STORAGE AND
						 		CLUSTERING ORDER BY (col_text DESC) AND
								ID='5a1c395e-b41f-11e5-9f22-ba0be0483c18'`,
			expectedSpannerStmt: "CREATE TABLE t_test (\n" +
				" pk_text STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_text STRING(MAX) OPTIONS (cassandra_type = 'text'),\n" +
				") PRIMARY KEY (pk_text)",
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "All NativeType without counter",
			cqlStmt: `CREATE TABLE ks.t_all_native (
									pk_text  		TEXT 		PRIMARY KEY,
									col_ascii 		ASCII,
									col_bigint 		BIGINT,
									col_blob 		BLOB,
									col_boolean 	BOOLEAN,
									col_date 		DATE,
									col_decimal 	DECIMAL,
									col_double 		DOUBLE,
									col_float 		FLOAT,
									col_inet 		INET,
									col_int 		INT,
									col_smallint 	SMALLINT,
									col_text 		TEXT,
									col_time 		TIME,
									col_timestamp 	TIMESTAMP,
									col_timeuuid 	TIMEUUID,
									col_tinyint 	TINYINT,
									col_uuid 		UUID,
									col_varchar 	VARCHAR,
									col_varint 		VARINT
								)`,
			expectedSpannerStmt: "CREATE TABLE t_all_native (\n" +
				" pk_text STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_ascii STRING(MAX) OPTIONS (cassandra_type = 'ascii'),\n" +
				" col_bigint INT64 OPTIONS (cassandra_type = 'bigint'),\n" +
				" col_blob BYTES(MAX) OPTIONS (cassandra_type = 'blob'),\n" +
				" col_boolean BOOL OPTIONS (cassandra_type = 'boolean'),\n" +
				" col_date DATE OPTIONS (cassandra_type = 'date'),\n" +
				" col_decimal NUMERIC OPTIONS (cassandra_type = 'decimal'),\n" +
				" col_double FLOAT64 OPTIONS (cassandra_type = 'double'),\n" +
				" col_float FLOAT32 OPTIONS (cassandra_type = 'float'),\n" +
				" col_inet STRING(MAX) OPTIONS (cassandra_type = 'inet'),\n" +
				" col_int INT64 OPTIONS (cassandra_type = 'int'),\n" +
				" col_smallint INT64 OPTIONS (cassandra_type = 'smallint'),\n" +
				" col_text STRING(MAX) OPTIONS (cassandra_type = 'text'),\n" +
				" col_time INT64 OPTIONS (cassandra_type = 'time'),\n" +
				" col_timestamp TIMESTAMP OPTIONS (cassandra_type = 'timestamp'),\n" +
				" col_timeuuid STRING(MAX) OPTIONS (cassandra_type = 'timeuuid'),\n" +
				" col_tinyint INT64 OPTIONS (cassandra_type = 'tinyint'),\n" +
				" col_uuid STRING(MAX) OPTIONS (cassandra_type = 'uuid'),\n" +
				" col_varchar STRING(MAX) OPTIONS (cassandra_type = 'varchar'),\n" +
				" col_varint NUMERIC OPTIONS (cassandra_type = 'varint'),\n" +
				") PRIMARY KEY (pk_text)",
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Use nonreserved keyword as column name",
			cqlStmt: `CREATE TABLE ks.t_all_native (
									TEXT  		TEXT 		PRIMARY KEY,
									ASCII 		ASCII,
									BIGINT 		BIGINT,
									BLOB 		BLOB,
									BOOLEAN 	BOOLEAN,
									DATE 		DATE,
									DECIMAL 	DECIMAL,
									DOUBLE 		DOUBLE,
									FLOAT 		FLOAT,
									INET 		INET,
									INT 		INT,
									SMALLINT 	SMALLINT,
									TEXT 		TEXT,
									TIME 		TIME,
									TIMESTAMP 	TIMESTAMP,
									TIMEUUID 	TIMEUUID,
									TINYINT 	TINYINT,
									UUID 		UUID,
									VARCHAR 	VARCHAR,
									VARINT 		VARINT
								)`,
			expectedSpannerStmt: "CREATE TABLE t_all_native (\n" +
				" TEXT STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" ASCII STRING(MAX) OPTIONS (cassandra_type = 'ascii'),\n" +
				" BIGINT INT64 OPTIONS (cassandra_type = 'bigint'),\n" +
				" BLOB BYTES(MAX) OPTIONS (cassandra_type = 'blob'),\n" +
				" BOOLEAN BOOL OPTIONS (cassandra_type = 'boolean'),\n" +
				" DATE DATE OPTIONS (cassandra_type = 'date'),\n" +
				" DECIMAL NUMERIC OPTIONS (cassandra_type = 'decimal'),\n" +
				" DOUBLE FLOAT64 OPTIONS (cassandra_type = 'double'),\n" +
				" FLOAT FLOAT32 OPTIONS (cassandra_type = 'float'),\n" +
				" INET STRING(MAX) OPTIONS (cassandra_type = 'inet'),\n" +
				" INT INT64 OPTIONS (cassandra_type = 'int'),\n" +
				" SMALLINT INT64 OPTIONS (cassandra_type = 'smallint'),\n" +
				" TEXT STRING(MAX) OPTIONS (cassandra_type = 'text'),\n" +
				" TIME INT64 OPTIONS (cassandra_type = 'time'),\n" +
				" TIMESTAMP TIMESTAMP OPTIONS (cassandra_type = 'timestamp'),\n" +
				" TIMEUUID STRING(MAX) OPTIONS (cassandra_type = 'timeuuid'),\n" +
				" TINYINT INT64 OPTIONS (cassandra_type = 'tinyint'),\n" +
				" UUID STRING(MAX) OPTIONS (cassandra_type = 'uuid'),\n" +
				" VARCHAR STRING(MAX) OPTIONS (cassandra_type = 'varchar'),\n" +
				" VARINT NUMERIC OPTIONS (cassandra_type = 'varint'),\n" +
				") PRIMARY KEY (TEXT)",
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Counter table",
			cqlStmt: `CREATE TABLE ks.t_counter (
						pk_text	TEXT 	PRIMARY KEY,
						c1		Counter,
						c2		COUNTER
					  )`,
			expectedSpannerStmt: "CREATE TABLE t_counter (\n" +
				" pk_text STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" c1 INT64 NOT NULL OPTIONS (cassandra_type = 'counter'),\n" +
				" c2 INT64 NOT NULL OPTIONS (cassandra_type = 'counter'),\n" +
				") PRIMARY KEY (pk_text)",
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "LIST with all native types excluding counter",
			cqlStmt: `CREATE TABLE ks.t_list (
									pk_text 			TEXT 			PRIMARY KEY,
									col_list_ascii 		LIST<ASCII>,
									col_list_bigint		LIST<BIGINT>,
									col_list_blob 		LIST<BLOB>,
									col_list_boolean 	LIST<BOOLEAN>,
									col_list_date 		LIST<DATE>,
									col_list_decimal	LIST<DECIMAL>,
									col_list_double 	LIST<DOUBLE>,
									col_list_float 		LIST<FLOAT>,
									col_list_inet 		LIST<INET>,
									col_list_int 		LIST<INT>,
									col_list_smallint 	LIST<SMALLINT>,
									col_list_text 		LIST<TEXT>,
									col_list_time 		LIST<TIME>,
									col_list_timestamp 	LIST<TIMESTAMP>,
									col_list_timeuuid 	LIST<TIMEUUID>,
									col_list_tinyint 	LIST<TINYINT>,
									col_list_uuid 		LIST<UUID>,
									col_list_varchar 	LIST<VARCHAR>,
									col_list_varint 	LIST<VARINT>
								)`,
			expectedSpannerStmt: "CREATE TABLE t_list (\n" +
				" pk_text STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_list_ascii ARRAY<STRING(MAX)> OPTIONS (cassandra_type = 'list<ascii>'),\n" +
				" col_list_bigint ARRAY<INT64> OPTIONS (cassandra_type = 'list<bigint>'),\n" +
				" col_list_blob ARRAY<BYTES(MAX)> OPTIONS (cassandra_type = 'list<blob>'),\n" +
				" col_list_boolean ARRAY<BOOL> OPTIONS (cassandra_type = 'list<boolean>'),\n" +
				" col_list_date ARRAY<DATE> OPTIONS (cassandra_type = 'list<date>'),\n" +
				" col_list_decimal ARRAY<NUMERIC> OPTIONS (cassandra_type = 'list<decimal>'),\n" +
				" col_list_double ARRAY<FLOAT64> OPTIONS (cassandra_type = 'list<double>'),\n" +
				" col_list_float ARRAY<FLOAT32> OPTIONS (cassandra_type = 'list<float>'),\n" +
				" col_list_inet ARRAY<STRING(MAX)> OPTIONS (cassandra_type = 'list<inet>'),\n" +
				" col_list_int ARRAY<INT64> OPTIONS (cassandra_type = 'list<int>'),\n" +
				" col_list_smallint ARRAY<INT64> OPTIONS (cassandra_type = 'list<smallint>'),\n" +
				" col_list_text ARRAY<STRING(MAX)> OPTIONS (cassandra_type = 'list<text>'),\n" +
				" col_list_time ARRAY<INT64> OPTIONS (cassandra_type = 'list<time>'),\n" +
				" col_list_timestamp ARRAY<TIMESTAMP> OPTIONS (cassandra_type = 'list<timestamp>'),\n" +
				" col_list_timeuuid ARRAY<STRING(MAX)> OPTIONS (cassandra_type = 'list<timeuuid>'),\n" +
				" col_list_tinyint ARRAY<INT64> OPTIONS (cassandra_type = 'list<tinyint>'),\n" +
				" col_list_uuid ARRAY<STRING(MAX)> OPTIONS (cassandra_type = 'list<uuid>'),\n" +
				" col_list_varchar ARRAY<STRING(MAX)> OPTIONS (cassandra_type = 'list<varchar>'),\n" +
				" col_list_varint ARRAY<NUMERIC> OPTIONS (cassandra_type = 'list<varint>'),\n" +
				") PRIMARY KEY (pk_text)",
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "SET with all native types excluding counter",
			cqlStmt: `CREATE TABLE ks.t_set (
									pk_text 			TEXT 			PRIMARY KEY,
									col_set_ascii 		SET<ASCII>,
									col_set_bigint 		SET<BIGINT>,
									col_set_blob 		SET<BLOB>,
									col_set_boolean 	SET<BOOLEAN>,
									col_set_date 		SET<DATE>,
									col_set_decimal 	SET<DECIMAL>,
									col_set_double 		SET<DOUBLE>,
									col_set_float 		SET<FLOAT>,
									col_set_inet 		SET<INET>,
									col_set_int 		SET<INT>,
									col_set_smallint 	SET<SMALLINT>,
									col_set_text 		SET<TEXT>,
									col_set_time 		SET<TIME>,
									col_set_timestamp 	SET<TIMESTAMP>,
									col_set_timeuuid 	SET<TIMEUUID>,
									col_set_tinyint 	SET<TINYINT>,
									col_set_uuid 		SET<UUID>,
									col_set_varchar 	SET<VARCHAR>,
									col_set_varint 		SET<VARINT>
								)`,
			expectedSpannerStmt: "CREATE TABLE t_set (\n" +
				" pk_text STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_set_ascii ARRAY<STRING(MAX)> OPTIONS (cassandra_type = 'set<ascii>'),\n" +
				" col_set_bigint ARRAY<INT64> OPTIONS (cassandra_type = 'set<bigint>'),\n" +
				" col_set_blob ARRAY<BYTES(MAX)> OPTIONS (cassandra_type = 'set<blob>'),\n" +
				" col_set_boolean ARRAY<BOOL> OPTIONS (cassandra_type = 'set<boolean>'),\n" +
				" col_set_date ARRAY<DATE> OPTIONS (cassandra_type = 'set<date>'),\n" +
				" col_set_decimal ARRAY<NUMERIC> OPTIONS (cassandra_type = 'set<decimal>'),\n" +
				" col_set_double ARRAY<FLOAT64> OPTIONS (cassandra_type = 'set<double>'),\n" +
				" col_set_float ARRAY<FLOAT32> OPTIONS (cassandra_type = 'set<float>'),\n" +
				" col_set_inet ARRAY<STRING(MAX)> OPTIONS (cassandra_type = 'set<inet>'),\n" +
				" col_set_int ARRAY<INT64> OPTIONS (cassandra_type = 'set<int>'),\n" +
				" col_set_smallint ARRAY<INT64> OPTIONS (cassandra_type = 'set<smallint>'),\n" +
				" col_set_text ARRAY<STRING(MAX)> OPTIONS (cassandra_type = 'set<text>'),\n" +
				" col_set_time ARRAY<INT64> OPTIONS (cassandra_type = 'set<time>'),\n" +
				" col_set_timestamp ARRAY<TIMESTAMP> OPTIONS (cassandra_type = 'set<timestamp>'),\n" +
				" col_set_timeuuid ARRAY<STRING(MAX)> OPTIONS (cassandra_type = 'set<timeuuid>'),\n" +
				" col_set_tinyint ARRAY<INT64> OPTIONS (cassandra_type = 'set<tinyint>'),\n" +
				" col_set_uuid ARRAY<STRING(MAX)> OPTIONS (cassandra_type = 'set<uuid>'),\n" +
				" col_set_varchar ARRAY<STRING(MAX)> OPTIONS (cassandra_type = 'set<varchar>'),\n" +
				" col_set_varint ARRAY<NUMERIC> OPTIONS (cassandra_type = 'set<varint>'),\n" +
				") PRIMARY KEY (pk_text)",
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "MAP with all native types as the key excluding counter",
			cqlStmt: `CREATE TABLE ks.t_map_key (
									pk_text 				TEXT 				PRIMARY KEY,
									col_map_key_ascii 		MAP<ASCII,TEXT>,
									col_map_key_bigint 		MAP<BIGINT,TEXT>,
									col_map_key_blob 		MAP<BLOB,TEXT>,
									col_map_key_boolean 	MAP<BOOLEAN,TEXT>,
									col_map_key_date 		MAP<DATE,TEXT>,
									col_map_key_decimal 	MAP<DECIMAL,TEXT>,
									col_map_key_double 		MAP<DOUBLE,TEXT>,
									col_map_key_float 		MAP<FLOAT,TEXT>,
									col_map_key_inet 		MAP<INET,TEXT>,
									col_map_key_int 		MAP<INT,TEXT>,
									col_map_key_smallint 	MAP<SMALLINT,TEXT>,
									col_map_key_text 		MAP<TEXT,TEXT>,
									col_map_key_time 		MAP<TIME,TEXT>,
									col_map_key_timestamp 	MAP<TIMESTAMP,TEXT>,
									col_map_key_timeuuid 	MAP<TIMEUUID,TEXT>,
									col_map_key_tinyint 	MAP<TINYINT,TEXT>,
									col_map_key_uuid 		MAP<UUID,TEXT>,
									col_map_key_varchar 	MAP<VARCHAR,TEXT>,
									col_map_key_varint 		MAP<VARINT,TEXT>
								)`,
			expectedSpannerStmt: "CREATE TABLE t_map_key (\n" +
				" pk_text STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_map_key_ascii JSON OPTIONS (cassandra_type = 'map<ascii,text>'),\n" +
				" col_map_key_bigint JSON OPTIONS (cassandra_type = 'map<bigint,text>'),\n" +
				" col_map_key_blob JSON OPTIONS (cassandra_type = 'map<blob,text>'),\n" +
				" col_map_key_boolean JSON OPTIONS (cassandra_type = 'map<boolean,text>'),\n" +
				" col_map_key_date JSON OPTIONS (cassandra_type = 'map<date,text>'),\n" +
				" col_map_key_decimal JSON OPTIONS (cassandra_type = 'map<decimal,text>'),\n" +
				" col_map_key_double JSON OPTIONS (cassandra_type = 'map<double,text>'),\n" +
				" col_map_key_float JSON OPTIONS (cassandra_type = 'map<float,text>'),\n" +
				" col_map_key_inet JSON OPTIONS (cassandra_type = 'map<inet,text>'),\n" +
				" col_map_key_int JSON OPTIONS (cassandra_type = 'map<int,text>'),\n" +
				" col_map_key_smallint JSON OPTIONS (cassandra_type = 'map<smallint,text>'),\n" +
				" col_map_key_text JSON OPTIONS (cassandra_type = 'map<text,text>'),\n" +
				" col_map_key_time JSON OPTIONS (cassandra_type = 'map<time,text>'),\n" +
				" col_map_key_timestamp JSON OPTIONS (cassandra_type = 'map<timestamp,text>'),\n" +
				" col_map_key_timeuuid JSON OPTIONS (cassandra_type = 'map<timeuuid,text>'),\n" +
				" col_map_key_tinyint JSON OPTIONS (cassandra_type = 'map<tinyint,text>'),\n" +
				" col_map_key_uuid JSON OPTIONS (cassandra_type = 'map<uuid,text>'),\n" +
				" col_map_key_varchar JSON OPTIONS (cassandra_type = 'map<varchar,text>'),\n" +
				" col_map_key_varint JSON OPTIONS (cassandra_type = 'map<varint,text>'),\n" +
				") PRIMARY KEY (pk_text)",
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "MAP with all native types as the value excluding counter",
			cqlStmt: `CREATE TABLE ks.t_map_value (
									pk_text 				TEXT 				PRIMARY KEY,
									col_map_value_ascii 	MAP<INT,ASCII>,
									col_map_value_bigint	MAP<INT,BIGINT>,
									col_map_value_blob 		MAP<INT,BLOB>,
									col_map_value_boolean 	MAP<INT,BOOLEAN>,
									col_map_value_date 		MAP<INT,DATE>,
									col_map_value_decimal	MAP<INT,DECIMAL>,
									col_map_value_double 	MAP<INT,DOUBLE>,
									col_map_value_float 	MAP<INT,FLOAT>,
									col_map_value_inet 		MAP<INT,INET>,
									col_map_value_int 		MAP<INT,INT>,
									col_map_value_smallint 	MAP<INT,SMALLINT>,
									col_map_value_text 		MAP<INT,TEXT>,
									col_map_value_time 		MAP<INT,TIME>,
									col_map_value_timestamp MAP<INT,TIMESTAMP>,
									col_map_value_timeuuid 	MAP<INT,TIMEUUID>,
									col_map_value_tinyint 	MAP<INT,TINYINT>,
									col_map_value_uuid 		MAP<INT,UUID>,
									col_map_value_varchar 	MAP<INT,VARCHAR>,
									col_map_value_varint 	MAP<INT,VARINT>
								)`,
			expectedSpannerStmt: "CREATE TABLE t_map_value (\n" +
				" pk_text STRING(MAX) NOT NULL OPTIONS (cassandra_type = 'text'),\n" +
				" col_map_value_ascii JSON OPTIONS (cassandra_type = 'map<int,ascii>'),\n" +
				" col_map_value_bigint JSON OPTIONS (cassandra_type = 'map<int,bigint>'),\n" +
				" col_map_value_blob JSON OPTIONS (cassandra_type = 'map<int,blob>'),\n" +
				" col_map_value_boolean JSON OPTIONS (cassandra_type = 'map<int,boolean>'),\n" +
				" col_map_value_date JSON OPTIONS (cassandra_type = 'map<int,date>'),\n" +
				" col_map_value_decimal JSON OPTIONS (cassandra_type = 'map<int,decimal>'),\n" +
				" col_map_value_double JSON OPTIONS (cassandra_type = 'map<int,double>'),\n" +
				" col_map_value_float JSON OPTIONS (cassandra_type = 'map<int,float>'),\n" +
				" col_map_value_inet JSON OPTIONS (cassandra_type = 'map<int,inet>'),\n" +
				" col_map_value_int JSON OPTIONS (cassandra_type = 'map<int,int>'),\n" +
				" col_map_value_smallint JSON OPTIONS (cassandra_type = 'map<int,smallint>'),\n" +
				" col_map_value_text JSON OPTIONS (cassandra_type = 'map<int,text>'),\n" +
				" col_map_value_time JSON OPTIONS (cassandra_type = 'map<int,time>'),\n" +
				" col_map_value_timestamp JSON OPTIONS (cassandra_type = 'map<int,timestamp>'),\n" +
				" col_map_value_timeuuid JSON OPTIONS (cassandra_type = 'map<int,timeuuid>'),\n" +
				" col_map_value_tinyint JSON OPTIONS (cassandra_type = 'map<int,tinyint>'),\n" +
				" col_map_value_uuid JSON OPTIONS (cassandra_type = 'map<int,uuid>'),\n" +
				" col_map_value_varchar JSON OPTIONS (cassandra_type = 'map<int,varchar>'),\n" +
				" col_map_value_varint JSON OPTIONS (cassandra_type = 'map<int,varint>'),\n" +
				") PRIMARY KEY (pk_text)",
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Non counter column in counter table error",
			cqlStmt: `CREATE TABLE ks.t_test (
						pk 			text	Primary key,
						non_pk1 	text,
						non_pk2 	counter
					)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "Cannot mix counter and non counter columns in the same table",
		},
		// Testcases for invalid number of pk.
		{
			name: "Has both table and column primary key error",
			cqlStmt: `CREATE TABLE ks.t_test (
						pk 		text	Primary key,
						non_pk 	text,
						primary key (pk)
					)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "Multiple PRIMARY KEY specified for table 't_test' (exactly one required)",
		},
		{
			name: "Has multiple column primary key error",
			cqlStmt: `CREATE TABLE ks.t_test (
						pk 		text 	Primary key,
						non_pk 	text	Primary key
					)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "Multiple PRIMARY KEY specified for table 't_test' (exactly one required)",
		},
		{
			name: "No primary key error",
			cqlStmt: `CREATE TABLE ks.t_no_pk (
						col1 	list<int>,
						col2 	TEXT
					)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "No PRIMARY KEY specified for table 't_no_pk' (exactly one required)",
		},
		// Testcases for invalud/unsupported cql data type.
		{
			name: "Invalid Cql JSON type error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text 	Primary key,
									non_pk 	JSON
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: mismatched input 'JSON' expecting {'ASCII', 'BIGINT', 'BLOB', 'BOOLEAN', 'COUNTER', 'DATE', 'DECIMAL', 'DOUBLE', 'FLOAT', 'INET', 'INT', 'LIST', 'MAP', 'SET', 'SMALLINT', 'TEXT', 'TIME', 'TIMESTAMP', 'TIMEUUID', 'TINYINT', 'UUID', 'VARCHAR', 'VARINT'} at line 3, column 17",
		},
		{
			name: "Invalid Cql INT64 type error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text 	Primary key,
									non_pk 	INT64
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: mismatched input 'INT64' expecting {'ASCII', 'BIGINT', 'BLOB', 'BOOLEAN', 'COUNTER', 'DATE', 'DECIMAL', 'DOUBLE', 'FLOAT', 'INET', 'INT', 'LIST', 'MAP', 'SET', 'SMALLINT', 'TEXT', 'TIME', 'TIMESTAMP', 'TIMEUUID', 'TINYINT', 'UUID', 'VARCHAR', 'VARINT'} at line 3, column 17",
		},
		{
			name: "Unsupported Cql duration type error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text 		Primary key,
									non_pk 	Duration
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: mismatched input 'Duration' expecting {'ASCII', 'BIGINT', 'BLOB', 'BOOLEAN', 'COUNTER', 'DATE', 'DECIMAL', 'DOUBLE', 'FLOAT', 'INET', 'INT', 'LIST', 'MAP', 'SET', 'SMALLINT', 'TEXT', 'TIME', 'TIMESTAMP', 'TIMEUUID', 'TINYINT', 'UUID', 'VARCHAR', 'VARINT'} at line 3, column 17",
		},
		{
			name: "Unsupported Cql tuple type error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text 		Primary key,
									non_pk 	tuple<text>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: mismatched input 'tuple' expecting {'ASCII', 'BIGINT', 'BLOB', 'BOOLEAN', 'COUNTER', 'DATE', 'DECIMAL', 'DOUBLE', 'FLOAT', 'INET', 'INT', 'LIST', 'MAP', 'SET', 'SMALLINT', 'TEXT', 'TIME', 'TIMESTAMP', 'TIMEUUID', 'TINYINT', 'UUID', 'VARCHAR', 'VARINT'} at line 3, column 17",
		},
		{
			name: "Unsupported Cql frozn type error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text 				Primary key,
									non_pk 	frozen<set<int>>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: mismatched input 'frozen' expecting {'ASCII', 'BIGINT', 'BLOB', 'BOOLEAN', 'COUNTER', 'DATE', 'DECIMAL', 'DOUBLE', 'FLOAT', 'INET', 'INT', 'LIST', 'MAP', 'SET', 'SMALLINT', 'TEXT', 'TIME', 'TIMESTAMP', 'TIMEUUID', 'TINYINT', 'UUID', 'VARCHAR', 'VARINT'} at line 3, column 17",
		},
		// Testcases for using counter as the type param of collection types.
		{
			name: "List of Counter error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text 			Primary key,
									non_pk 	list<counter>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "Counters are not allowed inside collections: list<counter>",
		},
		{
			name: "SET of Counter error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text 			Primary key,
									non_pk 	SET<counter>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "Counters are not allowed inside collections: SET<counter>",
		},
		{
			name: "Counter as Map key error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text 				Primary key,
									non_pk 	MAP<counter,int>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "Counters are not allowed inside collections: MAP<counter,int>",
		},
		{
			name: "Counter as Map value error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text 				Primary key,
									non_pk 	MAP<text,counter>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "Counters are not allowed inside collections: MAP<text,counter>",
		},
		// Testcases for nested collection.
		{
			name: "A list with a neseted set error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text 				Primary key,
									non_pk 	list<set<int>>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: mismatched input 'set' expecting {'ASCII', 'BIGINT', 'BLOB', 'BOOLEAN', 'COUNTER', 'DATE', 'DECIMAL', 'DOUBLE', 'FLOAT', 'INET', 'INT', 'SMALLINT', 'TEXT', 'TIME', 'TIMESTAMP', 'TIMEUUID', 'TINYINT', 'UUID', 'VARCHAR', 'VARINT'} at line 3, column 22",
		},
		{
			name: "A set with a neseted list error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text 				Primary key,
									non_pk 	set<list<int>>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: mismatched input 'list' expecting {'ASCII', 'BIGINT', 'BLOB', 'BOOLEAN', 'COUNTER', 'DATE', 'DECIMAL', 'DOUBLE', 'FLOAT', 'INET', 'INT', 'SMALLINT', 'TEXT', 'TIME', 'TIMESTAMP', 'TIMEUUID', 'TINYINT', 'UUID', 'VARCHAR', 'VARINT'} at line 3, column 21",
		},
		{
			name: "A map with a neseted list error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text 				Primary key,
									non_pk 	map<list<int>,text>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: mismatched input 'list' expecting {'ASCII', 'BIGINT', 'BLOB', 'BOOLEAN', 'COUNTER', 'DATE', 'DECIMAL', 'DOUBLE', 'FLOAT', 'INET', 'INT', 'SMALLINT', 'TEXT', 'TIME', 'TIMESTAMP', 'TIMEUUID', 'TINYINT', 'UUID', 'VARCHAR', 'VARINT'} at line 3, column 21",
		},
		{
			name: "A set with a nested frozen type is unsupported",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text 				Primary key,
									non_pk 	set<frozen <set<int>>>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: mismatched input 'frozen' expecting {'ASCII', 'BIGINT', 'BLOB', 'BOOLEAN', 'COUNTER', 'DATE', 'DECIMAL', 'DOUBLE', 'FLOAT', 'INET', 'INT', 'SMALLINT', 'TEXT', 'TIME', 'TIMESTAMP', 'TIMEUUID', 'TINYINT', 'UUID', 'VARCHAR', 'VARINT'} at line 3, column 21",
		},

		// Testcases for wrong number of type params of collection types.
		{
			name: "LIST too many type params error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text				Primary key,
									non_pk 	LIST<text,text>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: missing '>' at ',' at line 3, column 26",
		},
		{
			name: "LIST too few type params error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text				Primary key,
									non_pk 	LIST<>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: missing {'ASCII', 'BIGINT', 'BLOB', 'BOOLEAN', 'COUNTER', 'DATE', 'DECIMAL', 'DOUBLE', 'FLOAT', 'INET', 'INT', 'SMALLINT', 'TEXT', 'TIME', 'TIMESTAMP', 'TIMEUUID', 'TINYINT', 'UUID', 'VARCHAR', 'VARINT'} at '>' at line 3, column 22",
		},
		{
			name: "SET too many type params error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text				Primary key,
									non_pk 	SET<text,text>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: missing '>' at ',' at line 3, column 25",
		},
		{
			name: "SET too few type params error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text				Primary key,
									non_pk 	SET<>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: missing {'ASCII', 'BIGINT', 'BLOB', 'BOOLEAN', 'COUNTER', 'DATE', 'DECIMAL', 'DOUBLE', 'FLOAT', 'INET', 'INT', 'SMALLINT', 'TEXT', 'TIME', 'TIMESTAMP', 'TIMEUUID', 'TINYINT', 'UUID', 'VARCHAR', 'VARINT'} at '>' at line 3, column 21",
		},
		{
			name: "MAP too many type params error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text				Primary key,
									non_pk 	MAP<text,text,text>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: missing '>' at ',' at line 3, column 30",
		},
		{
			name: "MAP too few type params error",
			cqlStmt: `CREATE TABLE ks.t_test (
									pk 		text				Primary key,
									non_pk 	MAP<text>
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: mismatched input '>' expecting ',' at line 3, column 25",
		},
		// Testcases for invalid data type of the pk.
		{
			name: "Table-level pk is a counter error",
			cqlStmt: `CREATE TABLE ks.t_pk_counter (
						pk 		counter,
						non_pk 	TEXT,
						PRIMARY KEY (pk)
					)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "counter type is not supported for PRIMARY KEY column 'pk'",
		},
		{
			name: "Column-level pk is a counter error",
			cqlStmt: `CREATE TABLE ks.t_pk_counter (
									pk 		counter Primary key,
									non_pk 	TEXT
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "counter type is not supported for PRIMARY KEY column 'pk'",
		},
		{
			name: "Table-level pk is a map error",
			cqlStmt: `CREATE TABLE ks.t_pk_counter (
						pk 		map<text,smallint>,
						non_pk 	TEXT,
						Primary key (pk)
					)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "Invalid non-frozen collection type map<text,smallint> for PRIMARY KEY column 'pk'",
		},
		{
			name: "Column-level pk is a map error",
			cqlStmt: `CREATE TABLE ks.t_pk_counter (
									pk 		map<text,smallint> Primary key,
									non_pk 	TEXT
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "Invalid non-frozen collection type map<text,smallint> for PRIMARY KEY column 'pk'",
		},
		{
			name: "Table-level pk is a set error",
			cqlStmt: `CREATE TABLE ks.t_pk_counter (
						pk 		set<double>,
						non_pk 	TEXT,
						Primary key (pk)
					)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "Invalid non-frozen collection type set<double> for PRIMARY KEY column 'pk'",
		},
		{
			name: "Column-level pk is a set error",
			cqlStmt: `CREATE TABLE ks.t_pk_counter (
									pk 		set<double> Primary key,
									non_pk 	TEXT
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "Invalid non-frozen collection type set<double> for PRIMARY KEY column 'pk'",
		},
		{
			name: "Table-level pk is a list error",
			cqlStmt: `CREATE TABLE ks.t_pk_counter (
						pk 		list<int>,
						non_pk 	TEXT,
						primary  key (pk) 
					)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "Invalid non-frozen collection type list<int> for PRIMARY KEY column 'pk'",
		},
		{
			name: "Column-level pk is a list error",
			cqlStmt: `CREATE TABLE ks.t_pk_counter (
									pk 		list<int> Primary key,
									non_pk 	TEXT
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "Invalid non-frozen collection type list<int> for PRIMARY KEY column 'pk'",
		},
		// Testcases for missing required element in the stmt.
		{
			name:                "Empty query Syntax Error",
			cqlStmt:             ``,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: mismatched input '<EOF>' expecting {'ALTER', 'BEGIN', 'CREATE', 'DELETE', 'DROP', 'GRANT', 'INSERT', 'LIST', 'REVOKE', 'SELECT', 'TRUNCATE', 'UPDATE', 'USE'} at line 1",
		},
		{
			name:                "No table name Syntax Error",
			cqlStmt:             `Create table (pk text)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: extraneous input '(' expecting {'AGGREGATE', 'ASCII', 'BIGINT', 'BLOB', 'BOOLEAN', 'CALLED', 'CLUSTERING', 'COMPACT', 'COUNT', 'COUNTER', 'CUSTOM', 'DATE', 'DECIMAL', 'DOUBLE', 'FILTERING', 'FINALFUNC', 'FLOAT', 'FROZEN', 'FUNCTION', 'FUNCTIONS', 'IF', 'INET', 'INITCOND', 'INPUT', 'INT', 'JSON', 'KEY', 'KEYS', 'KEYSPACES', 'LANGUAGE', 'LIST', 'LOGIN', 'MAP', 'NOLOGIN', 'NOSUPERUSER', 'OPTIONS', 'PASSWORD', 'PERMISSION', 'PERMISSIONS', 'RETURNS', 'ROLE', 'ROLES', 'SFUNC', 'SMALLINT', 'STATIC', 'STORAGE', 'STYPE', 'SUPERUSER', 'TEXT', 'TIME', 'TIMESTAMP', 'TIMEUUID', 'TINYINT', 'TRIGGER', 'TTL', 'TUPLE', 'TYPE', 'USER', 'USERS', 'UUID', 'VALUES', 'VARCHAR', 'VARINT', 'WRITETIME', IDENTIFIER, IDENTIFIER_WITH_HYPHEN} at line 1, column 13",
		},
		{
			name:                "No column definition Syntax Error",
			cqlStmt:             `Create table ks.t ()`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: mismatched input ')' expecting {'AGGREGATE', 'ASCII', 'BIGINT', 'BLOB', 'BOOLEAN', 'CALLED', 'CLUSTERING', 'COMPACT', 'COUNT', 'COUNTER', 'CUSTOM', 'DATE', 'DECIMAL', 'DOUBLE', 'FILTERING', 'FINALFUNC', 'FLOAT', 'FROZEN', 'FUNCTION', 'FUNCTIONS', 'INET', 'INITCOND', 'INPUT', 'INT', 'JSON', 'KEY', 'KEYS', 'KEYSPACES', 'LANGUAGE', 'LIST', 'LOGIN', 'MAP', 'NOLOGIN', 'NOSUPERUSER', 'OPTIONS', 'PASSWORD', 'PERMISSION', 'PERMISSIONS', 'RETURNS', 'ROLE', 'ROLES', 'SFUNC', 'SMALLINT', 'STATIC', 'STORAGE', 'STYPE', 'SUPERUSER', 'TEXT', 'TIME', 'TIMESTAMP', 'TIMEUUID', 'TINYINT', 'TRIGGER', 'TTL', 'TUPLE', 'TYPE', 'USER', 'USERS', 'UUID', 'VALUES', 'VARCHAR', 'VARINT', 'WRITETIME', IDENTIFIER} at line 1, column 19",
		},
		// Testcases for invalid keyspace name.
		{
			name: "Keyspance does not match Database error",
			cqlStmt: `CREATE TABLE ks_2.t_test (
									pk_text  TEXT PRIMARY KEY,
									col_text text
								)`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "Keyspace 'ks_2' does not match the Spanner database 'ks'",
		},
		// Testcase for reserved keywords.
		{
			name: "Use reserved keyword ORDER as the table name",
			cqlStmt: `CREATE TABLE ks.ORDER (
						pk_text  TEXT PRIMARY KEY,
						col_text text
					  )`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: mismatched input '.' expecting '(' at line 1, column 15",
		},
		{
			name: "Use reserved keyword SCHEMA as the column name",
			cqlStmt: `CREATE TABLE ks.t_test (
						SCHEMA		TEXT PRIMARY KEY,
						col_text	text
					  )`,
			expectedSpannerStmt: "",
			expectError:         true,
			expectedErrorMsg:    "SyntaxException: extraneous input 'SCHEMA' expecting {'AGGREGATE', 'ASCII', 'BIGINT', 'BLOB', 'BOOLEAN', 'CALLED', 'CLUSTERING', 'COMPACT', 'COUNT', 'COUNTER', 'CUSTOM', 'DATE', 'DECIMAL', 'DOUBLE', 'FILTERING', 'FINALFUNC', 'FLOAT', 'FROZEN', 'FUNCTION', 'FUNCTIONS', 'INET', 'INITCOND', 'INPUT', 'INT', 'JSON', 'KEY', 'KEYS', 'KEYSPACES', 'LANGUAGE', 'LIST', 'LOGIN', 'MAP', 'NOLOGIN', 'NOSUPERUSER', 'OPTIONS', 'PASSWORD', 'PERMISSION', 'PERMISSIONS', 'RETURNS', 'ROLE', 'ROLES', 'SFUNC', 'SMALLINT', 'STATIC', 'STORAGE', 'STYPE', 'SUPERUSER', 'TEXT', 'TIME', 'TIMESTAMP', 'TIMEUUID', 'TINYINT', 'TRIGGER', 'TTL', 'TUPLE', 'TYPE', 'USER', 'USERS', 'UUID', 'VALUES', 'VARCHAR', 'VARINT', 'WRITETIME', IDENTIFIER} at line 2, column 6",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			spannerStmt, ignored, err := ToSpannerStmt(test.cqlStmt, "ks")
			assert.Equal(t, false, ignored)
			if test.expectError {
				assert.Equal(t, test.expectedErrorMsg, err.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.expectedSpannerStmt, spannerStmt)
		})
	}
}

func TestToSpannerStmt_NonCreateTable(t *testing.T) {
	tests := []struct {
		name    string
		cqlStmt string
	}{
		{
			name:    "use keyspace",
			cqlStmt: `USE my_keyspace;`,
		},
		{
			name:    "create keyspace",
			cqlStmt: `CREATE KEYSPACE IF NOT EXISTS my_keyspace WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3};`,
		},
		{
			name:    "alter keyspace",
			cqlStmt: `ALTER KEYSPACE my_keyspace WITH replication = {'class': 'NetworkTopologyStrategy', 'datacenter1': 2};`,
		},
		{
			name:    "drop keyspace",
			cqlStmt: `DROP KEYSPACE IF EXISTS my_keyspace;`,
		},
		{
			name:    "alter table",
			cqlStmt: `ALTER TABLE users ADD age INT;`,
		},
		{
			name:    "drop table",
			cqlStmt: `DROP TABLE IF EXISTS users;`,
		},
		{
			name:    "truncate",
			cqlStmt: `TRUNCATE users;`,
		},
		{
			name:    "select",
			cqlStmt: `SELECT * FROM users WHERE id = 123e4567-e89b-12d3-a456-426614174000;`,
		},
		{
			name:    "insert",
			cqlStmt: `INSERT INTO users (id, name, email) VALUES (uuid(), 'Alice', 'alice@example.com');`,
		},
		{
			name:    "update",
			cqlStmt: `UPDATE users SET email = 'newalice@example.com' WHERE id = 123e4567-e89b-12d3-a456-426614174000;`,
		},
		{
			name:    "delete",
			cqlStmt: `DELETE FROM users WHERE id = 123e4567-e89b-12d3-a456-426614174000;`,
		},
		{
			name: "batch",
			cqlStmt: `BEGIN BATCH
					    INSERT INTO users (id, name) VALUES (uuid(), 'Bob');
					    UPDATE users SET email = 'bob@example.com' WHERE name = 'Bob';
					  APPLY BATCH;`,
		},
		{
			name:    "create index",
			cqlStmt: `CREATE INDEX ON users (email);`,
		},
		{
			name:    "drop index",
			cqlStmt: `DROP INDEX users_email_idx;`,
		},
		{
			name: "create materialized view",
			cqlStmt: `CREATE MATERIALIZED VIEW users_by_email AS
						SELECT * FROM users
						WHERE email IS NOT NULL AND id IS NOT NULL
						PRIMARY KEY (email, id);`,
		},
		{
			name: "alter materialized view",
			cqlStmt: `ALTER MATERIALIZED VIEW cycling.cyclist_by_age 
					  WITH comment = 'A most excellent and useful view'
					  AND bloom_filter_fp_chance = 0.02;`,
		},
		{
			name:    "drop materialized view",
			cqlStmt: `DROP MATERIALIZED VIEW users_by_email;`,
		},
		{
			name:    "create role",
			cqlStmt: `CREATE ROLE analyst WITH PASSWORD = 'secret' AND LOGIN = true;`,
		},
		{
			name:    "alter role",
			cqlStmt: `ALTER ROLE bob WITH PASSWORD = 'PASSWORD_B' AND SUPERUSER = false;`,
		},
		{
			name:    "drop role",
			cqlStmt: `DROP ROLE analyst;`,
		},
		{
			name:    "grant role",
			cqlStmt: `GRANT report_writer TO alice;`,
		},
		{
			name:    "revoke role",
			cqlStmt: `REVOKE report_writer FROM alice;`,
		},
		{
			name:    "list roles",
			cqlStmt: `LIST ROLES;`,
		},
		{
			name:    "grant permission",
			cqlStmt: `GRANT SELECT ON ALL KEYSPACES TO data_reader;`,
		},
		{
			name:    "revoke permission",
			cqlStmt: `REVOKE DROP ON keyspace1.table1 FROM schema_owner;`,
		},
		{
			name:    "list permissions",
			cqlStmt: `LIST ALL PERMISSIONS OF analyst;`,
		},
		{
			name: "create user",
			cqlStmt: `CREATE USER alice WITH PASSWORD = 'password123';
`,
		},
		{
			name:    "alter user",
			cqlStmt: `ALTER USER alice WITH PASSWORD = 'newpass456';`,
		},
		{
			name:    "drop user",
			cqlStmt: `DROP USER alice;`,
		},
		{
			name:    "list users",
			cqlStmt: `LIST USERS;`,
		},
		{
			name: "create function",
			cqlStmt: `CREATE OR REPLACE FUNCTION cycling.fLog (input double) 
					  CALLED ON NULL INPUT 
					  RETURNS double LANGUAGE java AS
					  'return Double.valueOf(Math.log(input.doubleValue()));';`,
		},
		{
			name:    "drop function",
			cqlStmt: `DROP FUNCTION add_numbers;`,
		},
		{
			name: "create aggregate",
			cqlStmt: `CREATE AGGREGATE average(int)
					  SFUNC sum_and_count
					  STYPE tuple<int, int>
					  FINALFUNC divide_sum_by_count
					  INITCOND (0, 0);`,
		},
		{
			name:    "drop aggregate",
			cqlStmt: `DROP AGGREGATE average(int);`,
		},
		{
			name:    "create type",
			cqlStmt: `CREATE TYPE address (street text, city text, zip int);`,
		},
		{
			name:    "alter type",
			cqlStmt: `ALTER TYPE address ADD country text;`,
		},
		{
			name:    "drop type",
			cqlStmt: `DROP TYPE address;`,
		},
		{
			name:    "create trigger",
			cqlStmt: `CREATE TRIGGER user_trigger ON users USING 'com.example.TriggerClass';`,
		},
		{
			name:    "drop trigger",
			cqlStmt: `DROP TRIGGER user_trigger;`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			spannerStmt, ignored, err := ToSpannerStmt(test.cqlStmt, "ks")
			assert.Equal(t, true, ignored)
			assert.NoError(t, err)
			assert.Equal(t, "", spannerStmt)
		})
	}
}
