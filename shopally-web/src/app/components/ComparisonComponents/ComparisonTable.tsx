// components/compare/ComparisonTable.tsx
import type { ComparisonItem } from "@/types/Compare/Comparison";
import { FC } from "react";
import { IoIosStar } from "react-icons/io";
import { useDarkMode } from "@/app/components/ProfileComponents/DarkModeContext";
import { useLanguage } from "@/hooks/useLanguage";

interface ComparisonTableProps {
  comparison: ComparisonItem[];
}

export const ComparisonTable: FC<ComparisonTableProps> = ({ comparison }) => {
  const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();

  if (!comparison || comparison.length === 0) return null;

  return (
    <div
      className="overflow-x-auto rounded-xl sm:rounded-2xl border shadow-lg transition-colors"
      style={{
        backgroundColor: "var(--color-bg-card)",
        borderColor: "var(--color-border-primary)",
        boxShadow: "var(--color-shadow)",
      }}
    >
      <div className="p-3 sm:p-4 lg:p-5">
        <h1
          className="text-sm sm:text-base font-bold transition-colors"
          style={{ color: "var(--color-text-primary)" }}
        >
          {t("Detailed Comparison")}
        </h1>
      </div>
      <table className="w-full text-left border-collapse min-w-[400px] sm:min-w-[600px]">
        <thead>
          <tr style={{ backgroundColor: "var(--color-bg-secondary)" }}>
            <th
              className="px-2 sm:px-3 lg:px-4 py-2 sm:py-3 text-xs sm:text-sm font-semibold transition-colors"
              style={{ color: "var(--color-text-primary)" }}
            >
              {t("Feature")}
            </th>
            {comparison.map((c) => (
              <th
                key={c.product.id}
                className="px-2 sm:px-3 lg:px-4 py-2 sm:py-3 text-xs sm:text-sm font-semibold text-center transition-colors"
                style={{ color: "var(--color-text-primary)" }}
              >
                {c.product.title}
              </th>
            ))}
          </tr>
        </thead>
        <tbody
          className="divide-y transition-colors"
          style={{ borderColor: "var(--color-border-secondary)" }}
        >
          {/* Price */}
          <tr>
            <td
              className="px-2 sm:px-3 lg:px-4 py-2 sm:py-3 text-xs sm:text-sm font-semibold transition-colors"
              style={{ color: "var(--color-text-primary)" }}
            >
              {t("Price (USD)")}
            </td>
            {comparison.map((c) => (
              <td
                key={c.product.id}
                className="px-2 sm:px-3 lg:px-4 py-2 sm:py-3 text-center text-xs sm:text-sm transition-colors"
                style={{ color: "var(--color-text-secondary)" }}
              >
                ${c.product.price.usd}
              </td>
            ))}
          </tr>

          {/* Rating */}
          <tr>
            <td
              className="px-2 sm:px-3 lg:px-4 py-2 sm:py-3 text-xs sm:text-sm font-semibold transition-colors"
              style={{ color: "var(--color-text-primary)" }}
            >
              {t("Rating")}
            </td>
            {comparison.map((c) => (
              <td
                key={c.product.id}
                className="px-2 sm:px-3 lg:px-4 py-2 sm:py-3 text-center text-xs sm:text-sm transition-colors"
                style={{ color: "var(--color-text-secondary)" }}
              >
                <div className="flex items-center justify-center gap-1">
                  {c.product.productRating}
                  <IoIosStar
                    className="w-3 h-3 sm:w-4 sm:h-4 transition-colors"
                    style={{ fill: "var(--color-text-primary)" }}
                  />
                </div>
              </td>
            ))}
          </tr>

          {/* Delivery Estimate */}
          <tr>
            <td
              className="px-2 sm:px-3 lg:px-4 py-2 sm:py-3 text-xs sm:text-sm font-semibold transition-colors"
              style={{ color: "var(--color-text-primary)" }}
            >
              {t("Delivery Estimate")}
            </td>
            {comparison.map((c) => (
              <td
                key={c.product.id}
                className="px-2 sm:px-3 lg:px-4 py-2 sm:py-3 text-center text-xs sm:text-sm transition-colors"
                style={{ color: "var(--color-text-secondary)" }}
              >
                {c.product.deliveryEstimate}
              </td>
            ))}
          </tr>

          {/* Seller Score */}
          <tr>
            <td
              className="px-2 sm:px-3 lg:px-4 py-2 sm:py-3 text-xs sm:text-sm font-semibold transition-colors"
              style={{ color: "var(--color-text-primary)" }}
            >
              {t("Seller Score")}
            </td>
            {comparison.map((c) => (
              <td
                key={c.product.id}
                className="px-2 sm:px-3 lg:px-4 py-2 sm:py-3 text-center text-xs sm:text-sm transition-colors"
                style={{ color: "var(--color-text-secondary)" }}
              >
                {c.product.sellerScore}
              </td>
            ))}
          </tr>
        </tbody>
      </table>
    </div>
  );
};
