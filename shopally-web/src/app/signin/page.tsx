"use client";

import { signIn } from "next-auth/react";

export const dynamic = "force-dynamic";

export default function SignInPage() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-white">
      <div className="max-w-md w-full space-y-8 px-6">
        {/* Logo */}
        <div className="flex justify-center">
          <div className="w-16 h-16 bg-yellow-400 rounded-full flex items-center justify-center">
            <img
              src="/WebsiteLogo/Frame.png"
              alt="ShopAlly AI Logo"
              className="w-12 h-12"
            />
          </div>
        </div>

        {/* Title */}
        <div className="text-center">
          <h1 className="text-3xl font-bold text-gray-900 mb-2">
            Welcome to ShopAlly AI
          </h1>
          <p className="text-gray-600 text-lg">Your Smart AI Assistant for</p>
          <p className="text-gray-900 font-semibold text-lg">
            AliExpress Shopping
          </p>
          <p className="text-gray-500 text-sm mt-3">
            Discover the perfect products with AI-powered recommendations
            tailored to your needs.
          </p>
        </div>

        {/* Sign in Button */}
        <div className="space-y-4">
          <button
            onClick={() => signIn("google", { callbackUrl: "/home" })}
            className="w-full flex items-center justify-center px-4 py-3 border border-gray-300 rounded-lg bg-white hover:bg-gray-50 transition-colors duration-200"
          >
            <svg className="w-5 h-5 mr-3" viewBox="0 0 24 24">
              <path
                fill="#4285F4"
                d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
              />
              <path
                fill="#34A853"
                d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
              />
              <path
                fill="#FBBC05"
                d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
              />
              <path
                fill="#EA4335"
                d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
              />
            </svg>
            <span className="text-gray-700 font-medium">
              Sign in with Google
            </span>
          </button>
        </div>

        {/* Terms */}
        <p className="text-center text-xs text-gray-500">
          By signing in, you agree to our{" "}
          <a href="#" className="text-blue-600 hover:underline">
            Terms of Service
          </a>{" "}
          and{" "}
          <a href="#" className="text-blue-600 hover:underline">
            Privacy Policy
          </a>
        </p>

        {/* Benefits */}
        <div className="mt-8">
          <h3 className="text-lg font-semibold text-gray-900 mb-4">
            What you&apos;ll get:
          </h3>
          <ul className="space-y-2">
            <li className="flex items-center text-sm text-gray-600">
              <span className="w-2 h-2 bg-yellow-400 rounded-full mr-3"></span>
              AI-powered product recommendations
            </li>
            <li className="flex items-center text-sm text-gray-600">
              <span className="w-2 h-2 bg-yellow-400 rounded-full mr-3"></span>
              Smart product comparison tools
            </li>
            <li className="flex items-center text-sm text-gray-600">
              <span className="w-2 h-2 bg-yellow-400 rounded-full mr-3"></span>
              Personalized shopping experience
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
}
