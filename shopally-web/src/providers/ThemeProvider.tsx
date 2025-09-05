"use client";
import { ThemeProvider as NextThemesProvider } from "next-themes";
import * as React from "react";
export function ThemeProvider({
  children,
  ...props
}: React.ComponentProps<typeof NextThemesProvider>) {
  return <NextThemesProvider {...props}>{children}</NextThemesProvider>;
}
// src/app/components/ThemeToggle.tsx
("use client");
import { Moon, Sun } from "lucide-react";
import { useTheme } from "next-themes";
export function ThemeToggle() {
  const { theme, setTheme } = useTheme();
  return (
    <button
      onClick={() => setTheme(theme === "light" ? "dark" : "light")}
      className="flex items-center justify-center lg:justify-start gap-2 lg:gap-3 px-2 py-2 lg:px-3 lg:py-2 rounded-md w-full transition-colors bg-brand-white hover:bg-brand-yellow text-brand-yellow hover:text-brand-dark border border-gray-300 "
      aria-label="Toggle Theme"
    >
      {theme === "light" ? (
        <Sun className="w-5 h-5" />
      ) : (
        <Moon className="w-5 h-5" />
      )}
      <span className="hidden lg:block">
        {theme === "light" ? "Light Mode" : "Dark Mode"}
      </span>
    </button>
  );
}
