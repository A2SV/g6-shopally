"use client";

import { Search } from "lucide-react";

export default function Hero() {
  return (
    <section className="bg-white text-gray-900 py-16 px-6 md:px-12">
      <div className="max-w-6xl mx-auto flex flex-col md:flex-row items-center justify-between gap-10">
        
        {/* Left Content */}
        <div className="flex-1">
          <h1 className="text-3xl md:text-5xl font-bold leading-tight">
            Your Smart AI <br /> Shopping Assistant <br /> for Alibaba
          </h1>
          <p className="mt-4 text-gray-700 max-w-md">
            Ask in English or Amharic and get instant product recommendations
            with Ethiopian pricing.
          </p>

          {/* Search Bar */}
          <div className="mt-6 flex items-center bg-gray-100 rounded-md overflow-hidden w-full max-w-md">
            <input
              type="text"
              placeholder="What do you need to buy today?"
              aria-label="Search products"
              className="flex-1 px-4 py-2 text-gray-900 outline-none"
            />
            <button
              aria-label="Search"
              className="bg-yellow-500 px-4 py-2 flex items-center justify-center hover:bg-yellow-600 transition-colors"
            >
              <Search className="w-5 h-5 text-white" />
            </button>
          </div>

          {/* Buttons */}
          <div className="mt-6 flex gap-4">
            <button className="bg-yellow-500 text-white px-6 py-2 rounded-md font-medium hover:bg-yellow-600 transition-colors">
              Ask ShopAlly
            </button>
            <button className="border border-gray-900 px-6 py-2 rounded-md font-medium hover:bg-gray-900 hover:text-white transition-colors">
              See How It Works
            </button>
          </div>
        </div>

      </div>
    </section>
  );
}
