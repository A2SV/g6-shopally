"use client";

import { useState, useEffect } from "react";
import { MessageCircleMore, ArrowRight } from "lucide-react";
import { useDarkMode } from "@/app/components/ProfileComponents/DarkModeContext";
import { useLanguage } from "@/hooks/useLanguage";

export default function Home() {
  const { t } = useLanguage();
  const { isDarkMode } = useDarkMode();
  const [messages, setMessages] = useState<string[]>([]);
  const [input, setInput] = useState("");

  const handleSend = () => {
    if (input.trim() === "") return;
    setMessages([...messages, input]);
    setInput("");
  };

  // Load messages from localStorage
  useEffect(() => {
    const stored = localStorage.getItem("messages");
    setMessages(stored ? JSON.parse(stored) : []);
  }, []);

  // Save messages to localStorage
  useEffect(() => {
    localStorage.setItem("messages", JSON.stringify(messages));
  }, [messages]);

  return (
    <main
      className={`min-h-screen flex flex-col items-center pb-24 ${
        isDarkMode ? "bg-[#090C11]" : "bg-[#FFFFFF]"
      }`}
    >
      {/* Hero Section */}
      <section
        className={`max-w-3xl text-center mt-12 mb-4 px-4 ${
          isDarkMode ? "bg-[#090C11]" : "bg-[#FFFFFF]"
        }`}
      >
        <h1
          className={`text-2xl sm:text-3xl lg:text-4xl mb-4 font-bold leading-tight text-center ${
            messages.length === 0 ? "block" : "hidden"
          } ${isDarkMode ? "text-white" : "text-black"}`}
        >
          {t("Your Smart AI Assistant for Alibaba Shopping")}
        </h1>
        <p
          className={`mb-12 text-center leading-relaxed max-w-2xl mx-auto ${
            messages.length === 0 ? "block" : "hidden"
          } ${isDarkMode ? "text-[#FFF]" : "text-gray-400"}`}
        >
          {t(
            "Discover the perfect products on Alibaba with AI-powered recommendations tailored to your needs."
          )}
        </p>

        {/* Search Box */}
        <div
          className={`flex items-center border border-gray-300 dark:border-gray-600 rounded-xl overflow-hidden shadow-sm max-w-2xl mx-auto w-full ${
            isDarkMode
              ? "bg-[#262B32] text-[#999999]"
              : "bg-white text-[#999999]"
          } ${messages.length === 0 ? "hidden sm:flex" : "hidden"}`}
        >
          <input
            type="text"
            placeholder={t("Ask me anything about products you need...")}
            value={input}
            onChange={(e) => setInput(e.target.value)}
            className={`flex-1 px-4 py-3 text-sm sm:text-base focus:outline-none min-w-0 placeholder:text-sm ${
              isDarkMode ? "bg-gray-800 text-white" : "text-black"
            }`}
            style={{
              fontSize: "14px",
              lineHeight: "20px",
              height: "48px",
            }}
          />
          <button
            onClick={handleSend}
            className="rounded-xl p-3 mr-1 my-1 flex items-center justify-center transition-colors"
            style={{
              backgroundColor: "var(--color-accent-primary)",
              color: "var(--color-text-button)",
            }}
          >
            <ArrowRight size={20} />
          </button>
        </div>
      </section>

      {/* Message List or Suggestions */}
      {messages.length === 0 ? (
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
                onClick={() => setMessages([...messages, q])}
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
        <div className="w-full max-w-3xl px-4 flex flex-col gap-2 mt-6">
          {messages.map((msg, idx) => (
            <div key={idx} className="flex justify-end">
              <div className="max-w-1/2 break-words px-4 py-2 rounded-xl bg-yellow-400 text-black">
                {msg}
              </div>
            </div>
          ))}
        </div>
      )}

      {/* Mobile Search - Responsive and properly sized */}
      <section
        className={`fixed bottom-2 sm:bottom-4 left-0 right-0 px-3 sm:px-4 ${
          messages.length === 0 ? "block sm:hidden" : "block"
        } ${isDarkMode ? "bg-[#090C11]" : "bg-[#FFFFFF]"}`}
        style={{
          marginLeft: "4rem", // Account for mobile sidebar width
          marginRight: "1rem", // Right margin for spacing
        }}
      >
        <div
          className={`flex items-center border border-gray-300 dark:border-gray-600 rounded-xl overflow-hidden shadow-sm w-full ${
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
            className={`flex-1 px-3 sm:px-4 py-2.5 sm:py-3 text-sm sm:text-base focus:outline-none min-w-0 placeholder:text-xs sm:placeholder:text-sm ${
              isDarkMode ? "bg-gray-800 text-white" : "text-black"
            }`}
            style={{
              fontSize: "14px",
              lineHeight: "20px",
              height: "44px",
            }}
          />
          <button
            onClick={handleSend}
            className="bg-yellow-400 rounded-xl p-2.5 sm:p-3 mr-1 my-1 flex items-center justify-center flex-shrink-0"
          >
            <ArrowRight size={18} className="sm:w-5 sm:h-5 text-black" />
          </button>
        </div>
      </section>
    </main>
  );
}
