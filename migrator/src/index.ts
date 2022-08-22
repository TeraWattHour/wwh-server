import { ArgumentParser } from "argparse";
import { migrator } from "./utils/db";

(async () => {
  const parser = new ArgumentParser();

  parser.add_argument("-m", "--migrate", {
    choices: ["up", "down", "latest"],
    default: "up",
  });
  const { migrate } = parser.parse_args();
  console.log(migrate);

  switch (migrate) {
    case "up":
      await migrator.migrateUp();
      break;
    case "down":
      await migrator.migrateDown();
      break;
    case "latest":
      await migrator.migrateToLatest();
      break;
    default:
      break;
  }

  console.log("Migrated");
  process.exit(0);
})();
