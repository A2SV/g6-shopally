import AuthProvider from "@/providers/AuthProvider";
import { LanguageProvider } from "@/providers/LanguageProvider";
import ReduxProvider from "@/providers/ReduxProvider";
import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import FirebaseSWRegistrar from "./components/FirebaseSWRegistrar";
import { DarkModeProvider } from "./components/ProfileComponents/DarkModeContext";
import SharedLayout from "./components/SharedLayout";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "ShopAlly - Your Smart AI Assistant for AliExpress Shopping",
  description:
    "Discover the perfect products on AliExpress with AI-powered recommendations tailored to your needs.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased transition-colors duration-200`}
        style={{
          backgroundColor: "var(--color-bg-primary)",
          color: "var(--color-text-primary)",
        }}
      >
        <LanguageProvider>
          <DarkModeProvider>
            <AuthProvider>
              <ReduxProvider>
                <FirebaseSWRegistrar />
                <SharedLayout>{children}</SharedLayout>
              </ReduxProvider>
            </AuthProvider>
          </DarkModeProvider>
        </LanguageProvider>
      </body>
    </html>
  );
}
