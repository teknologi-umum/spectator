import i18n from "i18next";
import LanguageDetector from "i18next-browser-languagedetector";
import { initReactI18next } from "react-i18next";

import enTranslation from "./data/en/translations.json";
import enQuestion from "./data/en/questions.json";

import idTranslation from "./data/id/translations.json";
import idQuestion from "./data/id/questions.json";

i18n
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    fallbackLng: "en",
    debug: true,
    resources: {
      en: {
        translation: {
          ns1: enTranslation,
          ns2: enQuestion
        }
      },
      id: {
        translation: {
          ns1: idTranslation,
          ns2: idQuestion
        }
      }
    },
    interpolation: {
      escapeValue: false
    }
  });
export default i18n;