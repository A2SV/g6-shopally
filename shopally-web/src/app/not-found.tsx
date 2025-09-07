"use client";

import Link from "next/link";
import { useRouter } from "next/navigation";
import Image from "next/image";

export default function NotFound() {
  const router = useRouter();

  return (
    <div className="min-h-screen flex items-center justify-center bg-white">
      <div className="max-w-md w-full mx-4 text-center">
        {/* ShopAlly AI Logo */}
        <div className="flex justify-center mb-8">
          <div className="w-24 h-24 bg-yellow-400 rounded-full flex items-center justify-center shadow-lg">
            <Image
              src="/WebsiteLogo/Frame.png"
              alt="ShopAlly Logo"
              width={56}
              height={56}
              className="object-contain"
              onError={(e) => {
                // Fallback to text if image fails to load
                e.currentTarget.style.display = "none";
                const fallback = document.createElement("span");
                fallback.textContent = "S";
                fallback.className = "text-black font-bold text-3xl";
                e.currentTarget.parentNode?.appendChild(fallback);
              }}
            />
          </div>
        </div>

        {/* 404 Error Message */}
        <div className="mb-8">
          <h1 className="text-6xl font-bold text-gray-900 mb-4">404</h1>
          <h2 className="text-2xl font-semibold text-gray-800 mb-4">
            Page Not Found
          </h2>
          <p className="text-gray-600 text-lg mb-2">
            Oops! The page you&apos;re looking for
          </p>
          <p className="text-gray-600 text-lg mb-6">
            doesn&apos;t exist in our ShopAlly AI universe.
          </p>
        </div>

        {/* Action Buttons */}
        <div className="space-y-4">
          <button
            onClick={() => router.back()}
            className="w-full bg-gray-100 hover:bg-gray-200 text-gray-800 font-medium py-3 px-6 rounded-xl transition-colors duration-200"
          >
            ← Go Back
          </button>

          <Link
            href="/home"
            className="block w-full bg-yellow-400 hover:bg-yellow-500 text-black font-medium py-3 px-6 rounded-xl transition-colors duration-200"
          >
            Go to Home
          </Link>
        </div>
      </div>
    </div>
  );
}
