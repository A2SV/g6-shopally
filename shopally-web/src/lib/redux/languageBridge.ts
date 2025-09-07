// src/lib/redux/languageBridge.ts
let currentLang = "en-US";
export const setLanguage = (lang: string) => {
  currentLang = lang;
};
export const getLanguage = () => currentLang;