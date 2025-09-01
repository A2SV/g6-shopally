// //shopally-web/src/utils/SavedItems/cashing.ts
// "use client";

// import type { SavedItem } from "@/types/types";

// interface LocalDb {
//   savedItems: SavedItem[];
// }

// const loadLocalDb = (): LocalDb => {
//   if (typeof window === "undefined") {
//     return { savedItems: [] };
//   }
//   try {
//     return JSON.parse(
//       localStorage.getItem("itemsList") || '{"savedItems": []}'
//     );
//   } catch {
//     return { savedItems: [] };
//   }
// };

// export const saveItemToLocalDb = (item: SavedItem, maxItems = 50) => {
//   const db = loadLocalDb();

//   const exists = db.savedItems.find((i) => i.id === item.id);
//   if (!exists) {
//     if (db.savedItems.length >= maxItems) {
//       db.savedItems.shift();
//     }

//     db.savedItems.push(item);
//     localStorage.setItem("itemsList", JSON.stringify(db));
//   }
// };

// export const removeItemFromLocalDb = (itemId: string) => {
//   const db = loadLocalDb();
//   db.savedItems = db.savedItems.filter((item) => item.id !== itemId);
//   localStorage.setItem("itemsList", JSON.stringify(db));
// };

// export const getSavedItemsFromLocalDb = (): SavedItem[] => {
//   return loadLocalDb().savedItems;
// };

// export const updateItemInLocalDb = (
//   itemId: string,
//   Price: { etb: number; usd: number; fxTimestamp: string }
// ) => {
//   const db = loadLocalDb();
//   const itemIndex = db.savedItems.findIndex((item) => item.id === itemId);
//   if (itemIndex !== -1) {
//     db.savedItems[itemIndex].price = {
//       ...db.savedItems[itemIndex].price,
//       etb: newPrice.etb,
//       usd: newPrice.usd,
//       fxTimestamp: newPrice.fxTimestamp,
//     };
//     localStorage.setItem("itemsList", JSON.stringify(db));
//   }
// };

// export const clearLocalDb = () => {
//   localStorage.removeItem("itemsList");
// };
