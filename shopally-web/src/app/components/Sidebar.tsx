"use client";

import Link from "next/link";
import { useDarkMode } from "./ProfileComponents/DarkModeContext";
import { useLanguage } from "../../hooks/useLanguage";

interface SidebarProps {
  activePage?: string;
}

export default function Sidebar({ activePage = "profile" }: SidebarProps) {
  const { isDarkMode, toggleDarkMode } = useDarkMode();
  const { currentLanguage, switchLanguage, t } = useLanguage();

  const navigationItems = [
    {
      id: "home",
      label: t("Home"),
      path: "/home",
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
      path: "/comparison",
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
      <svg
        className="w-4 h-4 lg:w-5 lg:h-5"
        fill="currentColor"
        viewBox="0 0 20 20"
      >
        <path
          fillRule="evenodd"
          d="M7 2a1 1 0 011 1v1h3a1 1 0 110 2H9.578a18.87 18.87 0 01-1.724 4.78c.29.354.596.696.914 1.026a1 1 0 11-1.44 1.389c-.188-.196-.373-.396-.554-.6a19.098 19.098 0 01-3.107 3.567 1 1 0 01-1.334-1.49 17.087 17.087 0 003.13-3.733 18.992 18.992 0 01-1.487-2.494 1 1 0 111.79-.89c.234.47.489.928.764 1.372.417-.934.752-1.913.997-2.927H3a1 1 0 110-2h3V3a1 1 0 011-1zm6 6a1 1 0 01.894.553l2.991 5.982a.869.869 0 01.02.037l.99 1.98a1 1 0 11-1.79.895L15.383 16h-4.764l-.724 1.447a1 1 0 11-1.788-.894l.99-1.98.019-.038 2.99-5.982A1 1 0 0113 8zm-1.382 6h2.764L13 11.236 11.618 14z"
          clipRule="evenodd"
        />
      </svg>
    ),
  };

  return (
    <div
      className="fixed left-0 top-0 h-screen w-16 lg:w-64 border-r flex flex-col transition-colors z-50"
      style={{
        backgroundColor: isDarkMode
          ? "var(--color-bg-secondary)"
          : "var(--color-bg-primary)",
        borderColor: isDarkMode
          ? "var(--color-border-primary)"
          : "var(--color-border-primary)",
      }}
    >
      {/* Brand Header */}
      <div
        className="p-3 lg:p-6 border-b transition-colors"
        style={{
          borderColor: "var(--color-border-primary)",
        }}
      >
        <div className="flex items-center justify-center lg:justify-start gap-2">
          <div className="w-6 h-6 lg:w-8 lg:h-8 bg-[var(--color-brand-yellow)] rounded flex items-center justify-center">
            <img
              src="WebsiteLogo/Frame.png"
              alt="ShopAlly Logo"
              className="object-contain w-4 h-4 lg:w-5 lg:h-5"
              onError={(e) => {
                // Fallback to text if image fails to load
                e.currentTarget.style.display = "none";
                const fallback = document.createElement("span");
                fallback.textContent = "S";
                fallback.className = "text-black font-bold text-sm";
                e.currentTarget.parentNode?.appendChild(fallback);
              }}
            />
          </div>
          <span
            className="text-xl font-bold transition-colors hidden lg:block"
            style={{ color: "var(--color-text-primary)" }}
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
              <Link
                href={item.path}
                className={`flex items-center justify-center lg:justify-start gap-2 lg:gap-3 px-2 py-2 lg:px-3 lg:py-2 rounded-md transition-colors w-full ${
                  activePage === item.id ? "font-medium" : "hover:opacity-80"
                }`}
                style={{
                  backgroundColor:
                    activePage === item.id
                      ? "var(--color-accent-primary)"
                      : "transparent",
                  color:
                    activePage === item.id
                      ? "var(--color-text-button)"
                      : "var(--color-text-secondary)",
                }}
              >
                {item.icon}
                <span className="hidden lg:block">{item.label}</span>
              </Link>
            </li>
          ))}

          {/* Language Switch */}
          <li
            className="mt-4 lg:mt-8 pt-2 lg:pt-4 border-t transition-colors"
            style={{
              borderColor: "var(--color-border-primary)",
            }}
          >
            <button
              onClick={switchLanguage}
              className="flex items-center justify-center lg:justify-start gap-2 lg:gap-3 px-2 py-2 lg:px-3 lg:py-2 rounded-md w-full transition-colors hover:opacity-80"
              style={{
                color: "var(--color-text-secondary)",
              }}
            >
              {languageItem.icon}
              <span className="hidden lg:block">{languageItem.label}</span>
            </button>
          </li>

          {/* Dark Mode Toggle */}
          <li className="mt-2 lg:mt-4">
            <button
              onClick={toggleDarkMode}
              className="flex items-center justify-center lg:justify-start gap-2 lg:gap-3 px-2 py-2 lg:px-3 lg:py-2 rounded-md w-full transition-colors hover:opacity-80"
              style={{
                backgroundColor: isDarkMode
                  ? "var(--color-overlay)"
                  : "transparent",
                color: "var(--color-text-secondary)",
              }}
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
