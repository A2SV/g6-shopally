//src/providers/ThemeProvider.tsx

"use client";

import { ThemeProvider, useTheme } from "next-themes";
import { useEffect } from "react";

function ThemeManager({ children }: { children: React.ReactNode }) {
  const { resolvedTheme } = useTheme();

  useEffect(() => {
    const html = document.documentElement;
    console.log("ThemeManager: resolvedTheme =", resolvedTheme);

    // Remove all theme classes first
    html.classList.remove("light", "dark");

    if (resolvedTheme === "dark") {
      html.classList.add("dark");
      console.log("ThemeManager: Added 'dark' class");
    } else {
      // For light mode or undefined, no class needed (default Tailwind behavior)
      console.log("ThemeManager: Light mode - no class added");
    }

    console.log("ThemeManager: Final HTML class =", html.className);
    console.log(
      "ThemeManager: HTML element classes:",
      html.classList.toString(),
    );
  }, [resolvedTheme]);

  return <>{children}</>;
}

export function ThemeProviders({ children }: { children: React.ReactNode }) {
  return (
    <ThemeProvider
      attribute="class"
      defaultTheme="dark"
      enableSystem={false}
      disableTransitionOnChange={false}
      storageKey="shopally-theme"
    >
      <ThemeManager>{children}</ThemeManager>
    </ThemeProvider>
  );
}
