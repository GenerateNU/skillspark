import i18n from 'i18next'
import { initReactI18next } from 'react-i18next'
import en from './en.json'
import th from './th.json'
import AsyncStorage from '@react-native-async-storage/async-storage';


let savedLng: string = "en";
AsyncStorage.getItem('lng').then((savedLng) => {
  if (savedLng) i18n.changeLanguage(savedLng);
});

i18n.use(initReactI18next).init({
  resources: {
    en: { translation: en },
    th: { translation: th },
  },
  lng: savedLng ?? undefined,
  fallbackLng: 'en',
  debug: false,
  interpolation: {
    escapeValue: false,
  },
})

const storeData = async (value: string) => {
  try {
    await AsyncStorage.setItem('lng', value);
  } catch (e) {
    console.log('error saving language value');
  }
};

i18n.on('languageChanged', (lng) => {
    storeData(lng);
})



export default i18n
