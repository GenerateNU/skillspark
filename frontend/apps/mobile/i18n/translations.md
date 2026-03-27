# SkillSpark Mobile — Translation Guide

## Overview

The app uses [react-i18next](https://react.i18next.com/) for internationalization. Translation files live in `frontend/apps/mobile/i18n/`.

Currently supported languages:
- `en` — English (default)
- `th` — Thai

---

## Adding a New Language

### 1. Create a translation file

Create a new JSON file in `frontend/apps/mobile/i18n/` named `<code>.json` where `<code>` is the [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code (e.g. `ja.json` for Japanese).

Copy `en.json` as a starting point and translate every value. Do not translate the keys.

```json
{
  "nav": {
    "home": "ホーム",
    "map": "マップ",
    ...
  }
}
```

### 2. Register the language in `i18n/index.ts`

Import the new file and add it to the `resources` object:

```ts
import ja from './ja.json'

i18n.use(initReactI18next).init({
  resources: {
    en: { translation: en },
    th: { translation: th },
    ja: { translation: ja },  // add this
  },
  ...
})
```

### 3. Add the language to the selector in `language.tsx`

Add an entry to the `LANGUAGES` array:

```ts
const LANGUAGES = [
  { code: 'en', label: 'English', flag: '🇺🇸' },
  { code: 'th', label: 'Thai',    flag: '🇹🇭' },
  { code: 'ja', label: 'Japanese', flag: '🇯🇵' },  // add this
];
```

### 4. Add the language label translation to all language files

In every existing translation file (e.g. `en.json`, `th.json`), add the new language's display name under `settings.languages`:

```json
"settings": {
  "languages": {
    "en": "English",
    "th": "Thai",
    "ja": "Japanese"
  }
}
```

### 5. Add the `Accept-Language` header mapping in `apiClient.ts`

In `frontend/packages/api-client/src/apiClient.ts`, update the `languageHeader` line to handle the new code:

```ts
const langMap: Record<string, string> = {
  en: 'en-US',
  th: 'th-TH',
  ja: 'ja-JP',
};
const languageHeader = langMap[i18n.language] ?? 'en-US';
```

---

## Translation Key Reference

All keys follow the structure below. When adding new UI strings, add the key to **all** language files before using `t()` in a component.

```
nav
  home, map, profile, events, about

common
  submit, cancel, loading, loadingMapData, loadingEvents
  errorLoadingEvents, errorOccurred, errorFetchingEvents
  noEventsNearby, noEventsAvailable, noUpcomingEvents
  noChildrenFound, noChildProfilesAdded, error
  locationDenied, openSettings

home
  title, subtitle

dashboard
  title, searchPlaceholder, discoverWeekly, forYou
  basedOn, upcomingEvents, reserve, reserved, members
  categories: technology, art, music, math, sports

profile
  contact, family, born (interpolated: {{year}})
  myBookings, upcoming, previous, saved
  preferences, payment, familyInformation, settings

settings
  title, language
  languages: en, th
  termsAndConditions, privacyPolicy, logOut, deleteAccount

payment
  title, manageBilling, creditCard, name, updateBilling, delete

familyInformation
  title, childProfile, addProfile, interests
  emergencyContact, addContact

childProfile
  editTitle, createTitle, firstName, lastName
  month, year, saving, saveChanges, searchInterests
  deleteProfile, deleteConfirm, requiredFieldsError
  saveError, deleteError

interests
  science, math, music, art, sports, technology, language, other
```

---

## How Language Persistence Works

1. On app startup, `_layout.tsx` reads the saved language code from `AsyncStorage` (`'lng'` key) and calls `i18n.changeLanguage()` before rendering.
2. When the user selects a language in the language screen, `language.tsx` calls `i18n.changeLanguage()`, writes the code to `AsyncStorage`, invalidates the React Query cache (forcing API data to refetch with the new `Accept-Language` header), and updates the guardian's `language_preference` via the backend API.
3. `i18n/index.ts` listens to the `languageChanged` event and also persists the value to `AsyncStorage` as a safety net.

---

## How API Translations Work

Every API request passes an `Accept-Language` header set from `i18n.language` in `customInstance` (`api-client/src/apiClient.ts`). The backend uses this header to return translated content where supported.
