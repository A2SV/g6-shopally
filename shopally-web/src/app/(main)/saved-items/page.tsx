"use client";

import Sidebar from "@/app/components/Sidebar";
import {
  DarkModeProvider,
  useDarkMode,
} from "@/app/components/ProfileComponents/DarkModeContext"; 
import { useLanguage } from "@/hooks/useLanguage";
import { LanguageProvider } from "@/providers/LanguageProvider";
import SavedItemsContent from "@/app/components/saved-items/SavedItemsContent";
import { useSavedItems } from "@/hooks/useSavedItems";
import { ReactNode } from "react";

// Layout component
function SavedItemsLayout({ children }: { children: ReactNode }) {
  const { isDarkMode } = useDarkMode();

  return (
    <div
      className={`min-h-screen flex transition-colors ${
        isDarkMode ? "bg-[#090C11]" : "bg-gray-50"
      }`}
    >
      {children}
    </div>
  );
}

// Main content of the page
function SavedItemsPageContent() {
  const { t } = useLanguage();
  const { isDarkMode } = useDarkMode();
  const { savedItems, removeItem, updateItemPrice, alertChange } =
    useSavedItems();

  return (
    <SavedItemsLayout>
      <Sidebar activePage="saved-items" />
      <main className="flex-1 flex flex-col p-6 lg:p-10">
        {/* Page Header */}
        <div className="flex items-center justify-between mb-6">
          <h1
            className={`text-2xl font-bold transition-colors ${
              isDarkMode ? "text-white" : "text-[#0D2A4B]"
            }`}
          >
            {t("Saved Items")}{" "}
            <span className="text-gray-500 text-base">
              ({savedItems.length} items)
            </span>
          </h1>

          {savedItems.length > 0 && (
            <button
              className="px-4 py-2 rounded-lg bg-[#0D2A4B] text-white hover:bg-[#133864] transition"
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
        />
      </main>
    </SavedItemsLayout>
  );
}

// Full page wrapped with providers
export default function SavedItemsPage() {
  return (
     
      <LanguageProvider>
      <DarkModeProvider>
        <SavedItemsPageContent />
      </DarkModeProvider>
      </LanguageProvider>  
  );
}
