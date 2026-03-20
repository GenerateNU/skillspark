# SkillSpark Mobile Environment Setup

## Objective

Set up your local development environment to run the Skillspark mobile app on both iOS and Android platforms.

Please follow these steps to set everything up so that we can have a smooth mobile dev experience

## Prerequisites

- Bun package manager installed
- Access to the Skillspark monorepo
- *Note: macOS is required for iOS Simulator. Non-Mac users can test on physical iPhone devices.*

## Setup Instructions

### 1. Clone and Install Dependencies

```bash
cd frontend
bun install
```

### 2. iOS Setup

#### Option A: macOS Users - iOS Simulator

**Install Xcode:**

1. Install Xcode from the Mac App Store (this will take a while - it's ~15GB)
2. Open Xcode and accept the license agreement
3. Install Command Line Tools:

```bash
   xcode-select --install
```

**Install iOS Simulator:**

1. Open Xcode
2. Go to **Xcode** → **Settings** → **Platforms**
3. Click the **+** button and install the latest iOS simulator

**Verify Setup:**

```bash
cd apps/mobile
bun ios
```

The iOS Simulator should launch and the app should load.

#### Option B: Windows/Linux Users - Physical iPhone Device

**Install Expo Go on your iPhone:**

1. Download **Expo Go** from the App Store
2. Make sure your iPhone and computer are on the same WiFi network

**Run the app:**

```bash
cd apps/mobile
bun start
```

**Connect your iPhone:**

1. Scan the QR code displayed in terminal using your iPhone's Camera app
2. It will open in Expo Go
3. The app will load and hot reload as you make changes

*Note: You won't be able to build standalone iOS apps or test native modules without macOS, but you can develop and test 95% of features this way.*

### 3. Android Setup

**Install Android Studio:**

1. Download and install Android Studio from https://developer.android.com/studio
2. During installation, make sure to install:
   - Android SDK
   - Android SDK Platform
   - Android Virtual Device (AVD)

**Configure Environment Variables:**

*macOS/Linux* - Add to your `~/.zshrc` or `~/.bash_profile`:

```bash
export ANDROID_HOME=$HOME/Library/Android/sdk  # macOS
# export ANDROID_HOME=$HOME/Android/Sdk  # Linux
export PATH=$PATH:$ANDROID_HOME/emulator
export PATH=$PATH:$ANDROID_HOME/platform-tools
```

*Windows* - Add to System Environment Variables:

```
ANDROID_HOME=C:\Users\YourUsername\AppData\Local\Android\Sdk
Path=%Path%;%ANDROID_HOME%\emulator;%ANDROID_HOME%\platform-tools
```

Then reload your terminal (or restart it).

**Create an Android Emulator:**

1. Open Android Studio
2. Click **More Actions** → **Virtual Device Manager**
3. Click **Create Device**
4. Select **Pixel 6** (or any modern device)
5. Click **Next**
6. Select a system image (recommend **API 33 - Android 13**)
   - Click **Download** if not already installed
7. Click **Next** → **Finish**

**Verify Setup:**

```bash
# Start the emulator
# Either with a command with the emulator name or manually from android studio
emulator -avd Pixel_6_API_33 & ## use the name of your emulator

# In a new terminal, run the app
cd apps/mobile
bun android
```

### 4. Alternative: Physical Device Testing (All Platforms)

**For quick testing without emulators:**

1. Install **Expo Go** on your phone:
   - iOS: App Store
   - Android: Google Play Store

2. Ensure your phone and computer are on the **same WiFi network**

3. Run the app:

```bash
   cd apps/mobile
   bun start
```

4. Scan the QR code with:
   - iOS: Camera app → opens in Expo Go
   - Android: Expo Go app → tap "Scan QR Code"

## Verification Checklist

### Mac Users:

- [ ] iOS Simulator launches and app loads successfully
- [ ] Android Emulator launches and app loads successfully
- [ ] Hot reload works on both platforms (make a small change and save)
- [ ] No TypeScript errors in your IDE

### Windows/Linux Users:

- [ ] Android Emulator launches and app loads successfully
- [ ] Physical iPhone (with Expo Go) connects and loads app successfully
- [ ] Hot reload works on both platforms
- [ ] No TypeScript errors in your IDE

## Common Issues

**"Command not found: emulator"**

- Make sure you've added Android SDK to your PATH and reloaded your shell/terminal

**"No Android connected device found"**

- Start the emulator first before running `bun android`
- Check running emulators: `adb devices`

**iOS Simulator won't launch (Mac only)**

- Try: `sudo xcode-select --reset`
- Restart your computer

**Can't connect physical device to Expo**

- Ensure phone and computer are on the same WiFi network
- Try using tunnel mode: `bun start --tunnel`
- Check that your network allows device communication (not on guest network)

**Metro bundler cache issues**
```bash
cd apps/mobile
bun start --clear
```

**Windows: Android Studio can't find SDK**

- Make sure ANDROID_HOME path uses backslashes: `C:\Users\...\Android\Sdk`
- Run Android Studio as Administrator

## Platform-Specific Notes

### macOS

- ✅ Can develop for iOS and Android
- ✅ Full native module support
- ✅ Can build standalone apps for both platforms

### Windows/Linux

- ✅ Can develop for Android (full support)
- ⚠️  iOS development requires physical device + Expo Go
- ❌ Cannot build standalone iOS apps (requires macOS)
- 💡 95% of features testable via Expo Go on iPhone

## Getting Help

If you encounter issues:

1. Check the error message carefully
2. Search the [Expo documentation](https://docs.expo.dev)
3. Ask in slack
