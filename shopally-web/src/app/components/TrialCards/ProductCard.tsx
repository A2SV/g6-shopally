// ProductCard.tsx
"use client";

import type { Product, SavedItem } from "@/types/types";
import { useLanguage } from "@/hooks/useLanguage";

interface ProductCardProps {
  product: Product;
  savedItems: SavedItem[];
  saveItem: (item: SavedItem) => void;
  removeItem?: (id: string) => void;
}

export default function ProductCard({
  product,
  savedItems,
  saveItem,
  removeItem,
}: ProductCardProps) {
  const { t } = useLanguage();
  const isSaved = savedItems.some((i) => i.id === product.id);

  return (
    <div className="border rounded-lg shadow-sm p-4 flex flex-col items-center">
      <img
        src={product.imageUrl}
        alt={product.title}
        className="w-32 h-32 object-cover"
      />
      <h2 className="font-semibold text-lg mt-2">{product.title}</h2>
      <p className="text-sm text-gray-600">${product.price.usd}</p>
      <p className="text-xs text-gray-500">
        {product.deliveryEstimate} • ⭐ {product.productRating}
      </p>

      <button
        className={`p-2 rounded mt-2 ${
          isSaved
            ? "bg-gray-300 text-gray-700"
            : "bg-brand-yellow text-brand-dark"
        }`}
        disabled={isSaved}
        onClick={() => saveItem(product)}
      >
        {isSaved ? t("Saved") : t("Save Item")}
      </button>
    </div>
  );
}
