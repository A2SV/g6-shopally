// src/hooks/useSavedItems.ts
"use client";

import {
  useCreateAlertMutation,
  useDeleteAlertMutation,
} from "@/lib/redux/api/userApiSlice";
import { AlertCreateResponse } from "@/types/SavedItems/AlertCreateResponse";
import type { SavedItem, SavedItemUI } from "@/types/types";
import { useCallback, useEffect, useState } from "react";

const LOCAL_DB_KEY = "itemsList";
const ORDERS_DB_KEY = "ordersList";

interface Order {
  id: string;
  productId: string;
  productTitle: string;
  price: { etb: number; usd: number; fxTimestamp: string };
  orderDate: string;
}

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

export const useSavedItems = (maxItems = 50) => {
  const [savedItems, setSavedItems] = useState<SavedItemUI[]>([]);
  const [orders, setOrders] = useState<Order[]>([]);

  const [createAlert] = useCreateAlertMutation();
  const [deleteAlert] = useDeleteAlertMutation();

  useEffect(() => {
    setSavedItems(loadLocalDb().savedItems);

    if (typeof window !== "undefined") {
      try {
        const savedOrders = localStorage.getItem(ORDERS_DB_KEY);
        if (savedOrders) setOrders(JSON.parse(savedOrders));
      } catch {
        setOrders([]);
      }
    }
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

      let newList = [...prev.filter((i) => i.id !== item.id), uiItem];

      if (newList.length > maxItems) {
        newList = newList.slice(newList.length - maxItems);
      }

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

  const alertChange = useCallback(
    async (itemId: string) => {
      console.log("🚀 alertChange called for itemId:", itemId);

      const item = savedItems.find((i) => i.id === itemId);
      if (!item) {
        console.log("⚠️ Item not found in state");
        return;
      }

      console.log("Current item state:", item);

      const newStatus = !item.priceAlertOn;
      console.log("New toggle status will be:", newStatus);

      // Optimistic UI update
      setSavedItems((prev) =>
        prev.map((i) =>
          i.id === itemId ? { ...i, priceAlertOn: newStatus } : i
        )
      );
      console.log("✅ Optimistic UI update done");

      try {
        if (newStatus) {
          // TURN ON
          console.log("Turning ON alert...");

          const deviceId =
            typeof window !== "undefined"
              ? localStorage.getItem("deviceId") || "default-device"
              : "default-device";

          console.log("Using deviceId:", deviceId);
          console.log(item.title, item.price.usd)

          const res: AlertCreateResponse = await createAlert({
            productId: item.id,
            deviceId,
            productTitle: item.title,
            currentPriceETB: item.price.usd,
          }).unwrap();

          const alertId = res.data?.alertId;
          console.log("Alert created successfully:", res);
          console.log("alertId from response:", alertId);

          if (alertId) {
            setSavedItems((prev) => {
              const newList = prev.map((i) =>
                i.id === itemId ? { ...i, priceAlertOn: true, alertId } : i
              );
              localStorage.setItem(
                LOCAL_DB_KEY,
                JSON.stringify({ savedItems: newList })
              );
              console.log(
                "LocalStorage updated after creating alert:",
                newList
              );
              return newList;
            });
          } else {
            console.warn("⚠️ No alertId returned from createAlert API");
          }
        } else {
          // TURN OFF
          console.log("Turning OFF alert...");
          const alertId = item.alertId;
          console.log("Current alertId:", alertId);

          if (!alertId) {
            console.log("⚠️ No alertId found, skipping DELETE request");
            return;
          }

          const res = await deleteAlert({ id: alertId }).unwrap();
          console.log("Alert deleted successfully:", res);

          setSavedItems((prev) => {
            const newList = prev.map((i) =>
              i.id === itemId
                ? { ...i, priceAlertOn: false, alertId: undefined }
                : i
            );
            localStorage.setItem(
              LOCAL_DB_KEY,
              JSON.stringify({ savedItems: newList })
            );
            console.log("LocalStorage updated after deleting alert:", newList);
            return newList;
          });
        }
      } catch (err) {
        console.error("Alert API failed:", err);
        // revert UI
        setSavedItems((prev) =>
          prev.map((i) =>
            i.id === itemId ? { ...i, priceAlertOn: item.priceAlertOn } : i
          )
        );
      }
    },
    [savedItems, createAlert, deleteAlert]
  );

  const placeOrder = useCallback(
    (
      productId: string,
      productTitle: string,
      price: { etb: number; usd: number; fxTimestamp: string }
    ) => {
      const newOrder: Order = {
        id: `order_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`,
        productId,
        productTitle,
        price,
        orderDate: new Date().toISOString(),
      };

      setOrders((prev) => {
        const newOrders = [...prev, newOrder];
        localStorage.setItem(ORDERS_DB_KEY, JSON.stringify(newOrders));
        return newOrders;
      });
    },
    []
  );

  const clearAll = useCallback(() => {
    setSavedItems([]);
    localStorage.removeItem(LOCAL_DB_KEY);
  }, []);

  return {
    savedItems,
    orders,
    saveItem,
    removeItem,
    alertChange,
    updateItemPrice,
    placeOrder,
    clearAll,
  };
};
