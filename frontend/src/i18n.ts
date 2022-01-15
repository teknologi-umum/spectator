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
          translation: enTranslation,
          question: enQuestion
        }
      },
      id: {
        translation: {
          translation: idTranslation,
          question: idQuestion
        }
      }
    },
    interpolation: {
      escapeValue: false
    }
  });
export default i18n;