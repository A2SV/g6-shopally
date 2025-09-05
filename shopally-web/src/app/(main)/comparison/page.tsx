// src/app/(main)/comparison/page.tsx
"use client";

import { AIRecommendation } from "@/app/components/ComparisonComponents/AIRecommendation";
import { ComparisonTable } from "@/app/components/ComparisonComponents/ComparisonTable";
import { ProductCard } from "@/app/components/ComparisonComponents/ProductCard";
import type { ComparisonItem } from "@/types/Compare/Comparison";
import { fetchComparison } from "@/utils/Compare/CompareUtils";
import { useEffect, useState } from "react";
import { useDarkMode } from "@/app/components/ProfileComponents/DarkModeContext";
import { useLanguage } from "@/hooks/useLanguage";

export default function ComparePage() {
  const [comparison, setComparison] = useState<ComparisonItem[]>([]);
  const [loading, setLoading] = useState(true);
  const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();

  useEffect(() => {
    async function loadData() {
      try {
        const items = await fetchComparison();
        if (items && items.length > 0) {
          setComparison(items);
        } else {
          // Fallback to mock data if API fails or returns empty
          const mockComparison: ComparisonItem[] = [
            {
              product: {
                id: "1",
                title: "Advanced Smart Fitness Tracker",
                imageUrl:
                  "https://via.placeholder.com/300x300/4F46E5/FFFFFF?text=Product+1",
                price: {
                  usd: 45,
                  etb: 2650,
                  fxTimestamp: new Date().toISOString(),
                },
                productRating: 4.3,
                sellerScore: 95,
                deliveryEstimate: "7-14 days",
                aiMatchPercentage: 92,
                summaryBullets: [
                  "High quality product with advanced sensors",
                  "Fast shipping available worldwide",
                  "Water resistant and durable design",
                ],
                deeplinkUrl: "#",
              },
              synthesis: {
                pros: ["High quality", "Good price", "Fast delivery"],
                cons: ["Long delivery time", "Limited warranty"],
                isBestValue: true,
              },
            },
            {
              product: {
                id: "2",
                title: "Premium Wireless Headphones",
                imageUrl:
                  "https://via.placeholder.com/300x300/10B981/FFFFFF?text=Product+2",
                price: {
                  usd: 89,
                  etb: 5200,
                  fxTimestamp: new Date().toISOString(),
                },
                productRating: 4.7,
                sellerScore: 98,
                deliveryEstimate: "5-10 days",
                aiMatchPercentage: 88,
                summaryBullets: [
                  "Premium sound quality with noise cancellation",
                  "Long battery life up to 30 hours",
                  "Comfortable over-ear design",
                ],
                deeplinkUrl: "#",
              },
              synthesis: {
                pros: ["Excellent sound", "Long battery", "Comfortable"],
                cons: ["Higher price", "Bulky design"],
                isBestValue: false,
              },
            },
            {
              product: {
                id: "3",
                title: "Smart Home Security Camera",
                imageUrl:
                  "https://via.placeholder.com/300x300/F59E0B/FFFFFF?text=Product+3",
                price: {
                  usd: 67,
                  etb: 3900,
                  fxTimestamp: new Date().toISOString(),
                },
                productRating: 4.5,
                sellerScore: 92,
                deliveryEstimate: "10-15 days",
                aiMatchPercentage: 85,
                summaryBullets: [
                  "1080p HD video quality",
                  "Night vision and motion detection",
                  "Easy mobile app control",
                ],
                deeplinkUrl: "#",
              },
              synthesis: {
                pros: ["Good video quality", "Night vision", "Easy setup"],
                cons: ["Medium delivery time", "Basic features"],
                isBestValue: false,
              },
            },
          ];
          setComparison(mockComparison);
        }
      } catch (err) {
        console.error("Error fetching comparison:", err);
        // Use mock data as fallback
        const mockComparison: ComparisonItem[] = [
          {
            product: {
              id: "1",
              title: "Advanced Smart Fitness Tracker",
              imageUrl:
                "https://via.placeholder.com/300x300/4F46E5/FFFFFF?text=Product+1",
              price: {
                usd: 45,
                etb: 2650,
                fxTimestamp: new Date().toISOString(),
              },
              productRating: 4.3,
              sellerScore: 95,
              deliveryEstimate: "7-14 days",
              aiMatchPercentage: 92,
              summaryBullets: [
                "High quality product with advanced sensors",
                "Fast shipping available worldwide",
                "Water resistant and durable design",
              ],
              deeplinkUrl: "#",
            },
            synthesis: {
              pros: ["High quality", "Good price", "Fast delivery"],
              cons: ["Long delivery time", "Limited warranty"],
              isBestValue: true,
            },
          },
          {
            product: {
              id: "2",
              title: "Premium Wireless Headphones",
              imageUrl:
                "https://via.placeholder.com/300x300/10B981/FFFFFF?text=Product+2",
              price: {
                usd: 89,
                etb: 5200,
                fxTimestamp: new Date().toISOString(),
              },
              productRating: 4.7,
              sellerScore: 98,
              deliveryEstimate: "5-10 days",
              aiMatchPercentage: 88,
              summaryBullets: [
                "Premium sound quality with noise cancellation",
                "Long battery life up to 30 hours",
                "Comfortable over-ear design",
              ],
              deeplinkUrl: "#",
            },
            synthesis: {
              pros: ["Excellent sound", "Long battery", "Comfortable"],
              cons: ["Higher price", "Bulky design"],
              isBestValue: false,
            },
          },
          {
            product: {
              id: "3",
              title: "Smart Home Security Camera",
              imageUrl:
                "https://via.placeholder.com/300x300/F59E0B/FFFFFF?text=Product+3",
              price: {
                usd: 67,
                etb: 3900,
                fxTimestamp: new Date().toISOString(),
              },
              productRating: 4.5,
              sellerScore: 92,
              deliveryEstimate: "10-15 days",
              aiMatchPercentage: 85,
              summaryBullets: [
                "1080p HD video quality",
                "Night vision and motion detection",
                "Easy mobile app control",
              ],
              deeplinkUrl: "#",
            },
            synthesis: {
              pros: ["Good video quality", "Night vision", "Easy setup"],
              cons: ["Medium delivery time", "Basic features"],
              isBestValue: false,
            },
          },
        ];
        setComparison(mockComparison);
      } finally {
        setLoading(false);
      }
    }
    loadData();
  }, []);

  if (loading) {
    return (
      <div
        className="min-h-screen flex items-center justify-center transition-colors"
        style={{
          backgroundColor: "var(--color-bg-primary)",
          color: "var(--color-text-primary)",
        }}
      >
        <p
          className="text-lg transition-colors"
          style={{ color: "var(--color-text-secondary)" }}
        >
          {t("Loading comparison...")}
        </p>
      </div>
    );
  }

  return (
    <div
      className="min-h-screen flex flex-col p-4 lg:p-8 transition-colors"
      style={{
        backgroundColor: "var(--color-bg-primary)",
        color: "var(--color-text-primary)",
      }}
    >
      {/* Page Title */}
      <h1
        className="text-2xl sm:text-3xl font-bold mb-6 sm:mb-8 transition-colors"
        style={{ color: "var(--color-text-primary)" }}
      >
        {t("Product Comparison")}
      </h1>

      <p
        className="text-sm sm:text-base mb-6 sm:mb-8 transition-colors"
        style={{ color: "var(--color-text-secondary)" }}
      >
        {t("Compare selected products to make the best choice")}
      </p>

      {/* Product cards */}
      <div className="grid gap-4 sm:gap-6 lg:gap-8 md:grid-cols-2 lg:grid-cols-3 mb-8">
        {comparison.map((c) => (
          <ProductCard key={c.product.id} product={c.product} />
        ))}
      </div>

      {/* Pass comparison directly */}
      <ComparisonTable comparison={comparison} />
      <AIRecommendation comparison={comparison} />
    </div>
  );
}
