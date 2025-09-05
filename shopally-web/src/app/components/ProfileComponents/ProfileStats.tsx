"use client";

import { useState, useEffect } from "react";
import { useDarkMode } from "./DarkModeContext";
import { useLanguage } from "../../../hooks/useLanguage";
import { LoadingSpinner } from "../LoadingSpinner";
import { useSavedItems } from "@/hooks/useSavedItems";

export default function ProfileStats() {
  const [isLoading, setIsLoading] = useState(true);
  const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();
  const { savedItems, orders } = useSavedItems();

  // Simulate loading on component mount
  useEffect(() => {
    const timer = setTimeout(() => {
      setIsLoading(false);
    }, 1500);
    return () => clearTimeout(timer);
  }, []);

  // Calculate statistics from saved items and orders
  const stats = [
    {
      value: savedItems.length.toString(),
      label: t("Products Saved"),
    },
    {
      value: orders.length.toString(),
      label: t("Orders Placed"),
    },
    {
      value: `${savedItems
        .reduce((total, item) => total + (item.price.etb || 0), 0)
        .toFixed(0)} ETB`,
      label: t("Total Value"),
    },
    {
      value: `${Math.round(
        savedItems.reduce((total, item) => total + item.aiMatchPercentage, 0) /
          Math.max(savedItems.length, 1)
      )}%`,
      label: t("Avg AI Match Rate"),
    },
  ];

  // Show loading state
  if (isLoading) {
    return (
      <div
        className="flex-1 transition-colors"
        style={{
          backgroundColor: "var(--color-bg-primary)",
        }}
      >
        <div
          className="rounded-lg shadow-sm border p-4 sm:p-6 transition-colors"
          style={{
            backgroundColor: "var(--color-bg-card)",
            borderColor: "var(--color-border-primary)",
          }}
        >
          <div className="flex items-center justify-center py-12">
            <LoadingSpinner size="xl" color="yellow" />
          </div>
        </div>
      </div>
    );
  }

  return (
    <div
      className="flex-1 transition-colors"
      style={{
        backgroundColor: "var(--color-bg-primary)",
      }}
    >
      {/* Statistics Card */}
      <div
        className="rounded-lg shadow-sm border p-4 sm:p-6 transition-colors"
        style={{
          backgroundColor: "var(--color-bg-card)",
          borderColor: "var(--color-border-primary)",
        }}
      >
        {/* Header */}
        <h2
          className="text-lg sm:text-xl font-semibold mb-4 sm:mb-6 transition-colors"
          style={{ color: "var(--color-text-primary)" }}
        >
          {t("Account Statistics")}
        </h2>

        {/* Statistics Grid */}
        <div className="grid grid-cols-2 gap-3 sm:gap-4">
          {stats.map((stat, index) => (
            <div
              key={index}
              className="border rounded-lg p-3 sm:p-4 text-center shadow-sm transition-colors"
              style={{
                backgroundColor: "var(--color-bg-card)",
                borderColor: "var(--color-border-primary)",
              }}
            >
              <div
                className="text-xl sm:text-2xl font-bold mb-1"
                style={{ color: "var(--color-accent-primary)" }}
              >
                {stat.value}
              </div>
              <div
                className="text-sm transition-colors"
                style={{ color: "var(--color-text-secondary)" }}
              >
                {stat.label}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
