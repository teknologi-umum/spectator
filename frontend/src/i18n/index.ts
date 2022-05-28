import i18n from "i18next";
import LanguageDetector from "i18next-browser-languagedetector";
import { initReactI18next } from "react-i18next";

import enTranslation from "@/i18n/en/translations.json";
import enQuestion from "@/i18n/en/questions.json";

import idTranslation from "@/i18n/id/translations.json";
import idQuestion from "@/i18n/id/questions.json";

i18n
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    fallbackLng: "en",
    debug: import.meta.env.DEV,
    returnObjects: true,
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
