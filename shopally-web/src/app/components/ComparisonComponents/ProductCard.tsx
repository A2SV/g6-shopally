// //src/app/components/ComparePage/ProductCard.tsx
// import { useDarkMode } from "@/app/components/ProfileComponents/DarkModeContext";
import { useLanguage } from "@/hooks/useLanguage";
import { useSavedItems } from "@/hooks/useSavedItems";
import { Product, SavedItem } from "@/types/types";
import { FaHeart } from "react-icons/fa";
import { FaCheck } from "react-icons/fa6";
import { IoIosStar } from "react-icons/io";

export const ProductCard: React.FC<{ product: Product }> = ({ product }) => {
  // const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();
  const { savedItems, saveItem, removeItem } = useSavedItems();

  // Check if this product is already saved
  const isSaved = savedItems.some((item) => item.id === product.id);

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
        <img
          src={product.imageUrl}
          alt={product.title}
          className="w-full h-full object-cover"
        />
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
            {product.price.etb
              ? `${Number(product.price.etb).toFixed(2)} ETB`
              : `$${Number(product.price.usd).toFixed(2)}`}
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
          {product.summaryBullets.map((feature, i) => (
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
            onClick={() => window.open(product.deeplinkUrl, "_blank")}
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
