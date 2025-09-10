// src/lib/redux/languageBridge.ts
"use client";

import Cookies from "js-cookie";

let currentLang = "en"; // default

export const setLanguage = (lang: string) => {
  currentLang = lang === "am" ? "am" : "en";
  Cookies.set("lang", currentLang); // ✅ persist to cookie
};

export const getLanguage = () => {
  // read from cookie on client
  return Cookies.get("lang") || currentLang;
};
