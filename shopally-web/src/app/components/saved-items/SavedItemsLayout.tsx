"use client";

import { useDarkMode } from "@/app/components/saved-items/DarkModeContext";
import OfflineAlert from "@/app/components/saved-items/OfflineAlert";
import { useSavedItems } from "@/hooks/useSavedItems";
import React from "react";

export default function SavedItemsLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const { isDarkMode } = useDarkMode();
  const savedItemsHook = useSavedItems(); // single hook instance

  return (
    <div
      className={`min-h-screen flex transition-colors ${
        isDarkMode ? "bg-[#090C11]" : "bg-gray-50"
      }`}
    >
      <main className="flex-1 p-6 lg:p-10">
        <div className="flex items-center justify-between mb-4">
          <h1
            className={`text-2xl font-bold ${
              isDarkMode ? "text-white" : "text-[#0D2A4B]"
            }`}
          >
            Saved Items{" "}
            <span className="text-gray-500 text-base">
              ({savedItemsHook.savedItems.length} items saved)
            </span>
          </h1>
          <button
            className="px-4 py-2 rounded-lg bg-[#0D2A4B] text-white hover:bg-[#133864] transition"
            onClick={savedItemsHook.clearAll}
          >
            Clear All
          </button>
        </div>

        <OfflineAlert />

        {/* Pass hook state & methods to children */}
        {React.cloneElement(children as React.ReactElement, {
          ...savedItemsHook,
        })}
      </main>
    </div>
  );
}
