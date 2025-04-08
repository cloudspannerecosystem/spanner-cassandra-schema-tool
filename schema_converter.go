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

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
	"github.com/cloudspannerecosystem/spanner-cassandra-schema-tool/translator"
)

// Flags holds the command-line flags for configuring the Cassandra to Spanner schema conversion.
type Flags struct {
	projectID  string
	instanceID string
	databaseID string
	cqlFile    string
	dryRun     bool
	// TODO: add a flag for creating the database if it does not exist.
}

// parseFlags parses the command-line flags and populates the Flags struct.
func (f *Flags) parseFlags() {
	flag.StringVar(&f.projectID, "project", "", "The Google Cloud project ID")
	flag.StringVar(&f.instanceID, "instance", "", "The Spanner instance ID")
	flag.StringVar(&f.databaseID, "database", "", "The Spanner database ID")
	flag.StringVar(&f.cqlFile, "cql", "", "The path of the CQL file containing DDL statements to be converted and executed on the Spanner database")
	flag.BoolVar(&f.dryRun, "dry-run", false, "Only output converted Spanner DDL statements without execution")

	flag.Parse()
}

// validateRequiredFlags checks if all required command-line flags are provided.
// It returns true if all required flags are present, and false otherwise.
// If false, it also prints usage instructions to the console.
func (f *Flags) validateRequiredFlags() bool {
	missingFlags := []string{}
	if f.projectID == "" {
		missingFlags = append(missingFlags, "-project")
	}
	if f.instanceID == "" {
		missingFlags = append(missingFlags, "-instance")
	}
	if f.databaseID == "" {
		missingFlags = append(missingFlags, "-database")
	}
	if f.cqlFile == "" {
		missingFlags = append(missingFlags, "-cql")
	}

	if len(missingFlags) > 0 {
		fmt.Println("Missing required flags:", missingFlags)
		flag.PrintDefaults()
		return false
	}
	return true
}

// checkGCPCredentials checks if the GCP credentials are set in the environment.
func checkGCPCredentials() error {
	credentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credentials == "" {
		return fmt.Errorf("GCP credentials are not found. Set GOOGLE_APPLICATION_CREDENTIALS environment variable")
	}
	return nil
}

// parseCqlFile reads a CQL file, extracts statements delimited by semicolons,
// and ignores single-line (-- ... or // ...) and multi-line (/* ... */) comments.
// Returns a slice of CQL statements or an error if file operations fail.
//
// TODO: Considering refactoring this function to a struct that has a getNextStmt function.
// The streaming approach can reduce memory usage and it is crucial for handling large files.
func parseCqlFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var statements []string
	var currentStatement strings.Builder
	inMultilineComment := false
	inLineComment := false
	reader := bufio.NewReader(file)

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading file: %w", err)
		}

		currentChar := string(r)

		if inMultilineComment {
			if currentChar == "*" {
				nextRune, _, peekErr := reader.ReadRune()
				if peekErr == nil && string(nextRune) == "/" {
					inMultilineComment = false
				} else if peekErr == nil {
					err = reader.UnreadRune()
					if err != nil {
						return nil, fmt.Errorf("error unreading rune: %w", err)
					}
				}
			}
			continue
		}

		if inLineComment {
			if currentChar == "\n" {
				inLineComment = false
			}
			continue
		}

		if currentChar == "/" {
			nextRune, _, peekErr := reader.ReadRune()
			if peekErr == nil && string(nextRune) == "*" {
				inMultilineComment = true
			} else if peekErr == nil && string(nextRune) == "/" {
				inLineComment = true
			} else if peekErr == nil {
				currentStatement.WriteString(currentChar)
				err = reader.UnreadRune()
				if err != nil {
					return nil, fmt.Errorf("error unreading rune: %w", err)
				}
			} else {
				currentStatement.WriteString(currentChar)
			}
			continue
		}

		if currentChar == "-" {
			nextRune, _, peekErr := reader.ReadRune()
			if peekErr == nil && string(nextRune) == "-" {
				inLineComment = true
			} else if peekErr == nil {
				currentStatement.WriteString(currentChar)
				err = reader.UnreadRune()
				if err != nil {
					return nil, fmt.Errorf("error unreading rune: %w", err)
				}
			} else {
				currentStatement.WriteString(currentChar)
			}
			continue
		}

		currentStatement.WriteString(currentChar)

		if currentChar == ";" {
			statement := strings.TrimSpace(currentStatement.String())
			if statement != "" {
				statements = append(statements, statement)
			}
			currentStatement.Reset()
		}
	}

	// Handle any remaining statement if the file doesn't end with a semicolon
	remainingStatement := strings.TrimSpace(currentStatement.String())
	if remainingStatement != "" {
		statements = append(statements, remainingStatement)
	}
	return statements, nil
}

