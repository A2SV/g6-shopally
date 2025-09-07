// shopally-web/src/app/components/saved-items/SavedItemCard.tsx
"use client";

import { useDarkMode } from "@/app/components/ProfileComponents/DarkModeContext";
import ToggleSwitch from "@/app/components/saved-items/ToggleSwitch";
import { useLanguage } from "@/hooks/useLanguage";
import { Trash2 } from "lucide-react";
import { SavedItemUI } from "../../../types/types";
import { useState } from "react";
import { formatPriceForEthiopia } from "@/utils/priceUtils";
import { FaHeart } from "react-icons/fa";

// Extend SavedItemUI with optional callbacks
interface SavedItemCardProps extends SavedItemUI {
  onRemove?: (id: string) => void;
  onUpdatePrice?: (id: string, price: SavedItemUI["price"]) => void;
  onToggleAlert?: (id: string) => void;
  onPlaceOrder?: (
    productId: string,
    productTitle: string,
    price: SavedItemUI["price"]
  ) => void;
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
  onPlaceOrder,
}: SavedItemCardProps) {
  const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();

  // Image loading states
  const [imageLoaded, setImageLoaded] = useState(false);
  const [imageError, setImageError] = useState(false);
  const [retryCount, setRetryCount] = useState(0);

  // Function to get image URL with retry logic
  const getImageUrl = () => {
    if (retryCount > 0) {
      return `https://images.weserv.nl/?url=${encodeURIComponent(imageUrl)}`;
    }
    return imageUrl;
  };

  return (
    <div
      className={`rounded-2xl border shadow p-6 space-y-6 w-full max-w-md mx-auto sm:max-w-none transition-colors ${
        isDarkMode
          ? "bg-gray-800 border-gray-700 text-white"
          : "bg-white border-gray-200 text-gray-900"
      }`}
      style={{
        backgroundColor: "var(--color-bg-card)",
        borderColor: "var(--color-border-primary)",
        boxShadow: "var(--color-shadow)",
      }}
    >
      {/* Image container */}
      <div className="aspect-square rounded-xl overflow-hidden transition-colors bg-gray-200 relative">
        {/* Loading skeleton */}
        {!imageLoaded && !imageError && (
          <div className="w-full h-full bg-gray-300 dark:bg-gray-600 animate-pulse flex items-center justify-center">
            <div className="text-gray-500 dark:text-gray-400 text-sm">
              Loading...
            </div>
          </div>
        )}

        {/* Fallback UI */}
        {imageError && (
          <div className="w-full h-full bg-gray-300 dark:bg-gray-600 flex flex-col items-center justify-center text-gray-500 dark:text-gray-400">
            <div className="text-4xl mb-2">📷</div>
            <div className="text-sm text-center px-4">Image not available</div>
          </div>
        )}

        {/* Actual image */}
        <img
          src={getImageUrl()}
          alt={title}
          className={`w-full h-full object-cover transition-opacity duration-300 ${
            imageError ? "opacity-0" : "opacity-100"
          }`}
          loading="lazy"
          referrerPolicy="no-referrer"
          crossOrigin="anonymous"
          onLoad={() => setImageLoaded(true)}
          onError={() => {
            if (retryCount === 0) {
              setRetryCount(1);
            } else {
              setImageError(true);
            }
          }}
        />

        {/* Price Alert badge */}
        <div className="absolute top-2 left-2 bg-[#FFD300] text-white text-xs font-medium px-2 py-1 rounded-full shadow">
          {t("Price Alert")}
        </div>

        {/* Delete button */}
        <button
          className="absolute top-2 right-2 bg-red-500 hover:bg-red-600 text-white rounded-full p-2 shadow-lg transition-colors"
          onClick={() => onRemove?.(id)}
          title="Delete item"
        >
          <Trash2 className="w-5 h-5" />
        </button>
      </div>

      {/* Body */}
      <div className="space-y-4">
        <h3
          className="text-lg font-semibold transition-colors cursor-pointer"
          style={{ color: "var(--color-text-primary)" }}
          title={title} // full title on hover
        >
          {title.split(" ").slice(0, 5).join(" ")}
          {title.split(" ").length > 5 ? " ..." : ""}
        </h3>

        <div className="flex items-center justify-between">
          <span
            className="text-2xl font-bold transition-colors"
            style={{ color: "var(--color-accent-primary)" }}
          >
            {formatPriceForEthiopia(price)}
          </span>
          <div
            className="flex items-center gap-1 text-sm transition-colors"
            style={{ color: "var(--color-text-tertiary)" }}
          >
            <span>⭐</span>
            <span>
              {rating} ({ratingCount})
            </span>
          </div>
        </div>

        {/* Key Features */}
        <div className="space-y-3">
          <h4
            className="text-sm font-medium transition-colors"
            style={{ color: "var(--color-text-primary)" }}
          >
            {t("Key Features")}
          </h4>
          <div className="space-y-2">
            <div className="flex items-center gap-2">
              <span className="w-2 h-2 rounded-full bg-green-500"></span>
              <span
                className="text-sm transition-colors"
                style={{ color: "var(--color-text-secondary)" }}
              >
                {t("High quality product")}
              </span>
            </div>
            <div className="flex items-center gap-2">
              <span
                className="text-sm transition-colors"
                style={{ color: "var(--color-text-secondary)" }}
              >
                {t("Fast shipping available")}
              </span>
            </div>
          </div>
        </div>

        {/* Price Alert Toggle */}
        <div className="flex items-center justify-between">
          <span
            className="text-sm font-medium transition-colors"
            style={{ color: "var(--color-text-secondary)" }}
          >
            {t("Price Alert")}
          </span>
          <ToggleSwitch
            checked={priceAlertOn}
            onChange={() => onToggleAlert?.(id)}
          />
        </div>

        {/* Actions */}
        <div className="space-y-3">
          <button
            className="w-full font-medium py-3 px-6 rounded-xl hover:opacity-80 transition-colors border"
            style={{
              backgroundColor: "var(--color-bg-tertiary)",
              color: "var(--color-text-primary)",
              borderColor: "var(--color-border-primary)",
            }}
            onClick={() => onPlaceOrder?.(id, title, price)}
          >
            {t("Add To Compare")}
          </button>

          <div className="flex gap-3">
            <button
              className="flex-1 font-medium py-3 px-6 rounded-xl hover:opacity-80 transition-colors"
              style={{
                backgroundColor: "var(--color-accent-primary)",
                color: "var(--color-text-button)",
              }}
              onClick={() => {
                // Track the order when user clicks "Buy from AliExpress"
                onPlaceOrder?.(id, title, price);
                // Open AliExpress link in new tab
                window.open(deeplinkUrl, "_blank", "noopener,noreferrer");
              }}
            >
              {t("Buy from AliExpress")}
            </button>

            <button
              className="p-3 rounded-xl hover:opacity-80 transition-colors"
              style={{
                backgroundColor: "var(--color-bg-tertiary)",
                color: "var(--color-text-primary)",
              }}
              onClick={() => onRemove?.(id)}
            >
              <FaHeart
                className="w-5 h-5 transition-colors"
                style={{
                  color: "var(--color-accent-primary)",
                }}
              />
            </button>
          </div>
        </div>

        {/* Additional Info */}
        <div className="text-xs text-gray-500 space-y-1">
          <p>
            {t("Checked")} {checked} {t("ago")}
          </p>
          <p>
            {t("Sold by")} {seller}
          </p>
        </div>
      </div>
    </div>
  );
}
