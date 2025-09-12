"use client";

import { useLanguage } from "@/hooks/useLanguage";
import { useDarkMode } from "@/app/components/ProfileComponents/DarkModeContext";

export default function HowItWorks() {
  const { t } = useLanguage();
  const { isDarkMode } = useDarkMode();

  return (
    <section
      className={`py-12 px-6 md:px-12 transition-colors ${
        isDarkMode
          ? "bg-[var(--color-bg-primary)] text-[var(--color-text-primary)]"
          : "bg-white text-gray-900"
      }`}
    >
      <div className="max-w-4xl mx-auto text-center">
        {/* Title */}
        <h2 className="text-2xl lg:text-3xl font-bold text-[var(--color-text-primary)] mb-4">
          {t("How")} <span className="text-brand">ShopAlly</span> {t("Works")}
        </h2>

        <p className="text-gray-500 mb-10">
          {t("Get personalized product recommendations in three simple steps")}
        </p>

        {/* Steps */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          {/* Step 1 */}
          <div className="flex flex-col items-center text-center">
            <div className="w-12 h-12 flex items-center justify-center bg-[#FFD300] text-white rounded-lg mb-3">
              💬
            </div>
            <h3 className="font-semibold text-lg mb-2">{t("Ask a Question")}</h3>
            <p className="text-gray-500 text-sm">
              {t(
                "Type your question in English or Amharic about what you need to buy"
              )}
            </p>
          </div>

          {/* Step 2 */}
          <div className="flex flex-col items-center text-center">
            <div className="w-12 h-12 flex items-center justify-center bg-[#FFD300] text-white rounded-lg mb-3">
              💡
            </div>
            <h3 className="font-semibold text-lg mb-2">
              {t("Get Smart Recommendations")}
            </h3>
            <p className="text-gray-500 text-sm">
              {t(
                "Our AI analyzes your needs and suggests the best products with Ethiopian pricing"
              )}
            </p>
          </div>

          {/* Step 3 */}
          <div className="flex flex-col items-center text-center">
            <div className="w-12 h-12 flex items-center justify-center bg-[#FFD300] text-white rounded-lg mb-3">
              📊
            </div>
            <h3 className="font-semibold text-lg mb-2">{t("Compare & Buy")}</h3>
            <p className="text-gray-500 text-sm">
              {t(
                "Compare prices, reviews, and shipping options to make the best choice"
              )}
            </p>
          </div>
        </div>
      </div>
    </section>
  );
}
