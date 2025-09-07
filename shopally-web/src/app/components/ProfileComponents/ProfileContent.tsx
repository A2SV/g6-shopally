"use client";

import { useState, useRef, useEffect } from "react";
import Image from "next/image";
import { useDarkMode } from "./DarkModeContext";
import { useLanguage } from "../../../hooks/useLanguage";
import { LoadingSpinner } from "../LoadingSpinner";
import { useSession } from "next-auth/react";

export default function ProfileContent() {
  const [selectedCategories, setSelectedCategories] = useState([
    "Electronics",
    "Fashion",
  ]);
  const [currency, setCurrency] = useState("USD");
  const [region, setRegion] = useState("US");
  const [language, setLanguage] = useState("English");
  const [profileImage, setProfileImage] = useState<string | null>(null);
  const [isEditing, setIsEditing] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [isSaving, setIsSaving] = useState(false);
  const [isUploading, setIsUploading] = useState(false);
  const [userData, setUserData] = useState({
    fullName: "",
    email: "",
  });
  const [originalUserData, setOriginalUserData] = useState({
    fullName: "",
    email: "",
  });
  const [showMessage, setShowMessage] = useState<{
    type: "success" | "info" | "error" | null;
    message: string;
  }>({ type: null, message: "" });
  const fileInputRef = useRef<HTMLInputElement>(null);
  const { isDarkMode } = useDarkMode();
  const { t } = useLanguage();
  const { data: session } = useSession();

  // Initialize user data from session
  useEffect(() => {
    if (session?.user) {
      const userInfo = {
        fullName: session.user.name || "",
        email: session.user.email || "",
      };

      // Check localStorage first for custom user data, fallback to session data
      const savedUserData = localStorage.getItem(
        `userData_${session.user.email}`
      );
      if (savedUserData) {
        try {
          const parsedData = JSON.parse(savedUserData);
          setUserData(parsedData);
          setOriginalUserData(parsedData);
        } catch (error) {
          // If parsing fails, use session data
          setUserData(userInfo);
          setOriginalUserData(userInfo);
        }
      } else {
        setUserData(userInfo);
        setOriginalUserData(userInfo);
      }

      // Check localStorage first for custom uploaded image, fallback to session image
      const savedProfileImage = localStorage.getItem(
        `profileImage_${session.user.email}`
      );
      if (savedProfileImage) {
        setProfileImage(savedProfileImage);
      } else {
        setProfileImage(session.user.image || null);
      }
    }
  }, [session]);

  // Simulate loading on component mount
  useEffect(() => {
    const timer = setTimeout(() => {
      setIsLoading(false);
    }, 1500);
    return () => clearTimeout(timer);
  }, []);

  // Original category names (for selection logic)
  const originalCategories = [
    "Electronics",
    "Fashion",
    "Home & Garden",
    "Sports & Outdoors",
    "Books",
    "Beauty & Health",
    "Automotive",
    "Toys & Games",
  ];

  // Translated categories (for display)
  const categories = [
    t("Electronics"),
    t("Fashion"),
    t("Home & Garden"),
    t("Sports & Outdoors"),
    t("Books"),
    t("Beauty & Health"),
    t("Automotive"),
    t("Toys & Games"),
  ];

  const toggleCategory = (originalCategory: string) => {
    setSelectedCategories((prev) =>
      prev.includes(originalCategory)
        ? prev.filter((c) => c !== originalCategory)
        : [...prev, originalCategory]
    );
  };

  const handleSavePreferences = async () => {
    setIsSaving(true);
    try {
      // Simulate API call
      await new Promise((resolve) => setTimeout(resolve, 2000));
      setShowMessage({
        type: "success",
        message: t("Preferences saved successfully!"),
      });
    } catch (error) {
      setShowMessage({
        type: "error",
        message: t("Failed to save preferences. Please try again."),
      });
    } finally {
      setIsSaving(false);
      setTimeout(() => setShowMessage({ type: null, message: "" }), 3000);
    }
  };

  const handleImageUpload = async (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = event.target.files?.[0];
    if (file) {
      setIsUploading(true);
      try {
        // Simulate upload delay
        await new Promise((resolve) => setTimeout(resolve, 2000));
        const reader = new FileReader();
        reader.onload = (e) => {
          const imageDataUrl = e.target?.result as string;
          setProfileImage(imageDataUrl);

          // Save to localStorage with user email as key
          if (session?.user?.email) {
            localStorage.setItem(
              `profileImage_${session.user.email}`,
              imageDataUrl
            );
          }

          setShowMessage({
            type: "success",
            message: t("Profile image uploaded successfully!"),
          });
        };
        reader.readAsDataURL(file);
      } catch (error) {
        setShowMessage({
          type: "error",
          message: t("Failed to upload image. Please try again."),
        });
      } finally {
        setIsUploading(false);
        setTimeout(() => setShowMessage({ type: null, message: "" }), 3000);
      }
    }
  };

  const handleEditPhoto = () => {
    fileInputRef.current?.click();
  };

  const handleRemovePhoto = () => {
    if (session?.user?.email) {
      // Remove from localStorage
      localStorage.removeItem(`profileImage_${session.user.email}`);
      // Revert to email account image
      setProfileImage(session.user.image || null);
      setShowMessage({
        type: "success",
        message: t("Profile image removed successfully!"),
      });
      setTimeout(() => setShowMessage({ type: null, message: "" }), 3000);
    }
  };

  const handleEditToggle = () => {
    setIsEditing(!isEditing);
    setShowMessage({ type: null, message: "" });
  };

  const handleSaveChanges = async () => {
    // Check if anything has changed
    const hasChanges =
      userData.fullName !== originalUserData.fullName ||
      userData.email !== originalUserData.email;

    if (!hasChanges) {
      setShowMessage({
        type: "info",
        message: t("No changes detected. Nothing to update."),
      });
      setTimeout(() => setShowMessage({ type: null, message: "" }), 3000);
      return;
    }

    setIsSaving(true);
    try {
      // Simulate API call
      await new Promise((resolve) => setTimeout(resolve, 2000));

      // Save to localStorage
      if (session?.user?.email) {
        localStorage.setItem(
          `userData_${session.user.email}`,
          JSON.stringify(userData)
        );
      }

      // Update original data to match current data
      setOriginalUserData(userData);

      setShowMessage({
        type: "success",
        message: t("Profile updated successfully!"),
      });

      setIsEditing(false);
    } catch (error) {
      setShowMessage({
        type: "error",
        message: t("Failed to update profile. Please try again."),
      });
    } finally {
      setIsSaving(false);
      setTimeout(() => setShowMessage({ type: null, message: "" }), 3000);
    }
  };

  const handleCancelEdit = () => {
    // Reset to original values
    setUserData(originalUserData);
    setIsEditing(false);
    setShowMessage({ type: null, message: "" });
  };

  const handleResetToSessionData = () => {
    if (session?.user) {
      const sessionData = {
        fullName: session.user.name || "",
        email: session.user.email || "",
      };

      // Remove from localStorage
      localStorage.removeItem(`userData_${session.user.email}`);
      localStorage.removeItem(`profileImage_${session.user.email}`);

      // Reset to session data
      setUserData(sessionData);
      setOriginalUserData(sessionData);
      setProfileImage(session.user.image || null);

      setShowMessage({
        type: "success",
        message: t("Profile reset to original data!"),
      });
      setTimeout(() => setShowMessage({ type: null, message: "" }), 3000);
    }
  };

  const handleInputChange = (field: string, value: string) => {
    setUserData((prev) => ({
      ...prev,
      [field]: value,
    }));
  };

  // Show loading state
  if (isLoading) {
    return (
      <div
        className="flex-1 transition-colors"
        style={{
          backgroundColor: "var(--color-bg-primary)",
        }}
      >
        <div className="space-y-6 sm:space-y-8">
          <div
            className="rounded-lg shadow-sm border p-4 sm:p-6 transition-colors"
            style={{
              backgroundColor: "var(--color-bg-card)",
              borderColor: "var(--color-border-primary)",
            }}
          >
            <div className="flex items-center justify-center py-12">
              <LoadingSpinner size="xl" color="yellow" />
            </div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div
      className="flex-1 transition-colors"
      style={{
        backgroundColor: "var(--color-bg-primary)",
      }}
    >
      <div className="space-y-6 sm:space-y-8">
        {/* User Information Card */}
        <div
          className="rounded-lg shadow-sm border p-4 sm:p-6 transition-colors"
          style={{
            backgroundColor: "var(--color-bg-card)",
            borderColor: "var(--color-border-primary)",
          }}
        >
          <div className="flex items-center justify-between mb-6">
            <h2
              className="text-lg sm:text-xl font-semibold transition-colors"
              style={{ color: "var(--color-text-primary)" }}
            >
              {t("User Information")}
            </h2>
            {!isEditing ? (
              <button
                onClick={handleEditToggle}
                className="flex items-center gap-2 font-medium hover:opacity-80 transition-colors"
                style={{
                  color: "var(--color-accent-primary)",
                }}
              >
                <svg
                  className={`w-4 h-4 transition-colors ${
                    isDarkMode ? "text-black" : "text-gray-900"
                  }`}
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />
                </svg>
                {t("Edit")}
              </button>
            ) : (
              <button
                onClick={handleCancelEdit}
                className="px-3 py-1 font-medium rounded-md hover:opacity-80 transition-colors text-sm"
                style={{
                  backgroundColor: "var(--color-bg-tertiary)",
                  color: "var(--color-text-primary)",
                }}
              >
                {t("Cancel")}
              </button>
            )}
          </div>

          {/* Message Display */}
          {showMessage.type && (
            <div
              className={`mb-4 p-3 rounded-md text-sm font-medium ${
                showMessage.type === "success"
                  ? "bg-green-100 text-green-800 border border-green-200"
                  : showMessage.type === "error"
                  ? "bg-red-100 text-red-800 border border-red-200"
                  : "bg-blue-100 text-blue-800 border border-blue-200"
              }`}
            >
              {showMessage.message}
            </div>
          )}

          <div className="space-y-4 sm:space-y-6">
            {/* Avatar */}
            <div className="flex justify-center">
              <div className="relative">
                <div
                  className="w-20 h-20 rounded-full flex items-center justify-center overflow-hidden"
                  style={{ backgroundColor: "var(--color-accent-primary)" }}
                >
                  {profileImage ? (
                    <img
                      src={profileImage}
                      alt="Profile"
                      className="w-full h-full object-cover"
                    />
                  ) : (
                    <span
                      className="text-2xl font-bold"
                      style={{ color: "var(--color-text-button)" }}
                    >
                      {userData.fullName
                        ? userData.fullName
                            .split(" ")
                            .map((n) => n[0])
                            .join("")
                            .toUpperCase()
                        : "U"}
                    </span>
                  )}
                </div>
                <button
                  onClick={handleEditPhoto}
                  disabled={isUploading}
                  className={`absolute -bottom-1 -right-1 w-5.5 h-5.5 rounded-full flex items-center justify-center transition-colors shadow-sm ${
                    isUploading
                      ? "bg-gray-400 cursor-not-allowed"
                      : "bg-[var(--color-brand-yellow)] hover:bg-[#E6BE00]"
                  }`}
                >
                  {isUploading ? (
                    <LoadingSpinner size="sm" color="gray" />
                  ) : (
                    <Image
                      src="/edit-img.png"
                      alt="Edit Photo"
                      width={16}
                      height={16}
                      className="w-4 h-4 object-contain"
                    />
                  )}
                </button>
                <input
                  ref={fileInputRef}
                  type="file"
                  accept="image/*"
                  onChange={handleImageUpload}
                  className="hidden"
                />
              </div>
            </div>

            {/* User Details */}
            <div className="space-y-3 sm:space-y-4">
              <div>
                <label
                  className="block text-sm font-medium mb-1 transition-colors"
                  style={{ color: "var(--color-text-secondary)" }}
                >
                  {t("Full Name")}
                </label>
                <input
                  type="text"
                  value={userData.fullName}
                  onChange={(e) =>
                    handleInputChange("fullName", e.target.value)
                  }
                  disabled={!isEditing}
                  className={`w-full px-3 py-2 border rounded-md transition-colors ${
                    !isEditing ? "cursor-not-allowed" : "cursor-text"
                  }`}
                  style={{
                    borderColor: "var(--color-border-primary)",
                    backgroundColor: isEditing
                      ? "var(--color-bg-input)"
                      : "var(--color-bg-tertiary)",
                    color: "var(--color-text-primary)",
                  }}
                />
              </div>

              <div>
                <label
                  className="block text-sm font-medium mb-1 transition-colors"
                  style={{ color: "var(--color-text-secondary)" }}
                >
                  {t("Email Address")}
                </label>
                <input
                  type="email"
                  value={userData.email}
                  onChange={(e) => handleInputChange("email", e.target.value)}
                  disabled={!isEditing}
                  className={`w-full px-3 py-2 border rounded-md transition-colors ${
                    !isEditing ? "cursor-not-allowed" : "cursor-text"
                  }`}
                  style={{
                    borderColor: "var(--color-border-primary)",
                    backgroundColor: isEditing
                      ? "var(--color-bg-input)"
                      : "var(--color-bg-tertiary)",
                    color: "var(--color-text-primary)",
                  }}
                />
              </div>

              {/* Save Changes Button */}
              {isEditing && (
                <div className="pt-4">
                  <button
                    onClick={handleSaveChanges}
                    disabled={isSaving}
                    className={`w-full font-semibold py-3 px-6 rounded-lg transition-all duration-200 ${
                      isSaving
                        ? "bg-gray-400 text-gray-600 cursor-not-allowed"
                        : "hover:opacity-80"
                    }`}
                    style={{
                      backgroundColor: isSaving
                        ? "var(--color-text-muted)"
                        : "var(--color-accent-primary)",
                      color: "var(--color-text-button)",
                    }}
                  >
                    {isSaving ? (
                      <div className="flex items-center justify-center gap-2">
                        <LoadingSpinner size="sm" color="gray" />
                        {t("Saving...")}
                      </div>
                    ) : (
                      t("Save Changes")
                    )}
                  </button>
                </div>
              )}
            </div>
          </div>
        </div>

        {/* ShopAlly AI Personalization Card */}
        <div
          className={`rounded-lg shadow-sm border p-4 sm:p-6 transition-colors ${
            isDarkMode
              ? "bg-gray-800/20 border-gray-700"
              : "bg-[var(--color-brand-white)] border-gray-200"
          }`}
        >
          {/* Header */}
          <h2
            className={`text-lg sm:text-xl font-semibold mb-4 sm:mb-6 transition-colors ${
              isDarkMode ? "text-[var(--color-brand-white)]" : "text-gray-900"
            }`}
          >
            {t("ShopAlly AI Personalization")}
          </h2>

          <div className="space-y-6 sm:space-y-8">
            {/* Preferred Product Categories */}
            <div>
              <label
                className={`block text-sm font-semibold mb-3 transition-colors ${
                  isDarkMode ? "text-gray-300" : "text-gray-700"
                }`}
              >
                {t("Preferred Product Categories")}
              </label>
              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-2 sm:gap-3">
                {categories.map((translatedCategory, index) => {
                  const originalCategory = originalCategories[index];
                  return (
                    <button
                      key={originalCategory}
                      onClick={() => toggleCategory(originalCategory)}
                      className={`px-3 py-2 rounded-lg text-sm font-medium transition-all duration-200 border ${
                        selectedCategories.includes(originalCategory)
                          ? "bg-[var(--color-brand-yellow)] text-gray-900 border-[#FFD300]"
                          : isDarkMode
                          ? "bg-[#0000001A] text-gray-300 hover:bg-[#00000033] border-gray-600 hover:border-gray-500"
                          : "bg-gray-50 text-gray-700 hover:bg-gray-100 border-gray-200 hover:border-gray-300"
                      }`}
                    >
                      {translatedCategory}
                    </button>
                  );
                })}
              </div>
            </div>

            {/* Settings Grid */}
            <div className="space-y-4 sm:space-y-6">
              {/* Default Currency and Shipping Region */}
              <div className="grid grid-cols-1 lg:grid-cols-2 gap-4 sm:gap-6">
                {/* Default Currency */}
                <div>
                  <label
                    className={`block text-sm font-semibold mb-2 transition-colors ${
                      isDarkMode ? "text-gray-300" : "text-gray-700"
                    }`}
                  >
                    {t("Default Currency")}
                  </label>
                  <select
                    value={currency}
                    onChange={(e) => setCurrency(e.target.value)}
                    className={`w-full px-3 py-2.5 border rounded-lg focus:outline-none focus:ring-2 focus:ring-[#FFD300] focus:border-transparent transition-colors text-sm ${
                      isDarkMode
                        ? "border-gray-600 bg-[var(--color-brand-white)] text-gray-900"
                        : "border-gray-300 text-gray-900 bg-[#D9D9D9]"
                    }`}
                  >
                    <option
                      value="USD"
                      className={
                        isDarkMode
                          ? "text-[var(--color-brand-white)] bg-[var(--color-brand-dark)]"
                          : "text-gray-900 bg-[var(--color-brand-white)]"
                      }
                    >
                      USD
                    </option>
                    <option
                      value="EUR"
                      className={
                        isDarkMode
                          ? "text-[var(--color-brand-white)] bg-[var(--color-brand-dark)]"
                          : "text-gray-900 bg-[var(--color-brand-white)]"
                      }
                    >
                      EUR
                    </option>
                    <option
                      value="GBP"
                      className={
                        isDarkMode
                          ? "text-[var(--color-brand-white)] bg-[var(--color-brand-dark)]"
                          : "text-gray-900 bg-[var(--color-brand-white)]"
                      }
                    >
                      GBP
                    </option>
                    <option
                      value="CAD"
                      className={
                        isDarkMode
                          ? "text-[var(--color-brand-white)] bg-[var(--color-brand-dark)]"
                          : "text-gray-900 bg-[var(--color-brand-white)]"
                      }
                    >
                      CAD
                    </option>
                  </select>
                </div>

                {/* Shipping Region */}
                <div>
                  <label
                    className={`block text-sm font-semibold mb-2 transition-colors ${
                      isDarkMode ? "text-gray-300" : "text-gray-700"
                    }`}
                  >
                    {t("Shipping Region")}
                  </label>
                  <select
                    value={region}
                    onChange={(e) => setRegion(e.target.value)}
                    className={`w-full px-3 py-2.5 border rounded-lg focus:outline-none focus:ring-2 focus:ring-[#FFD300] focus:border-transparent transition-colors text-sm ${
                      isDarkMode
                        ? "border-gray-600 bg-[var(--color-brand-white)] text-gray-900"
                        : "border-gray-300 text-gray-900 bg-[#D9D9D9]"
                    }`}
                  >
                    <option
                      value="US"
                      className={
                        isDarkMode
                          ? "text-[var(--color-brand-white)] bg-[var(--color-brand-dark)]"
                          : "text-gray-900 bg-[var(--color-brand-white)]"
                      }
                    >
                      US
                    </option>
                    <option
                      value="CA"
                      className={
                        isDarkMode
                          ? "text-[var(--color-brand-white)] bg-[var(--color-brand-dark)]"
                          : "text-gray-900 bg-[var(--color-brand-white)]"
                      }
                    >
                      CA
                    </option>
                    <option
                      value="UK"
                      className={
                        isDarkMode
                          ? "text-[var(--color-brand-white)] bg-[var(--color-brand-dark)]"
                          : "text-gray-900 bg-[var(--color-brand-white)]"
                      }
                    >
                      UK
                    </option>
                    <option
                      value="EU"
                      className={
                        isDarkMode
                          ? "text-[var(--color-brand-white)] bg-[var(--color-brand-dark)]"
                          : "text-gray-900 bg-[var(--color-brand-white)]"
                      }
                    >
                      EU
                    </option>
                  </select>
                </div>
              </div>

              {/* Language Preference */}
              <div>
                <label
                  className={`block text-sm font-semibold mb-2 transition-colors ${
                    isDarkMode ? "text-gray-300" : "text-gray-700"
                  }`}
                >
                  {t("Language Preference")}
                </label>
                <select
                  value={language}
                  onChange={(e) => setLanguage(e.target.value)}
                  className={`w-full px-3 py-2.5 border rounded-lg focus:outline-none focus:ring-2 focus:ring-[#FFD300] focus:border-transparent transition-colors text-sm ${
                    isDarkMode
                      ? "border-gray-600 bg-[var(--color-brand-white)] text-gray-900"
                      : "border-gray-300 text-gray-900 bg-[#D9D9D9]"
                  }`}
                >
                  <option
                    value="English"
                    className={
                      isDarkMode
                        ? "text-[var(--color-brand-white)] bg-[var(--color-brand-dark)]"
                        : "text-gray-900 bg-[var(--color-brand-white)]"
                    }
                  >
                    English
                  </option>
                  <option
                    value="Spanish"
                    className={
                      isDarkMode
                        ? "text-[var(--color-brand-white)] bg-[var(--color-brand-dark)]"
                        : "text-gray-900 bg-[var(--color-brand-white)]"
                    }
                  >
                    Spanish
                  </option>
                  <option
                    value="French"
                    className={
                      isDarkMode
                        ? "text-[var(--color-brand-white)] bg-[var(--color-brand-dark)]"
                        : "text-gray-900 bg-[var(--color-brand-white)]"
                    }
                  >
                    French
                  </option>
                  <option
                    value="German"
                    className={
                      isDarkMode
                        ? "text-[var(--color-brand-white)] bg-[var(--color-brand-dark)]"
                        : "text-gray-900 bg-[var(--color-brand-white)]"
                    }
                  >
                    German
                  </option>
                </select>
              </div>
            </div>

            {/* Save Preferences Button */}
            <div className="pt-4">
              <button
                onClick={handleSavePreferences}
                disabled={isSaving}
                className={`w-full font-semibold py-3 px-6 rounded-lg transition-all duration-200 ${
                  isSaving
                    ? "bg-gray-400 text-gray-600 cursor-not-allowed"
                    : "bg-[var(--color-brand-yellow)] text-gray-900 hover:bg-[#E6BE00]"
                }`}
              >
                {isSaving ? (
                  <div className="flex items-center justify-center gap-2">
                    <LoadingSpinner size="sm" color="gray" />
                    {t("Saving...")}
                  </div>
                ) : (
                  t("Save Preferences")
                )}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
