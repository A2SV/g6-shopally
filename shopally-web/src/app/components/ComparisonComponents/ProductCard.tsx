// //src/app/components/ComparePage/ProductCard.tsx
// import { useDarkMode } from "@/app/components/ProfileComponents/DarkModeContext";
import { useLanguage } from "@/hooks/useLanguage";
import { useSavedItems } from "@/hooks/useSavedItems";
import { Product, SavedItem } from "@/types/types";
import { FaHeart } from "react-icons/fa";
import { FaCheck } from "react-icons/fa6";
import { IoIosStar } from "react-icons/io";
import { useState } from "react";
import { formatPriceForEthiopia } from "@/utils/priceUtils";

export const ProductCard: React.FC<{ product: Product }> = ({ product }) => {
  // const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();
  const { savedItems, saveItem, removeItem, placeOrder } = useSavedItems();
  const [imageLoaded, setImageLoaded] = useState(false);
  const [imageError, setImageError] = useState(false);
  const [retryCount, setRetryCount] = useState(0);

  // Check if this product is already saved
  const isSaved = savedItems.some((item) => item.id === product.id);

  // Function to get alternative image URL
  const getImageUrl = () => {
    if (retryCount === 0) {
      return product.imageUrl;
    }
    // Try alternative URL format or proxy
    return `https://images.weserv.nl/?url=${encodeURIComponent(
      product.imageUrl
    )}`;
  };

  const handleHeartClick = () => {
    if (isSaved) {
      // Remove from saved items
      removeItem(product.id);
    } else {
      // Add to saved items
      const savedItem: SavedItem = {
        id: product.id,
        title: product.title,
        imageUrl: product.imageUrl,
        aiMatchPercentage: product.aiMatchPercentage,
        price: product.price,
        productRating: product.productRating ?? 0,
        sellerScore: product.numberSold ?? 0, // adjust if you have a correct sellerScore field
        deliveryEstimate: product.deliveryEstimate,
        summaryBullets: product.summaryBullets,
        deeplinkUrl: product.deeplinkUrl,
      };
      saveItem(savedItem); // saves it to localStorage via your hook
    }
  };

  return (
    <div
      className="rounded-2xl border shadow p-6 space-y-6 w-full max-w-md mx-auto lg:max-w-none transition-colors"
      style={{
        backgroundColor: "var(--color-bg-card)",
        borderColor: "var(--color-border-primary)",
        boxShadow: "var(--color-shadow)",
      }}
    >
      <div
        className="aspect-square rounded-xl overflow-hidden transition-colors relative"
        style={{
          backgroundColor: "var(--color-bg-tertiary)",
        }}
      >
        {/* Loading skeleton - only show when image is not loaded and no error */}
        {!imageLoaded && !imageError && (
          <div className="absolute inset-0 bg-gray-200 animate-pulse flex items-center justify-center">
            <div className="w-8 h-8 border-2 border-gray-300 border-t-blue-500 rounded-full animate-spin"></div>
          </div>
        )}

        <img
          src={getImageUrl()}
          alt={product.title}
          className={`w-full h-full object-cover transition-opacity duration-300 ${
            imageError ? "opacity-0" : "opacity-100"
          }`}
          loading="lazy"
          referrerPolicy="no-referrer"
          crossOrigin="anonymous"
          onLoad={() => {
            setImageLoaded(true);
            setImageError(false);
          }}
          onError={(e) => {
            if (retryCount < 1) {
              setRetryCount((prev) => prev + 1);
              setImageError(false);
            } else {
              setImageError(true);
              setImageLoaded(false);
            }
          }}
        />

        {/* Fallback image - only show when there's an error */}
        {imageError && (
          <div className="absolute inset-0 bg-gray-200 flex items-center justify-center">
            <div className="text-center text-gray-500">
              <div className="w-12 h-12 mx-auto mb-2 bg-gray-300 rounded-full flex items-center justify-center">
                <span className="text-2xl">📷</span>
              </div>
              <p className="text-sm">Image not available</p>
            </div>
          </div>
        )}
      </div>
      <div className="space-y-4">
        <h3
          className="text-lg font-semibold transition-colors line-clamp-2"
          style={{ color: "var(--color-text-primary)" }}
        >
          {product.title}
        </h3>
        <div className="flex items-center justify-between">
          <span
            className="text-2xl font-bold transition-colors"
            style={{ color: "var(--color-accent-primary)" }}
          >
            {formatPriceForEthiopia(product.price)}
          </span>
          <div
            className="flex items-center gap-1 text-sm transition-colors"
            style={{ color: "var(--color-text-tertiary)" }}
          >
            <IoIosStar
              className="w-4 h-4 transition-colors"
              style={{ fill: "var(--color-text-primary)" }}
            />
            <span>
              {product.productRating} ({product.numberSold} {t("reviews")})
            </span>
          </div>
        </div>
        <div className="space-y-3">
          <h4
            className="text-sm font-medium transition-colors"
            style={{ color: "var(--color-text-primary)" }}
          >
            {t("Key Features")}
          </h4>
          {product.summaryBullets.map((feature: string, i: number) => (
            <div key={i} className="flex items-center gap-2">
              <FaCheck
                className="w-4 h-4 transition-colors"
                style={{ color: "var(--color-text-primary)" }}
              />
              <span
                className="text-sm transition-colors line-clamp-2"
                style={{ color: "var(--color-text-secondary)" }}
              >
                {feature}
              </span>
            </div>
          ))}
        </div>
        <div className="flex gap-3">
          <button
            className="flex-1 font-medium py-3 px-6 rounded-xl hover:opacity-80 transition-colors"
            style={{
              backgroundColor: "var(--color-accent-primary)",
              color: "var(--color-text-button)",
            }}
            onClick={() => {
              // Track the order when user clicks "Buy from AliExpress"
              placeOrder(product.id, product.title, product.price);
              // Open AliExpress link in new tab
              window.open(product.deeplinkUrl, "_blank", "noopener,noreferrer");
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
            onClick={handleHeartClick}
          >
            <FaHeart
              className="w-5 h-5 transition-colors"
              style={{
                color: isSaved
                  ? "var(--color-accent-primary)"
                  : "var(--color-text-primary)",
              }}
            />
          </button>
        </div>
      </div>
    </div>
  );
};
