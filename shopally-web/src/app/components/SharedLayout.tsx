"use client";

import { ReactNode } from "react";
import Sidebar from "./Sidebar";
import { usePathname } from "next/navigation";

interface SharedLayoutProps {
  children: ReactNode;
}

export default function SharedLayout({ children }: SharedLayoutProps) {
  const pathname = usePathname();

  // Determine if sidebar should be shown (only on main app pages)
  const shouldShowSidebar = () => {
    const mainAppPages = [
      "/",
      "/home",
      "/comparison",
      "/saved-items",
      "/profile",
      "/how-it-works",
    ];
    return mainAppPages.includes(pathname);
  };

  // Determine active page based on current pathname
  const getActivePage = () => {
    if (pathname === "/" || pathname === "/home") return "home";
    if (pathname === "/comparison") return "compare";
    if (pathname === "/saved-items") return "saved-items";
    if (pathname === "/profile") return "profile";
    if (pathname === "/how-it-works") return "how-it-works";
    return "home";
  };

  // If it's an auth page or other non-main pages, render without sidebar
  if (!shouldShowSidebar()) {
    return <div className="min-h-screen">{children}</div>;
  }

  // For main app pages, render with sidebar
  return (
    <div className="min-h-screen flex relative">
      <Sidebar activePage={getActivePage()} />
      <div className="flex-1 ml-16 lg:ml-64 relative">{children}</div>
    </div>
  );
}
