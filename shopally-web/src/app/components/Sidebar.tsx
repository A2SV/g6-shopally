"use client";

import Image from "next/image";
import { useRouter } from "next/navigation";
import { useDarkMode } from "./ProfileComponents/DarkModeContext";
import { useLanguage } from "./LanguageContext";

interface SidebarProps {
  activePage?: string;
}

export default function Sidebar({ activePage = "profile" }: SidebarProps) {
  const { isDarkMode, toggleDarkMode } = useDarkMode();
  const { currentLanguage, switchLanguage, t } = useLanguage();
  const router = useRouter();

  const navigationItems = [
    {
      id: "home",
      label: t("Home"),
      path: "/",
      icon: (
        <svg
          className="w-4 h-4 lg:w-5 lg:h-5"
          fill="currentColor"
          viewBox="0 0 20 20"
        >
          <path d="M10.707 2.293a1 1 0 00-1.414 0l-7 7a1 1 0 001.414 1.414L4 10.414V17a1 1 0 001 1h2a1 1 0 001-1v-2a1 1 0 011-1h2a1 1 0 011 1v2a1 1 0 001 1h2a1 1 0 001-1v-6.586l.293.293a1 1 0 001.414-1.414l-7-7z" />
        </svg>
      ),
    },
    {
      id: "how-it-works",
      label: t("How It Works"),
      path: "/how-it-works",
      icon: (
        <svg
          className="w-4 h-4 lg:w-5 lg:h-5"
          fill="currentColor"
          viewBox="0 0 20 20"
        >
          <path d="M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464a1 1 0 10-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.476.859h4.002z" />
        </svg>
      ),
    },
    {
      id: "compare",
      label: t("Compare"),
      path: "/compare",
      icon: (
        <svg
          className="w-4 h-4 lg:w-5 lg:h-5"
          fill="currentColor"
          viewBox="0 0 20 20"
        >
          <path d="M2 11a1 1 0 011-1h2a1 1 0 011 1v5a1 1 0 01-1 1H3a1 1 0 01-1-1v-5zM8 7a1 1 0 011-1h2a1 1 0 011 1v9a1 1 0 01-1 1H9a1 1 0 01-1-1V7zM14 4a1 1 0 011-1h2a1 1 0 011 1v12a1 1 0 01-1 1h-2a1 1 0 01-1-1V4z" />
        </svg>
      ),
    },
    {
      id: "saved-items",
      label: t("Saved Items"),
      path: "/saved-items",
      icon: (
        <svg
          className="w-4 h-4 lg:w-5 lg:h-5"
          fill="currentColor"
          viewBox="0 0 20 20"
        >
          <path
            fillRule="evenodd"
            d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z"
            clipRule="evenodd"
          />
        </svg>
      ),
    },
    {
      id: "profile",
      label: t("Profile"),
      path: "/profile",
      icon: (
        <svg
          className="w-4 h-4 lg:w-5 lg:h-5"
          fill="currentColor"
          viewBox="0 0 20 20"
        >
          <path
            fillRule="evenodd"
            d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z"
            clipRule="evenodd"
          />
        </svg>
      ),
    },
  ];

  const languageItem = {
    id: "switch-language",
    label:
      currentLanguage === "English"
        ? t("Switch to Amharic")
        : t("Switch to English"),
    path: "#",
    icon: (
      <Image
        src="/translate_icon.svg"
        alt="Translate"
        width={20}
        height={20}
        className={`w-4 h-4 lg:w-5 lg:h-5 object-contain transition-all ${
          isDarkMode ? "brightness-0 invert" : ""
        }`}
      />
    ),
  };

  return (
    <div
      className={`w-16 lg:w-64 border-r flex flex-col transition-colors ${
        isDarkMode
          ? "bg-gray-800/20 border-gray-700"
          : "bg-white border-gray-200"
      }`}
    >
      {/* Brand Header */}
      <div
        className={`p-3 lg:p-6 border-b transition-colors ${
          isDarkMode ? "border-gray-700" : "border-gray-200"
        }`}
      >
        <div className="flex items-center justify-center lg:justify-start gap-2">
          <div className="w-6 h-6 lg:w-8 lg:h-8 bg-[#FFD300] rounded flex items-center justify-center">
            <Image
              src="/WebsiteLogo/Frame.png"
              alt="ShopAlly Logo"
              width={24}
              height={24}
              className="object-contain w-4 h-4 lg:w-5 lg:h-5"
            />
          </div>
          <span
            className={`text-xl font-bold transition-colors hidden lg:block ${
              isDarkMode ? "text-white" : "text-gray-900"
            }`}
          >
            ShopAlly
          </span>
        </div>
      </div>

      {/* Navigation Links */}
      <nav className="flex-1 p-2 lg:p-4">
        <ul className="space-y-1 lg:space-y-2">
          {navigationItems.map((item) => (
            <li key={item.id}>
              <button
                onClick={() => router.push(item.path)}
                className={`flex items-center justify-center lg:justify-start gap-2 lg:gap-3 px-2 py-2 lg:px-3 lg:py-2 rounded-md transition-colors w-full ${
                  activePage === item.id
                    ? "bg-[#FFD300] text-gray-900 font-medium"
                    : isDarkMode
                    ? "text-gray-300 hover:bg-[#0000001A]"
                    : "text-gray-700 hover:bg-gray-100"
                }`}
              >
                {item.icon}
                <span className="hidden lg:block">{item.label}</span>
              </button>
            </li>
          ))}

          {/* Language Switch */}
          <li
            className={`mt-4 lg:mt-8 pt-2 lg:pt-4 border-t transition-colors ${
              isDarkMode ? "border-gray-700" : "border-gray-200"
            }`}
          >
            <button
              onClick={switchLanguage}
              className={`flex items-center justify-center lg:justify-start gap-2 lg:gap-3 px-2 py-2 lg:px-3 lg:py-2 rounded-md w-full transition-colors ${
                activePage === languageItem.id
                  ? "bg-[#FFD300] text-gray-900 font-medium"
                  : isDarkMode
                  ? "text-gray-300 hover:bg-[#0000001A]"
                  : "text-gray-700 hover:bg-gray-100"
              }`}
            >
              {languageItem.icon}
              <span className="hidden lg:block">{languageItem.label}</span>
            </button>
          </li>

          {/* Dark Mode Toggle */}
          <li className="mt-2 lg:mt-4">
            <button
              onClick={toggleDarkMode}
              className={`flex items-center justify-center lg:justify-start gap-2 lg:gap-3 px-2 py-2 lg:px-3 lg:py-2 rounded-md w-full transition-colors ${
                isDarkMode
                  ? "bg-[#0000001A] text-white hover:bg-[#00000033]"
                  : "text-gray-700 hover:bg-gray-100"
              }`}
            >
              <svg
                className="w-4 h-4 lg:w-5 lg:h-5"
                fill="currentColor"
                viewBox="0 0 20 20"
              >
                {isDarkMode ? (
                  <path d="M10 2a1 1 0 011 1v1a1 1 0 11-2 0V3a1 1 0 011-1zm4 8a4 4 0 11-8 0 4 4 0 018 0zm-.464 4.95l.707.707a1 1 0 001.414-1.414l-.707-.707a1 1 0 00-1.414 1.414zm2.12-10.607a1 1 0 010 1.414l-.706.707a1 1 0 11-1.414-1.414l.707-.707a1 1 0 011.414 0zM17 11a1 1 0 100-2h-1a1 1 0 100 2h1zm-7 4a1 1 0 011 1v1a1 1 0 11-2 0v-1a1 1 0 011-1zM5.05 6.464A1 1 0 106.465 5.05l-.708-.707a1 1 0 00-1.414 1.414l.707.707zm1.414 8.486l-.707.707a1 1 0 01-1.414-1.414l.707-.707a1 1 0 011.414 1.414zM4 11a1 1 0 100-2H3a1 1 0 000 2h1z" />
                ) : (
                  <path d="M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z" />
                )}
              </svg>
              <span className="hidden lg:block">
                {isDarkMode ? t("Light Mode") : t("Dark Mode")}
              </span>
            </button>
          </li>
        </ul>
      </nav>
    </div>
  );
}
