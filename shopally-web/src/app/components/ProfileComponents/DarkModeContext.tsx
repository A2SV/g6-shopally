"use client";

import {
  createContext,
  useContext,
  useState,
  useEffect,
  ReactNode,
} from "react";

interface DarkModeContextType {
  isDarkMode: boolean;
  toggleDarkMode: () => void;
  resetToLightMode: () => void;
}

const DarkModeContext = createContext<DarkModeContextType | undefined>(
  undefined
);

export function DarkModeProvider({ children }: { children: ReactNode }) {
  const [isDarkMode, setIsDarkMode] = useState(false);

  // Initialize theme from localStorage, default to light mode for new users
  useEffect(() => {
    const savedTheme = localStorage.getItem("theme");

    // Only apply dark mode if explicitly saved as "dark"
    // Default to light mode for new users (no saved theme)
    if (savedTheme === "dark") {
      setIsDarkMode(true);
      document.documentElement.classList.add("dark");
    } else {
      setIsDarkMode(false);
      document.documentElement.classList.remove("dark");

      // If no theme is saved, set light mode as default for new users
      if (!savedTheme) {
        localStorage.setItem("theme", "light");
      }
    }
  }, []);

  const toggleDarkMode = () => {
    const newDarkMode = !isDarkMode;
    setIsDarkMode(newDarkMode);

    if (newDarkMode) {
      document.documentElement.classList.add("dark");
      localStorage.setItem("theme", "dark");
    } else {
      document.documentElement.classList.remove("dark");
      localStorage.setItem("theme", "light");
    }
  };

  const resetToLightMode = () => {
    setIsDarkMode(false);
    document.documentElement.classList.remove("dark");
    localStorage.setItem("theme", "light");
  };

  return (
    <DarkModeContext.Provider
      value={{ isDarkMode, toggleDarkMode, resetToLightMode }}
    >
      {children}
    </DarkModeContext.Provider>
  );
}

export function useDarkMode() {
  const context = useContext(DarkModeContext);
  if (context === undefined) {
    throw new Error("useDarkMode must be used within a DarkModeProvider");
  }
  return context;
}
