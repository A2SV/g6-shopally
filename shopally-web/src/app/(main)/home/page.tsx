//src/app/(main)/home/page.tsx
"use client";

import CardComponent from "@/app/components/home-page-component/HomeCard";
import { useDarkMode } from "@/app/components/ProfileComponents/DarkModeContext";
import { useLanguage } from "@/hooks/useLanguage";
import {
  useCompareProductsMutation,
  useSearchProductsMutation,
} from "@/lib/redux/api/userApiSlice";
import { ComparePayload } from "@/types/Compare/Comparison";
import { ApiErrorResponse, Product } from "@/types/types";
import { SerializedError } from "@reduxjs/toolkit";
import { FetchBaseQueryError } from "@reduxjs/toolkit/query";
import { ArrowRight, Loader2, MessageCircleMore } from "lucide-react";
import { useEffect, useState } from "react";

interface ConversationMessage {
  id: string;
  type: "user" | "ai";
  content: string;
  products?: Product[];
  timestamp: number;
}

export default function Home() {
  const { t } = useLanguage();
  const { isDarkMode } = useDarkMode();

  const [conversation, setConversation] = useState<ConversationMessage[]>([]);
  const [isClient, setIsClient] = useState(false);
  const [input, setInput] = useState("");
  const [compareList, setCompareList] = useState<Product[]>([]);
  const [expandedMessages, setExpandedMessages] = useState<Set<string>>(
    new Set()
  );

  const [searchProducts, { isLoading }] = useSearchProductsMutation();
  const [compareProducts, { isLoading: isComparing }] =
    useCompareProductsMutation();

  const toggleExpanded = (messageId: string) => {
    setExpandedMessages((prev) => {
      const newSet = new Set(prev);
      if (newSet.has(messageId)) {
        newSet.delete(messageId);
      } else {
        newSet.add(messageId);
      }
      return newSet;
    });
  };

  const handleSend = async () => {
    if (input.trim() === "") return;

    const userMessage: ConversationMessage = {
      id: `user-${Date.now()}`,
      type: "user",
      content: input,
      timestamp: Date.now(),
    };

    // Add user message to conversation
    setConversation((prev) => [...prev, userMessage]);

    const userInput = input;
    setInput("");

    try {
      const res = await searchProducts({
        query: userInput,
        priceMaxETB: null,
        minRating: null,
      }).unwrap();

      const products = res?.data?.products || [];

      const aiMessage: ConversationMessage = {
        id: `ai-${Date.now()}`,
        type: "ai",
        content: `Found ${products.length} products for "${userInput}"`,
        products: products,
        timestamp: Date.now(),
      };

      // Add AI response to conversation
      setConversation((prev) => [...prev, aiMessage]);
    } catch (err) {
      console.error("❌ Search failed:", err);

      const errorMessage: ConversationMessage = {
        id: `ai-error-${Date.now()}`,
        type: "ai",
        content: `Sorry, I couldn't find products for "${userInput}". Please try again.`,
        products: [],
        timestamp: Date.now(),
      };

      setConversation((prev) => [...prev, errorMessage]);
    }
  };

  // Set client flag and load conversation from localStorage
  useEffect(() => {
    setIsClient(true);

    // Load conversation from localStorage
    const stored = localStorage.getItem("conversation");
    if (stored) {
      try {
        const parsedConversation = JSON.parse(stored);
        setConversation(parsedConversation);

        // Auto-scroll to bottom after loading conversation
        setTimeout(() => {
          window.scrollTo({
            top: document.body.scrollHeight,
            behavior: "smooth",
          });
        }, 100);
      } catch (err) {
        console.error("Failed to parse stored conversation:", err);
      }
    }
  }, []);

  useEffect(() => {
    if (isClient) {
      localStorage.setItem("conversation", JSON.stringify(conversation));

      // Auto-scroll to bottom when new messages are added
      if (conversation.length > 0) {
        setTimeout(() => {
          window.scrollTo({
            top: document.body.scrollHeight,
            behavior: "smooth",
          });
        }, 100);
      }
    }
  }, [conversation, isClient]);

  // Initialize compare list
  useEffect(() => {
    const list = localStorage.getItem("compareProduct");
    setCompareList(list ? JSON.parse(list) : []);
  }, []);

  // Listen for compare list changes
  useEffect(() => {
    const updateCompare = () => {
      const list = localStorage.getItem("compareProduct");
      setCompareList(list ? JSON.parse(list) : []);
    };

    window.addEventListener("storage", updateCompare);

    return () => window.removeEventListener("storage", updateCompare);
  }, []);

  const handleCompare = async () => {
    if (compareList.length < 2 || compareList.length > 4) {
      alert("Please select 2 to 4 products for comparison.");
      return;
    }

    try {
      console.log("Sending compare list:", compareList);

      const payload: ComparePayload = {
        products: compareList.map((p) => ({
          id: p.id,
          title: p.title,
          imageUrl: p.imageUrl,
          aiMatchPercentage: p.aiMatchPercentage,
          price: p.price,
          productRating: p.productRating,
          sellerScore: p.sellerScore,
          deliveryEstimate: p.deliveryEstimate,
          summaryBullets: p.summaryBullets,
          deeplinkUrl: p.deeplinkUrl,
        })),
      };

      const res = await compareProducts(payload).unwrap();
      console.log("Comparison result:", res);

      // ✅ Save the response to localStorage
      localStorage.setItem(
        "comparisonResults",
        JSON.stringify(res.data.products) // full ComparisonItem objects
      );

      // ✅ Clear compare list storage after success
      localStorage.removeItem("compareProduct");
      setCompareList([]);

      // ✅ Redirect to compare page
      window.location.href = "/comparison";
    } catch (err: unknown) {
      console.error("❌ Compare API failed:", err);

      if (typeof err === "object" && err !== null) {
        const e = err as FetchBaseQueryError | SerializedError;

        if ("data" in e) {
          // 👇 cast `data` to your API error type
          const data = e.data as ApiErrorResponse | undefined;
          console.error("Error data:", data);

          alert(`Compare failed: ${data?.error?.message ?? "Unknown error"}`);
          return;
        }

        if ("status" in e) {
          console.error("Error status:", (e as FetchBaseQueryError).status);
        }
      }

      alert("Compare failed due to an unexpected error. Check console.");
    }
  };
  return (
    <main
      className={`min-h-screen flex flex-col items-center pb-24 ${
        isDarkMode ? "bg-[#090C11]" : "bg-[#FFFFFF]"
      }`}
    >
      {/* Compare Products Button - moved above chat section */}
      {compareList.length >= 2 && (
        <div className="w-full max-w-3xl px-4 mt-6 sticky top-4 z-50">
          <button
            onClick={handleCompare}
            className="w-full bg-yellow-400 text-black py-3 rounded-xl font-semibold shadow-lg"
            disabled={isComparing}
          >
            {isComparing
              ? "Comparing..."
              : `Compare Products (${compareList.length})`}
          </button>
        </div>
      )}

      {/* Hero Section */}
      <section className="max-w-3xl text-center mt-12 mb-4 px-4">
        <h1
          className={`text-2xl sm:text-3xl lg:text-4xl mb-4 font-bold leading-tight text-center ${
            isClient && conversation.length === 0 ? "block" : "hidden"
          } ${isDarkMode ? "text-white" : "text-black"}`}
        >
          {t("Your Smart AI Assistant for AliExpress Shopping")}
        </h1>
        <p
          className={`mb-12 text-center leading-relaxed max-w-2xl mx-auto ${
            isClient && conversation.length === 0 ? "block" : "hidden"
          } ${isDarkMode ? "text-[#FFF]" : "text-gray-400"}`}
        >
          {t(
            "Discover the perfect products on AliExpress with AI-powered recommendations tailored to your needs."
          )}
        </p>

        {/* Desktop Search Box */}
        <div
          className={`flex items-center border border-gray-300 dark:border-gray-600 rounded-xl overflow-hidden shadow-sm max-w-2xl mx-auto w-full ${
            isDarkMode
              ? "bg-[#262B32] text-[#999999]"
              : "bg-white text-[#999999]"
          } ${
            isClient && conversation.length === 0 ? "hidden sm:flex" : "hidden"
          }`}
        >
          <input
            type="text"
            placeholder={t("Ask me anything about products you need...")}
            value={input}
            onChange={(e) => setInput(e.target.value)}
            onKeyDown={(e) => e.key === "Enter" && handleSend()}
            className={`flex-1 px-4 py-3 text-sm sm:text-base focus:outline-none min-w-0 placeholder:text-sm ${
              isDarkMode ? "bg-gray-800 text-white" : "text-black"
            }`}
          />
          <button
            onClick={handleSend}
            disabled={isLoading}
            className="rounded-xl p-3 mr-1 my-1 flex items-center justify-center transition-colors bg-yellow-400 text-black"
          >
            {isLoading ? (
              <Loader2 size={20} className="animate-spin" />
            ) : (
              <ArrowRight size={20} />
            )}
          </button>
        </div>
      </section>

      {/* Suggestions or Results */}
      {isClient && conversation.length === 0 ? (
        <section className="mt-10 px-4 max-w-3xl w-full">
          <h2
            className={`text-lg font-semibold mb-4 text-center ${
              isDarkMode ? "text-gray-600" : "text-gray-700"
            }`}
          >
            {t("Try asking:")}
          </h2>
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
            {[
              t("What do I need to bake bread?"),
              t("Best kitchen appliances for small apartments"),
              t("Essential tools for home gardening"),
              t("What should I buy for a home office setup?"),
            ].map((q, i) => (
              <button
                key={i}
                onClick={() => {
                  setInput(q);
                  // We'll trigger handleSend after setting input
                  setTimeout(() => {
                    const userMessage: ConversationMessage = {
                      id: `user-${Date.now()}`,
                      type: "user",
                      content: q,
                      timestamp: Date.now(),
                    };

                    setConversation((prev) => [...prev, userMessage]);

                    // Simulate the search
                    searchProducts({
                      query: q,
                      priceMaxETB: null,
                      minRating: null,
                    })
                      .unwrap()
                      .then((res) => {
                        const products = res?.data?.products || [];

                        const aiMessage: ConversationMessage = {
                          id: `ai-${Date.now()}`,
                          type: "ai",
                          content: `Found ${products.length} products for "${q}"`,
                          products: products,
                          timestamp: Date.now(),
                        };

                        setConversation((prev) => [...prev, aiMessage]);
                      })
                      .catch((err) => {
                        console.error("❌ Search failed:", err);

                        const errorMessage: ConversationMessage = {
                          id: `ai-error-${Date.now()}`,
                          type: "ai",
                          content: `Sorry, I couldn't find products for "${q}". Please try again.`,
                          products: [],
                          timestamp: Date.now(),
                        };

                        setConversation((prev) => [...prev, errorMessage]);
                      });
                  }, 0);
                }}
                className={`flex items-center gap-2 border rounded-lg px-3 py-2 text-sm transition hover:bg-gray-50 ${
                  isDarkMode
                    ? "border-gray-700 hover:bg-gray-700 text-white"
                    : "border-gray-300 hover:bg-gray-50"
                }`}
              >
                <MessageCircleMore className="w-4 h-4 rounded-full text-white bg-black" />
                <span>{q}</span>
              </button>
            ))}
          </div>
        </section>
      ) : (
        <div className="w-full flex flex-col gap-4 mt-6 pb-24 sm:pb-8">
          {/* Conversation Messages */}
          {conversation.map((message) => (
            <div key={message.id} className="flex flex-col gap-2">
              {/* User Message */}
              {message.type === "user" && (
                <div className="flex justify-end w-full">
                  <div className="max-w-xs sm:max-w-md break-words px-4 py-2 rounded-xl bg-yellow-400 text-black mr-2">
                    {message.content}
                  </div>
                </div>
              )}

              {/* AI Response */}
              {message.type === "ai" && (
                <div className="flex flex-col gap-3 w-full">
                  <div className="flex justify-start w-full">
                    <div
                      className={`max-w-xs sm:max-w-md break-words px-4 py-2 rounded-xl ml-2 ${
                        isDarkMode
                          ? "bg-gray-700 text-white"
                          : "bg-gray-100 text-gray-800"
                      }`}
                    >
                      {message.content}
                    </div>
                  </div>

                  {/* Product Results */}
                  {message.products && message.products.length > 0 && (
                    <div className="mt-2">
                      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 sm:gap-5 lg:gap-6">
                        {(expandedMessages.has(message.id)
                          ? message.products
                          : message.products.slice(0, 4)
                        ).map((product, i) => (
                          <CardComponent
                            key={`${message.id}-product-${i}`}
                            product={product}
                          />
                        ))}
                      </div>

                      {/* See More Button */}
                      {message.products.length > 4 && (
                        <div className="flex justify-center mt-4">
                          <button
                            onClick={() => toggleExpanded(message.id)}
                            className={`px-4 py-2 rounded-lg text-sm font-medium transition-colors ${
                              isDarkMode
                                ? "bg-gray-700 text-white hover:bg-gray-600"
                                : "bg-gray-100 text-gray-800 hover:bg-gray-200"
                            }`}
                          >
                            {expandedMessages.has(message.id)
                              ? `Show Less (${
                                  message.products.length - 4
                                } hidden)`
                              : `See More (${
                                  message.products.length - 4
                                } more products)`}
                          </button>
                        </div>
                      )}
                    </div>
                  )}
                </div>
              )}
            </div>
          ))}
        </div>
      )}

      {/* Mobile Search Box */}
      <section
        className={`fixed bottom-0 left-16 lg:left-64 right-0 z-50 ${
          isClient && conversation.length === 0 ? "block sm:hidden" : "block"
        }`}
      >
        {/* Full backdrop to prevent any content from showing through */}
        <div
          className={`absolute inset-0 backdrop-blur-sm ${
            isDarkMode ? "bg-[#090C11]/95" : "bg-[#FFFFFF]/95"
          }`}
        ></div>
        <div className="relative px-4 sm:px-8 py-2 sm:py-4 flex justify-center">
          <div
            className={`relative flex items-center border border-gray-300 dark:border-gray-600 rounded-xl overflow-hidden shadow-sm w-full max-w-4xl ${
              isDarkMode
                ? "bg-[#262B32] text-[#999999]"
                : "bg-white text-[#999999]"
            }`}
          >
            <input
              type="text"
              placeholder={t("Ask me anything about products you need...")}
              value={input}
              onChange={(e) => setInput(e.target.value)}
              onKeyDown={(e) => e.key === "Enter" && handleSend()}
              className={`flex-1 px-3 sm:px-4 py-2.5 sm:py-3 text-sm sm:text-base focus:outline-none min-w-0 ${
                isDarkMode ? "bg-gray-800 text-white" : "text-black"
              }`}
            />
            <button
              onClick={handleSend}
              disabled={isLoading}
              className="bg-yellow-400 rounded-xl p-2.5 sm:p-3 mr-1 my-1 flex items-center justify-center flex-shrink-0"
            >
              {isLoading ? (
                <Loader2 size={18} className="animate-spin text-black" />
              ) : (
                <ArrowRight size={18} className="sm:w-5 sm:h-5 text-black" />
              )}
            </button>
          </div>
        </div>
      </section>
    </main>
  );
}
