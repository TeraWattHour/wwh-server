import { Kysely, sql } from "kysely";

export async function up(db: Kysely<any>) {
  await db.schema
    .createTable("user")
    .addColumn("id", "uuid", (c) => c.defaultTo(sql`gen_random_uuid()`).primaryKey())
    .addColumn("name", "varchar(24)", (c) => c.notNull().check(sql`LENGTH(name) >= 4 and LENGTH(name) <= 24`))
    .addColumn("discriminator", "integer", (c) => c.notNull().check(sql`discriminator >= 0 and discriminator <= 9999`))
    .addColumn("email", "varchar(255)", (c) => c.notNull())
    .addColumn("created_at", "timestamptz", (c) => c.defaultTo(sql`now()`).notNull())
    .addColumn("updated_at", "timestamptz", (c) => c.defaultTo(sql`now()`).notNull())
    .addColumn("deleted_at", "timestamptz")
    .addUniqueConstraint("uq_name", ["discriminator", "name"])
    .addUniqueConstraint("uq_email", ["email"])
    .execute();

  await db.schema
    .createTable("session")
    .addColumn("id", "uuid", (c) => c.defaultTo(sql`gen_random_uuid()`).primaryKey())
    .addColumn("user_agent", "varchar(255)", (c) => c.notNull())
    .addColumn("ip", "varchar(16)", (c) => c.notNull())
    .addColumn("user_id", "uuid", (c) => c.notNull().references("user.id").onDelete("cascade").onUpdate("cascade"))
    .addColumn("created_at", "timestamptz", (c) => c.defaultTo(sql`now()`).notNull())
    .addColumn("updated_at", "timestamptz", (c) => c.defaultTo(sql`now()`).notNull())
    .addColumn("expires_at", "timestamptz")
    .execute();

  await db.schema.createType("provider_type").asEnum(["google", "discord"]).execute();

  await db.schema
    .createTable("social_profile")
    .addColumn("id", "serial", (c) => c.primaryKey())
    .addColumn("provider_name", sql`provider_type`, (c) => c.notNull())
    .addColumn("provider_id", "varchar(255)", (c) => c.notNull())
    .addColumn("user_id", "uuid", (c) => c.notNull().references("user.id").onDelete("cascade").onUpdate("cascade"))
    .addUniqueConstraint("uq_provider", ["provider_id", "provider_name"])
    .execute();
}

export async function down(db: Kysely<any>) {
  await db.schema.dropTable("social_profile").execute();
  await db.schema.dropType("provider_type").execute();
  await db.schema.dropTable("session").execute();
  await db.schema.dropTable("user").execute();
}
