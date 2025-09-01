"use client";

import { useState } from "react";
import { X, SlidersHorizontal } from "lucide-react";

interface FilterSidebarProps {
  darkMode: boolean;
}

export default function FilterSidebar({ darkMode }: FilterSidebarProps) {
  const [isOpen, setIsOpen] = useState(false);

  // ✅ Keep track of selected filters
  const [filters, setFilters] = useState({
    brandNames: "",
    minPrice: "",
    maxPrice: "",
    minRating: 0,
  });

  // ✅ Validation state
  const [errors, setErrors] = useState<{ minPrice?: string; maxPrice?: string }>(
    {}
  );

  const handleRatingClick = (rating: number) => {
    setFilters((prev) => ({ ...prev, minRating: rating }));
  };

  const validateInputs = () => {
    const newErrors: { minPrice?: string; maxPrice?: string } = {};
    const min = Number(filters.minPrice);
    const max = Number(filters.maxPrice);

    if (filters.minPrice && isNaN(min)) {
      newErrors.minPrice = "Min price must be a number";
    } else if (min < 0) {
      newErrors.minPrice = "Min price cannot be negative";
    }

    if (filters.maxPrice && isNaN(max)) {
      newErrors.maxPrice = "Max price must be a number";
    } else if (max < 0) {
      newErrors.maxPrice = "Max price cannot be negative";
    }

    if (!newErrors.minPrice && !newErrors.maxPrice && min && max && min > max) {
      newErrors.maxPrice = "Max price must be greater than min price";
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const applyFilters = () => {
    if (validateInputs()) {
      console.log("✅ Filters to send:", filters);
      // here you can call fetch/axios to send filters to backend
    }
  };

  return (
    <div className="relative">
      {/* Button to open sidebar (right side) */}
      <div className="absolute right-0 top-0">
        <button
          onClick={() => setIsOpen(true)}
          className="flex items-center gap-2 px-4 py-2 rounded-lg shadow bg-[#0A2640] text-white"
        >
          <SlidersHorizontal className="w-4 h-4" />
          <span className="text-sm font-medium">Filter</span>
        </button>
      </div>

      {/* Overlay */}
      {isOpen && (
        <div
          onClick={() => setIsOpen(false)}
          className="fixed inset-0 bg-black/40 z-40"
        />
      )}

      {/* Sidebar */}
      <div
        className={`fixed top-0 right-0 h-full w-80 shadow-2xl z-50 transform transition-transform duration-300 overflow-y-auto
        ${isOpen ? "translate-x-0" : "translate-x-full"}
        ${darkMode ? "bg-[#262B32] text-white" : "bg-white text-black"}`}
      >
        {/* Header */}
        <div className="flex items-center justify-between p-4 border-b">
          <h2 className="text-lg font-semibold">Filters</h2>
          <button
            onClick={() => setIsOpen(false)}
            className="flex items-center gap-1 px-3 py-1 rounded-full border text-sm text-gray-700 hover:bg-gray-100 transition"
          >
            <X className="w-4 h-4" />
            <span>Close</span>
          </button>
        </div>

        {/* Content */}
        <div className="p-4 space-y-6">
          {/* Brand names */}
          <div>
            <label className="block text-sm font-medium mb-1">
              Brand names
            </label>
            <input
              type="text"
              value={filters.brandNames}
              onChange={(e) =>
                setFilters((prev) => ({ ...prev, brandNames: e.target.value }))
              }
              placeholder="Add multiple brands separated by commas"
              className={`w-full border rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 
              ${
                darkMode
                  ? "bg-[#1E1E1E] border-gray-600 text-white placeholder-gray-400 focus:ring-yellow-500"
                  : "border-gray-300 text-black placeholder-gray-500 focus:ring-blue-500"
              }`}
            />
          </div>

          {/* Price range */}
          <div>
            <label className="block text-sm font-medium mb-1">Price range</label>
            <div className="flex gap-2">
              <div className="w-1/2">
                <input
                  type="number"
                  value={filters.minPrice}
                  onChange={(e) =>
                    setFilters((prev) => ({
                      ...prev,
                      minPrice: e.target.value,
                    }))
                  }
                  placeholder="Min $50"
                  className={`w-full border rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 
                  ${
                    darkMode
                      ? "bg-[#1E1E1E] border-gray-600 text-white placeholder-gray-400 focus:ring-yellow-500"
                      : "border-gray-300 text-black placeholder-gray-500 focus:ring-blue-500"
                  }`}
                />
                {errors.minPrice && (
                  <p className="text-xs text-red-500 mt-1">{errors.minPrice}</p>
                )}
              </div>

              <div className="w-1/2">
                <input
                  type="number"
                  value={filters.maxPrice}
                  onChange={(e) =>
                    setFilters((prev) => ({
                      ...prev,
                      maxPrice: e.target.value,
                    }))
                  }
                  placeholder="Max $200"
                  className={`w-full border rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 
                  ${
                    darkMode
                      ? "bg-[#1E1E1E] border-gray-600 text-white placeholder-gray-400 focus:ring-yellow-500"
                      : "border-gray-300 text-black placeholder-gray-500 focus:ring-blue-500"
                  }`}
                />
                {errors.maxPrice && (
                  <p className="text-xs text-red-500 mt-1">{errors.maxPrice}</p>
                )}
              </div>
            </div>
          </div>

          {/* Minimum rating */}
          <div>
            <label className="block text-sm font-medium mb-1">
              Minimum rating
            </label>
            <div className="flex gap-2">
              {[1, 2, 3, 4, 5].map((rating) => (
                <button
                  key={rating}
                  onClick={() => handleRatingClick(rating)}
                  className={`w-8 h-8 flex items-center justify-center rounded-full border focus:ring-2 transition
                  ${
                    filters.minRating === rating
                      ? "bg-yellow-400 text-black border-yellow-500"
                      : darkMode
                      ? "border-gray-600 text-white hover:bg-[#333]"
                      : "border-gray-300 text-black hover:bg-blue-50 focus:ring-blue-500"
                  }`}
                >
                  {rating}
                </button>
              ))}
            </div>
          </div>

          {/* Submit */}
          
        </div>
      </div>
    </div>
  );
}
