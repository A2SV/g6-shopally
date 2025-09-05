"use client";

import SavedItemsContent from "@/app/components/saved-items/SavedItemsContent";
import { useSavedItems } from "@/hooks/useSavedItems";
import { useDarkMode } from "@/app/components/ProfileComponents/DarkModeContext";
import { useLanguage } from "@/hooks/useLanguage";

export default function SavedItemsPage() {
  const { t } = useLanguage();
  const { isDarkMode } = useDarkMode();
  const { savedItems, removeItem, updateItemPrice, alertChange, placeOrder } =
    useSavedItems();

  return (
    <div
      className={`min-h-screen flex flex-col p-3 sm:p-4 lg:p-6 xl:p-10 transition-colors ${
        isDarkMode
          ? "bg-[var(--color-bg-primary)]"
          : "bg-[var(--color-bg-secondary)]"
      }`}
    >
      {/* Page Header */}
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3 sm:gap-0 mb-4 sm:mb-6">
        <h1
          className={`text-xl sm:text-2xl font-bold transition-colors ${
            isDarkMode
              ? "text-[var(--color-text-primary)]"
              : "text-[var(--color-text-primary)]"
          }`}
        >
          {t("Saved Items")}{" "}
          <span className="text-gray-500 text-sm sm:text-base">
            ({savedItems.length} items)
          </span>
        </h1>

        {savedItems.length > 0 && (
          <button
            className="px-3 sm:px-4 py-2 sm:py-2.5 rounded-lg transition-colors hover:opacity-80 font-medium text-sm sm:text-base w-full sm:w-auto"
            style={{
              backgroundColor: "var(--color-accent-primary)",
              color: "var(--color-text-button)",
            }}
            onClick={() => {
              savedItems.forEach((item) => removeItem(item.id));
            }}
          >
            {t("Clear All")}
          </button>
        )}
      </div>

      {/* Saved Items Content */}
      <SavedItemsContent
        savedItems={savedItems}
        removeItem={removeItem}
        updateItemPrice={updateItemPrice}
        onToggleAlert={alertChange}
        onPlaceOrder={placeOrder}
      />
    </div>
  );
}
