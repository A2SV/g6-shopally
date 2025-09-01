"use client";

import { useEffect, useState } from "react";
import { Sun, Moon, MessageCircleMore, ArrowRight } from "lucide-react";
import Image from "next/image";

export default function Home() {
  const [darkMode, setDarkMode] = useState(false);
  const [lang, setLang] = useState("EN");
  const [message, setMessage] = useState<string[]>([]);
  const [input, setinput] = useState("");

  const toggleDarkMode = () => {
    setDarkMode(!darkMode);
    if (typeof document !== "undefined") {
      document.documentElement.classList.toggle("dark");
    }
  };

  const handleSend = () => {
    if (input.trim() === "") return;
    setMessage([...message, input]); // add message
    setinput(""); // clear input
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
              <Image src="/images/frame.png" alt="" />
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
            {/* Language Dropdown */}
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

            {/* Dark/Light Mode Toggle */}
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
          </div>
        </div>
      </header>

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
