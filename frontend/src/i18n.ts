import i18n from "i18next";
import LanguageDetector from 'i18next-browser-languagedetector';
import { initReactI18next } from "react-i18next"
import Backend from "i18next-http-backend"

// import en from '../src/data/en/translations.json'
// import id from '../src/data/id/translations.json'

// const resource = {
//   en: en,
//   id: id
// }

i18n
  // load translation using http -> see /public/locales
  // learn more: https://github.com/i18next/i18next-http-backend
  .use(Backend)
  // detect user language
  // learn more: https://github.com/i18next/i18next-browser-languageDetector
  .use(LanguageDetector)
  // pass the i18n instance to react-i18next.
  .use(initReactI18next)
  // init i18next
  // for all options read: https://www.i18next.com/overview/configuration-options
  .init({
    fallbackLng: 'en',
    debug: true,
    backend: {
      loadPath () {
        const path = window.location.pathname;
        return path === "/coding-test" ? `./src/data/{{lng}}/questions.json` : `./src/data/{{lng}}/translations.json`
      }
    },
    interpolation: {
      escapeValue: false, // not needed for react as it escapes by default
    },
  });
export default i18n