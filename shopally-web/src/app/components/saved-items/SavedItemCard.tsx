// shopally-web/src/app/components/saved-items/SavedItemCard.tsx
"use client";

import { useDarkMode } from "@/app/components/saved-items/DarkModeContext";
import ToggleSwitch from "@/app/components/saved-items/ToggleSwitch";
import { Trash2 } from "lucide-react";
import { SavedItemUI } from "../../../types/types";
import OpenLinkButton from "./AlibabaLinkButton";
import Rating from "./Rating";

// Extend SavedItemUI with optional callbacks
interface SavedItemCardProps extends SavedItemUI {
  onRemove?: (id: string) => void;
  onUpdatePrice?: (id: string, price: SavedItemUI["price"]) => void;
  onToggleAlert?: (id: string) => void;
}

export default function SavedItemCard({
  title,
  imageUrl,
  rating,
  ratingCount,
  price,
  oldPrice,
  seller,
  checked,
  priceAlertOn,
  placeholderText,
  id,
  deeplinkUrl,
  onRemove,
  onUpdatePrice,
  onToggleAlert,
}: SavedItemCardProps) {
  const { isDarkMode } = useDarkMode();

  return (
    <div
      className={`rounded-xl shadow border overflow-hidden mb-8 transition-colors ${
        isDarkMode
          ? "bg-gray-800 border-gray-700 text-white"
          : "bg-white border-gray-200 text-gray-900"
      }`}
    >
      {/* Image container */}
      <div className="relative w-full h-44 bg-gray-200 flex items-center justify-center text-3xl font-bold text-gray-400">
        <img
          src={imageUrl}
          alt={title}
          className="w-full h-full object-cover"
        />

        {/* Price Alert badge */}
        <div className="absolute top-2 left-2 bg-[#FFD300] text-white text-xs font-medium px-2 py-1 rounded-full shadow">
          Price Alert
        </div>

        {/* Delete button */}
        <button
          className="absolute top-2 right-2 bg-white/80 hover:bg-white rounded-full p-2 shadow"
          onClick={() => onRemove?.(id)}
        >
          <Trash2 className="w-5 h-5 text-black" />
        </button>
      </div>

      {/* Body */}
      <div className="p-5">
        <h2 className="text-lg font-semibold mb-2">{title}</h2>

        {/* Rating */}
        <Rating value={rating} count={ratingCount} />

        {/* Price */}
        <div className="flex items-center gap-2 mt-2">
          <p className="text-xl font-bold">
            {price.etb
              ? `${price.etb} ETB`
              : price.usd
              ? `$${price.usd}`
              : "N/A"}
          </p>
          {oldPrice && (
            <span className="line-through text-gray-400 text-sm">
              {typeof oldPrice === "object"
                ? oldPrice.etb
                  ? `${oldPrice.etb} ETB`
                  : `$${oldPrice.usd}`
                : oldPrice}
            </span>
          )}
        </div>

        <p className="text-xs text-gray-500 mt-1">Checked {checked} ago</p>

        {/* Toggle */}
        <div className="flex items-center mt-3">
          <span className="text-sm font-medium mr-2 text-gray-500">
            Price Alert
          </span>
          <ToggleSwitch
            checked={priceAlertOn} // controlled
            onChange={() => onToggleAlert?.(id)}
          />
        </div>

        {/* Actions */}
        <div className="flex items-center gap-2 mt-4">
          <OpenLinkButton deeplinkUrl={deeplinkUrl} />
          <button className="px-4 py-2 rounded-lg border text-sm hover:bg-gray-100 dark:hover:bg-gray-700 transition">
            Compare
          </button>
        </div>

        {/* Seller */}
        <p className="text-xs text-gray-500 mt-3">Sold by {seller}</p>
      </div>
    </div>
  );
}
