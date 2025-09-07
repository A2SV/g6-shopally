"use client";

import { ClientSafeProvider, getProviders, signIn } from "next-auth/react";
import Image from "next/image";
import { useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";

export default function CustomSignIn() {
  // const router = useRouter();
  const searchParams = useSearchParams();
  const callbackUrl = searchParams.get("callbackUrl") || "/home";
  const [isLoading, setIsLoading] = useState(false);
  const [providers, setProviders] = useState<Record<
    string,
    ClientSafeProvider
  > | null>(null);

  useEffect(() => {
    const fetchProviders = async () => {
      const res = await getProviders();
      setProviders(res);
    };
    fetchProviders();
  }, []);

  const handleGoogleSignIn = async () => {
    setIsLoading(true);
    try {
      await signIn("google", { callbackUrl });
    } catch (error) {
      console.error("Sign in error:", error);
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-white">
      <div className="max-w-md w-full mx-4">
        {/* ShopAlly AI Logo and Welcome Section */}
        <div className="text-center mb-8">
          {/* Logo */}
          <div className="flex justify-center mb-6">
            <div className="w-20 h-20 bg-yellow-400 rounded-full flex items-center justify-center shadow-lg">
              <Image
                src="/WebsiteLogo/Frame.png"
                alt="ShopAlly Logo"
                width={48}
                height={48}
                className="object-contain"
                onError={(e) => {
                  // Fallback to text if image fails to load
                  e.currentTarget.style.display = "none";
                  const fallback = document.createElement("span");
                  fallback.textContent = "S";
                  fallback.className = "text-black font-bold text-2xl";
                  e.currentTarget.parentNode?.appendChild(fallback);
                }}
              />
            </div>
          </div>

          {/* Welcome Text */}
          <h1 className="text-3xl font-bold text-gray-900 mb-2">
            Welcome to ShopAlly AI
          </h1>
          <p className="text-gray-600 text-lg mb-2">
            Your Smart AI Assistant for
          </p>
          <p className="text-gray-800 font-semibold text-lg">
            AliExpress Shopping
          </p>
          <p className="text-gray-500 text-sm mt-4 max-w-sm mx-auto">
            Discover the perfect products with AI-powered recommendations
            tailored to your needs.
          </p>
        </div>

        {/* Sign In Button */}
        <div className="space-y-4">
          {providers && (
            <button
              onClick={handleGoogleSignIn}
              disabled={isLoading}
              className="w-full flex items-center justify-center px-6 py-4 border border-gray-300 rounded-xl shadow-sm bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-yellow-400 focus:ring-offset-2 transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {isLoading ? (
                <div className="flex items-center">
                  <div className="animate-spin rounded-full h-5 w-5 border-b-2 border-gray-600 mr-3"></div>
                  <span className="text-gray-700 font-medium">
                    Signing in...
                  </span>
                </div>
              ) : (
                <div className="flex items-center">
                  {/* Google Logo */}
                  <svg
                    className="w-5 h-5 mr-3"
                    viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg"
                  >
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
                  <span className="text-gray-700 font-medium text-base">
                    Sign in with Google
                  </span>
                </div>
              )}
            </button>
          )}

          {/* Additional Info */}
          <div className="text-center">
            <p className="text-xs text-gray-500">
              By signing in, you agree to our Terms of Service and Privacy
              Policy
            </p>
          </div>
        </div>

        {/* Features Preview */}
        <div className="mt-12 space-y-3">
          <h3 className="text-center text-gray-700 font-semibold mb-4">
            What you will get:
          </h3>
          <div className="space-y-2">
            <div className="flex items-center text-sm text-gray-600">
              <div className="w-2 h-2 bg-yellow-400 rounded-full mr-3"></div>
              AI-powered product recommendations
            </div>
            <div className="flex items-center text-sm text-gray-600">
              <div className="w-2 h-2 bg-yellow-400 rounded-full mr-3"></div>
              Smart product comparison tools
            </div>
            <div className="flex items-center text-sm text-gray-600">
              <div className="w-2 h-2 bg-yellow-400 rounded-full mr-3"></div>
              Personalized shopping experience
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
