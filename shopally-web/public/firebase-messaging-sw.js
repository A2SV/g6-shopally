// public/firebase-messaging-sw.js
importScripts(
  "https://www.gstatic.com/firebasejs/9.23.0/firebase-app-compat.js"
);
importScripts(
  "https://www.gstatic.com/firebasejs/9.23.0/firebase-messaging-compat.js"
);

// Your Firebase config again
firebase.initializeApp({
  apiKey: "AIzaSyAlycppLAEkwEUAB_OtYvzfazz7P-VQ94s",
  authDomain: "shopally-944a4.firebaseapp.com",
  projectId: "shopally-944a4",
  storageBucket: "shopally-944a4.firebasestorage.app",
  messagingSenderId: "999033337399",
  appId: "1:999033337399:web:1a51e362f1a4f67ac131b2",
  measurementId: "G-E4ZGHQT7R0",
});

const messaging = firebase.messaging();

// Handle background messages
messaging.onBackgroundMessage(function (payload) {
  console.log(
    "[firebase-messaging-sw.js] Received background message ",
    payload
  );
  const notificationTitle = payload.notification?.title || "Notification";
  const notificationOptions = {
    body: payload.notification?.body || "products info outdated",
    icon: "/WebsiteLogo/Frame.png",
  };

  self.registration.showNotification(notificationTitle, notificationOptions);
});
