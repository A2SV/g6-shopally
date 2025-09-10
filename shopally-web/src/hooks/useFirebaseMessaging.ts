// src/hooks/useFirebaseMessaging.ts
"use client";

import { messaging } from "@/lib/firebase";
import { getToken, onMessage } from "firebase/messaging";
import { useEffect, useState } from "react";

const VAPID_KEY = process.env.NEXT_PUBLIC_FIREBASE_VAPID_KEY!;

interface FCMNotification {
  title?: string;
  body?: string;
  icon?: string;
}

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

function saveTokenToCookies(token: string) {
  document.cookie = `deviceId=${token}; path=/; max-age=${
    60 * 60 * 24 * 30
  }; secure; samesite=strict`;
  console.log("🍪 Saved FCM token into cookies");
}

export function useFirebaseMessaging() {
  const [token, setToken] = useState<string | null>(null);
  const [notification, setNotification] = useState<FCMNotification | null>(
    null
  );

  useEffect(() => {
    if (!messaging) {
      console.warn("⚠️ Firebase messaging is not available.");
      return;
    }
    const messagingInstance = messaging;

    const initFCM = async () => {
      const permission = await requestPermission();
      if (permission !== "granted") return;

      try {
        const currentToken = await getToken(messagingInstance, {
          vapidKey: VAPID_KEY,
        });
        if (currentToken) {
          console.log("✅ FCM Token:", currentToken);
          setToken(currentToken);
          saveTokenToCookies(currentToken);
        } else {
          console.warn("⚠️ No registration token available.");
        }
      } catch (err) {
        console.error("❌ Error retrieving FCM token:", err);
      }
    };

    initFCM();

    const unsubscribeOnMessage = onMessage(messagingInstance, (payload) => {
      console.log("📩 Foreground message received:", payload);
      setNotification(payload.notification ?? null);

      if (payload.notification?.title) {
        new Notification(payload.notification.title, {
          body: payload.notification.body,
          icon: payload.notification.icon ?? "/WebsiteLogo/Frame.png",
        });
      }
    });

    return () => unsubscribeOnMessage();
  }, []);

  return { token, notification };
}