// writeStmtsToFile writes a slice of statements to the specified file.
// If the file already exists, it will be removed before writing.
// Returns an error if any operation fails.
func writeStmtsToFile(filename string, stmts []string) error {
	err := os.Remove(filename)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove existing file '%s': %w", filename, err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %w", filename, err)
	}
	defer file.Close()

	for _, stmt := range stmts {
		_, err := file.WriteString(stmt + ";\n\n")
		if err != nil {
			return fmt.Errorf("failed to write to file '%s': %w", filename, err)
		}
	}
	return nil
}

func main() {
	var flags Flags
	flags.parseFlags()
	if !flags.validateRequiredFlags() {
		os.Exit(1)
	}

	if err := checkGCPCredentials(); err != nil {
		log.Fatalf("Error: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create admin client: %v", err)
	}
	defer adminClient.Close()

	// Translates the Cassandra DDL stmts in the CQL file to corresponding Spanner Stmts.
	log.Printf("Starting Cassandra to Spanner schema conversion.\n")
	log.Printf("Reading input CQL file: %s\n", flags.cqlFile)
	var spannerCreateTableStmts []string
	stmts, err := parseCqlFile(flags.cqlFile)
	if err != nil {
		log.Fatalf("Failed to read the file: %v\n", err)
	}

	log.Printf("----------------------------------------------")
	for _, stmt := range stmts {
		log.Printf("[Cassandra statement]\n%s\n", stmt)
		spannerStmt, ignored, err := translator.ToSpannerStmt(stmt, flags.databaseID)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		if ignored {
			log.Printf("Skipping non-CREATE TABLE statement.\n")
		} else {
			log.Printf("[Converted Spanner statement]\n%s\n", spannerStmt)
		}
		log.Printf("----------------------------------------------")
		spannerCreateTableStmts = append(spannerCreateTableStmts, spannerStmt)
	}

	outputFile := "schema.txt"
	log.Printf("Writing converted Spanner schema to: %s\n", outputFile)
	err = writeStmtsToFile(outputFile, spannerCreateTableStmts)
	if err != nil {
		log.Println("Error writing to file:", err)
		return
	}

	if flags.dryRun {
		log.Println("Dry run enabled. Skipping schema execution.")
	} else {
		log.Println("Executing generated DDL on Spanner")
		op, err := adminClient.UpdateDatabaseDdl(ctx, &databasepb.UpdateDatabaseDdlRequest{
			Database:   fmt.Sprintf("projects/%s/instances/%s/databases/%s", flags.projectID, flags.instanceID, flags.databaseID),
			Statements: spannerCreateTableStmts,
		})
		if err != nil {
			log.Fatalf("Failed to execute DDL statements on the Cloud Spanner database: %v", err)
		}
		if err := op.Wait(ctx); err != nil {
			log.Fatalf("Failed to execute DDL statements on the Cloud Spanner database: %v", err)
		}
	}
	log.Println("Schema conversion completed!")
}
