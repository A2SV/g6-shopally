"use client";

import type { SavedItem, SavedItemUI } from "@/types/types";
import { useCallback, useEffect, useState } from "react";

const LOCAL_DB_KEY = "itemsList";

const loadLocalDb = (): { savedItems: SavedItemUI[] } => {
  if (typeof window === "undefined") return { savedItems: [] };
  try {
    return JSON.parse(
      localStorage.getItem(LOCAL_DB_KEY) || '{"savedItems":[] }'
    );
  } catch {
    return { savedItems: [] };
  }
};

export const useSavedItems = () => {
  const [savedItems, setSavedItems] = useState<SavedItemUI[]>([]);

  useEffect(() => {
    setSavedItems(loadLocalDb().savedItems);
  }, []);

  const saveItem = useCallback((item: SavedItem) => {
    setSavedItems((prev) => {
      const prevItem = prev.find((i) => i.id === item.id);
      const uiItem: SavedItemUI = {
        ...item,
        rating: item.productRating,
        ratingCount: prevItem?.ratingCount || 0,
        oldPrice: prevItem?.oldPrice,
        seller: prevItem?.seller || "Unknown",
        checked: prevItem?.checked || "N/A",
        priceAlertOn: prevItem?.priceAlertOn ?? false,
        placeholderText: "IMG",
      };

      const newList = prev.filter((i) => i.id !== item.id);
      newList.push(uiItem);

      localStorage.setItem(
        LOCAL_DB_KEY,
        JSON.stringify({ savedItems: newList })
      );
      return newList;
    });
  }, []);

  const removeItem = useCallback((itemId: string) => {
    setSavedItems((prev) => {
      const newList = prev.filter((item) => item.id !== itemId);
      localStorage.setItem(
        LOCAL_DB_KEY,
        JSON.stringify({ savedItems: newList })
      );
      return newList;
    });
  }, []);

  const updateItemPrice = useCallback(
    (
      itemId: string,
      newPrice: { etb: number; usd: number; fxTimestamp: string }
    ) => {
      setSavedItems((prev) => {
        const newList = prev.map((item) =>
          item.id === itemId ? { ...item, price: { ...newPrice } } : item
        );
        localStorage.setItem(
          LOCAL_DB_KEY,
          JSON.stringify({ savedItems: newList })
        );
        return newList;
      });
    },
    []
  );

  const alertChange = useCallback((itemId: string) => {
    setSavedItems((prev) => {
      const newList = prev.map((item) =>
        item.id === itemId
          ? { ...item, priceAlertOn: !item.priceAlertOn }
          : item
      );
      localStorage.setItem(
        LOCAL_DB_KEY,
        JSON.stringify({ savedItems: newList })
      );
      return newList;
    });
  }, []);

  const clearAll = useCallback(() => {
    setSavedItems([]);
    localStorage.removeItem(LOCAL_DB_KEY);
  }, []);

  return {
    savedItems,
    saveItem,
    removeItem,
    alertChange,
    updateItemPrice,
    clearAll,
  };
};
