"use client";

import Image from "next/image";
import { Facebook, Twitter, Linkedin, Instagram } from "lucide-react";
import { useDarkMode } from "@/app/components/ProfileComponents/DarkModeContext";
import { useLanguage } from "@/hooks/useLanguage";

export default function Footer() {
  const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();

  return (
    <footer
      className={`mt-12 py-10 px-6 md:px-12 transition-colors ${
        isDarkMode ? "bg-[#0D1117] text-gray-300" : "bg-gray-50 text-gray-700"
      }`}
    >
      <div className="max-w-6xl mx-auto grid grid-cols-1 md:grid-cols-3 gap-8">
        {/* Logo + Description */}
        <div>
          <div className="flex items-center gap-3 mb-3">
            <div className="bg-[#F59E0B] p-2 flex items-center justify-center w-12 h-12">
              <Image
                src="/WebsiteLogo/Frame.png"
                alt="ShopAlly Logo"
                width={30}
                height={30}
                className="object-contain"
              />
            </div>
            <h2
              className={`text-2xl font-bold ${
                isDarkMode ? "text-white" : "text-gray-900"
              }`}
            >
              ShopAlly
            </h2>
          </div>
          <p className="text-sm">
            {t(
              "Your AI-powered shopping assistant for Alibaba, helping you shop smarter with Ethiopian pricing."
            )}
          </p>
        </div>

        {/* Quick Links */}
        <div>
          <h3
            className={`font-semibold mb-3 ${
              isDarkMode ? "text-white" : "text-gray-900"
            }`}
          >
            {t("Quick Links")}
          </h3>
          <ul className="space-y-2">
            <li>
              <a href="/about" className="hover:text-[#F59E0B] transition-colors">
                {t("About Us")}
              </a>
            </li>
            <li>
              <a
                href="/how-it-works"
                className="hover:text-[#F59E0B] transition-colors"
              >
                {t("How It Works")}
              </a>
            </li>
            <li>
              <a href="/contact" className="hover:text-[#F59E0B] transition-colors">
                {t("Contact")}
              </a>
            </li>
          </ul>
        </div>

        {/* Social Media */}
        <div>
          <h3
            className={`font-semibold mb-3 ${
              isDarkMode ? "text-white" : "text-gray-900"
            }`}
          >
            {t("Follow Us")}
          </h3>
          <div className="flex space-x-4">
            <a href="#" aria-label="Facebook" className="hover:opacity-80">
              <Facebook className="w-5 h-5 text-[#F59E0B]" />
            </a>
            <a href="#" aria-label="Twitter" className="hover:opacity-80">
              <Twitter className="w-5 h-5 text-[#F59E0B]" />
            </a>
            <a href="#" aria-label="LinkedIn" className="hover:opacity-80">
              <Linkedin className="w-5 h-5 text-[#F59E0B]" />
            </a>
            <a href="#" aria-label="Instagram" className="hover:opacity-80">
              <Instagram className="w-5 h-5 text-[#F59E0B]" />
            </a>
          </div>
        </div>
      </div>

      {/* Bottom */}
      <div
        className={`mt-8 pt-6 border-t text-center text-sm ${
          isDarkMode
            ? "border-gray-700 text-gray-400"
            : "border-gray-300 text-gray-600"
        }`}
      >
        © {new Date().getFullYear()}{" "}
        <span className="font-semibold text-[#F59E0B]">ShopAlly</span>.{" "}
        {t("All rights reserved.")}.
      </div>
    </footer>
  );
}
