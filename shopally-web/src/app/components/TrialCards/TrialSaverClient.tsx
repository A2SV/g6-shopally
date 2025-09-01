// TrialSaverClient.tsx
"use client";

import { useSavedItems } from "@/hooks/useSavedItems";
import type { Product } from "@/types/types";
import ProductCard from "./ProductCard";

interface TrialSaverClientProps {
  products: Product[];
}

export default function TrialSaverClient({ products }: TrialSaverClientProps) {
  const { savedItems, saveItem, removeItem } = useSavedItems();

  return (
    <div className="grid grid-cols-3 gap-4 p-4">
      {products.map((p) => (
        <ProductCard
          key={p.id}
          product={p}
          savedItems={savedItems}
          saveItem={saveItem}
          removeItem={removeItem}
        />
      ))}
    </div>
  );
}
