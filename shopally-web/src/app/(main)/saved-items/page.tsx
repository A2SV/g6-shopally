// saved-items/page.tsx
"use client";

import { DarkModeProvider } from "@/app/components/saved-items/DarkModeContext";
import { useSavedItems } from "@/hooks/useSavedItems";
import SavedItemsContent from "../../components/saved-items/SavedItemsContent";
import SavedItemsLayout from "../../components/saved-items/SavedItemsLayout";

export default function SavedItemsPage() {
  const { savedItems, removeItem, updateItemPrice, alertChange } =
    useSavedItems();
  return (
    <DarkModeProvider>
      <SavedItemsLayout>
        <SavedItemsContent
          savedItems={savedItems}
          removeItem={removeItem}
          updateItemPrice={updateItemPrice}
          onToggleAlert={alertChange} // now it works
        />
      </SavedItemsLayout>
    </DarkModeProvider>
  );
}
