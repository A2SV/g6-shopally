import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:shopallymobile/notification/local/local_notification.dart' show LocalNotifications;

class FcmHandler {
  static final _messaging = FirebaseMessaging.instance;

  static Future<void> init() async {
    await _messaging.requestPermission();
    // Foreground messages
    FirebaseMessaging.onMessage.listen((message) {
      final notification = message.notification;
      if (notification != null) {
        LocalNotifications.show(
          title: notification.title ?? 'Price Alert',
          body: notification.body ?? '',
        );
      }
    });
      
  

    

    // App opened from notification (foreground/background)
    FirebaseMessaging.onMessageOpenedApp.listen((message) {
      // TODO: handle deep links or navigate to product details using message.data
    });



    // Token refresh listener (inform backend or re-register alerts if needed)
    _messaging.onTokenRefresh.listen((token) {
      // TODO: notify backend about new token if you persist deviceId on server
    });




    // Background messages handled by Firebase automatically
    FirebaseMessaging.onBackgroundMessage(firebaseMessagingBackgroundHandler);
  }

  static Future<String?> getToken() async {
    return await _messaging.getToken();
  }
}




// Background handler (must be a top-level function)
Future<void> firebaseMessagingBackgroundHandler(RemoteMessage message) async {
  final notification = message.notification;
  if (notification != null) {
    await LocalNotifications.show(
      title: notification.title ?? 'Price Alert',
      body: notification.body ?? '',
    );
  }
}
