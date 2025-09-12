"use client";

import React, { useRef } from "react";
import { useDarkMode } from "@/app/components/ProfileComponents/DarkModeContext";
import { useLanguage } from "@/hooks/useLanguage";
import { FaChevronLeft, FaChevronRight } from "react-icons/fa";

// Mock data
const deals = [
  { id: 1, leftTags: ["Deal"], rightTags: ["Ships in 3 days"] },
  { id: 2, leftTags: ["Deal"], rightTags: ["Limited stock"] },
  { id: 3, leftTags: [], rightTags: ["Fast delivery"] },
  { id: 4, leftTags: ["Deal"], rightTags: ["Ships in 5 days"] },
];

export default function FeaturedDeals() {
  const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();
  const scrollRef = useRef<HTMLDivElement>(null);

  const scroll = (direction: "left" | "right") => {
    if (scrollRef.current) {
      scrollRef.current.scrollBy({
        left: direction === "left" ? -320 : 320,
        behavior: "smooth",
      });
    }
  };

  return (
    <section
      className={`py-12 px-6 md:px-12 transition-colors ${
        isDarkMode
          ? "bg-[var(--color-bg-primary)] text-[var(--color-text-primary)]"
          : "bg-white text-gray-900"
      }`}
    >
      <div className="max-w-6xl mx-auto">
        {/* Title + Arrows */}
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-2xl font-bold">{t("Featured Deals")}</h2>
          <div className="flex gap-2">
            <button
              onClick={() => scroll("left")}
              className={`p-2 rounded-full ${
                isDarkMode ? "bg-gray-700" : "bg-gray-200"
              } hover:bg-opacity-80 transition`}
            >
              <FaChevronLeft />
            </button>
            <button
              onClick={() => scroll("right")}
              className={`p-2 rounded-full ${
                isDarkMode ? "bg-gray-700" : "bg-gray-200"
              } hover:bg-opacity-80 transition`}
            >
              <FaChevronRight />
            </button>
          </div>
        </div>

        {/* Cards */}
        <div
          ref={scrollRef}
          className="flex gap-4 overflow-x-auto scroll-smooth scrollbar-hide pb-1"
        >
          {deals.map((deal) => (
            <div
              key={deal.id}
              className="min-w-[300px] max-w-[300px] bg-white dark:bg-[var(--color-card-bg)] rounded-lg overflow-hidden shadow-sm flex-shrink-0"
            >
              {/* Image + Tags */}
              <div className="relative w-full h-[300px] bg-gray-300 dark:bg-gray-200 flex items-center justify-center">
                <span className="text-gray-500 text-sm">300 × 300</span>

                {/* Left tags */}
                <div className="absolute top-2 left-2 flex flex-wrap gap-2 z-10">
                  {deal.leftTags.map((tag, i) => (
                    <span
                      key={i}
                      className={`text-xs font-medium px-2 py-1 rounded-full text-white ${
                        tag === "Deal" ? "bg-green-500" : "bg-yellow-500"
                      }`}
                    >
                      {t(tag)}
                    </span>
                  ))}
                </div>

                {/* Right tags */}
                <div className="absolute top-2 right-2 flex flex-wrap gap-2 z-10">
                  {deal.rightTags.map((tag, i) => (
                    <span
                      key={i}
                      className={`text-xs font-medium px-2 py-1 rounded-full text-white ${
                        tag === "Deal" ? "bg-green-500" : "bg-yellow-500"
                      }`}
                    >
                      {t(tag)}
                    </span>
                  ))}
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
