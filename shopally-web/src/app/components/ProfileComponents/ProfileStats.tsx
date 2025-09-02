"use client";

import { useState, useEffect } from "react";
import { useDarkMode } from "./DarkModeContext";
import { useLanguage } from "../../components/LanguageContext";
import { LoadingSpinner } from "../LoadingSpinner";

export default function ProfileStats() {
  const [isLoading, setIsLoading] = useState(true);
  const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();

  // Simulate loading on component mount
  useEffect(() => {
    const timer = setTimeout(() => {
      setIsLoading(false);
    }, 1500);
    return () => clearTimeout(timer);
  }, []);

  const stats = [
    {
      value: "47",
      label: t("Products Compared"),
    },
    {
      value: "12",
      label: t("Orders Placed"),
    },
    {
      value: "$1,247",
      label: t("Total Saved"),
    },
    {
      value: "89%",
      label: t("AI Match Rate"),
    },
  ];

  // Show loading state
  if (isLoading) {
    return (
      <div
        className={`flex-1 transition-colors ${
          isDarkMode ? "bg-[var(--color-brand-dark)]" : "bg-gray-50"
        }`}
      >
        <div
          className={`rounded-lg shadow-sm border p-4 sm:p-6 transition-colors ${
            isDarkMode
              ? "bg-gray-800/20 border-gray-700"
              : "bg-[var(--color-brand-white)] border-gray-200"
          }`}
        >
          <div className="flex items-center justify-center py-12">
            <LoadingSpinner size="xl" color="yellow" />
          </div>
        </div>
      </div>
    );
  }

  return (
    <div
      className={`flex-1 transition-colors ${
        isDarkMode ? "bg-[var(--color-brand-dark)]" : "bg-gray-50"
      }`}
    >
      {/* Statistics Card */}
      <div
        className={`rounded-lg shadow-sm border p-4 sm:p-6 transition-colors ${
          isDarkMode
            ? "bg-gray-800/20 border-gray-700"
            : "bg-[var(--color-brand-white)] border-gray-200"
        }`}
      >
        {/* Header */}
        <h2
          className={`text-lg sm:text-xl font-semibold mb-4 sm:mb-6 transition-colors ${
            isDarkMode ? "text-white" : "text-gray-900"
          }`}
        >
          {t("Account Statistics")}
        </h2>

        {/* Statistics Grid */}
        <div className="grid grid-cols-2 gap-3 sm:gap-4">
          {stats.map((stat, index) => (
            <div
              key={index}
              className={`border rounded-lg p-3 sm:p-4 text-center shadow-sm transition-colors ${
                isDarkMode
                  ? "bg-[var(--color-brand-dark)] border-gray-700"
                  : "bg-[var(--color-brand-white)] border-gray-200"
              }`}
            >
              <div className="text-xl sm:text-2xl font-bold text-[var(--color-brand-yellow)] mb-1">
                {stat.value}
              </div>
              <div
                className={`text-sm transition-colors ${
                  isDarkMode ? "text-gray-300" : "text-gray-600"
                }`}
              >
                {stat.label}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
