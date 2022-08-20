import { FileMigrationProvider, Kysely, Migrator, PostgresDialect } from "kysely";
import { Pool } from "pg";
import { promises as fs } from "fs";
import * as path from "path";
require("dotenv").config();

export const db = new Kysely<any>({
  dialect: new PostgresDialect({
    pool: new Pool({
      host: process.env.POSTGRES_HOST,
      database: process.env.POSTGRES_DB,
      port: parseInt(process.env.POSTGRES_PORT!),
      user: process.env.POSTGRES_USER,
      password: process.env.POSTGRES_PASSWORD,
    }),
  }),
});

export const migrator = new Migrator({
  db,
  provider: new FileMigrationProvider({
    fs,
    path,
    migrationFolder: path.resolve(__dirname, "..", "migrations"),
  }),
});
