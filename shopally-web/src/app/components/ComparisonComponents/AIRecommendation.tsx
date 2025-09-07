// components/compare/AIRecommendation.tsx
// import { useDarkMode } from "@/app/components/ProfileComponents/DarkModeContext";
import { useLanguage } from "@/hooks/useLanguage";
// import { useSavedItems } from "@/hooks/useSavedItems";
import type { ComparisonItem } from "@/types/Compare/Comparison";
import { Check, X } from "lucide-react";
import { BsLightningChargeFill } from "react-icons/bs";
import { FaCircleCheck } from "react-icons/fa6";
import { IoIosStar } from "react-icons/io";
import { TiDelete } from "react-icons/ti";

type AIRecommendationProps = {
  comparison: ComparisonItem[];
};

export const AIRecommendation: React.FC<AIRecommendationProps> = ({
  comparison,
}) => {
  // const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();
  // const { placeOrder } = useSavedItems();

  if (!comparison || comparison.length === 0) return null;

  // Mock AI analysis - in real app, this would come from backend
  const productAnalysis = comparison.map((item) => ({
    product: item.product,
    pros: item.synthesis.pros,
    cons: item.synthesis.cons,
    score: item.product.aiMatchPercentage,
  }));

  // Pick top product by aiMatchPercentage, tie-breaker: sellerScore
  // Find the recommended product by isBestValue
  const recommendedItem = comparison.find((item) => item.synthesis.isBestValue);

  if (!recommendedItem) return null; // fallback in case API doesn't mark any

  const recommendedProduct = recommendedItem.product;

  return (
    <div
      className="rounded-2xl border shadow-lg p-6 space-y-6 transition-colors"
      style={{
        backgroundColor: "var(--color-bg-card)",
        borderColor: "var(--color-border-primary)",
        boxShadow: "var(--color-shadow)",
      }}
    >
      <div className="flex items-center gap-4">
        <div className="w-8 h-8 rounded-full bg-brand-saffron flex items-center justify-center">
          <IoIosStar
            className="w-5 h-5 transition-colors"
            style={{ color: "var(--color-text-primary)" }}
          />
        </div>
        <h2
          className="text-xl font-semibold transition-colors"
          style={{ color: "var(--color-text-primary)" }}
        >
          {t("AI Recommendation")}
        </h2>
      </div>

      <div
        className="rounded-xl p-6 transition-colors"
        style={{
          backgroundColor: "var(--color-bg-tertiary)",
        }}
      >
        <div className="flex items-start gap-4 mb-6">
          <div
            className="w-12 h-12 rounded-full flex items-center justify-center transition-colors"
            style={{
              backgroundColor: "var(--color-accent-primary)",
            }}
          >
            <BsLightningChargeFill
              className="w-6 h-6 transition-colors"
              style={{ color: "var(--color-text-button)" }}
            />
          </div>
          <div className="flex-1">
            <div className="flex items-center gap-2 mb-2">
              <span
                className="font-semibold transition-colors"
                style={{ color: "var(--color-text-primary)" }}
              >
                {t("Our Top Pick:")}
              </span>
              <span
                className="font-semibold transition-colors"
                style={{ color: "var(--color-text-primary)" }}
              >
                {recommendedProduct.title}
              </span>
            </div>
            <p
              className="text-sm mb-4 transition-colors"
              style={{ color: "var(--color-text-secondary)" }}
            >
              {t(
                "Based on price, ratings, and versatility, this product offers a strong balance of value and everyday utility"
              )}
            </p>
            <div className="flex items-center gap-4">
              <span
                className="text-sm font-medium transition-colors"
                style={{ color: "var(--color-accent-primary)" }}
              >
                ${recommendedProduct.price.usd}
              </span>
              <div className="flex items-center gap-1">
                <span
                  className="text-sm font-medium transition-colors"
                  style={{ color: "var(--color-text-primary)" }}
                >
                  {recommendedProduct.productRating}
                </span>
                <IoIosStar
                  className="w-4 h-4 transition-colors"
                  style={{ fill: "var(--color-text-primary)" }}
                />
                <span
                  className="text-sm transition-colors"
                  style={{ color: "var(--color-text-secondary)" }}
                >
                  ({recommendedProduct.numberSold} {t("sold")})
                </span>
              </div>
            </div>
          </div>
        </div>

        <div className="flex gap-3">
          <button
            className="px-4 py-2 rounded-lg font-medium hover:opacity-80 transition-colors"
            style={{
              backgroundColor: "var(--color-accent-primary)",
              color: "var(--color-text-primary)",
            }}
            onClick={() =>
              window.open(recommendedProduct.deeplinkUrl, "_blank")
            }
          >
            {t("Buy from AliExpress")}
          </button>
        </div>
      </div>

      {/* Product Analysis */}
      <div className="space-y-6">
        <h3
          className="text-lg font-semibold transition-colors"
          style={{ color: "var(--color-text-primary)" }}
        >
          {t("Product Analysis")}
        </h3>

        {productAnalysis.map((analysis) => (
          <div key={analysis.product.id} className="space-y-4">
            <h3
              className="font-semibold transition-colors"
              style={{ color: "var(--color-text-primary)" }}
            >
              {analysis.product.title}
            </h3>

            <div className="space-y-3">
              <div className="flex items-center gap-2">
                <FaCircleCheck
                  className="w-4 h-4 transition-colors"
                  style={{ color: "var(--color-text-primary)" }}
                />
                <span
                  className="text-sm font-medium transition-colors"
                  style={{ color: "var(--color-text-primary)" }}
                >
                  {t("Pros")}
                </span>
              </div>
              {analysis.pros.map((pro, i) => (
                <div key={i} className="flex items-center gap-2 ml-6">
                  <Check
                    className="w-3 h-3 transition-colors"
                    style={{ color: "var(--color-text-primary)" }}
                  />
                  <span
                    className="text-sm transition-colors"
                    style={{ color: "var(--color-text-secondary)" }}
                  >
                    {pro}
                  </span>
                </div>
              ))}
            </div>

            <div className="space-y-3">
              <div className="flex items-center gap-2">
                <TiDelete
                  className="w-4 h-4 transition-colors"
                  style={{ color: "var(--color-text-primary)" }}
                />
                <span
                  className="text-sm font-medium transition-colors"
                  style={{ color: "var(--color-text-primary)" }}
                >
                  {t("Cons")}
                </span>
              </div>
              {analysis.cons.map((con, i) => (
                <div key={i} className="flex items-center gap-2 ml-6">
                  <X
                    className="w-3 h-3"
                    style={{ color: "var(--color-accent-error)" }}
                  />
                  <span
                    className="text-sm transition-colors"
                    style={{ color: "var(--color-text-secondary)" }}
                  >
                    {con}
                  </span>
                </div>
              ))}
            </div>

            <div className="space-y-2">
              <span
                className="text-xs font-medium transition-colors"
                style={{ color: "var(--color-text-primary)" }}
              >
                {t("AI Score")}
              </span>
              <div
                className="w-full rounded-full h-2 transition-colors"
                style={{
                  backgroundColor: "var(--color-bg-tertiary)",
                }}
              >
                <div
                  className="h-2 rounded-full transition-colors"
                  style={{
                    backgroundColor:
                      analysis.product.id &&
                      comparison.find(
                        (c) => c.product.id === analysis.product.id
                      )?.synthesis.isBestValue
                        ? "var(--color-accent-primary)" // yellow
                        : "var(--color-text-tertiary)", // gray
                    width: `${analysis.score}%`,
                  }}
                />
              </div>
              <span
                className="text-xs transition-colors"
                style={{ color: "var(--color-text-tertiary)" }}
              >
                {analysis.score}% {t("match")}
              </span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};
