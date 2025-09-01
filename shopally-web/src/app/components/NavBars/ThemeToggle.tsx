//src/app/components/ThemeToggle.tsx
"use client";

import { useTheme } from "next-themes";
import { useEffect, useState } from "react";
import { FiMoon, FiSun } from "react-icons/fi";

export default function ThemeToggle() {
  const { setTheme, resolvedTheme } = useTheme();
  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

  const toggleTheme = () => {
    const newTheme = resolvedTheme === "dark" ? "light" : "dark";
    console.log("ThemeToggle: Switching from", resolvedTheme, "to", newTheme);
    setTheme(newTheme);
  };

  // Don't render anything until mounted to prevent hydration mismatch
  if (!mounted) {
    return null;
  }

  const isDark = resolvedTheme === "dark";
  console.log(
    "ThemeToggle: Current theme is",
    resolvedTheme,
    "isDark:",
    isDark,
  );

  return (
    <button
      className={`p-2 rounded-lg transition-colors duration-200 ${
        isDark
          ? "bg-gray-800 hover:bg-gray-700 text-yellow-400 hover:text-yellow-300"
          : "bg-white hover:bg-gray-100 text-gray-800 hover:text-gray-600 border border-gray-300"
      }`}
      onClick={toggleTheme}
      aria-label={isDark ? "Switch to light mode" : "Switch to dark mode"}
    >
      {isDark ? <FiSun className="w-5 h-5" /> : <FiMoon className="w-5 h-5" />}
    </button>
  );
}
