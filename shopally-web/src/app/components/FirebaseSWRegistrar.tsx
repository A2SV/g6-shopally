"use client";

import { useEffect } from "react";

export default function FirebaseSWRegistrar() {
  useEffect(() => {
    if ("serviceWorker" in navigator) {
      navigator.serviceWorker
        .register("/firebase-messaging-sw.js")
        .then((registration) => {
          console.log("✅ Firebase Service Worker registered:", registration);
        })
        .catch((err) => {
          console.error("❌ Firebase Service Worker registration failed:", err);
        });
    }
  }, []);

  return null; // nothing visible, just runs on mount
}
