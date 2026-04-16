const { getDefaultConfig } = require("expo/metro-config");
const { withNativeWind } = require("nativewind/metro");
const path = require("path");

const projectRoot = __dirname;
const monorepoRoot = path.resolve(projectRoot, "../..");
const repoRoot = path.resolve(monorepoRoot, "..");

const config = getDefaultConfig(projectRoot);

config.resolver.assetExts = [...(config.resolver.assetExts ?? []), "svg"];

// Watch both the frontend monorepo and the repo root, because bun may hoist
// packages into the repo root's node_modules/.bun store.
config.watchFolders = [monorepoRoot, repoRoot];

config.resolver.nodeModulesPaths = [
  path.resolve(projectRoot, "node_modules"),
  path.resolve(monorepoRoot, "node_modules"),
  path.resolve(repoRoot, "node_modules"),
];

module.exports = withNativeWind(config, { input: "./global.css" });
