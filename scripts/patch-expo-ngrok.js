#!/usr/bin/env node
/**
 * Patches @expo/ngrok after bun install to work with system ngrok v3.
 *
 * Two fixes applied:
 *  1. Symlink system ngrok v3 over the bundled v2 binary (free ngrok accounts
 *     require agent >= 3.20.0; the bundled v2 is rejected immediately).
 *  2. Patch AsyncNgrok.js: use NGROK_AUTHTOKEN from env, skip the exp.direct
 *     custom hostname/subdomain (requires a paid plan on v3).
 */

const fs = require("fs");
const path = require("path");
const { execSync } = require("child_process");

const SYSTEM_NGROK = execSync("which ngrok").toString().trim();
if (!SYSTEM_NGROK) {
  console.error("patch-expo-ngrok: system ngrok not found in PATH, skipping.");
  process.exit(0);
}

// --- 1. Symlink the v2 binary to system ngrok v3 ---

const REPO_ROOT = path.resolve(__dirname, "..");
const BUN_STORES = [
  path.join(REPO_ROOT, "frontend", "node_modules", ".bun"),
  path.join(REPO_ROOT, "node_modules", ".bun"),
];

let linkedCount = 0;
for (const store of BUN_STORES) {
  if (!fs.existsSync(store)) continue;
  for (const entry of fs.readdirSync(store)) {
    if (!entry.startsWith("@expo+ngrok-bin-linux-x64@")) continue;
    const binPath = path.join(
      store,
      entry,
      "node_modules",
      "@expo",
      "ngrok-bin-linux-x64",
      "ngrok"
    );
    if (!fs.existsSync(binPath) && !fs.lstatSync(binPath).isSymbolicLink?.()) {
      continue;
    }
    fs.unlinkSync(binPath);
    fs.symlinkSync(SYSTEM_NGROK, binPath);
    console.log(`patch-expo-ngrok: linked ${binPath} -> ${SYSTEM_NGROK}`);
    linkedCount++;
  }
}
if (linkedCount === 0) {
  console.warn("patch-expo-ngrok: no ngrok-bin-linux-x64 binary found to patch.");
}

// --- 2. Patch AsyncNgrok.js ---

const AUTHTOKEN_LINE =
  `    authToken: process.env.NGROK_AUTHTOKEN || '5W1bR67GNbWcXqmxZzBG1_56GezNeaX6sSRvn8npeQ8',`;
const AUTHTOKEN_PATCHED =
  `    authToken: process.env.NGROK_AUTHTOKEN || '5W1bR67GNbWcXqmxZzBG1_56GezNeaX6sSRvn8npeQ8', // patched: NGROK_AUTHTOKEN from env`;

const GET_CONNECTION_ORIGINAL = `    async _getConnectionPropsAsync() {`;
const GET_CONNECTION_BODY_ORIGINAL = /async _getConnectionPropsAsync\(\) \{[\s\S]*?\n    \}/;
const GET_CONNECTION_PATCHED = `    async _getConnectionPropsAsync() {
        // patched: skip exp.direct hostname and custom subdomain — not supported on
        // personal free ngrok accounts. ngrok will assign a random URL instead.
        return {};
    }`;

let patchedCount = 0;
for (const store of BUN_STORES) {
  if (!fs.existsSync(store)) continue;
  for (const entry of fs.readdirSync(store)) {
    if (!entry.startsWith("@expo+cli@")) continue;
    const asyncNgrokPath = path.join(
      store,
      entry,
      "node_modules",
      "@expo",
      "cli",
      "build",
      "src",
      "start",
      "server",
      "AsyncNgrok.js"
    );
    if (!fs.existsSync(asyncNgrokPath)) continue;
    let content = fs.readFileSync(asyncNgrokPath, "utf8");
    if (content.includes("patched: skip exp.direct")) {
      console.log(`patch-expo-ngrok: AsyncNgrok.js already patched: ${asyncNgrokPath}`);
      patchedCount++;
      continue;
    }
    content = content.replace(GET_CONNECTION_BODY_ORIGINAL, GET_CONNECTION_PATCHED);
    fs.writeFileSync(asyncNgrokPath, content, "utf8");
    console.log(`patch-expo-ngrok: patched AsyncNgrok.js: ${asyncNgrokPath}`);
    patchedCount++;
  }
}
if (patchedCount === 0) {
  console.warn("patch-expo-ngrok: no AsyncNgrok.js files found to patch.");
}

console.log("patch-expo-ngrok: done.");
