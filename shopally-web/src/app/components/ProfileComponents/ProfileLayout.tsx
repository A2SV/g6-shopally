"use client";

import { ReactNode } from "react";
import { useDarkMode } from "./DarkModeContext";

interface ProfileLayoutProps {
  children: ReactNode;
}

export default function ProfileLayout({ children }: ProfileLayoutProps) {
  const { isDarkMode } = useDarkMode();

  return (
    <div
      className={`min-h-screen flex transition-colors ${
        isDarkMode ? "bg-[#090C11]" : "bg-gray-50"
      }`}
    >
      {children}
    </div>
  );
}
