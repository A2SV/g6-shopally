"use client";

import { useDarkMode } from "./DarkModeContext";
import { useLanguage } from "../LanguageContext";
import { useProfile } from "@/hooks/useProfile";

export default function ProfileStats() {
  const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();
  const { statistics, isLoading } = useProfile();

  const stats = [
    {
      value: statistics?.productsCompared?.toString() || "0",
      label: t("Products Compared"),
    },
    {
      value: statistics?.ordersPlaced?.toString() || "0",
      label: t("Orders Placed"),
    },
    {
      value: `$${statistics?.totalSaved?.toLocaleString() || "0"}`,
      label: t("Total Saved"),
    },
    {
      value: `${statistics?.aiMatchRate || 0}%`,
      label: t("AI Match Rate"),
    },
  ];

  return (
    <div
      className={`flex-1 transition-colors ${
        isDarkMode ? "bg-[#090C11]" : "bg-gray-50"
      }`}
    >
      {/* Statistics Card */}
      <div
        className={`rounded-lg shadow-sm border p-4 sm:p-6 transition-colors ${
          isDarkMode
            ? "bg-gray-800/20 border-gray-700"
            : "bg-white border-gray-200"
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

        {/* Loading State */}
        {isLoading && (
          <div className="flex items-center justify-center py-8">
            <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-[#FFD300]"></div>
          </div>
        )}

        {/* Statistics Grid */}
        {!isLoading && (
          <div className="grid grid-cols-2 gap-3 sm:gap-4">
            {stats.map((stat, index) => (
              <div
                key={index}
                className={`border rounded-lg p-3 sm:p-4 text-center shadow-sm transition-colors ${
                  isDarkMode
                    ? "bg-[#090C11] border-gray-700"
                    : "bg-white border-gray-200"
                }`}
              >
                <div className="text-xl sm:text-2xl font-bold text-[#FFD300] mb-1">
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
        )}
      </div>
    </div>
  );
}
