"use client";

import ProfileContent from "@/app/components/ProfileComponents/ProfileContent";
import ProfileStats from "@/app/components/ProfileComponents/ProfileStats";
// import { useDarkMode } from "@/app/components/ProfileComponents/DarkModeContext";
import { useLanguage } from "@/hooks/useLanguage";

export default function ProfilePage() {
  // const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();

  return (
    <div
      className="min-h-screen flex flex-col p-4 lg:p-8 transition-colors"
      style={{
        backgroundColor: "var(--color-bg-primary)",
        color: "var(--color-text-primary)",
      }}
    >
      {/* Page Title */}
      <h1
        className="text-2xl sm:text-3xl font-bold mb-6 sm:mb-8 transition-colors"
        style={{ color: "var(--color-text-primary)" }}
      >
        {t("Profile Settings")}
      </h1>

      {/* Cards Container */}
      <div className="flex flex-col xl:flex-row gap-4 sm:gap-6 lg:gap-8">
        {/* User Information and Personalization - Always on top on mobile */}
        <div className="flex-1 order-1">
          <ProfileContent />
        </div>
        {/* Account Statistics - Below on mobile, right side on desktop */}
        <div className="flex-1 order-2 xl:order-2">
          <ProfileStats />
        </div>
      </div>
    </div>
  );
}
