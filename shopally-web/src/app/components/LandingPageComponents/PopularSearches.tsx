"use client";

import { Search } from "lucide-react";
import { useDarkMode } from "@/app/components/ProfileComponents/DarkModeContext";
import { useLanguage } from "@/hooks/useLanguage";

const searches = [
  { id: 1, text: "Best laptop for students" },
  { id: 2, text: "Kitchen tools for baking bread" },
  { id: 3, text: "Affordable fitness equipment" },
  { id: 4, text: "Smartphone under 5000 ETB" },
  { id: 5, text: "Home office setup" },
  { id: 6, text: "Baby care essentials" },
];

export default function PopularSearches() {
  const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();

  return (
    <section
      className={`py-12 px-6 md:px-12 transition-colors ${
        isDarkMode
          ? "bg-[var(--color-bg-primary)] text-[var(--color-text-primary)]"
          : "bg-white text-gray-900"
      }`}
    >
      <div className="max-w-6xl mx-auto text-center">
        {/* Title */}
        <h2 className="text-2xl font-bold">{t("Try These Popular Searches")}</h2>
        <p className="mt-2 text-gray-500">
          {t("Get started with these example questions")}
        </p>

        {/* Search Cards */}
        <div className="mt-8 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
          {searches.map((item) => (
            <button
              key={item.id}
              className={`flex items-center gap-3 border rounded-lg px-4 py-3 text-left transition-colors ${
                isDarkMode
                  ? "border-gray-700 hover:bg-gray-800"
                  : "border-gray-300 hover:bg-gray-100"
              }`}
            >
              <Search className="w-5 h-5 text-gray-500" />
              <div>
                <p className="font-medium">{t(item.text)}</p>
                <p className="text-sm text-gray-500">{t("Click to search")}</p>
              </div>
            </button>
          ))}
        </div>
      </div>
    </section>
  );
}
