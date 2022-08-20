import { migrator } from "./utils/db";

(async () => {
  const args = process.argv;
  if (args.includes("--migrate-up")) {
    await migrator.migrateUp();
  } else if (args.includes("--migrate-down")) {
    await migrator.migrateDown();
  }
  console.log("Migrated");
})();
