"use client";

import { useEffect, useState } from "react";
import {
  Sun,
  Moon,
  MessageCircleMore,
  ArrowRight,
  SlidersHorizontal,
  X,
} from "lucide-react";

export default function Home() {
  const [darkMode, setDarkMode] = useState(false);
  const [lang, setLang] = useState("EN");
  const [message, setMessage] = useState<string[]>([]);
  const [input, setinput] = useState("");
  // Removed unused bottomSearch

  // Filter sidebar state
  const [isFilterOpen, setIsFilterOpen] = useState(false);

  const toggleDarkMode = () => {
    setDarkMode(!darkMode);
    if (typeof document !== "undefined") {
      document.documentElement.classList.toggle("dark");
    }
  };

  const handleSend = () => {
    if (input.trim() === "") return;
    setMessage([...message, input]);
    setinput("");
  };

  useEffect(() => {
    const mes = localStorage.getItem("message");
    setMessage(mes ? JSON.parse(mes) : []);
  }, []);

  useEffect(() => {
    localStorage.setItem("message", JSON.stringify(message));
  }, [message]);

  return (
    <main
      className={`min-h-screen ${
        darkMode ? "bg-[#090C11]" : "bg-[#FFFFFF]"
      } flex flex-col items-center pb-24`}
    >
      <header
        className={`${
          darkMode ? "bg-[#262B32]" : "bg-[#FFFFFF]"
        } w-full border-b ${
          darkMode ? "border-gray-600" : "border-gray-200"
        } dark:border-gray-700 `}
      >
        <div className="max-w-7xl mx-auto flex items-center justify-between px-4 sm:px-6 lg:px-8 h-16">
          <div className="flex items-center gap-2">
            <div className="w-6 h-6 rounded-full bg-yellow-400">
              {/* Replace img with Next.js Image for lint compliance */}
              {/* <img src="/images/frame.png" alt="" /> */}
            </div>
            <span
              className={`font-bold text-lg ${
                darkMode ? "text-white" : "text-black"
              } `}
            >
              ShopAlly
            </span>
          </div>

          {/* Language & Theme Controls */}
          <div className="flex items-center gap-4">
            <select
              value={lang}
              onChange={(e) => setLang(e.target.value)}
              className={`border border-gray-300 dark:border-gray-600 rounded-md px-2 py-1 text-sm ${
                darkMode ? "bg-[#757B81]" : "bg-[#F3F4F6]"
              } dark:text-white focus:outline-none`}
            >
              <option value="EN">EN</option>
              <option value="AM">AM</option>
            </select>
            <button
              onClick={toggleDarkMode}
              className={`p-2 rounded-md border border-gray-300 ${
                darkMode ? "bg-[#757B81]" : "bg-[#F3F4F6]"
              } hover:bg-gray-700 `}
            >
              {darkMode ? (
                <Sun size={18} className="text-yellow-400" />
              ) : (
                <Moon size={18} className="text-black dark:text-black" />
              )}
            </button>
            {/* Filter Button to right side */}
            <button
              onClick={() => setIsFilterOpen(true)}
              className="flex items-center gap-2 px-4 py-2 rounded-lg shadow bg-[#0A2640] text-white"
            >
              <SlidersHorizontal className="w-4 h-4" />
              <span className="text-sm font-medium">Filter</span>
            </button>
          </div>
        </div>
      </header>

      {/* Filter Sidebar */}
      {isFilterOpen && (
        <>
          <div
            onClick={() => setIsFilterOpen(false)}
            className="fixed inset-0 bg-black/40 z-40"
          />
          <div
            className={`fixed top-0 right-0 h-full w-80 shadow-2xl z-50 transition-transform duration-300 overflow-y-auto
            ${darkMode ? "bg-[#262B32] text-white" : "bg-white text-black"}`}
          >
            <div
              className={`flex items-center justify-between p-4 border-b ${
                darkMode ? "border-gray-700" : "border-gray-200"
              }`}
            >
              <h2 className="text-lg font-semibold">Filters</h2>
              <button
                onClick={() => setIsFilterOpen(false)}
                className="flex items-center gap-1 px-3 py-1 rounded-full border text-sm hover:bg-gray-100 transition"
              >
                <X
                  className={`w-4 h-4 ${darkMode ? "text-white" : "text-black"}`}
                />
                <span className={`${darkMode ? "text-white" : "text-black"}`}>
                  Close
                </span>
              </button>
            </div>
            <div className="p-4 space-y-6">
              {/* Brand names */}
              <div>
                <label className="block text-sm font-medium mb-1">
                  Brand names
                </label>
                <input
                  type="text"
                  placeholder="Add multiple brands separated by commas"
                  className={`w-full border rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 
                  ${
                    darkMode
                      ? "bg-[#1E1E1E] border-gray-600 text-white placeholder-gray-400 focus:ring-yellow-500"
                      : "border-gray-300 text-black placeholder-gray-500 focus:ring-blue-500"
                  }`}
                />
                <p
                  className={`text-xs mt-1 ${darkMode ? "text-gray-400" : "text-gray-500"}`}
                >
                  e.g., Nike, Adidas, Asics
                </p>
              </div>
              {/* Price range */}
              <div>
                <label className="block text-sm font-medium mb-1">
                  Price range
                </label>
                <div className="flex gap-2">
                  <input
                    type="number"
                    placeholder="Min $50"
                    className={`w-1/2 border rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 
                    ${
                      darkMode
                        ? "bg-[#1E1E1E] border-gray-600 text-white placeholder-gray-400 focus:ring-yellow-500"
                        : "border-gray-300 text-black placeholder-gray-500 focus:ring-blue-500"
                    }`}
                  />
                  <input
                    type="number"
                    placeholder="Max $200"
                    className={`w-1/2 border rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 
                    ${
                      darkMode
                        ? "bg-[#1E1E1E] border-gray-600 text-white placeholder-gray-400 focus:ring-yellow-500"
                        : "border-gray-300 text-black placeholder-gray-500 focus:ring-blue-500"
                    }`}
                  />
                </div>
                <p
                  className={`text-xs mt-1 ${darkMode ? "text-gray-400" : "text-gray-500"}`}
                >
                  Enter minimum and maximum price in USD.
                </p>
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
                      className={`w-8 h-8 flex items-center justify-center rounded-full border hover:bg-blue-50 focus:ring-2 
                      ${
                        darkMode
                          ? "border-gray-600 text-white hover:bg-[#333]"
                          : "border-gray-300 text-black hover:bg-blue-50 focus:ring-blue-500"
                      }`}
                    >
                      {rating}
                    </button>
                  ))}
                </div>
              </div>
            </div>
          </div>
        </>
      )}

      {/* Hero Section */}
      <section
        className={`max-w-3xl text-center mt-12 mb-4 px-4 ${
          darkMode ? "bg-[#090C11]" : "bg-[#FFFFFF]"
        }`}
      >
        <h1
          className={`text-2xl sm:text-3xl lg:text-4xl lg:w-[100%] mb-4 ${
            message.length === 0 ? "block" : "hidden"
          }  font-bold  ${darkMode ? "text-white" : "text-black"}`}
        >
          Your Smart AI Assistant for Alibaba Shopping
        </h1>
        <p
          className={` ${darkMode ? "text-[#FFF] " : "text-gray-400 "} ${
            message.length === 0 ? "block" : "hidden"
          } mb-12`}
        >
          Discover the perfect products on Alibaba with AI-powered
          recommendations tailored to your needs.
        </p>
        {/* Search Box */}
        <div
          className={`flex items-center ${
            message.length === 0 ? "hidden sm:flex" : "hidden"
          }  border-gray-300 dark:border-gray-600 rounded-xl overflow-hidden shadow-sm ${
            darkMode
              ? "bg-[#262B32] text-[#999999] "
              : "bg-white text-[#999999] "
          } `}
        >
          <input
            type="text"
            placeholder="Ask me anything about products you need..."
            onChange={(e) => setinput(e.target.value)}
            value={input}
            className={` ${
              darkMode ? "bg-gray-800 text-white" : "text-black"
            } flex-1 px-4 py-3 text-sm sm:text-base focus:outline-none `}
          />
          <button
            onClick={handleSend}
            className="bg-yellow-400 rounded-xl p-3 mr-1 my-1  flex items-center justify-center"
          >
            <ArrowRight size={20} className="text-black" />
          </button>
        </div>
      </section>

      {message.length === 0 ? (
        <section className="mt-10 px-4 max-w-3xl w-full">
          <h2
            className={` ${
              darkMode ? "text-gray-600" : "text-gray-700"
            } text-lg font-semibold mb-4 text-center`}
          >
            Try asking:
          </h2>
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
            {[
              "What do I need to bake bread?",
              "Best kitchen appliances for small apartments",
              "Essential tools for home gardening",
              "What should I buy for a home office setup?",
            ].map((q, i) => (
              <button
                key={i}
                onClick={() => setMessage([...message, q])}
                className={`flex items-center gap-2 border ${
                  darkMode
                    ? "border-gray-700 hover:bg-gray-700 text-[#FFFFFF]"
                    : "hover:bg-gray-50 border-gray-300"
                } rounded-lg px-3 py-2 text-sm text-left hover:bg-gray-50 transition`}
              >
                <MessageCircleMore className="w-4 h-4 rounded-full text-white bg-black" />
                <span>{q}</span>
              </button>
            ))}
          </div>
        </section>
      ) : (
        <div className="w-full max-w-3xl px-4 flex flex-col gap-2 mt-6">
          {message.map((msg, idx) => (
            <div key={idx} className="flex justify-end">
              <div
                className={`max-w-1/2 break-words px-4 py-2 rounded-xl bg-yellow-400 text-black`}
              >
                {msg}
              </div>
            </div>
          ))}
        </div>
      )}

      {/* {search bar if there is message} */}
      <section
        className={`max-w-3xl w-full fixed bottom-1 ${
          message.length === 0 ? "block sm:hidden" : "block"
        }   text-center mt-12  px-4 ${
          darkMode ? "bg-[#090C11]" : "bg-[#FFFFFF]"
        }`}
      >
        <div
          className={`flex items-center border  border-gray-300 dark:border-gray-600 rounded-xl overflow-hidden shadow-sm ${
            darkMode
              ? "bg-[#262B32] text-[#999999] "
              : "bg-white text-[#999999] "
          } `}
        >
          <input
            type="text"
            placeholder="Ask me anything about products you need..."
            onChange={(e) => setinput(e.target.value)}
            value={input}
            className={` ${
              darkMode ? "bg-gray-800 text-white" : "text-black"
            } flex-1 px-4 py-3 text-sm sm:text-base focus:outline-none `}
          />
          <button
            onClick={handleSend}
            className="bg-yellow-400 rounded-xl p-3 mr-1 my-1  flex items-center justify-center"
          >
            <ArrowRight size={20} className="text-black" />
          </button>
        </div>
      </section>
    </main>
  );
}
