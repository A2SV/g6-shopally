"use client";

import CardComponent from "@/app/components/home-page-component/page";
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
import { useSavedItems } from "@/hooks/useSavedItems";
import {
  ArrowRight,
  Loader2,
  Star,
  X,
  Heart,
  MessageCircleMore,
} from "lucide-react";
import { useEffect, useRef, useState } from "react";

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
  const [selectedProduct, setSelectedProduct] = useState<Product | null>(null);
  const { saveItem } = useSavedItems();
  const [expandedMessages, setExpandedMessages] = useState<Set<string>>(
    new Set()
  );

  const [searchProducts, { isLoading }] = useSearchProductsMutation();
  const [compareProducts, { isLoading: isComparing }] =
    useCompareProductsMutation();

  const bottomRef = useRef<HTMLDivElement | null>(null);

  const generateId = () => `msg-${Math.random().toString(36).slice(2, 10)}`;

  const handleSend = async () => {
    if (input.trim() === "") return;

    const userMessage: ConversationMessage = {
      id: generateId(),
      type: "user",
      content: input,
      timestamp: Date.now(),
    };

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
        id: generateId(),
        type: "ai",
        content: `Found ${products.length} products for "${userInput}"`,
        products: products,
        timestamp: Date.now(),
      };

      setConversation((prev) => [...prev, aiMessage]);
    } catch (err) {
      console.error("❌ Search failed:", err);

      const errorMessage: ConversationMessage = {
        id: generateId(),
        type: "ai",
        content: `Sorry, I couldn't find products for "${userInput}". Please try again.`,
        products: [],
        timestamp: Date.now(),
      };

      setConversation((prev) => [...prev, errorMessage]);
    }
  };

  const handleSaveItem = () => {
    saveItem(selectedProduct);
    alert("Item saved!");
  };

  // Load stored conversation
  useEffect(() => {
    setIsClient(true);
    const stored = localStorage.getItem("conversation");
    if (stored) {
      try {
        const parsedConversation = JSON.parse(stored);
        setConversation(parsedConversation);
      } catch (err) {
        console.error("Failed to parse stored conversation:", err);
      }
    }
  }, []);

  // Persist conversation & auto-scroll
  useEffect(() => {
    if (isClient) {
      localStorage.setItem("conversation", JSON.stringify(conversation));
      if (conversation.length > 0) {
        setTimeout(() => {
          bottomRef.current?.scrollIntoView({ behavior: "smooth" });
        }, 100);
      }
    }
  }, [conversation, isClient]);

  // Prevent background scroll when sidebar open
  useEffect(() => {
    if (selectedProduct) {
      document.body.style.overflow = "hidden";
    } else {
      document.body.style.overflow = "";
    }
  }, [selectedProduct]);

  // Track compare list
  useEffect(() => {
    const updateCompare = () => {
      const list = localStorage.getItem("compareProduct");
      setCompareList(list ? JSON.parse(list) : []);
    };

    updateCompare();
    window.addEventListener("storage", updateCompare);
    return () => window.removeEventListener("storage", updateCompare);
  }, []);

  const handleCompare = async () => {
    if (compareList.length < 2 || compareList.length > 4) {
      alert("Please select 2 to 4 products for comparison.");
      return;
    }

    try {
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

      localStorage.setItem(
        "comparisonResults",
        JSON.stringify(res.data.products)
      );

      localStorage.removeItem("compareProduct");
      setCompareList([]);

      window.location.href = "/comparison";
    } catch (err: unknown) {
      console.error("❌ Compare API failed:", err);

      if (typeof err === "object" && err !== null) {
        const e = err as FetchBaseQueryError | SerializedError;
        if ("data" in e) {
          const data = e.data as ApiErrorResponse | undefined;
          alert(`Compare failed: ${data?.error?.message ?? "Unknown error"}`);
          return;
        }
      }

      alert("Compare failed due to an unexpected error. Check console.");
    }
  };

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

  return (
    <main
      className={`min-h-screen flex flex-col items-center relative pb-24 ${
        isDarkMode ? "bg-[#090C11]" : "bg-[#FFFFFF]"
      }`}
    >
      {/* Compare Button */}
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

      {/* Hero + Search */}
      <section className="max-w-3xl text-center mt-12 mb-4 px-4">
        <h1
          className={`text-2xl sm:text-3xl lg:text-4xl mb-4 font-bold leading-tight ${
            isClient && conversation.length === 0 ? "block" : "hidden"
          } ${isDarkMode ? "text-white" : "text-black"}`}
        >
          {t("Your Smart AI Assistant for AliExpress Shopping")}
        </h1>
        <p
          className={`mb-12 leading-relaxed max-w-2xl mx-auto ${
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
            isDarkMode ? "bg-[#262B32] text-[#999999]" : "bg-white text-[#999999]"
          } ${isClient && conversation.length === 0 ? "hidden sm:flex" : "hidden"}`}
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

      {/* Suggestions or Conversation */}
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
                  setTimeout(() => handleSend(), 0);
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
        <div className="w-full max-w-3xl px-4 flex flex-col gap-4 mt-6 pb-24">
          {conversation.map((message) => (
            <div key={message.id} className="flex flex-col gap-2">
              {/* User */}
              {message.type === "user" && (
                <div className="flex justify-end">
                  <div className="max-w-3/4 px-4 py-2 rounded-xl bg-yellow-400 text-black">
                    {message.content}
                  </div>
                </div>
              )}

              {/* AI */}
              {message.type === "ai" && (
                <div className="flex flex-col gap-3">
                  <div className="flex justify-start">
                    <div
                      className={`max-w-3/4 px-4 py-2 rounded-xl ${
                        isDarkMode
                          ? "bg-gray-700 text-white"
                          : "bg-gray-100 text-gray-800"
                      }`}
                    >
                      {message.content}
                    </div>
                  </div>
                  {message.products && message.products.length > 0 && (
                    <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 mt-2">
                      {(expandedMessages.has(message.id)
                        ? message.products
                        : message.products.slice(0, 4)
                      ).map((product, i) => (
                        <CardComponent
                          key={`${message.id}-product-${i}`}
                          mode={isDarkMode ? "dark" : "light"}
                          product={product}
                          onClick={() => setSelectedProduct(product)}
                        />
                      ))}
                    </div>
                  )}
                  {message.products && message.products.length > 4 && (
                    <div className="flex justify-center mt-2">
                      <button
                        onClick={() => toggleExpanded(message.id)}
                        className={`px-4 py-2 rounded-lg text-sm font-medium transition-colors ${
                          isDarkMode
                            ? "bg-gray-700 text-white hover:bg-gray-600"
                            : "bg-gray-100 text-gray-800 hover:bg-gray-200"
                        }`}
                      >
                        {expandedMessages.has(message.id)
                          ? `Show Less (${message.products.length - 4} hidden)`
                          : `See More (${message.products.length - 4} more products)`}
                      </button>
                    </div>
                  )}
                </div>
              )}
            </div>
          ))}
          <div ref={bottomRef} />
        </div>
      )}

      {/* 🔍 Mobile Search Box */}
      <section
        className={`fixed bottom-0 left-16 lg:left-64 right-0 z-50 ${
          isClient && conversation.length === 0 ? "block sm:hidden" : "block"
        }`}
      >
        <div
          className={`absolute inset-0 backdrop-blur-sm ${
            isDarkMode ? "bg-[#090C11]/95" : "bg-[#FFFFFF]/95"
          }`}
        />
        <div className="relative px-4 sm:px-8 py-2 sm:py-4 flex justify-center">
          <div
            className={`relative flex items-center border border-gray-300 dark:border-gray-600 rounded-xl overflow-hidden shadow-sm w-full max-w-4xl ${
              isDarkMode ? "bg-[#262B32] text-[#999999]" : "bg-white text-[#999999]"
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

      {/* 🔥 Sidebar Product Description */}
      {selectedProduct && (
        <>
          <div
            onClick={() => setSelectedProduct(null)}
            className="fixed inset-0 bg-black/40 backdrop-blur-sm z-40"
          />
          <aside
            className={`fixed top-0 right-0 h-full w-full sm:w-96 shadow-2xl z-50 
              transition-transform duration-300 ease-in-out transform translate-x-0
              ${isDarkMode ? "bg-gray-900 text-white" : "bg-white text-gray-900"}
              rounded-l-2xl flex flex-col`}
          >
            <div
              className={`flex items-start justify-between gap-3 p-4 border-b ${
                isDarkMode ? "border-gray-700" : "border-gray-200"
              }`}
            >
              <div>
                <h2 className="font-semibold text-lg line-clamp-2">
                  {selectedProduct.title}
                </h2>
              </div>
              <button
                onClick={() => setSelectedProduct(null)}
                className="hover:bg-gray-200 dark:hover:bg-gray-700 rounded-full p-1"
              >
                <X className="w-5 h-5" />
              </button>
            </div>

            <div className="flex-1 p-5 space-y-6 overflow-y-auto">
              <div className="w-full aspect-square bg-gray-100 dark:bg-gray-800 rounded-xl overflow-hidden flex items-center justify-center shadow-md">
                <img
                  src={selectedProduct.imageUrl}
                  alt={selectedProduct.title}
                  className="object-contain max-h-full max-w-full hover:scale-105 transition"
                />
              </div>

              <div className="flex justify-between items-center">
                <p className="text-2xl font-bold text-yellow-400">
                  ${selectedProduct.price?.usd ?? selectedProduct.price}
                </p>
                <div className="flex items-center gap-1 text-sm text-gray-500">
                  <Star className="w-4 h-4 text-yellow-400" />
                  {selectedProduct.productRating} (230 reviews)
                </div>
              </div>

              {selectedProduct.summaryBullets && (
                <div>
                  <h3 className="text-sm font-semibold mb-2">Key Features</h3>
                  <ul className="space-y-2 text-sm leading-relaxed">
                    {selectedProduct.summaryBullets.map((b, i) => (
                      <li key={i} className="flex items-start gap-2">
                        <span className="text-yellow-400">✔</span>
                        <span>{b}</span>
                      </li>
                    ))}
                  </ul>
                </div>
              )}
            </div>

            <div
              className={`p-4 border-t ${
                isDarkMode
                  ? "border-gray-700 bg-gray-900"
                  : "border-gray-100 bg-white"
              }`}
            >
              <div className="flex gap-3">
                <a
                  href={selectedProduct.deeplinkUrl}
                  target="_blank"
                  rel="noreferrer"
                  className="flex-1 bg-yellow-400 text-black text-center py-2 rounded-md font-semibold shadow hover:opacity-90 transition"
                >
                  Buy Now
                </a>
                <button 
                  onClick={(e) => {
                    e.stopPropagation();
                    handleSaveItem();
                  }}
                className="flex items-center justify-center px-4 py-2 rounded-md border border-gray-300 dark:border-gray-600 hover:bg-gray-100 dark:hover:bg-gray-700 transition">
                  <Heart className="w-5 h-5 text-red-500" />
                </button>
              </div>
            </div>
          </aside>
        </>
      )}
    </main>
  );
}
