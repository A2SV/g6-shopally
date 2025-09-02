"use client";

import ProfileLayout from "@/app/components/ProfileComponents/ProfileLayout";
import Sidebar from "@/app/components/Sidebar";
import ProfileContent from "@/app/components/ProfileComponents/ProfileContent";
import ProfileStats from "@/app/components/ProfileComponents/ProfileStats";
import {
  DarkModeProvider,
  useDarkMode,
} from "@/app/components/ProfileComponents/DarkModeContext";
import {
  LanguageProvider,
  useLanguage,
} from "@/app/components/LanguageContext";

function ProfilePageContent() {
  const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();

  return (
    <ProfileLayout>
      <Sidebar activePage="profile" />
      <div className="flex-1 flex flex-col p-4 lg:p-8">
        {/* Page Title */}
        <h1
          className={`text-2xl sm:text-3xl font-bold mb-6 sm:mb-8 transition-colors ${
            isDarkMode ? "text-white" : "text-gray-900"
          }`}
        >
          {t("Profile Settings")}
        </h1>

        {/* Cards Container */}
        <div className="flex flex-col lg:flex-row gap-4 lg:gap-8">
          <div className="flex-1">
            <ProfileContent />
          </div>
          <div className="flex-1">
            <ProfileStats />
          </div>
        </div>
      </div>
    </ProfileLayout>
  );
}

export default function ProfilePage() {
  return (
    <LanguageProvider>
      <DarkModeProvider>
        <ProfilePageContent />
      </DarkModeProvider>
    </LanguageProvider>
  );
}
