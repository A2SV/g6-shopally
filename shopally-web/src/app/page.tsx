"use client";

import Image from "next/image";
import ThemeToggle from "./components/NavBars/ThemeToggle";
import { useTheme } from "next-themes";
import { useEffect, useState } from "react";

export default function Home() {
  const { resolvedTheme } = useTheme();
  const [mounted, setMounted] = useState(false);
  const [htmlClasses, setHtmlClasses] = useState("");
  const [forceUpdate, setForceUpdate] = useState(0);

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      // Update htmlClasses whenever resolvedTheme changes
      setHtmlClasses(document.documentElement.className);

      // Force a re-render to ensure styles update
      setForceUpdate((prev) => prev + 1);

      // Also set up an interval to check for changes
      const interval = setInterval(() => {
        setHtmlClasses(document.documentElement.className);
      }, 100);

      return () => clearInterval(interval);
    }
  }, [resolvedTheme, mounted]);

  if (!mounted) return null;

  const isDark = resolvedTheme === "dark";

  return (
    <div
      className={`min-h-screen ${isDark ? "bg-brand-dark" : "bg-brand-white"}`}
    >
      <div className="container mx-auto px-4 py-8">
        <div className="flex items-center justify-between mb-8">
          <Image
            className="dark:invert"
            src="/next.svg"
            alt="Next.js logo"
            width={180}
            height={38}
            priority
          />
          <ThemeToggle />
        </div>

        <div className="max-w-2xl mx-auto text-center">
          <h1
            className={`text-4xl font-bold mb-6 ${
              isDark ? "text-brand-yellow" : "text-brand-yellow"
            }`}
          >
            Welcome to Shopally
          </h1>

          {/* Debug Info */}
          <div className="mb-6 p-4 bg-yellow-100 border-2 border-yellow-500">
            <h3 className="text-lg font-semibold text-yellow-800 mb-2">
              Debug Info
            </h3>
            <p className="text-yellow-700">
              Current Theme: <strong>{resolvedTheme}</strong>
            </p>
            <p className="text-yellow-700">
              HTML Classes: <strong>{htmlClasses}</strong>
            </p>
            <p className="text-yellow-700">
              Has 'dark' class:{" "}
              <strong>{htmlClasses.includes("dark") ? "YES" : "NO"}</strong>
            </p>
            <p className="text-yellow-700">
              Real-time check:{" "}
              <strong>{document.documentElement.className}</strong>
            </p>
            <p className="text-yellow-700">
              Force Update: <strong>{forceUpdate}</strong>
            </p>
          </div>

          {/* Simple Dark Mode Test */}
          <div
            className={`mb-6 p-4 border-2 border-brand-yellow ${
              isDark
                ? "bg-brand-dark text-brand-white"
                : "bg-brand-white text-brand-dark"
            }`}
          >
            <h3 className="text-lg font-semibold mb-2">
              Dark Mode Test (Simple)
            </h3>
            <p>
              This should be white background in light mode, dark background in
              dark mode.
            </p>
          </div>

          {/* Simple Tailwind Test */}
          <div className="mb-6 p-4 bg-blue-500 text-white rounded-lg">
            <h3 className="text-lg font-semibold mb-2">Basic Tailwind Test</h3>
            <p className="text-yellow-300">
              This should have a blue background and yellow text if Tailwind is
              working.
            </p>
          </div>

          {/* Simple Theme Test */}
          <div
            className={`mb-6 p-4 rounded-lg border border-brand-yellow/30 ${
              isDark ? "bg-brand-yellow/10" : "bg-brand-yellow/20"
            }`}
          >
            <h3
              className={`text-lg font-semibold mb-2 ${
                isDark ? "text-brand-white" : "text-brand-dark"
              }`}
            >
              Theme Test
            </h3>
            <p className={isDark ? "text-brand-white" : "text-brand-dark"}>
              This box should be light yellow in light mode and dark yellow in
              dark mode.
            </p>
            <p
              className={`text-sm mt-2 ${
                isDark ? "text-brand-gray" : "text-brand-gray"
              }`}
            >
              If you see this text in different colors when toggling, the theme
              is working!
            </p>
          </div>

          <div
            className={`p-6 rounded-lg mb-8 border border-brand-gray/20 ${
              isDark ? "bg-brand-darkGray" : "bg-brand-white"
            }`}
          >
            <h2
              className={`text-2xl font-semibold mb-4 ${
                isDark ? "text-brand-white" : "text-brand-dark"
              }`}
            >
              Tailwind CSS Test
            </h2>
            <p
              className={`mb-4 ${
                isDark ? "text-brand-gray" : "text-brand-gray"
              }`}
            >
              This page tests if Tailwind CSS is working properly. Try clicking
              the theme toggle!
            </p>

            <div className="mb-4">
              <p className="text-red-600 font-bold">
                This text should be RED (Tailwind: text-red-600)
              </p>
              <p className="text-blue-600 font-bold">
                This text should be BLUE (Tailwind: text-blue-600)
              </p>
              <p className="text-green-600 font-bold">
                This text should be GREEN (Tailwind: text-green-600)
              </p>
            </div>

            <ul
              className={`text-left space-y-2 ${
                isDark ? "text-brand-gray" : "text-brand-gray"
              }`}
            >
              <li className="flex items-center">
                <span className="w-3 h-3 bg-green-500 rounded-full mr-2"></span>
                Green dot (Tailwind class: bg-green-500)
              </li>
              <li className="flex items-center">
                <span className="w-3 h-3 bg-blue-500 rounded-full mr-2"></span>
                Blue dot (Tailwind class: bg-blue-500)
              </li>
              <li className="flex items-center">
                <span className="w-3 h-3 bg-purple-500 rounded-full mr-2"></span>
                Purple dot (Tailwind class: bg-purple-500)
              </li>
              <li className="flex items-center">
                <span className="w-3 h-3 bg-brand-yellow rounded-full mr-2"></span>
                Yellow dot (Brand color: bg-brand-yellow)
              </li>
            </ul>
          </div>

          <div className="grid md:grid-cols-2 gap-6">
            <div
              className={`p-4 rounded-lg border border-brand-gray/20 ${
                isDark ? "bg-brand-darkGray" : "bg-brand-white"
              }`}
            >
              <h3
                className={`text-lg font-semibold mb-2 ${
                  isDark ? "text-brand-white" : "text-brand-dark"
                }`}
              >
                Light Mode Card
              </h3>
              <p className={isDark ? "text-brand-gray" : "text-brand-gray"}>
                This card should have a light background in light mode and dark
                background in dark mode.
              </p>
            </div>

            <div
              className={`p-4 rounded-lg border border-brand-yellow/30 ${
                isDark ? "bg-brand-yellow/5" : "bg-brand-yellow/10"
              }`}
            >
              <h3
                className={`text-lg font-semibold mb-2 ${
                  isDark ? "text-brand-white" : "text-brand-dark"
                }`}
              >
                Dark Mode Card
              </h3>
              <p className={isDark ? "text-brand-gray" : "text-brand-gray"}>
                This card should have a light yellow background in light mode
                and dark yellow in dark mode.
              </p>
            </div>
          </div>

          <div
            className={`mt-8 p-4 rounded-lg border border-brand-yellow/30 ${
              isDark ? "bg-brand-yellow/10" : "bg-brand-yellow/20"
            }`}
          >
            <h3
              className={`text-lg font-semibold mb-2 ${
                isDark ? "text-brand-white" : "text-brand-dark"
              }`}
            >
              Dark Mode Test
            </h3>
            <p className={isDark ? "text-brand-white" : "text-brand-dark"}>
              If you see this text in dark and the background changes when you
              toggle the theme, then Tailwind CSS is working!
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
