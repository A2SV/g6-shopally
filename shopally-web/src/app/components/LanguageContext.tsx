"use client";

import React, {
  createContext,
  useContext,
  useState,
  ReactNode,
  useCallback,
  useEffect,
} from "react";

type Language = "English" | "Amharic";

interface LanguageContextType {
  currentLanguage: Language;
  switchLanguage: () => void;
  t: (key: string) => string;
  translateText: (text: string) => Promise<string>;
}

const LanguageContext = createContext<LanguageContextType | undefined>(
  undefined
);

// Cache for translations to avoid repeated API calls
const translationCache = new Map<string, string>();

// Common text translations that we'll pre-load
const commonTexts = [
  "Home",
  "How It Works",
  "Compare",
  "Saved Items",
  "Profile",
  "Switch to Amharic",
  "Switch to English",
  "Dark Mode",
  "Light Mode",
  "Profile Settings",
  "User Information",
  "Edit",
  "Cancel",
  "Full Name",
  "Email Address",
  "Password",
  "Change",
  "Preferred Product Categories",
  "Default Currency",
  "Shipping Region",
  "Language Preference",
  "Save Preferences",
  "Account Statistics",
  "Total Orders",
  "Reviews",
  "Wishlist",
  "Electronics",
  "Fashion",
  "Home & Garden",
  "Sports & Outdoors",
  "Books",
  "Beauty & Health",
  "Automotive",
  "Toys & Games",
  "Products Compared",
  "Orders Placed",
  "Total Saved",
  "AI Match Rate",
  "No changes detected. Nothing to update.",
  "Profile updated successfully!",
];

export function LanguageProvider({ children }: { children: ReactNode }) {
  const [currentLanguage, setCurrentLanguage] = useState<Language>("English");
  const [translations, setTranslations] = useState<Record<string, string>>({});

  // Google Translate API function
  const translateText = useCallback(
    async (text: string): Promise<string> => {
      if (currentLanguage === "English") {
        return text; // No translation needed for English
      }

      // Check cache first
      const cacheKey = `${text}_${currentLanguage}`;
      if (translationCache.has(cacheKey)) {
        return translationCache.get(cacheKey)!;
      }

      try {
        // Special handling for "ShopAlly AI Personalization"
        if (text === "ShopAlly AI Personalization") {
          if (currentLanguage === "Amharic") {
            return "ShopAlly AI የግል ማስተካከያ";
          }
          return text;
        }

        // Special handling for "Personalization" (for the dropdown)
        if (text === "Personalization") {
          if (currentLanguage === "Amharic") {
            return "የግል ማስተካከያ";
          }
          return text;
        }

        // Using Google Translate API
        const targetLang = currentLanguage === "Amharic" ? "am" : "en";
        const sourceLang = currentLanguage === "Amharic" ? "en" : "am";

        const response = await fetch(
          `https://translate.googleapis.com/translate_a/single?client=gtx&sl=${sourceLang}&tl=${targetLang}&dt=t&q=${encodeURIComponent(
            text
          )}`
        );

        if (!response.ok) {
          console.warn("Translation failed, returning original text");
          return text;
        }

        const data = await response.json();
        const translatedText = data[0]?.[0]?.[0] || text;

        // Cache the result
        translationCache.set(cacheKey, translatedText);

        return translatedText;
      } catch (error) {
        console.error("Translation error:", error);
        return text; // Return original text if translation fails
      }
    },
    [currentLanguage]
  );

  // Clear translations when switching to English
  useEffect(() => {
    if (currentLanguage === "English") {
      setTranslations({});
    }
  }, [currentLanguage]);

  const switchLanguage = () => {
    setCurrentLanguage((prev) => (prev === "English" ? "Amharic" : "English"));
  };

  // Synchronous translation function with on-demand translation
  const t = useCallback(
    (key: string): string => {
      if (currentLanguage === "English") {
        return key;
      }

      // Check if we have a cached translation
      if (translations[key]) {
        return translations[key];
      }

      // If not cached, trigger translation and return key for now
      // The translation will be cached and component will re-render
      translateText(key).then((translated) => {
        setTranslations((prev) => ({
          ...prev,
          [key]: translated,
        }));
      });

      return key;
    },
    [currentLanguage, translations, translateText]
  );

  return (
    <LanguageContext.Provider
      value={{ currentLanguage, switchLanguage, t, translateText }}
    >
      {children}
    </LanguageContext.Provider>
  );
}

export function useLanguage() {
  const context = useContext(LanguageContext);
  if (context === undefined) {
    throw new Error("useLanguage must be used within a LanguageProvider");
  }
  return context;
}
