// src/app/(main)/comparison/page.tsx
"use client";

import { AIRecommendation } from "@/app/components/ComparisonComponents/AIRecommendation";
import { ComparisonTable } from "@/app/components/ComparisonComponents/ComparisonTable";
import { ProductCard } from "@/app/components/ComparisonComponents/ProductCard";
import type { ComparisonItem } from "@/types/Compare/Comparison";
import { fetchComparison } from "@/utils/Compare/CompareUtils";
import { useEffect, useState } from "react";

export default function ComparePage() {
  const [comparison, setComparison] = useState<ComparisonItem[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function loadData() {
      try {
        const items = await fetchComparison();
        setComparison(items);
      } catch (err) {
        console.error("Error fetching comparison:", err);
      } finally {
        setLoading(false);
      }
    }
    loadData();
  }, []);

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <p className="text-brand-darkGray">Loading comparison...</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-brand-white">
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8 space-y-8">
        <div className="text-center space-y-4">
          <h1 className="text-2xl md:text-3xl font-bold text-brand-dark">
            Product Comparison
          </h1>
          <p className="text-brand-darkGray">
            Compare selected products to make the best choice
          </p>
        </div>

        {/* Product cards */}
        <div className="grid gap-8 md:grid-cols-2 lg:grid-cols-3">
          {comparison.map((c) => (
            <ProductCard key={c.product.id} product={c.product} />
          ))}
        </div>

        {/* Pass comparison directly */}
        <ComparisonTable comparison={comparison} />
        <AIRecommendation comparison={comparison} />
      </main>
    </div>
  );
}
