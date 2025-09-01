"use client";

import { useState } from "react";
import { X } from "lucide-react";

export default function FilterSidebar() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div>
      {/* Button to open sidebar */}
      <button
        onClick={() => setIsOpen(true)}
        className="px-4 py-2 bg-blue-600 text-white rounded-lg shadow"
      >
        Open Filters
      </button>

      {/* Overlay */}
      {isOpen && (
        <div
          onClick={() => setIsOpen(false)}
          className="fixed inset-0 bg-black/40 z-40"
        />
      )}

      {/* Sidebar */}
      <div
        className={`fixed top-0 right-0 h-full w-80 bg-white shadow-2xl z-50 transform transition-transform duration-300 ${
          isOpen ? "translate-x-0" : "translate-x-full"
        }`}
      >
        {/* Header */}
        <div className="flex items-center justify-between p-4 border-b">
          <h2 className="text-lg font-semibold">Filters</h2>
          <button onClick={() => setIsOpen(false)}>
            <X className="w-5 h-5" />
          </button>
        </div>

        {/* Content */}
        <div className="p-4 space-y-6">
          {/* Brand names */}
          <div>
            <label className="block text-sm font-medium mb-1">Brand names</label>
            <input
              type="text"
              placeholder="Add multiple brands separated by commas"
              className="w-full border rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
            <p className="text-xs text-gray-500 mt-1">e.g., Nike, Adidas, Asics</p>
          </div>

          {/* Price range */}
          <div>
            <label className="block text-sm font-medium mb-1">Price range</label>
            <div className="flex gap-2">
              <input
                type="number"
                placeholder="Min $50"
                className="w-1/2 border rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
              <input
                type="number"
                placeholder="Max $200"
                className="w-1/2 border rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
            <p className="text-xs text-gray-500 mt-1">
              Enter minimum and maximum price in USD.
            </p>
          </div>

          {/* Minimum rating */}
          <div>
            <label className="block text-sm font-medium mb-1">Minimum rating</label>
            <div className="flex gap-2">
              {[1, 2, 3, 4, 5].map((rating) => (
                <button
                  key={rating}
                  className="w-8 h-8 flex items-center justify-center rounded-full border border-gray-300 hover:bg-blue-50 focus:ring-2 focus:ring-blue-500"
                >
                  {rating}
                </button>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
