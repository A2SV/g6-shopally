"use client";

import Sidebar from "@/app/components/Sidebar";


export default function HowItWorksPage() {
  return (
    <div className="flex min-h-screen">
      {/* Sidebar */}
      <Sidebar activePage="how-it-works" />

      {/* Main Content */}
      <main className="flex-1 p-6 lg:p-12 bg-[var(--color-bg-primary)]">
        <div className="max-w-4xl mx-auto text-center">
          <h2 className="text-2xl lg:text-3xl font-bold text-[var(--color-text-primary)] mb-4">
            How ShopAlly Works
          </h2>
          <p className="text-gray-500 mb-10">
            Get personalized product recommendations in three simple steps
          </p>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {/* Step 1 */}
            <div className="flex flex-col items-center text-center">
              <div className="w-12 h-12 flex items-center justify-center bg-green-500 text-white rounded-lg mb-3">
                💬
              </div>
              <h3 className="font-semibold text-lg mb-2">Ask a Question</h3>
              <p className="text-gray-500 text-sm">
                Type your question in English or Amharic about what you need to buy
              </p>
            </div>

            {/* Step 2 */}
            <div className="flex flex-col items-center text-center">
              <div className="w-12 h-12 flex items-center justify-center bg-green-500 text-white rounded-lg mb-3">
                💡
              </div>
              <h3 className="font-semibold text-lg mb-2">Get Smart Recommendations</h3>
              <p className="text-gray-500 text-sm">
                Our AI analyzes your needs and suggests the best products with Ethiopian pricing
              </p>
            </div>

            {/* Step 3 */}
            <div className="flex flex-col items-center text-center">
              <div className="w-12 h-12 flex items-center justify-center bg-green-500 text-white rounded-lg mb-3">
                📊
              </div>
              <h3 className="font-semibold text-lg mb-2">Compare & Buy</h3>
              <p className="text-gray-500 text-sm">
                Compare prices, reviews, and shipping options to make the best choice
              </p>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}
