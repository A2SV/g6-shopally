// src/hooks/useFirebaseMessaging.ts
"use client";

import { messaging } from "@/lib/firebase";
import { getToken, onMessage } from "firebase/messaging";
import { useEffect, useState } from "react";

const VAPID_KEY = process.env.NEXT_PUBLIC_FIREBASE_VAPID_KEY!;

// Type for Firebase notification payload
interface FCMNotification {
  title?: string;
  body?: string;
  icon?: string;
}

/**
 * Ask user for browser notification permission.
 */
async function requestPermission(): Promise<NotificationPermission> {
  console.log("🔄 Requesting notification permission...");
  const permission = await Notification.requestPermission();

  if (permission === "granted")
    console.log("✅ Notification permission granted.");
  else if (permission === "denied")
    console.warn("❌ Notification permission denied.");
  else console.log("⚠️ Notification permission dismissed.");

  return permission;
}

/**
 * Custom React hook for Firebase Cloud Messaging
 */
export function useFirebaseMessaging() {
  const [token, setToken] = useState<string | null>(null);
  const [notification, setNotification] = useState<FCMNotification | null>(
    null
  );

  useEffect(() => {
    if (!messaging) return;

    const initFCM = async () => {
      const permission = await requestPermission();
      if (permission !== "granted") return;

      try {
        if (!messaging) return;

        const currentToken = await getToken(messaging, { vapidKey: VAPID_KEY });
        if (currentToken) {
          console.log("✅ FCM Token:", currentToken);
          setToken(currentToken);
        } else {
          console.warn("⚠️ No registration token available.");
        }
      } catch (err) {
        console.error("❌ Error retrieving FCM token:", err);
      }
    };

    initFCM();

    // Listen for foreground messages
    const unsubscribe = onMessage(messaging, (payload) => {
      console.log("📩 Foreground message received:", payload);
      setNotification(payload.notification ?? null);

      if (payload.notification?.title) {
        new Notification(payload.notification.title, {
          body: payload.notification.body,
          icon: payload.notification.icon ?? "/WebsiteLogo/Frame.png",
        });
      }
    });

    return () => unsubscribe();
  }, []);

  return { token, notification };
}
