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

// Apply NativeWind first so its resolver is in place, then wrap it.
// Bun's virtual store can hoist multiple versions of singleton packages when
// different deps declare different peer requirements. Metro sees all versions
// via watchFolders and registers native views twice → crash, or React context
// is split across two instances → QueryClient/hook errors.
const nativeWindConfig = withNativeWind(config, { input: "./global.css" });

// Packages that must resolve to a single instance.
// Match exact name OR subpath (e.g. react-native-css-interop/jsx-dev-runtime)
// because nativewind/babel sets importSource: "react-native-css-interop" which
// causes every JSX file to import react-native-css-interop/jsx-dev-runtime.
// If two versions coexist the style registry is split and className has no effect.
const DEDUPLICATED_PACKAGES = [
  "react",
  "react-native-css-interop",
  "react-native-safe-area-context",
  "react-native-screens",
  "react-native-reanimated",
  "@tanstack/react-query",
  "@tanstack/query-core",
];

function getDeduplicatedPackage(moduleName) {
  return DEDUPLICATED_PACKAGES.find(
    (pkg) => moduleName === pkg || moduleName.startsWith(pkg + "/")
  );
}

const nativeWindResolver = nativeWindConfig.resolver.resolveRequest;

nativeWindConfig.resolver.resolveRequest = (context, moduleName, platform) => {
  if (getDeduplicatedPackage(moduleName)) {
    try {
      const resolved = require.resolve(moduleName, { paths: [projectRoot] });
      return { filePath: resolved, type: "sourceFile" };
    } catch {
      // fall through to normal resolution if subpath can't be resolved via Node
    }
  }
  if (nativeWindResolver) {
    return nativeWindResolver(context, moduleName, platform);
  }
  return context.resolveRequest(context, moduleName, platform);
};

module.exports = nativeWindConfig;
