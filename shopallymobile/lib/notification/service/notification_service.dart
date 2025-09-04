import 'package:shopallymobile/notification/local/local_notification.dart';

import '../datasource/remote/fcm_handler.dart';
import 'package:firebase_messaging/firebase_messaging.dart';

class NotificationsService {
  static Future<void> init() async {
    await LocalNotifications.init();
    await FcmHandler.init();

    // Register background handler
    FirebaseMessaging.onBackgroundMessage(firebaseMessagingBackgroundHandler);
  }
}
